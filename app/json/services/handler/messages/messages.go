package messages

import (
	"context"
	"fmt"
	"math"
	"open.chat/app/json/consts"
	"open.chat/app/json/db/dbo"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/handler"
	"open.chat/app/json/services/handler/messages/core"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"strconv"
	"strings"
)

type cls struct {
	*core.MessagesCore
}

type BatchSendMessageJson struct {
	ID      int32   `json:"ID"`
	Message string  `json:"message"`
	Users   []int32 `json:"users"`
}

func New(s *svc.Service) {
	service := &cls{
		MessagesCore: core.New(nil),
	}
	s.AppendServices(handler.RegisterMessages(service))
}

func (s *cls) doToJson(do *dbo.BatchSendMessageDo) *BatchSendMessageJson {
	l := strings.Split(do.ToUsers, ",")
	out := &BatchSendMessageJson{
		ID:      int32(do.ID),
		Message: do.Message,
	}
	out.Users = make([]int32, len(l))
	for i, s := range l {
		u, err := strconv.Atoi(s)
		if err == nil {
			out.Users[i] = int32(u)
		}
	}
	return out
}

func (s *cls) BatchSend(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TBatchSend) *helper.ResultJSON {
	if len(r.Users) > 200 || len(r.Users) == 0 {
		return &helper.ResultJSON{Code: 400, Msg: "to many users"}
	}

	if len(r.Message) > 200 || len(r.Message) == 0 {
		return &helper.ResultJSON{Code: 401, Msg: "message to long"}
	}

	do := &dbo.BatchSendMessageDo{
		UID:     uint32(md.UserId),
		Message: r.Message,
		ToUsers: strings.Trim(strings.Replace(fmt.Sprint(r.Users), " ", ",", -1), "[]"),
	}
	lastID, _, err := s.MessagesCore.InsertBatchSendMessage(ctx, do)

	if err != nil {
		log.Errorf("BatchSend, save failed, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "save batch send message failed"}
	}
	do.ID = uint32(lastID)

	push := &helper.PushUpdate{
		Action: consts.ActionMessageBatchSendOnSend,
		From:   uint32(md.UserId),
	}
	_ = push.ToUsers(ctx, []uint32{uint32(md.UserId)}, s.doToJson(do))

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetBatchSend(ctx context.Context, md *grpc_util.RpcMetadata, request *handler.TGetBatchSend) *helper.ResultJSON {
	fromId := request.FromId
	if fromId == 0 {
		fromId = math.MaxInt32
	}
	doList, err := s.MessagesCore.SelectBatchSendMessage(ctx, md.UserId, fromId, request.Limit)
	if err != nil {
		log.Errorf("GetBatchSend, query failed, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "get batch send message failed"}
	}

	l := make([]*BatchSendMessageJson, len(doList))
	for i, do := range doList {
		l[i] = s.doToJson(do)
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: l}
}

func (s *cls) ClearBatchSend(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	_, err := s.MessagesCore.DeleteBatchSendMessageDelete(ctx, uint32(md.UserId))
	if err != nil {
		log.Errorf("ClearBatchSend, delete failed, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "clear batch send message failed"}
	}

	push := &helper.PushUpdate{Action: consts.ActionMessageBatchSendDelete, From: uint32(md.UserId)}
	_ = push.ToUsers(ctx, []uint32{uint32(md.UserId)}, nil)

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) SendReaction(ctx context.Context, md *grpc_util.RpcMetadata, request *handler.TSendReaction) *helper.ResultJSON {
	log.Debugf("messages.sendReaction userId:%d, request:%v", md.UserId, request)

	if request.Type == 1 {
		return s.sendReactionPrivate(ctx, request.ReactionId, request.Type, request.ChatId, md.UserId, request.MessageId)
	} else {
		return s.sendReactionSuperGroup(ctx, request.ReactionId, request.Type, request.ChatId, md.UserId, request.MessageId)
	}
}

func (s *cls) sendReactionSuperGroup(ctx context.Context, reactionId, chatType int8, chatId int64, userId int32, messageId int32) *helper.ResultJSON {
	reaction, err := s.MessagesCore.SendReactionSuperGroup(ctx, reactionId, chatType, chatId, userId, messageId)
	if err != nil {
		log.Errorf("messages.sendReaction, sendReactionSuperGroup failed: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "send failed"}
	}

	push := &helper.PushUpdate{
		Action: consts.ActionMessageReactionOnUpdate,
		From:   uint32(userId),
	}

	err = push.ToChannel(ctx, uint32(chatId), reaction)

	if err != nil {
		log.Errorf("messages.sendReaction, push failed: %v", err)
		return &helper.ResultJSON{Code: -2, Msg: "push failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) sendReactionPrivate(ctx context.Context, reactionId, chatType int8, chatId int64, userId int32, messageId int32) *helper.ResultJSON {
	reactionList, err := s.MessagesCore.SendReactionPrivate(ctx, reactionId, chatType, chatId, userId, messageId)
	if err != nil {
		log.Errorf("messages.sendReaction, sendReactionPrivate failed: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "send failed"}
	}

	push := func(from, to uint32, reaction *core.MessageReaction) error {
		push := &helper.PushUpdate{
			Action: consts.ActionMessageReactionOnUpdate,
			From:   from,
		}
		return push.ToUsers(ctx, []uint32{to}, reaction)
	}
	err = push(uint32(chatId), uint32(userId), reactionList[0])
	if err != nil {
		log.Errorf("messages.sendReaction, push failed: %v", err)
		return &helper.ResultJSON{Code: -2, Msg: "push failed"}
	}

	err = push(uint32(userId), uint32(chatId), reactionList[1])
	if err != nil {
		log.Errorf("messages.sendReaction, push failed: %v", err)
		return &helper.ResultJSON{Code: -3, Msg: "push failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetMessagesReactions(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TGetMessagesReactions) *helper.ResultJSON {
	var (
		l   []*core.MessageReaction
		err error
	)
	if r.Type == 1 {
		l, err = s.MessagesCore.GetMessagesReactionsPrivate(ctx, r.Type, r.ChatId, md.UserId, r.MessageIds)
	} else {
		l, err = s.MessagesCore.GetMessagesReactionsSuperGroup(ctx, r.Type, r.ChatId, md.UserId, r.MessageIds)
	}

	if err != nil {
		log.Errorf("messages.getMessagesReactions, get failed: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "get failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: l}
}
