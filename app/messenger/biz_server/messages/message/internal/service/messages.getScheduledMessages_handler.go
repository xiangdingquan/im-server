package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetScheduledMessages(ctx context.Context, request *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getScheduledMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_SELF:
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
		channel, err := s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId)
		if channel == nil {
			log.Errorf("messages.getScheduledMessages - error: %v", err)
			return nil, err
		}

		if channel.Channel.IsBroadcast() {
			me := channel.GetImmutableChannelParticipant(md.UserId)
			if me == nil || !me.IsCreatorOrAdmin() {
				err := mtproto.ErrChatAdminRequired
				log.Errorf("messages.getScheduledMessages - error: %v", err)
				return nil, err
			}
		}
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getScheduledMessages - error: %v", err)
		return nil, err
	}

	msgBoxList := s.MessageFacade.GetScheduledMessageListByIdList(ctx, md.UserId, peer, request.Id)
	messages, users, chats := msgBoxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	messagesMessages := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: messages,
		Users:    users,
		Chats:    chats,
	}).To_Messages_Messages()

	log.Debugf("messages.getScheduledMessages - reply: %s", messagesMessages.DebugString())
	return messagesMessages, nil
}
