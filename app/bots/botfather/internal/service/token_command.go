package service

import (
	"context"
	"fmt"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func init() {
	cmdHandlers["token"] = func(c botCallback) commandInterface {
		return &tokenCommand{
			botCallback: c,
		}
	}
}

type tokenCommand struct {
	botCallback
}

const (
	tokenMessage = "Choose a bot to generate a new token."
)

func (m *tokenCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	buttons := make([]*mtproto.KeyboardButton, 0)
	m.getDao().WalkMyBots(ctx, fromUserId, func(botId int32, botUserName string) {
		buttons = append(buttons, mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
			Text: "@" + botUserName,
			Data: hack.Bytes(fmt.Sprintf("bots/%d", botId)),
		}).To_KeyboardButton())
	})

	if len(buttons) == 0 {
		botMessage = makeNoBotsMessage(model.BotFatherID, fromUserId)
	} else {
		botMessage = makeSetMessage(model.BotFatherID, fromUserId, tokenMessage, buttons)
	}
	return
}

func (m *tokenCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
