package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneAcceptCall(ctx context.Context, request *mtproto.TLPhoneAcceptCall) (*mtproto.Phone_PhoneCall, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.acceptCall - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_ALREADY_ACCEPTED	The call was already accepted
	// 400	CALL_ALREADY_DECLINED	The call was already declined
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	// 400	CALL_PROTOCOL_FLAGS_INVALID	Call protocol flags invalid
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.acceptCall - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callSession == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	callSession.ParticipantProtocol = request.Protocol
	callSession.ParticipantAuthKeyId = md.AuthId
	callSession.GB = request.GB

	users := s.UserFacade.GetMutableUsers(ctx, md.UserId, md.UserId, callSession.AdminId)
	if len(users) != 2 {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	libraryVersion := model.CalcPhoneCallLibraryVersion(
		callSession.ParticipantProtocol.LibraryVersions,
		callSession.AdminProtocol.LibraryVersions)
	if libraryVersion == model.MinPhoneCallLibraryVersion {
		callSession.Video = false
	}

	err = s.Dao.PutPhoneCallSession(ctx, peer.GetId(), callSession)
	if err != nil {
		err := mtproto.ErrCallPeerInvalid
		log.Errorf("phone.acceptCall - error: %v", err)
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
			UdpP2P:          callSession.ParticipantProtocol.UdpP2P,
			UdpReflector:    callSession.ParticipantProtocol.UdpReflector,
			MinLayer:        callSession.ParticipantProtocol.MinLayer,
			MaxLayer:        callSession.ParticipantProtocol.MaxLayer,
			LibraryVersions: []string{libraryVersion},
		}).To_PhoneCallProtocol(),
		ReceiveDate: &types.Int32Value{Value: int32(time.Now().Unix())},
	}).To_PhoneCall()

	phoneCall := mtproto.MakeTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: callWaiting,
		Users:     users.GetUsersByIdList(md.UserId, []int32{md.UserId, callSession.AdminId}),
	}).To_Phone_PhoneCall()

	log.Debugf("phone.acceptCall#3bd2b4a0 - reply: {%v}", phoneCall)
	return model.WrapperGoFunc(phoneCall, func() {
		syncNotMe := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
			PhoneCall: mtproto.MakeTLPhoneCallDiscarded(&mtproto.PhoneCall{
				NeedRating: false,
				NeedDebug:  false,
				Video:      callSession.Video,
				Id:         callSession.Id,
				Reason:     nil,
				Duration:   nil,
			}).To_PhoneCall(),
		}).To_Update())
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, syncNotMe)

		pushUpdates := model.MakeUpdatesByUpdatesUsers(
			users.GetUsersByIdList(callSession.AdminId, []int32{md.UserId, callSession.AdminId}),
			mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
				PhoneCall: mtproto.MakeTLPhoneCallAccepted(&mtproto.PhoneCall{
					Video:         callSession.Video,
					Id:            callSession.Id,
					AccessHash:    callSession.AccessHash,
					Date:          int32(callSession.Date),
					AdminId:       callSession.AdminId,
					ParticipantId: callSession.ParticipantId,
					GB:            callSession.GB,
					Protocol: mtproto.MakeTLPhoneCallProtocol(&mtproto.PhoneCallProtocol{
						UdpP2P:          callSession.AdminProtocol.UdpP2P,
						UdpReflector:    callSession.AdminProtocol.UdpReflector,
						MinLayer:        callSession.AdminProtocol.MinLayer,
						MaxLayer:        callSession.AdminProtocol.MaxLayer,
						LibraryVersions: []string{libraryVersion},
					}).To_PhoneCallProtocol(),
				}).To_PhoneCall(),
			}).To_Update())
		pushUpdates.Users = users.GetUsersByIdList(callSession.AdminId, []int32{md.UserId, callSession.AdminId})
		sync_client.SyncUpdatesMe(
			context.Background(),
			callSession.AdminId,
			callSession.AdminAuthKeyId,
			0,
			"",
			pushUpdates)
	}).(*mtproto.Phone_PhoneCall), nil
}
