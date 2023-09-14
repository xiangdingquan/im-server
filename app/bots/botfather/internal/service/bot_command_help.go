package service

import (
	"context"
	"time"

	message_facade "open.chat/app/service/biz_service/message/facade"

	model2 "open.chat/model"

	"open.chat/app/bots/botfather/internal/dao"
	"open.chat/app/bots/botfather/internal/model"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
	"open.chat/mtproto"
)

const (
	botFatherHelpMessage = `I can help you create and manage Nebuchat bots. If you're new to the Bot API, please see the manual.

You can control me by sending these commands:

/newbot - create a new bot
/mybots - edit your bots [beta]

Edit Bots
/setname - change a bot's name
/setdescription - change bot description
/setabouttext - change bot about info
/setuserpic - change bot profile photo
/setcommands - change the list of commands
/deletebot - delete a bot

Bot Settings
/token - generate authorization token
/revoke - revoke bot access token
/setinline - toggle inline mode
/setinlinegeo - toggle inline location requests
/setinlinefeedback - change inline feedback settings
/setjoingroups - can your bot be added to groups?
/setprivacy - toggle privacy mode in groups

Games
/mygames - edit your games [beta]
/newgame - create a new game
/listgames - get a list of your games
/editgame - edit a game
/deletegame - delete an existing game`

	botFatherUnrecognizedMessage = "Unrecognized command. Say what?"
	noBotsMessage                = "You don't have any bots yet. Use the /newbot command to create a new bot first."
)

const (
	OpNone   = 0
	OpDelete = 1
	OpSave   = 2
)

type commandInterface interface {
	onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (*mtproto.Message, int)
	onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (*mtproto.Message, int)
}

type botCallback interface {
	getUsername() username_facade.UsernameFacade
	getDao() *dao.Dao
	getUser() user_client.UserFacade
	getMessage() message_facade.MessageFacade
}

var cmdHandlers = make(map[string]func(c botCallback) commandInterface)

func NewCommandHandler(cmdName string, c botCallback) commandInterface {
	if v, ok := cmdHandlers[cmdName]; !ok {
		return nil
	} else {
		return v(c)
	}
}

func makeBotMessage(botId, toId int32, msg string) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         msg,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
	}).To_Message()
}

func makeNoBotsMessage(botId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         noBotsMessage,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 37,
				Length: 7,
			}).To_MessageEntity(),
		},
	}).To_Message()
}

func makeSetMessage(botId, toId int32, message string, buttons []*mtproto.KeyboardButton) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         message,
		ReplyMarkup: mtproto.MakeTLReplyKeyboardMarkup(&mtproto.ReplyMarkup{
			Resize: true,
			Rows: []*mtproto.KeyboardButtonRow{
				mtproto.MakeTLKeyboardButtonRow(&mtproto.KeyboardButtonRow{
					Buttons: buttons,
				}).To_KeyboardButtonRow(),
			},
		}).To_ReplyMarkup(),
	}).To_Message()
}

const (
	noGamesMessage = "You currently have no games. Use /newgame command to create a first game."
)

func makeNoGamesMessage(botId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         noGamesMessage,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 33,
				Length: 8,
			}).To_MessageEntity(),
		},
	}).To_Message()
}

const (
	noInlineMessage = "You don't have any inline bots setup yet. Use /setinline to enable inline mode for a bot first."
)

func makeNoInlineMessage(botId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         noInlineMessage,
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: 46,
				Length: 10,
			}).To_MessageEntity(),
		},
	}).To_Message()
}
