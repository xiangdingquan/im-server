package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (s *Service) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getAllChats - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	messagesChats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: nil,
	}).To_Messages_Chats()

	chats := s.ChatFacade.GetAllChats(ctx, md.UserId)
	for _, chat := range chats {
		if ok, _ := util.Contains(chat.GetId(), request.ExceptIds); !ok {
			messagesChats.Chats = append(messagesChats.Chats, chat)
		}
	}
	channels := s.ChannelFacade.GetAllChannels(ctx, md.UserId)

	for _, channel := range channels {
		if ok, _ := util.Contains(channel.GetId(), request.ExceptIds); !ok {
			messagesChats.Chats = append(messagesChats.Chats, channel.ToUnsafeChat(md.UserId))
		}
	}

	log.Debugf("messages.getAllChats - reply: %s", messagesChats.DebugString())
	return messagesChats, nil
}
