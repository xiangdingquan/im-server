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

func (s *Service) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.createChat - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	var (
		chatUserIdList []int32
		chatTitle      = request.GetTitle()
		users          model.MutableUsers
		chat           *model.MutableChat
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.createChat - error: %v", err)
		return
	}

	// check chat title
	if chatTitle == "" {
		err = mtproto.ErrChatTitleEmpty
		log.Errorf("messages.createChat - error: %v", err)
		return
	}

	if len(request.Users) == 0 {
		err = mtproto.ErrUsersTooFew
		log.Errorf("messages.createChat - error: %v", err)
		return
	}

	// check user too much
	if len(request.GetUsers()) > 200-1 {
		err = mtproto.ErrUsersTooMuch
		log.Errorf("messages.createChat - error: %v", err)
		return
	}

	// s.UserFacade.GetMutableUsers(ctx, ...)
	// check len(users)
	chatUserIdList = make([]int32, 0, len(request.GetUsers()))
	for _, u := range request.Users {
		if u.PredicateName != mtproto.Predicate_inputUser {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("messages.createChat - error: %v", err)
			return
		} else {
			chatUserIdList = append(chatUserIdList, u.UserId)
		}
	}

	users = s.UserFacade.GetMutableUsers(ctx, append([]int32{md.UserId}, chatUserIdList...)...)
	if me, _ := users[md.UserId]; me.Restricted() {
		err = mtproto.ErrUserRestricted
		return
	}

	for _, u := range request.Users {
		if user, ok := users[u.UserId]; !ok {
			err = mtproto.ErrInputUserDeactivated
			log.Errorf("messages.createChat - error: %v", err)
			return
		} else {
			_ = user
		}
	}

	chat, err = s.ChatFacade.CreateChat2(ctx, md.UserId, chatUserIdList, chatTitle)
	if err != nil {
		log.Errorf("createChat duplicate: %v", err)
		return
	}

	reply, err = s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChatPeerUtil(chat.Chat.Id),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatCreate(chatTitle, chatUserIdList)),
			ScheduleDate: nil,
		})
	if err != nil {
		log.Errorf("messages.createChat - error: %v", err)
		return
	}

	log.Debugf("messages.createChat - reply: {%s}", reply.DebugString())
	return
}
