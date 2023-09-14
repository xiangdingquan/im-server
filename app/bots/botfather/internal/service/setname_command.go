package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func init() {
	cmdHandlers["setname"] = func(c botCallback) commandInterface {
		return &setNameCommand{
			botCallback: c,
		}
	}
}

type setNameCommand struct {
	botCallback
}

const (
	setNameMessage = "Choose a bot to change name."
)

func (m *setNameCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
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
		botMessage = makeSetMessage(model.BotFatherID, fromUserId, setNameMessage, buttons)
	}

	states.MainCmd = "setname"
	states.NextSubCmd = "username"
	r = OpSave

	return
}

func (m *setNameCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
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

		states.NextSubCmd = "setname"
		states.CacheSubCmdResults["selected_bot_id"] = strconv.Itoa(int(sBotId))

		botMessage = mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
			ToId:            model2.MakePeerUser(fromUserId),
			Date:            int32(time.Now().Unix()),
			Message:         "OK. Send me the new name for your bot.",
			ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		}).To_Message()
		r = OpSave
	case "setname":
		sBotId, _ := strconv.Atoi(states.CacheSubCmdResults["selected_bot_id"])
		if editMsgId, ok := states.CacheSubCmdResults["edit_msg_id"]; ok {
			id, _ := strconv.ParseInt(editMsgId, 10, 64)
			editMsg2, err := m.getMessage().GetUserMessage(ctx, model.BotFatherID, int32(id))
			if err != nil {
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
			rows := []*mtproto.KeyboardButtonRow{
				mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
					Buttons: []*mtproto.KeyboardButton{
						mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
							Text: "« Back to Bot",
							Data: hack.Bytes(fmt.Sprintf("bots/%d", sBotId)),
						}).To_KeyboardButton(),
						mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
							Text: "« Back to Bots List",
							Data: []byte("bots"),
						}).To_KeyboardButton(),
					},
				}).To_KeyboardButtonRow(),
			}

			replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
				Rows: rows,
			}).To_ReplyMarkup()

			editMsg := editMsg2.ToMessage(model.BotFatherID)
			botMessage = mtproto.MakeTLMessage(&mtproto.Message{
				Out:             editMsg.Out,
				Id:              editMsg.Id,
				FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
				PeerId:          editMsg.PeerId,
				ToId:            editMsg.ToId,
				Date:            editMsg.Date,
				Message:         "Success! Name updated. /help",
				Entities: []*mtproto.MessageEntity{
					mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
						Offset: 23,
						Length: 5,
					}).To_MessageEntity(),
				},
				ReplyMarkup: replyMarkup,
				EditDate:    editMsg.EditDate,
			}).To_Message()
			r = sendMessage<<16 | OpDelete
		} else {
			botMessage = mtproto.MakeTLMessage(&mtproto.Message{
				Out:             true,
				FromId_FLAGPEER: model2.MakePeerUser(model.BotFatherID),
				ToId:            model2.MakePeerUser(fromUserId),
				Date:            int32(time.Now().Unix()),
				Message:         "Success! Name updated. /help",
				Entities: []*mtproto.MessageEntity{
					mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
						Offset: 23,
						Length: 5,
					}).To_MessageEntity(),
				},
			}).To_Message()
			r = OpDelete
		}
		m.getUser().UpdateFirstAndLastName(ctx, int32(sBotId), msg.Message, "")
	}
	return
}
