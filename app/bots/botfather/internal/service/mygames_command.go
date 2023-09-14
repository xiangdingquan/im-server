package service

import (
	"context"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["mygames"] = func(c botCallback) commandInterface {
		return &myGamesCommand{
			botCallback: c,
		}
	}
}

type myGamesCommand struct {
	botCallback
}

func (m *myGamesCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	botMessage = makeNoGamesMessage(model.BotFatherID, fromUserId)
	return
}

func (m *myGamesCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
