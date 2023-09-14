package service

import (
	"context"
	"fmt"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/app/pkg/env2"
	model2 "open.chat/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["cancel"] = func(c botCallback) commandInterface {
		return &cancelCommand{
			botCallback: c,
		}
	}
}

type cancelCommand struct {
	botCallback
}

func (m *cancelCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	var mBuildHelper MessageBuildHelper = []MessageBuildItem{
		{
			Text:       "The command deletebot has been cancelled. Anything else I can do for you?\n\nSend ",
			Param:      "/help",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text: fmt.Sprintf(" for a list of commands. To learn more about %s Bots, see https://%s/bots",
				env2.MY_APP_NAME,
				env2.MY_WEB_SITE),
		},
	}

	botMessage = mBuildHelper.MakeMessage(mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
		ToId:            model2.MakePeerUser(fromUserId),
		Date:            int32(time.Now().Unix()),
		Message:         "",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities:        nil,
	}).To_Message())

	r = OpDelete
	return
}

func (m *cancelCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	return
}
