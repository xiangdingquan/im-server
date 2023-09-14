package service

import (
	"context"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["deletegame"] = func(c botCallback) commandInterface {
		return &deleteGameCommand{
			botCallback: c,
		}
	}
}

type deleteGameCommand struct {
	botCallback
}

const (
	deleteGameMessage = "To delete a game, please send me the game or a game link (e.g., t.me/bot?game=game)"
)

func makeDeleteGameMessage(botId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         deleteGameMessage,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 64,
				Length: 18,
			}).To_MessageEntity(),
		},
	}).To_Message()
}

func (m *deleteGameCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	botMessage = makeDeleteGameMessage(model.BotFatherID, fromUserId)
	return
}

func (m *deleteGameCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
