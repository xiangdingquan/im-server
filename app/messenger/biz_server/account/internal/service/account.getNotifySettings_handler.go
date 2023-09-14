package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getNotifySettings#12b3ad31 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err      error
		settings *mtproto.PeerNotifySettings
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.unregisterDevice - error: %v", err)
		return nil, err
	}

	peer := model.FromInputNotifyPeer(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_USER:
		// check peerUser Exists
	case model.PEER_CHAT:
		// check peerChat exists
	case model.PEER_CHANNEL:
		// check peerChannel exists
	case model.PEER_USERS:
	case model.PEER_CHATS:
	case model.PEER_BROADCASTS:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("account.updateNotifySettings - error: %v", err)
		return nil, err
	}

	settings, err = s.UserFacade.GetNotifySettings(ctx, md.UserId, peer)
	if err != nil {
		log.Errorf("getNotifySettings error - %v", err)
		return nil, err
	}

	log.Debugf("account.getNotifySettings#12b3ad31 - reply: %s", settings.DebugString())
	return settings, err
}
