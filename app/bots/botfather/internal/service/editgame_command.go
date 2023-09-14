package service

import (
	"context"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["editgame"] = func(c botCallback) commandInterface {
		return &editGameCommand{
			botCallback: c,
		}
	}
}

type editGameCommand struct {
	botCallback
}

const (
	editGameMessage = "Editing games. Please send me a game or a game link (e.g., t.me/bot?game=game)."
)

func makeEditGameMessage(botId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         editGameMessage,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 59,
				Length: 18,
			}).To_MessageEntity(),
		},
	}).To_Message()
}

func (m *editGameCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	botMessage = makeEditGameMessage(model.BotFatherID, fromUserId)
	return
}

func (m *editGameCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
