package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getFullChat#3b831c66 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	chat, err := s.ChatFacade.GetMutableChat(ctx, request.ChatId, md.UserId)
	if err != nil {
		log.Errorf("messages.getFullChat - error: %v", err)
		return nil, err
	}

	messagesChatFull := mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: chat.ToUnsafeChatFull(md.UserId),
		Chats:    []*mtproto.Chat{chat.ToUnsafeChat(md.UserId)},
		Users:    nil,
	}).To_Messages_ChatFull()

	messagesChatFull.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, append(chat.ToChatParticipantIdList(), md.UserId))

	log.Debugf("messages.getFullChat#3b831c66 - reply: %s", messagesChatFull.DebugString())
	return messagesChatFull, nil
}
