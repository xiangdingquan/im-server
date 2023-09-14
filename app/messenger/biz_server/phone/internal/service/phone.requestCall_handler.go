package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneRequestCall(ctx context.Context, request *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.requestCall - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		ok            bool
		users         model.MutableUsers
		user          *model.ImmutableUser
		err           error
		participantId int32
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	peer := model.FromInputUser(md.UserId, request.UserId)
	switch peer.PeerType {
	case model.PEER_USER:
		participantId = peer.PeerId
	default:
		err = mtproto.ErrUserIdInvalid
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	//400	BOT_METHOD_INVALID	This method can't be used by a bot
	//400	CALL_PROTOCOL_FLAGS_INVALID	Call protocol flags invalid
	//400	PARTICIPANT_VERSION_OUTDATED	The other participant does not use an up to date telegram client with support for calls
	//400	USER_ID_INVALID	The provided user ID is invalid
	//403	USER_IS_BLOCKED	You were blocked by this user
	//403	USER_PRIVACY_RESTRICTED	The user's privacy settings do not allow you to do this

	protocol := request.GetProtocol()
	// 400	CALL_PROTOCOL_FLAGS_INVALID	Call protocol flags invalid
	// 400	PARTICIPANT_VERSION_OUTDATED	The other participant does not use an up to date telegram client with support for calls
	if protocol.MinLayer < model.MinPhoneCallLayer || protocol.MaxLayer > model.MaxPhoneCallLayer {
		err = mtproto.ErrParticipantVersionOutdated
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	users = s.UserFacade.GetMutableUsers(ctx, md.UserId, participantId)

	if user, ok = users.GetImmutableUser(participantId); !ok {
		err = mtproto.ErrUserIdInvalid
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	if s.UserFacade.IsBlockedByUser(ctx, md.UserId, participantId) {
		err = mtproto.ErrUserIsBlocked
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	// 403	USER_PRIVACY_RESTRICTED	The user's privacy settings do not allow you to do this
	restricted := s.UserFacade.CheckPrivacy(ctx, int(model.PHONE_CALL), participantId, md.UserId, user.CheckContact(md.UserId))
	if !restricted {
		err = mtproto.ErrUserPrivacyRestricted
		log.Errorf("phone.requestCall - error: %v", err)
		return nil, err
	}

	callSession, err := s.Dao.CreatePhoneCallSession(
		ctx,
		request.Video,
		md.UserId,
		md.AuthId,
		participantId,
		int64(request.RandomId),
		request.GAHash,
		request.Protocol)

	if err != nil {
		log.Errorf("phone.requestCall - createPhoneCallSession error: %v", err)
		return nil, err
	}

	callWaiting := mtproto.MakeTLPhoneCallWaiting(&mtproto.PhoneCall{
		Video:         callSession.Video,
		Id:            callSession.Id,
		AccessHash:    callSession.AccessHash,
		Date:          int32(callSession.Date),
		AdminId:       callSession.AdminId,
		ParticipantId: callSession.ParticipantId,
		Protocol: mtproto.MakeTLPhoneCallProtocol(&mtproto.PhoneCallProtocol{
			UdpP2P:          callSession.AdminProtocol.UdpP2P,
			UdpReflector:    callSession.AdminProtocol.UdpReflector,
			MinLayer:        callSession.AdminProtocol.MinLayer,
			MaxLayer:        callSession.AdminProtocol.MaxLayer,
			LibraryVersions: []string{model.MinPhoneCallLibraryVersion},
		}).To_PhoneCallProtocol(),
		ReceiveDate: nil,
	}).To_PhoneCall()
	phoneCall := mtproto.MakeTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: callWaiting,
		Users:     users.GetUsersByIdList(md.UserId, []int32{md.UserId, participantId}),
	}).To_Phone_PhoneCall()

	log.Debugf("phone.requestCall - reply: {%v}", phoneCall)
	return model.WrapperGoFunc(phoneCall, func() {
		callRequested := mtproto.MakeTLPhoneCallRequested(&mtproto.PhoneCall{
			Video:         callSession.Video,
			Id:            callSession.Id,
			AccessHash:    callSession.AccessHash,
			Date:          int32(callSession.Date),
			AdminId:       callSession.AdminId,
			ParticipantId: callSession.ParticipantId,
			GAHash:        callSession.GAHash,
			Protocol: mtproto.MakeTLPhoneCallProtocol(&mtproto.PhoneCallProtocol{
				UdpP2P:          protocol.UdpP2P,
				UdpReflector:    protocol.UdpReflector,
				MinLayer:        protocol.MinLayer,
				MaxLayer:        protocol.MaxLayer,
				LibraryVersions: []string{model.MinPhoneCallLibraryVersion},
			}).To_PhoneCallProtocol(),
		}).To_PhoneCall()

		pushUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
				PhoneCall: callRequested,
			}).To_Update()},
			Users: users.GetUsersByIdList(participantId, []int32{md.UserId, participantId}),
		}).To_Updates()
		sync_client.PushUpdates(context.Background(), participantId, pushUpdates)
	}).(*mtproto.Phone_PhoneCall), nil
}
