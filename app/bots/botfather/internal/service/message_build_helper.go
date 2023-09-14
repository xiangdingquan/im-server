package service

import (
	"fmt"

	"time"

	"open.chat/app/pkg/env2"
	"open.chat/model"
	"open.chat/mtproto"
)

type MessageBuildItem struct {
	Text           string
	Param          string
	EntityType     string
	EntityUrl      string
	EntityUserId   int32
	EntityLanguage string
}

type MessageBuildHelper []MessageBuildItem

func (m MessageBuildHelper) MakeMessage(message *mtproto.Message) *mtproto.Message {
	if message == nil {
		message = mtproto.MakeTLMessage(&mtproto.Message{
			//
		}).To_Message()
	}

	if len(m) > 0 {
		var (
			offset int
			length int
		)
		for i := 0; i < len(m); i++ {
			message.Message = message.Message + m[i].Text
			offset = len(message.Message)
			length = len(m[i].Param)
			if length > 0 {
				switch m[i].EntityType {
				case mtproto.Predicate_messageEntityUnknown:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityUnknown(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityMention:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityMention(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityHashtag:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityHashtag(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityBotCommand:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityUrl:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityUrl(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityEmail:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityEmail(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityBold:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityItalic:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityItalic(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityCode:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityCode(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityPre:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityPre(&mtproto.MessageEntity{
						Offset:   int32(offset),
						Length:   int32(length),
						Language: m[i].EntityLanguage,
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityTextUrl:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityTextUrl(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
						Url:    m[i].EntityUrl,
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityMentionName:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityMentionName(&mtproto.MessageEntity{
						Offset:       int32(offset),
						Length:       int32(length),
						UserId_INT32: m[i].EntityUserId,
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityPhone:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityPhone(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityCashtag:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityCashtag(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityUnderline:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityUnderline(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityStrike:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityStrike(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				case mtproto.Predicate_messageEntityBlockquote:
					message.Entities = append(message.Entities, mtproto.MakeTLMessageEntityBlockquote(&mtproto.MessageEntity{
						Offset: int32(offset),
						Length: int32(length),
					}).To_MessageEntity())
				}
			}
			message.Message = message.Message + m[i].Param
		}
	}

	return message
}

func makeBotHelpMessage(botId, toId int32) *mtproto.Message {
	var mBuildHelper MessageBuildHelper = []MessageBuildItem{
		{
			Text:       "I can help you create and manage " + env2.MY_APP_NAME + " bots. If you're new to the Bot API, please ",
			Param:      "see the manual",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots", env2.MY_WEB_SITE),
		},
		{
			Text:       ".\n\nYou can control me by sending these commands:\n\n",
			Param:      "/newbot",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - create a new bot\n",
			Param:      "/mybots",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - edit your bots ",
			Param:      "[beta]",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       "\n\n",
			Param:      "Edit Bots",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       "\n",
			Param:      "/setname",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change a bot's name\n",
			Param:      "/setdescription",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change bot description\n",
			Param:      "/setabouttext",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change bot about info\n",
			Param:      "/setuserpic",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change bot profile photo\n",
			Param:      "/setcommands",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change the list of commands\n",
			Param:      "/deletebot",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - delete a bot\n\n",
			Param:      "Bot Settings",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       "\n",
			Param:      "/token",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - generate authorization token\n",
			Param:      "/revoke",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - revoke bot access token\n",
			Param:      "/setinline",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - toggle ",
			Param:      "inline mode",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots/inline", env2.MY_WEB_SITE),
		},
		{
			Text:       "\n",
			Param:      "/setinlinegeo",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - toggle inline ",
			Param:      "location requests",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots/inline#location-based-results", env2.MY_WEB_SITE),
		},
		{
			Text:       "\n",
			Param:      "/setinlinefeedback",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - change ",
			Param:      "inline feedback",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots/inline#collecting-feedback", env2.MY_WEB_SITE),
		},
		{
			Text:       " settings\n",
			Param:      "/setjoingroups",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - can your bot be added to groups?\n",
			Param:      "/setprivacy",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - toggle ",
			Param:      "privacy mode",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots#privacy-mode", env2.MY_WEB_SITE),
		},
		{
			Text:       " in groups\n\n",
			Param:      "Games",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       "\n",
			Param:      "/mygames",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - edit your ",
			Param:      "games",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots/games", env2.MY_WEB_SITE),
		},
		{
			Text:       " ",
			Param:      "[beta]",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       "\n",
			Param:      "/newgame",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - create a new ",
			Param:      "game",
			EntityType: mtproto.Predicate_messageEntityTextUrl,
			EntityUrl:  fmt.Sprintf("https://%s/bots/games", env2.MY_WEB_SITE),
		},
		{
			Text:       "\n",
			Param:      "/listgames",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - get a list of your games\n",
			Param:      "/editgame",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " - edit a game\n",
			Param:      "/deletegame",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text: " - delete an existing game",
		},
	}

	return mBuildHelper.MakeMessage(mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model.MakePeerUser(botId),
		ToId:            model.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         "",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities:        nil,
	}).To_Message())
}
