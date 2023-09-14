package core

import (
	"context"
	"github.com/pkg/errors"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
	"sort"
)

type UserMessageReaction struct {
	UserId     int32 `json:"userId"`
	ReactionId int8  `json:"reactionId"`
}

type MessageReaction struct {
	Type         int8                  `json:"type"`
	ChatId       int64                 `json:"chatId"`
	MessageId    int32                 `json:"messageId"`
	ReactionList []UserMessageReaction `json:"reactionList"`
}

func (c *MessagesCore) SendReactionPrivate(ctx context.Context, reactionId, chatType int8, chatId int64, userId, messageId int32) ([]*MessageReaction, error) {
	var (
		serverMID = messageIdFromRequest(messageId)
		err       error
	)

	messages, err := c.MessagesDao.SelectUserMessages(ctx, userId, []int32{serverMID})
	if err != nil {
		return nil, err
	}
	if messages == nil {
		log.Errorf("SendReactionPrivate, message not found, userId:%d, messageId:%d", userId, messageId)
		return nil, errors.New("my message not found")
	}
	message := messages[0]

	do := &dbo.MessageReactionDo{
		Type:       chatType,
		ChatID:     message.DialogId,
		MessageId:  message.DialogMessageId,
		UserId:     userId,
		ReactionId: reactionId,
	}

	_, _, err = c.Insert(ctx, do)

	out := []*MessageReaction{
		{
			Type:      chatType,
			ChatId:    chatId,
			MessageId: messageId,
			ReactionList: []UserMessageReaction{
				{UserId: userId, ReactionId: reactionId},
			},
		},
	}

	messages, err = c.MessagesDao.SelectDialogMessages(ctx, int32(chatId), message.DialogId, []int32{message.DialogMessageId})
	if err != nil {
		return nil, err
	}
	if messages == nil {
		log.Errorf("SendReactionPrivate, message not found, his userId:%d, messageId:%d", int32(chatId), messageId)
		return nil, errors.New("his message not found")
	}
	message = messages[0]
	log.Debugf("message: %v", message)

	out = append(out, &MessageReaction{
		Type:      chatType,
		ChatId:    int64(userId),
		MessageId: messageIdToResponse(message.UserMessageBoxId),
		ReactionList: []UserMessageReaction{
			{UserId: userId, ReactionId: reactionId},
		},
	})

	return out, err
}

func (c *MessagesCore) SendReactionSuperGroup(ctx context.Context, reactionId, chatType int8, chatId int64, userId, messageId int32) (*MessageReaction, error) {
	var (
		serverMID = messageIdFromRequest(messageId)
		err       error
	)

	do := &dbo.MessageReactionDo{
		Type:       chatType,
		ChatID:     chatId,
		MessageId:  serverMID,
		UserId:     userId,
		ReactionId: reactionId,
	}

	_, _, err = c.Insert(ctx, do)
	out := &MessageReaction{
		Type:      chatType,
		ChatId:    chatId,
		MessageId: messageId,
		ReactionList: []UserMessageReaction{
			{UserId: userId, ReactionId: reactionId},
		},
	}
	return out, err
}

func (c *MessagesCore) GetMessagesReactionsPrivate(ctx context.Context, chatType int8, chatId int64, userId int32, messageIds []int32) ([]*MessageReaction, error) {
	var err error

	ids := make([]int32, len(messageIds))
	for i, v := range messageIds {
		ids[i] = messageIdFromRequest(v)
	}

	messages, err := c.MessagesDao.SelectUserMessages(ctx, userId, ids)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	dialogId := messages[0].DialogId
	ids = make([]int32, len(messages))
	for i, v := range messages {
		ids[i] = v.DialogMessageId
	}

	mrList, err := c.getMessagesReactions(ctx, chatType, dialogId, ids)

	dToM := make(map[int32]int32)
	for _, v := range messages {
		dToM[v.DialogMessageId] = v.UserMessageBoxId
	}

	for _, v := range mrList {
		v.MessageId = messageIdToResponse(dToM[v.MessageId])
	}
	return mrList, nil
}

func (c *MessagesCore) GetMessagesReactionsSuperGroup(ctx context.Context, chatType int8, chatId int64, userId int32, messageIds []int32) ([]*MessageReaction, error) {
	ids := make([]int32, len(messageIds))
	for i, v := range messageIds {
		ids[i] = messageIdFromRequest(v)
	}

	return c.getMessagesReactions(ctx, chatType, chatId, ids)
}

func (c *MessagesCore) getMessagesReactions(ctx context.Context, chatType int8, chatId int64, ids []int32) ([]*MessageReaction, error) {
	//log.Debugf("GetMessagesReactions chatType:%d, chatId:%d, ids:%v", chatType, chatId, ids)
	doList, err := c.SelectReaction(ctx, chatType, chatId, ids)
	if err != nil {
		return nil, err
	}

	for _, v := range doList {
		v.MessageId = messageIdToResponse(v.MessageId)
	}

	type key struct {
		chatType  int8
		chatId    int64
		messageId int32
	}
	m := make(map[key]*MessageReaction)

	for _, do := range doList {
		k := key{
			chatType:  do.Type,
			chatId:    do.ChatID,
			messageId: do.MessageId,
		}
		v, ok := m[k]
		if ok {
			v.ReactionList = append(v.ReactionList, UserMessageReaction{
				UserId:     do.UserId,
				ReactionId: do.ReactionId,
			})
		} else {
			m[k] = &MessageReaction{
				Type:      do.Type,
				ChatId:    do.ChatID,
				MessageId: do.MessageId,
				ReactionList: []UserMessageReaction{{
					UserId:     do.UserId,
					ReactionId: do.ReactionId,
				}},
			}
		}
	}

	out := make([]*MessageReaction, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out, nil
}

func messageIdFromRequest(messageId int32) int32 {
	return messageId >> 20
}

func messageIdToResponse(messageId int32) int32 {
	return messageId << 20
}

func (c *MessagesCore) getDialogMessages(ctx context.Context, userId int32, messageIds []int32) (map[int32]*dbo.MessagesDo, error) {
	if len(messageIds) == 0 {
		return nil, nil
	}

	sort.Sort(util.Int32Slice(messageIds))

	lowerBound, err := c.SelectLowerBound(ctx, userId, messageIds[0])
	if err != nil {
		return nil, err
	}

	upperBound, err := c.SelectUpperBound(ctx, userId, messageIds[len(messageIds)-1])
	if err != nil {
		return nil, err
	}

	messageMap, err := c.getDialogMessageBetween(ctx, userId, lowerBound, upperBound)
	if err != nil {
		return nil, err
	}

	out := make(map[int32]*dbo.MessagesDo)
	for _, v := range messageIds {
		m, ok := messageMap[v]
		if ok {
			out[v] = m
		}
	}

	log.Debugf("getDialogMessages, messageIds:%v", messageIds)
	log.Debugf("getDialogMessages, messages:%v", out)

	return out, nil
}

func (c *MessagesCore) getDialogMessageBetween(ctx context.Context, userId, lowerBound, upperBound int32) (map[int32]*dbo.MessagesDo, error) {
	sent, err := c.SelectUserSentMessages(ctx, userId, lowerBound, upperBound)
	if err != nil {
		return nil, err
	}

	received, err := c.SelectUserReceivedMessages(ctx, userId, lowerBound, upperBound)

	messages := append(sent, received...)
	if len(messages) == 0 {
		return nil, nil
	}

	startMessageId := messages[0].UserMessageBoxId
	out := make(map[int32]*dbo.MessagesDo)
	for i, v := range messages {
		out[startMessageId+int32(i)] = v
	}

	return out, nil
}
