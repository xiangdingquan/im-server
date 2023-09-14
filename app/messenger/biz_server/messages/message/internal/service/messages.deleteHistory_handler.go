package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.deleteHistory#1c015b09 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer = model.FromInputPeer2(md.UserId, request.Peer)
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.deleteHistory - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF:
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.deleteHistory - error: %v", err)
		return nil, err
	}

	affectedHistory, err := s.MsgFacade.DeleteHistory(ctx,
		md.UserId,
		md.AuthId,
		request.JustClear,
		request.Revoke,
		peer,
		request.MaxId)

	if err != nil {
		log.Errorf("messages.deleteHistory - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.deleteHistory - reply: %s", affectedHistory.DebugString())
	return affectedHistory, nil
}
