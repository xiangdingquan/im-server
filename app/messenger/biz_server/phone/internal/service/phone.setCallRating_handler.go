package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneSetCallRating(ctx context.Context, request *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.setCallRating - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.setCallRating - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callSession == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	s.Dao.SetCallRating(ctx, callSession.Id, md.UserId, md.AuthId,
		request.GetUserInitiative(), request.GetRating(), request.GetComment())

	rUpdates := model.MakeEmptyUpdates()
	log.Debugf("phone.setCallRating#1c536a34 - reply: {%s}", rUpdates.DebugString())
	return rUpdates, nil
}
