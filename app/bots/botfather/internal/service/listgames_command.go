package service

import (
	"context"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["listgames"] = func(c botCallback) commandInterface {
		return &listGamesCommand{
			botCallback: c,
		}
	}
}

type listGamesCommand struct {
	botCallback
}

func (m *listGamesCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	botMessage = makeNoBotsMessage(model.BotFatherID, fromUserId)
	return
}

func (m *listGamesCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
