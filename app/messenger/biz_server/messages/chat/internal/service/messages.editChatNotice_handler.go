package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditChatNotice(ctx context.Context, request *mtproto.TLMessagesEditChatNotice) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editChatNotice - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.editChatNotice - error: %v", err)
		return nil, err
	}

	peer := model.FromInputPeer2(md.UserId, request.Peer)

	switch peer.PeerType {
	case model.PEER_CHAT:
		_, err := s.ChatFacade.EditChatNotice(ctx, peer.PeerId, md.UserId, request.Notice)
		if err != nil {
			log.Errorf("messages.editChatNotice - error: %v", err)
			return nil, err
		}
	case model.PEER_CHANNEL:
		_, err := s.ChannelFacade.EditNotice(ctx, peer.PeerId, md.UserId, request.Notice)
		if err != nil {
			log.Errorf("messages.editChatNotice - error: %v", err)
			return nil, err
		}
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	log.Debug("messages.editChatNotice - reply: {true}")
	return mtproto.ToBool(true), nil
}
