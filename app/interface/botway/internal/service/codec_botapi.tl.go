package service

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gogo/protobuf/types"
	"open.chat/app/interface/botway/botapi"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func EncodeToTLMethod(m botapi.BotApiMessage) mtproto.TLObject {
	var r mtproto.TLObject

	switch bm := m.(type) {
	case *botapi.GetUpdates2:
	case *botapi.SetWebhook2:
	case *botapi.DeleteWebhook2:
	case *botapi.GetWebhookInfo2:

	case *botapi.GetMe2:
		r = encodeToGetMe(bm)
	case *botapi.SendMessage2:
		r = encodeToSendMessage(bm)
	}
	return r
}

func encodeToGetMe(bm *botapi.GetMe2) *mtproto.TLUsersGetUsers {
	_ = bm
	return &mtproto.TLUsersGetUsers{
		Constructor: mtproto.CRC32_users_getUsers,
		Id:          []*mtproto.InputUser{mtproto.MakeTLInputUserSelf(nil).To_InputUser()},
	}
}

func encodeToSendMessage(bm *botapi.SendMessage2) *mtproto.TLMessagesSendMessage {
	r := &mtproto.TLMessagesSendMessage{
		Constructor:  mtproto.CRC32_messages_sendMessage_520c3870,
		NoWebpage:    bm.DisableWebPagePreview,
		Silent:       bm.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         nil,
		ReplyToMsgId: nil, //
		Message:      bm.Text,
		RandomId:     rand.Int63(),
		ReplyMarkup:  nil,
		Entities:     nil,
		ScheduleDate: nil,
	}
	if bm.ReplyToMessageId != 0 {
		r.ReplyToMsgId = &types.Int32Value{Value: bm.ReplyToMessageId}
	}
	if r.ReplyMarkup != nil {
	}
	return r
}

func encodeToReplyMarkup(bm *botapi.ReplyMarkup) *mtproto.ReplyMarkup {
	if bm == nil {
		return nil
	}

	var (
		r *mtproto.ReplyMarkup
	)

	if bm.InlineKeyboard != nil {
		var rows []*mtproto.KeyboardButtonRow
		for _, kRows := range bm.InlineKeyboard {
			buttons := make([]*mtproto.KeyboardButton, 0, len(bm.InlineKeyboard))
			for _, k := range kRows {
				buttons = append(buttons, mtproto.MakeTLKeyboardButtonCallback(&mtproto.KeyboardButton{
					Text: k.Text,
					Data: hack.Bytes(k.CallbackData),
				}).To_KeyboardButton())
			}
			rows = append(rows, mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
				Buttons: buttons,
			}).To_KeyboardButtonRow())
		}
		r = mtproto.MakeTLReplyInlineMarkup(&mtproto.ReplyMarkup{
			Rows: rows,
		}).To_ReplyMarkup()
	} else if bm.Keyboard != nil {
		var rows []*mtproto.KeyboardButtonRow
		for _, kRows := range bm.Keyboard {
			buttons := make([]*mtproto.KeyboardButton, 0, len(bm.Keyboard))
			for _, k := range kRows {
				buttons = append(buttons, mtproto.MakeTLKeyboardButton(&mtproto.KeyboardButton{
					Text: k.Text,
				}).To_KeyboardButton())
			}
			rows = append(rows, mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
				Buttons: buttons,
			}).To_KeyboardButtonRow())
		}
		r = mtproto.MakeTLReplyKeyboardMarkup(&mtproto.ReplyMarkup{
			Resize:    bm.ResizeKeyboard,
			SingleUse: bm.OneTimeKeyboard,
			Selective: bm.Selective,
			Rows:      rows,
		}).To_ReplyMarkup()
	} else if bm.ForceReply {
		r = mtproto.MakeTLReplyKeyboardForceReply(&mtproto.ReplyMarkup{
			SingleUse: bm.OneTimeKeyboard,
			Selective: bm.Selective,
		}).To_ReplyMarkup()
	} else if bm.RemoveKeyboard {
	}

	return r
}

func DecodeToBotApi(m mtproto.TLObject) botapi.BotApiMessage {
	var r botapi.BotApiMessage

	switch o := m.(type) {
	case *mtproto.User:
		r = decodeToUser(o)
	case *mtproto.Updates:
		r = updatesToUpdateList(0, o)
	}
	//
	return r
}

func decodeToUser(m *mtproto.User) *botapi.User {
	return &botapi.User{
		Id:           m.Id,
		IsBot:        m.Bot,
		FirstName:    m.GetFirstName().GetValue(),
		LastName:     m.GetLastName().GetValue(),
		Username:     m.GetUsername().GetValue(),
		LanguageCode: m.GetLangCode().GetValue(),
	}
}

func decodeToChatByUser(m *mtproto.User) *botapi.Chat {
	var chat *botapi.Chat
	if m != nil {
		chat = &botapi.Chat{
			Id:        botapi.ChatIdPrivate.ToChatId(m.Id),
			Type:      "private",
			Username:  m.GetUsername().GetValue(),
			FirstName: m.GetFirstName().GetValue(),
			LastName:  m.GetLastName().GetValue(),
		}
	}
	return chat
}

func decodeToChatByChat(m *mtproto.Chat) *botapi.Chat {
	var chat *botapi.Chat

	switch m.GetPredicateName() {
	case mtproto.Predicate_chat:
		chat = &botapi.Chat{
			Id:       botapi.ChatIdGroup.ToChatId(m.Id),
			Type:     "group",
			Title:    m.Title,
			Username: m.GetUsername().GetValue(),
		}
	case mtproto.Predicate_channel:
		chat = &botapi.Chat{
			Type:     "",
			Title:    m.Title,
			Username: m.GetUsername().GetValue(),
		}
		if m.Broadcast {
			// Id:       m.Id,
			chat.Id = botapi.ChatIdChannel.ToChatId(m.Id)
			chat.Type = "channel"
		} else {
			chat.Id = botapi.ChatIdSuperGroup.ToChatId(m.Id)
			chat.Type = "supergroup"
		}
	}

	return chat
}

func findUserByUserList(id int32, users []*mtproto.User) *mtproto.User {
	var user *mtproto.User
	for _, u := range users {
		if u.Id == id {
			user = u
		}
	}
	return user
}

func findChatByChatList(id int32, chats []*mtproto.Chat) *mtproto.Chat {
	var chat *mtproto.Chat
	for _, c := range chats {
		if c.Id == id {
			chat = c
		}
	}
	return chat
}

func messageToMessage(m *mtproto.Message, users []*mtproto.User, chats []*mtproto.Chat) *botapi.Message {
	message := &botapi.Message{
		MessageId:             m.Id,
		From:                  nil,
		Date:                  m.Date,
		Chat:                  nil,
		ForwardFrom:           nil,
		ForwardFromChat:       nil,
		ForwardFromMessageId:  0,
		ForwardSignature:      "",
		ForwardSenderName:     "",
		ForwardDate:           0,
		ReplyToMessage:        nil,
		EditDate:              0,
		MediaGroupId:          "",
		AuthorSignature:       "",
		Text:                  m.Message,
		Entities:              nil,
		CaptionEntities:       nil,
		Audio:                 nil,
		Document:              nil,
		Animation:             nil,
		Game:                  nil,
		Photo:                 nil,
		Sticker:               nil,
		Video:                 nil,
		Voice:                 nil,
		VideoNote:             nil,
		Caption:               "",
		Contact:               nil,
		Location:              nil,
		Venue:                 nil,
		Poll:                  nil,
		NewChatMembers:        nil,
		LeftChatMember:        nil,
		NewChatTitle:          "",
		NewChatPhoto:          nil,
		DeleteChatPhoto:       false,
		GroupChatCreated:      false,
		SupergroupChatCreated: false,
		ChannelChatCreated:    false,
		MigrateToChatId:       0,
		MigrateFromChatId:     0,
		PinnedMessage:         nil,
		Invoice:               nil,
		SuccessfulPayment:     nil,
		ConnectedWebsite:      "",
		PassportData:          nil,
		ReplyMarkup:           nil,
	}

	message.From = decodeToUser(findUserByUserList(m.GetFromId_FLAGPEER().GetUserId(), users))
	if m.ToId != nil {
		toId := model.FromPeer(m.ToId)
		switch toId.PeerType {
		case model.PEER_USER:
			if m.Out {
				message.Chat = decodeToChatByUser(findUserByUserList(m.ToId.UserId, users))
			} else {
				message.Chat = decodeToChatByUser(findUserByUserList(m.GetFromId_FLAGPEER().GetUserId(), users))
			}
		case model.PEER_CHAT:
			message.Chat = decodeToChatByChat(findChatByChatList(toId.PeerId, chats))
		case model.PEER_CHANNEL:
			message.Chat = decodeToChatByChat(findChatByChatList(toId.PeerId, chats))
		}
	}

	if len(m.Entities) > 0 {
		message.Entities = make([]*botapi.MessageEntity, 0, len(m.Entities))
		for _, e := range m.Entities {
			switch e.PredicateName {
			case mtproto.Predicate_messageEntityMention:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "mention",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityHashtag:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "hashtag",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityBotCommand:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "bot_command",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityUrl:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "url",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityEmail:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "email",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityBold:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "bold",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityItalic:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "italic",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityCode:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "code",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityPre:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "pre",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityTextUrl:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "text_link",
					Offset: e.Offset,
					Length: e.Length,
					Url:    e.Url,
				})
			case mtproto.Predicate_messageEntityMentionName:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "text_mention",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityPhone:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "phone_number",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityCashtag:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "cashtag",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityUnderline:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "underline",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityStrike:
				message.Entities = append(message.Entities, &botapi.MessageEntity{
					Type:   "strikethrough",
					Offset: e.Offset,
					Length: e.Length,
				})

			default:
			}
		}
	}

	if model.IsPhoto(m) {
		message.Photo = ToPhotoSize(m.Media.Photo_FLAGPHOTO)
	}
	return message
}

func channelMessageToMessage(m *mtproto.Message, users []*mtproto.User, chats []*mtproto.Chat) *botapi.Message {
	return nil
}

func updatesToUpdateList(botId int32, updates *mtproto.Updates) []*botapi.Update {
	var res []*botapi.Update

	model.VisitUpdates(botId, updates, map[string]model.UpdateVisitedFunc{
		mtproto.Predicate_updateNewMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			res = append(res, &botapi.Update{
				UpdateId: update.Pts_INT32,
				Message:  messageToMessage(update.GetMessage_MESSAGE(), users, chats),
			})
		},

		mtproto.Predicate_updateNewChannelMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			res = append(res, &botapi.Update{
				UpdateId: update.Pts_INT32,
				Message:  channelMessageToMessage(update.GetMessage_MESSAGE(), users, chats),
			})
		},

		mtproto.Predicate_updateBotCallbackQuery: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			var botUser *botapi.User
			user := findUserByUserList(update.UserId, users)
			if user != nil {
				botUser = decodeToUser(user)
			} else {
				log.Debugf("not found user: (%d, %v)", update.UserId, users)
			}

			res = append(res, &botapi.Update{
				UpdateId: update.Pts_INT32,
				CallbackQuery: &botapi.CallbackQuery{
					Id:              strconv.Itoa(int(update.QueryId)),
					From:            botUser,
					Message:         messageToMessage(update.GetMessage_MESSAGE(), users, chats),
					InlineMessageId: "",
					ChatInstance:    "",
					Data:            hack.String(update.Data_FLAGBYTES),
					GameShortName:   "",
				},
			})
		},
	})

	return res
}

func ToPhotoSize(photo *mtproto.Photo) (szList []*botapi.PhotoSize) {
	for _, sz := range photo.GetSizes() {
		switch sz.GetPredicateName() {
		case mtproto.Predicate_photoSize:
			szList = append(szList, &botapi.PhotoSize{
				FileId:   fmt.Sprintf("%d@%d.%d", photo.Id, sz.Location.GetVolumeId(), sz.Location.GetLocalId()),
				Width:    sz.W,
				Height:   sz.H,
				FileSize: sz.Size2,
			})
		}
	}

	return
}
