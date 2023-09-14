package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	model2 "open.chat/app/bots/botfather/internal/model"
	"open.chat/pkg/random2"

	"github.com/gogo/protobuf/types"

	"math/rand"

	"open.chat/app/bots/botpb"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/util"
)

func (s *Service) MessagesGetInlineBotResults(ctx context.Context, r *botpb.GetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	log.Debugf("messages.getInlineBotResults - request: {%s}", logger.JsonDebugData(r))

	err := mtproto.ErrMethodNotImpl

	log.Errorf("messages.getInlineBotResults - reply: {%v}", err)
	return nil, err
}

func (s *Service) MessagesQueryInlineBotResult(ctx context.Context, r *botpb.QueryInlineBotResult) (*mtproto.BotInlineResult, error) {
	log.Debugf("messages.getInlineBotResults - request: {%s}", logger.JsonDebugData(r))

	err := mtproto.ErrMethodNotImpl

	log.Errorf("messages.queryInlineBotResult - reply: {%v}", err)
	return nil, err
}

func (s *Service) MessagesGetBotCallbackAnswer(ctx context.Context, r *botpb.GetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	log.Debugf("messages.getBotCallbackAnswer - request: {%s}", logger.JsonDebugData(r))

	var (
		cbAnswer *mtproto.Messages_BotCallbackAnswer
		err      error
	)

	if r.IsGame {
		cbAnswer, err = s.DoGameCommand(ctx, r.UserId, r.BotId, r.Message, r.Data)
	} else {
		cbAnswer, err = s.DoBotCommand(ctx, r.UserId, r.BotId, r.Message, r.Data)
	}

	if err != nil {
		log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.getBotCallbackAnswer - reply: %s", cbAnswer.DebugString())
	return cbAnswer, nil
}

const (
	sendMessage = 0
	editMessage = 1
)

func (s *Service) DoBotCommand(ctx context.Context, fromUserId, botId int32, botMsg *mtproto.Message, data string) (answer *mtproto.Messages_BotCallbackAnswer, err error) {
	var (
		boxMsg   *mtproto.Message
		opResult = editMessage
	)

	cbAnswer := mtproto.MakeTLMessagesBotCallbackAnswer(&mtproto.Messages_BotCallbackAnswer{
		NativeUi: true,
	}).To_Messages_BotCallbackAnswer()

	log.Debugf("cmdLines: %v", data)
	cmdLines := strings.Split(data, "/")
	switch cmdLines[0] {
	case "bots":
		if len(cmdLines) == 1 {
			boxMsg, err = s.onBotsListCommand(ctx, fromUserId, botMsg)
			if err != nil {
				return
			}
		} else if len(cmdLines) == 2 {
			cmdBotId, _ := util.StringToInt32(cmdLines[1])
			boxMsg, err = s.onBotsCommand(ctx, fromUserId, cmdBotId, botMsg)
			if err != nil {
				return
			}
		} else if len(cmdLines) == 3 {
			cmdBotId, _ := util.StringToInt32(cmdLines[1])
			switch cmdLines[2] {
			case "tokn":
				boxMsg, err = s.onBotsTokenCommand(ctx, fromUserId, cmdBotId, botMsg)
				if err != nil {
					return
				}
			case "edit":
				boxMsg, err = s.onBotsEditListCommand(ctx, fromUserId, cmdBotId, botMsg)
				if err != nil {
					return
				}
			case "set":
				boxMsg, err = s.onBotsSettingsCommand(ctx, fromUserId, cmdBotId, botMsg)
				if err != nil {
					return
				}
			case "pay":
				boxMsg, err = s.onBotsPaymentsCommand(ctx, fromUserId, cmdBotId, botMsg)
				if err != nil {
					return
				}
			case "del":
				boxMsg, err = s.onBotsDelCommand(ctx, fromUserId, cmdBotId, botMsg)
				if err != nil {
					return
				}
			default:
				return
			}
		} else if len(cmdLines) == 4 {
			cmdBotId, _ := util.StringToInt32(cmdLines[1])
			switch cmdLines[2] {
			case "tokn":
				switch cmdLines[3] {
				case "revoke":
					boxMsg, err = s.onBotsTokenRevokeCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("revoke - %s", boxMsg.DebugString())
				default:
					return
				}
			case "edit":
				switch cmdLines[3] {
				case "name":
					boxMsg, err = s.onBotsEditNameCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
					cbAnswer.Message = &types.StringValue{Value: "OK. Send me the new name for your bot."}
					opResult = sendMessage
				case "desc":
					boxMsg, err = s.onBotsEditDescriptionCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
					cbAnswer.Message = &types.StringValue{Value: "OK. Send me the new description for the bot. People will see this description when they open a chat with your bot, in a block titled 'What can this bot do?'."}
					opResult = sendMessage
				case "about":
					boxMsg, err = s.onBotsEditAboutCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
					cbAnswer.Message = &types.StringValue{Value: "OK. Send me the new 'About' text. People will see this text on the bot's profile page and it will be sent together with a link to your bot when they share it with someone."}
					opResult = sendMessage
				case "pic":
					boxMsg, err = s.onBotsEditBotpicCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}

					log.Debugf("%s - %s", data, boxMsg.DebugString())
					cbAnswer.Message = &types.StringValue{Value: "OK. Send me the new profile photo for the bot."}
					opResult = sendMessage
				case "comm":
					boxMsg, err = s.onBotsEditCommandsCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}

					log.Debugf("%s - %s", data, boxMsg.DebugString())
					cbAnswer.Message = &types.StringValue{Value: "OK. Send me a list of commands for your bot. Please use this format:\n\ncommand1 - Description\ncommand2 - Another description\n\nSend /empty to keep the list empty."}
					opResult = sendMessage
				default:
					return
				}
			case "set":
				switch cmdLines[3] {
				case "inln":
					boxMsg, err = s.onBotsSettingsInlineCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
				case "grps":
					boxMsg, err = s.onBotsSettingsAllowGroupsCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
				case "priv":
					boxMsg, err = s.onBotsSettingsGroupPrivacyCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
				case "pay":
					boxMsg, err = s.onBotsSettingsPaymentsCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
				case "dmn":
					boxMsg, err = s.onBotsSettingsDomainCommand(ctx, fromUserId, cmdBotId, botMsg)
					if err != nil {
						return
					}
					log.Debugf("%s - %s", data, boxMsg.DebugString())
				default:
					return
				}
			default:
				return
			}
		}
	default:
		return
	}

	if botMsg != nil {
		if opResult == editMessage {
			s.MsgFacade.EditMessage(ctx,
				model.BotFatherId,
				0,
				model.MakeUserPeerUtil(fromUserId),
				&msgpb.OutboxMessage{
					NoWebpage:    false,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      boxMsg,
					ScheduleDate: nil,
				})
		} else {
			s.MsgFacade.SendMessage(ctx,
				model.BotFatherId,
				0,
				model.MakeUserPeerUtil(fromUserId),
				&msgpb.OutboxMessage{
					NoWebpage:    true,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      boxMsg,
					ScheduleDate: nil,
				})
		}
	}

	answer = cbAnswer
	return
}

func (s *Service) DoGameCommand(ctx context.Context, fromUserId, botId int32, msg *mtproto.Message, data string) (cbAnswer *mtproto.Messages_BotCallbackAnswer, err error) {
	return
}

func (s *Service) onBotsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if !botUser.GetBot() {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       fmt.Sprintf("Here it is: %s ", botUser.GetLastName().GetValue()),
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "API Token",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/tokn", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Bot Settings",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Payments",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/pay", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Delete Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/del", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bots List",
					Data: hack.Bytes("bots"),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsListCommand(ctx context.Context, fromId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	buttons := make([]*mtproto.KeyboardButton, 0)

	err := s.Dao.WalkMyBots(ctx, fromId, func(botId int32, botUserName string) {
		buttons = append(buttons, mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
			Text: "@" + botUserName,
			Data: hack.Bytes(fmt.Sprintf("bots/%d", botId)),
		}).To_KeyboardButton())
	})

	if err != nil {
		return nil, err
	}

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Choose a bot from the list below:",
		ReplyMarkup: mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
			Rows: []*mtproto.KeyboardButtonRow{
				mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
					Buttons: buttons,
				}).To_KeyboardButtonRow(),
			},
		}).To_ReplyMarkup(),
		Entities: nil,
		EditDate: editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsTokenCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	token, err := s.Dao.GetBotToken(ctx, fromId, cmdBotId)
	if err != nil {
		return nil, err
	}

	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if err != nil {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       fmt.Sprintf("Here is the token for bot: %s ", botUser.GetLastName().GetValue()),
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text:       "\n\n",
			Param:      fmt.Sprintf("%d:%s", cmdBotId, token),
			EntityType: mtproto.Predicate_messageEntityCode,
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Revoke current token",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/tokn/revoke", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup: mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
			Rows: rows,
		}).To_ReplyMarkup(),
		Entities: entities,
		EditDate: editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsTokenRevokeCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if err != nil {
		return nil, err
	}

	token := random2.RandomAlphanumeric(35)
	s.Dao.UpdateBotToken(ctx, cmdBotId, token)

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       fmt.Sprintf("Token for the bot %s ", botUser.GetUsername().GetValue()),
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text:       " has been revoked. New token is:\n\n",
			Param:      fmt.Sprintf("%d:%s", cmdBotId, token),
			EntityType: mtproto.Predicate_messageEntityCode,
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsEditListCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if !botUser.GetBot() {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       "Edit ",
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text:       " info.\n\n",
			Param:      "Name",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       fmt.Sprintf(": %s\n", botUser.GetFirstName().GetValue()),
			Param:      "Description",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       ": 🚫\n",
			Param:      "About",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       ": 🚫\n",
			Param:      "Botpic",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       ": 🚫 no botpic\n",
			Param:      "Commands",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text: ": no commands yet",
		},
	})
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit Name",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit/name", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit Description",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit/desc", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit About",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit/about", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit Botpic",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit/pic", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Edit Commands",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/edit/comm", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsEditNameCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	states := &model2.BotFatherCommandStates{
		MainCmd:    "setname",
		NextSubCmd: "setname",
		CacheSubCmdResults: map[string]string{
			"selected_bot_id": strconv.Itoa(int(cmdBotId)),
			"edit_msg_id":     strconv.Itoa(int(editMsg.Id)),
		},
	}

	s.Dao.PutBotFatherCommandStates(ctx, fromId, states)

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(model.BotFatherId),
		ToId:            model.MakePeerUser(fromId),
		Date:            int32(time.Now().Unix()),
		Message:         "OK. Send me the new name for your bot.",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message(), nil
}

func (s *Service) onBotsEditDescriptionCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	states := &model2.BotFatherCommandStates{
		MainCmd:    "setdescription",
		NextSubCmd: "setdescription",
		CacheSubCmdResults: map[string]string{
			"selected_bot_id": strconv.Itoa(int(cmdBotId)),
			"edit_msg_id":     strconv.Itoa(int(editMsg.Id)),
		},
	}

	s.Dao.PutBotFatherCommandStates(ctx, fromId, states)

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(model.BotFatherId),
		ToId:            model.MakePeerUser(fromId),
		Date:            int32(time.Now().Unix()),
		Message:         "OK. Send me the new description for the bot. People will see this description when they open a chat with your bot, in a block titled 'What can this bot do?'.",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message(), nil
}

func (s *Service) onBotsEditAboutCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	states := &model2.BotFatherCommandStates{
		MainCmd:    "setabouttext",
		NextSubCmd: "setabouttext",
		CacheSubCmdResults: map[string]string{
			"selected_bot_id": strconv.Itoa(int(cmdBotId)),
			"edit_msg_id":     strconv.Itoa(int(editMsg.Id)),
		},
	}

	s.Dao.PutBotFatherCommandStates(ctx, fromId, states)

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(model.BotFatherId),
		ToId:            model.MakePeerUser(fromId),
		Date:            int32(time.Now().Unix()),
		Message:         "OK. Send me the new 'About' text. People will see this text on the bot's profile page and it will be sent together with a link to your bot when they share it with someone.",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message(), nil
}

func (s *Service) onBotsEditBotpicCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	states := &model2.BotFatherCommandStates{
		MainCmd:    "setuserpic",
		NextSubCmd: "setuserpic",
		CacheSubCmdResults: map[string]string{
			"selected_bot_id": strconv.Itoa(int(cmdBotId)),
			"edit_msg_id":     strconv.Itoa(int(editMsg.Id)),
		},
	}

	s.Dao.PutBotFatherCommandStates(ctx, fromId, states)

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(model.BotFatherId),
		ToId:            model.MakePeerUser(fromId),
		Date:            int32(time.Now().Unix()),
		Message:         "OK. Send me the new profile photo for the bot.",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message(), nil
}

func (s *Service) onBotsEditCommandsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	states := &model2.BotFatherCommandStates{
		MainCmd:    "setcommands",
		NextSubCmd: "setcommands",
		CacheSubCmdResults: map[string]string{
			"selected_bot_id": strconv.Itoa(int(cmdBotId)),
			"edit_msg_id":     strconv.Itoa(int(editMsg.Id)),
		},
	}

	s.Dao.PutBotFatherCommandStates(ctx, fromId, states)

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(model.BotFatherId),
		ToId:            model.MakePeerUser(fromId),
		Date:            int32(time.Now().Unix()),
		Message:         "OK. Send me a list of commands for your bot. Please use this format:\n\ncommand1 - Description\ncommand2 - Another description\n\nSend /empty to keep the list empty.",
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 130,
				Length: 6,
			}).To_MessageEntity(),
		},
		ReplyMarkup: mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if !botUser.GetBot() {
		return nil, err
	}
	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       "Settings for ",
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text: ".\n",
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Inline Mode",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set/inln", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Allow Groups?",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set/grps", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Group Privacy",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set/priv", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Payments",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/pay", cmdBotId)),
				}).To_KeyboardButton(),
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Domain",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set/dmn", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsInlineCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if err != nil {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       "",
			Param:      "Inline",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  "https://core.telegram.org/bots/inline",
		},
		{
			Text:       fmt.Sprintf(" mode is currently disabled for %s ", botUser.GetFirstName().GetValue()),
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Settings",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/set", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsAllowGroupsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Allows Groups Disabled",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsInlineLocationDataCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Inline Location Data Disabled.",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsInlineFeedbackCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Inline Feedback Disabled",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsGroupPrivacyCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Group Privacy Disabled.",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsPaymentsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Payments Disabled",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsSettingsDomainCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         "Domain Disabled.",
		ReplyMarkup:     replyMarkup,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsPaymentsCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if !botUser.GetBot() {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       "Payment providers for ",
			Param:      botUser.GetFirstName().GetValue(),
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       " ",
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text: ".\n\nNo payment methods connected. Use the buttons below to add a new integration.",
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}

func (s *Service) onBotsDelCommand(ctx context.Context, fromId, cmdBotId int32, editMsg *mtproto.Message) (*mtproto.Message, error) {
	botUser, err := s.UserFacade.GetUserById(ctx, fromId, cmdBotId)
	if !botUser.GetBot() {
		return nil, err
	}

	message, entities := model.MakeTextAndMessageEntities([]model.MessageBuildEntry{
		{
			Text:       "You are about to delete your bot ",
			Param:      botUser.GetFirstName().GetValue(),
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       " ",
			Param:      "@" + botUser.GetUsername().GetValue(),
			EntityType: mtproto.Predicate_messageEntityMention,
		},
		{
			Text: ". Is that correct?",
		},
	})

	rows := []*mtproto.KeyboardButtonRow{
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Yes, delete the bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d/del/yes", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "Nope, nevermind",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "No",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
		mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
			Buttons: []*mtproto.KeyboardButton{
				mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: "« Back to Bot",
					Data: hack.Bytes(fmt.Sprintf("bots/%d", cmdBotId)),
				}).To_KeyboardButton(),
			},
		}).To_KeyboardButtonRow(),
	}

	replyMarkup := mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
		Rows: rows,
	}).To_ReplyMarkup()

	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             editMsg.Out,
		Id:              editMsg.Id,
		FromId_FLAGPEER: editMsg.FromId_FLAGPEER,
		PeerId:          editMsg.PeerId,
		ToId:            editMsg.ToId,
		Date:            editMsg.Date,
		Message:         message,
		ReplyMarkup:     replyMarkup,
		Entities:        entities,
		EditDate:        editMsg.EditDate,
	}).To_Message(), nil
}
