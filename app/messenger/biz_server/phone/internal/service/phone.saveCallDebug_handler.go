package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneSaveCallDebug(ctx context.Context, request *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.saveCallDebug - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	// 400	DATA_JSON_INVALID	The provided JSON data is invalid	//
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.saveCallDebug - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callSession == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	s.Dao.SetCallDebug(ctx, callSession.Id, md.UserId, md.AuthId, request.GetDebug().GetData())

	log.Debugf("phone.saveCallDebug#277add7e - reply: {true}")
	return mtproto.ToBool(true), nil
}
