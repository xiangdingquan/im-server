package service

import (
	"context"
	"fmt"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func init() {
	cmdHandlers["mybots"] = func(c botCallback) commandInterface {
		return &myBotsCommand{
			botCallback: c,
		}
	}
}

type myBotsCommand struct {
	botCallback
}

func (m *myBotsCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	buttons := make([]*mtproto.KeyboardButton, 0)
	//
	m.getDao().WalkMyBots(ctx, fromUserId, func(botId int32, botUserName string) {
		buttons = append(buttons, mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
			Text: "@" + botUserName,
			Data: hack.Bytes(fmt.Sprintf("bots/%d", botId)),
		}).To_KeyboardButton())
	})

	if len(buttons) == 0 {
		botMessage = makeBotMessage(model.BotFatherID, fromUserId, "You have currently no bots")
	} else {
		botMessage = makeMyBotsMessage(model.BotFatherID, fromUserId, buttons)
	}
	return
}

func (m *myBotsCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}

const (
	myBotsMessage = "Choose a bot from the list below:"
)

func makeMyBotsMessage(botId, toId int32, buttons []*mtproto.KeyboardButton) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         myBotsMessage,
		ReplyMarkup: mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
			Rows: []*mtproto.KeyboardButtonRow{
				mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
					Buttons: buttons,
				}).To_KeyboardButtonRow(),
			},
		}).To_ReplyMarkup(),
	}).To_Message()
}
