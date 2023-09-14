package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
)

func init() {
	cmdHandlers["setcommands"] = func(c botCallback) commandInterface {
		return &setCommandsCommand{
			botCallback: c,
		}
	}
}

type setCommandsCommand struct {
	botCallback
}

const (
	setCommandsMessage = "Choose a bot to change the list of commands."
)

func (m *setCommandsCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
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
		botMessage = makeSetMessage(model.BotFatherID, fromUserId, setCommandsMessage, buttons)
	}

	states.MainCmd = "setcommands"
	states.NextSubCmd = "username"
	r = OpSave

	return
}

func (m *setCommandsCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	switch states.NextSubCmd {
	case "username":
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

		var sBotId int
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

		states.NextSubCmd = "setcommands"
		states.CacheSubCmdResults["selected_bot_id"] = strconv.Itoa(int(sBotId))

		botMessage = mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
			ToId:            model2.MakePeerUser(fromUserId),
			Date:            int32(time.Now().Unix()),
			Message:         "OK. Send me a list of commands for your bot. Please use this format:\n\ncommand1 - Description\ncommand2 - Another description\n\nSend /empty to keep the list empty.",
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
					Offset: 130,
					Length: 6,
				}).To_MessageEntity(),
			},
			ReplyMarkup: mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		}).To_Message()
		r = OpSave
	case "setcommands":
		var cmdList []model.CommandInfo
		lines := strings.Split(msg.Message, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			idx := strings.Index(line, " - ")
			if idx < 1 {
				botMessage = mtproto.MakeTLMessage(&mtproto.Message{
					Out:             true,
					FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
					ToId:            model2.MakePeerUser(fromUserId),
					Date:            int32(time.Now().Unix()),
					Message:         "Sorry, the list of commands is invalid. Please use this format:\n\ncommand1 - Description\ncommand2 - Another description",
					ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
				}).To_Message()
				return
			}
			cmdList = append(cmdList, model.CommandInfo{
				CmdName:     line[:idx],
				Description: line[idx+3:],
			})
		}

		sBotId, _ := strconv.Atoi(states.CacheSubCmdResults["selected_bot_id"])
		m.getDao().SetBotCommands(ctx, int32(sBotId), cmdList)

		botMessage = mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
			ToId:            model2.MakePeerUser(fromUserId),
			Date:            int32(time.Now().Unix()),
			Message:         "Success! Command list updated. /help",
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
					Offset: 31,
					Length: 5,
				}).To_MessageEntity(),
			},
		}).To_Message()
		r = OpDelete
	}
	return
}
