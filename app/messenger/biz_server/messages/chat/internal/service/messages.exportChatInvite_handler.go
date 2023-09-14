package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.exportChatInvite - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer *model.PeerUtil
		link string
		err  error
	)

	switch request.Constructor {
	case mtproto.CRC32_messages_exportChatInvite_7d885289:
		peer = model.MakeChatPeerUtil(request.ChatId)
	case mtproto.CRC32_messages_exportChatInvite_df7534c:
		peer = model.FromInputPeer2(md.UserId, request.Peer)
	default:
		err := mtproto.ErrTypeConstructorInvalid
		log.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_CHAT:
		link, err = s.ChatFacade.ExportChatInvite(ctx, peer.PeerId, md.UserId)
		if err != nil {
			log.Errorf("messages.exportChatInvite - error: ", err)
			return nil, err
		}
	case model.PEER_CHANNEL:
		channel, err := s.ChannelFacade.ExportChannelInvite(ctx, peer.PeerId, md.UserId)
		if err != nil {
			log.Errorf("messages.exportChatInvite - error: ", err)
			return nil, err
		}
		link = channel.Channel.Link
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	exportedChatInvite := mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
		Link: link,
	}).To_ExportedChatInvite()

	log.Debugf("messages.exportChatInvite - reply: %s", exportedChatInvite.DebugString())
	return exportedChatInvite, nil
}
