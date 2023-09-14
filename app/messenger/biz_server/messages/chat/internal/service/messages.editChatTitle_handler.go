package service

import (
	"context"

	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("messages.editChatTitle#dc452855 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Title == "" {
		err := mtproto.ErrChatTitleEmpty
		log.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	chat, err := s.ChatFacade.EditChatTitle(ctx, request.ChatId, md.UserId, request.Title)
	if err != nil {
		log.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	replyUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChatPeerUtil(request.ChatId),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatEditTitle(request.Title)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("messages.editChatTitle - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.editChatTitle - reply: {%s}", replyUpdates.DebugString())
	return replyUpdates, nil
}
