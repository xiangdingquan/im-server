package service

import (
	"context"
	"open.chat/app/pkg/env2"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["newgame"] = func(c botCallback) commandInterface {
		return &newGameCommand{
			botCallback: c,
		}
	}
}

type newGameCommand struct {
	botCallback
}

func (m *newGameCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	var newGameMessage = "Serving specific pages to particular " + env2.MY_APP_NAME + " users is an extremely powerful tool that allows you to build cool HTML5 games and integrated interfaces. But power of this kind also requires a lot of responsibility on the part of bot developers. Please read our Rules carefully, and accept them only if you agree with them."

	buttons := []*mtproto.KeyboardButton{
		mtproto.MakeTLKeyboardButton(&mtproto.KeyboardButton{
			Text: "OK",
		}).To_KeyboardButton(),
	}
	botMessage = makeSetMessage(model.BotFatherID, fromUserId, newGameMessage, buttons)
	return
}

func (m *newGameCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
