package service

import (
	"context"

	"time"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.toggleChatAdmins - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		enabled = mtproto.FromBool(request.Enabled)
	)

	chat, err := s.ChatFacade.ToggleChatAdmins(ctx, request.ChatId, md.UserId, enabled)
	if err != nil {
		log.Errorf("messages.toggleChatAdmins - error: ", err)
		return nil, err
	}

	replyUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{chat.ToUnsafeChat(md.UserId)},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	log.Debugf("messages.toggleChatAdmins - reply: {%s}", replyUpdates.DebugString())
	return replyUpdates, nil
}
