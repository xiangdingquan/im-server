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

func (s *Service) PhoneReceivedCall(ctx context.Context, request *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.receivedCall - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_ALREADY_DECLINED	The call was already declined
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.receivedCall - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()
	callSession, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callSession == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	// 检查是否在通话中
	if callSession.State > model.CallStateReceived {
		log.Warnf("call received: {%v}", peer)
		err = mtproto.ErrCallAlreadyDeclined
		return nil, err
	} else {
		callSession.State = model.CallStateReceived
		s.Dao.PutPhoneCallSession(ctx, callSession.Id, callSession)
	}

	/////////////////////////////////////////////////////////////////////////////////
	log.Debugf("phone.receivedCall#17d54f61 - reply {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
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
			ReceiveDate: &types.Int32Value{Value: int32(time.Now().Unix())},
		}).To_PhoneCall()

		pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
			PhoneCall: callWaiting,
		}).To_Update())

		pushUpdates.Users = s.UserFacade.GetUserListByIdList(
			context.Background(),
			callSession.AdminId,
			[]int32{md.UserId, callSession.AdminId})
		sync_client.SyncUpdatesMe(context.Background(),
			callSession.AdminId,
			callSession.AdminAuthKeyId,
			0,
			"",
			pushUpdates)
	}).(*mtproto.Bool), nil
}
