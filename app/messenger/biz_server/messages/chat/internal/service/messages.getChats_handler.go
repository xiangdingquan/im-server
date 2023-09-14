package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getChats#3c6aa187 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	chats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: make([]*mtproto.Chat, 0, len(request.Id)),
	}).To_Messages_Chats()

	for _, id := range request.Id {
		chat, _ := s.ChatFacade.GetMutableChat(ctx, id, md.UserId)
		if chat != nil {
			chats.Chats = append(chats.Chats, chat.ToUnsafeChat(md.UserId))
		}
	}

	log.Debugf("messages.getChats#3c6aa187 - reply: %s", chats.DebugString())
	return chats, nil
}
