package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/gogo/protobuf/types"
	relay_client "open.chat/app/interface/relay/client"
	"open.chat/app/interface/relay/relaypb"
	sync_client "open.chat/app/messenger/sync/client"

	"github.com/golang/protobuf/proto"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func makePhoneCall(userSelfId int32, phoneCallData *mtproto.PhoneCall) *mtproto.TLPhoneCall {

	var (
		callData = proto.Clone(phoneCallData).(*mtproto.PhoneCall)
	)

	if userSelfId == phoneCallData.AdminId {
		callData.GAOrB = phoneCallData.GB
	} else {
		callData.GAOrB = phoneCallData.GAHash
	}

	return &mtproto.TLPhoneCall{
		Data2: callData,
	}
}

func (s *Service) PhoneConfirmCall(ctx context.Context, request *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.confirmCall - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_ALREADY_DECLINED	The call was already declined
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.confirmCall - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()
	callSession, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callSession == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	callSession.GA = request.GA
	callSession.KeyFingerprint = request.KeyFingerprint

	libraryVersion := model.CalcPhoneCallLibraryVersion(
		callSession.ParticipantProtocol.LibraryVersions,
		callSession.AdminProtocol.LibraryVersions)

	var callConnections *relaypb.CallConnections
	if libraryVersion == model.MaxPhoneCallLibraryVersion {
		callConnections = &relaypb.CallConnections{
			Id:      rand.Int63(),
			PeerTag: []byte{},
			Connection: mtproto.MakeTLPhoneConnectionWebrtc(&mtproto.PhoneConnection{
				Turn:     true,
				Stun:     true,
				Id:       callSession.Id,
				Ip:       "127.0.0.1",
				Ipv6:     "",
				Port:     3478,
				Username: "dong.chat",
				Password: "12345678",
			}).To_PhoneConnection(),
			AlternativeConnections: []*mtproto.PhoneConnection{},
		}
	} else {
		callConnections, err = relay_client.CreateCallSession(ctx, callSession.Id)
		if err != nil {
			log.Errorf(err.Error())
			return nil, err
		}
		if callConnections.AlternativeConnections == nil {
			callConnections.AlternativeConnections = []*mtproto.PhoneConnection{}
		}
	}

	users := s.UserFacade.GetMutableUsers(ctx, md.UserId, md.UserId, callSession.ParticipantId)
	if len(users) != 2 {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	s.Dao.PutPhoneCallSession(ctx, peer.GetId(), callSession)

	phoneCall := mtproto.MakeTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: mtproto.MakeTLPhoneCall(&mtproto.PhoneCall{
			P2PAllowed:     false,
			Video:          callSession.Video,
			Id:             callSession.Id,
			AccessHash:     callSession.AccessHash,
			Date:           int32(callSession.Date),
			AdminId:        callSession.AdminId,
			ParticipantId:  callSession.ParticipantId,
			GAOrB:          callSession.GB,
			KeyFingerprint: request.KeyFingerprint,
			Protocol: mtproto.MakeTLPhoneCallProtocol(&mtproto.PhoneCallProtocol{
				UdpP2P:          callSession.AdminProtocol.UdpP2P,
				UdpReflector:    callSession.AdminProtocol.UdpReflector,
				MinLayer:        callSession.AdminProtocol.MinLayer,
				MaxLayer:        callSession.AdminProtocol.MaxLayer,
				LibraryVersions: []string{libraryVersion},
			}).To_PhoneCallProtocol(),
			ReceiveDate:            &types.Int32Value{Value: int32(time.Now().Unix())},
			Connections:            []*mtproto.PhoneConnection{callConnections.Connection},
			Connection:             callConnections.Connection,
			AlternativeConnections: callConnections.AlternativeConnections,
			StartDate:              0,
		}).To_PhoneCall(),
		Users: users.GetUsersByIdList(md.UserId, []int32{md.UserId, callSession.ParticipantId}),
	}).To_Phone_PhoneCall()
	// phoneCall.PhoneCall.Connections = []*mtproto.PhoneConnection{phoneConnection}
	// phoneCall.PhoneCall.Connection = phoneConnection
	// phoneCall.PhoneCall.AlternativeConnections = callConnections.AlternativeConnections

	log.Debugf("phone.acceptCall#3bd2b4a0 - reply: %s", phoneCall.DebugString())
	return model.WrapperGoFunc(phoneCall, func() {
		// phoneCall := callSession.ToPhoneCall(callSession.ParticipantId, request.GetKeyFingerprint())
		// phoneCall.Connection = callConnections.Connection
		// phoneCall.Connections = []*mtproto.PhoneConnection{phoneConnection}
		// phoneCall.AlternativeConnections = callConnections.AlternativeConnections

		pushUpdates := model.MakeUpdatesByUpdatesUsers(
			users.GetUsersByIdList(callSession.ParticipantId, []int32{md.UserId, callSession.AdminId}),
			mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
				PhoneCall: mtproto.MakeTLPhoneCall(
					&mtproto.PhoneCall{
						P2PAllowed:     false,
						Video:          callSession.Video,
						Id:             callSession.Id,
						AccessHash:     callSession.AccessHash,
						Date:           int32(callSession.Date),
						AdminId:        callSession.AdminId,
						ParticipantId:  callSession.ParticipantId,
						GAOrB:          callSession.GA,
						KeyFingerprint: request.KeyFingerprint,
						Protocol: mtproto.MakeTLPhoneCallProtocol(&mtproto.PhoneCallProtocol{
							UdpP2P:          callSession.ParticipantProtocol.UdpP2P,
							UdpReflector:    callSession.ParticipantProtocol.UdpReflector,
							MinLayer:        callSession.ParticipantProtocol.MinLayer,
							MaxLayer:        callSession.ParticipantProtocol.MaxLayer,
							LibraryVersions: []string{libraryVersion},
						}).To_PhoneCallProtocol(),
						ReceiveDate:            &types.Int32Value{Value: int32(time.Now().Unix())},
						Connections:            []*mtproto.PhoneConnection{callConnections.Connection},
						Connection:             callConnections.Connection,
						AlternativeConnections: callConnections.AlternativeConnections,
						StartDate:              0,
					}).To_PhoneCall(),
			}).To_Update())
		//pushUpdates.Users = users.GetUsersByIdList(callSession.ParticipantId, []int32{md.UserId, callSession.AdminId})
		sync_client.SyncUpdatesMe(
			context.Background(),
			callSession.ParticipantId,
			callSession.ParticipantAuthKeyId,
			0,
			"",
			pushUpdates)
	}).(*mtproto.Phone_PhoneCall), nil
}
