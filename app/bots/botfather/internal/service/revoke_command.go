package service

import (
	"context"
	"fmt"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/random2"
)

func init() {
	cmdHandlers["revoke"] = func(c botCallback) commandInterface {
		return &revokeCommand{
			botCallback: c,
		}
	}
}

type revokeCommand struct {
	botCallback
}

const (
	revokeMessage = "Choose a bot to generate a new token. Warning: your old token will stop working."
)

func (m *revokeCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	buttons := make([]*mtproto.KeyboardButton, 0)
	//
	m.getDao().WalkMyBots(ctx, fromUserId, func(botId int32, botUserName string) {
		buttons = append(buttons, mtproto.MakeTLKeyboardButton(&mtproto.KeyboardButton{
			Text: "@" + botUserName,
		}).To_KeyboardButton())
	})

	if len(buttons) == 0 {
		botMessage = makeNoBotsMessage(model.BotFatherID, fromUserId)
	} else {
		botMessage = makeSetMessage(model.BotFatherID, fromUserId, revokeMessage, buttons)
	}

	states.MainCmd = "revoke"
	states.NextSubCmd = "revoke"
	r = OpSave

	return
}

func (m *revokeCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	switch states.NextSubCmd {
	case "revoke":
		if len(msg.Message) == 0 || msg.Message[0] != '@' {
			botMessage = mtproto.MakeTLMessage(&mtproto.Message{
				Out:             true,
				FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
				ToId:            model2.MakePeerUser(fromUserId),
				Date:            int32(time.Now().Unix()),
				Message:         "Invalid bot selected.",
				ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
			}).To_Message()
			return
		}

		var (
			sBotId int
		)
		m.getDao().WalkMyBots(ctx, fromUserId, func(botId int32, botUserName string) {
			if botUserName == msg.Message[1:] {
				sBotId = int(botId)
			}
		})
		if sBotId == 0 {
			botMessage = mtproto.MakeTLMessage(&mtproto.Message{
				Out:             true,
				FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
				ToId:            model2.MakePeerUser(fromUserId),
				Date:            int32(time.Now().Unix()),
				Message:         "Invalid bot selected.",
				ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
			}).To_Message()
			return
		}

		token := random2.RandomAlphanumeric(35)
		m.getDao().UpdateBotToken(ctx, int32(sBotId), token)

		botMessage = mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
			ToId:            model2.MakePeerUser(fromUserId),
			Date:            int32(time.Now().Unix()),
			Message:         fmt.Sprintf("Your token was replaced with a new one. You can use this token to access HTTP API:\n%d:%s", sBotId, token),
			ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		}).To_Message()
		r = OpDelete
	}
	return
}
