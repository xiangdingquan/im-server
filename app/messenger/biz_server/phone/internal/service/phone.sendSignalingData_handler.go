package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneSendSignalingData(ctx context.Context, request *mtproto.TLPhoneSendSignalingData) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.sendSignalingData - metadata: %s, request: %s", md.DebugString(), request.DebugString())

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

	// peerId
	pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePhoneCallSignalingData(&mtproto.Update{
		PhoneCallId:    callSession.Id,
		Data_FLAGBYTES: request.GetData(),
	}).To_Update())

	if md.UserId == callSession.AdminId {
		sync_client.SyncUpdatesMe(ctx, callSession.ParticipantId, callSession.ParticipantAuthKeyId, 0, "", pushUpdates)
	} else {
		sync_client.SyncUpdatesMe(ctx, callSession.AdminId, callSession.AdminAuthKeyId, 0, "", pushUpdates)
	}

	log.Debugf("phone.sendSignalingData - reply: {true})")
	return mtproto.BoolTrue, nil
}
