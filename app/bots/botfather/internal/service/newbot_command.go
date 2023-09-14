package service

import (
	"context"
	"strings"
	"time"

	"open.chat/app/pkg/env2"

	"fmt"

	"open.chat/app/bots/botfather/internal/model"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
)

func init() {
	cmdHandlers["newbot"] = func(c botCallback) commandInterface {
		return &newBotCommand{
			botCallback: c,
		}
	}
}

type newBotCommand struct {
	botCallback
}

func (m *newBotCommand) onDoMainCmd(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, params []string) (botMessage *mtproto.Message, r int) {
	states.MainCmd = "newbot"
	states.NextSubCmd = "name"
	botMessage = makeBotMessage(model.BotFatherID, fromUserId, "Alright, a new bot. How are we going to call it? Please choose a name for your bot.")
	r = OpSave
	return
}

func (m *newBotCommand) onDoNextCall(ctx context.Context, fromUserId int32, states *model.BotFatherCommandStates, msg *mtproto.Message) (botMessage *mtproto.Message, r int) {
	switch states.NextSubCmd {
	case "name":
		states.NextSubCmd = "username"
		states.CacheSubCmdResults["name"] = msg.Message
		botMessage = makeBotMessage(model.BotFatherID, fromUserId, "Good. Now let's choose a username for your bot. It must end in `bot`. Like this, for example: TetrisBot or tetris_bot.")
		r = OpSave
	case "username":
		if !strings.HasSuffix(strings.ToLower(msg.Message), "bot") {
			botMessage = makeBotMessage(model.BotFatherID, fromUserId, "Sorry, the username must end in 'bot'. E.g. 'Tetris_bot' or 'Tetrisbot'")
		} else {
			checked, err := m.getUsername().CheckUsername(ctx, msg.Message)
			if err != nil || checked != model2.UsernameNotExisted {
				botMessage = makeBotMessage(model.BotFatherID, fromUserId, "Sorry, this username is already taken. Please try something different.")
				return
			}

			var (
				newBotId int32
				token    = random2.RandomAlphanumeric(35)
			)
			newBotId, err = m.getDao().CreateNewBot(ctx, fromUserId, states.CacheSubCmdResults["name"], msg.Message, token)
			if err != nil {
				log.Errorf("createNewBot error: %v", err)
				return
			}

			m.getUsername().UpdateUsername(ctx, model2.PEER_USER, newBotId, msg.Message)
			botMessage = makeNewBotDoneMessage(model.BotFatherID, fromUserId, msg.Message, newBotId, token)
			r = OpDelete
		}
	}
	return
}
func makeNewBotDoneMessage(botId, toId int32, newBotUserName string, newBotId int32, newBotToken string) *mtproto.Message {
	var messageBuilderHelper MessageBuildHelper = []MessageBuildItem{
		{
			Text:       "Done! Congratulations on your new bot. You will find it at ",
			Param:      env2.T_ME + "/" + newBotUserName,
			EntityType: mtproto.Predicate_messageEntityUrl,
		},
		{
			Text:       ". You can now add a description, about section and profile picture for your bot, see ",
			Param:      "/help",
			EntityType: mtproto.Predicate_messageEntityBotCommand,
		},
		{
			Text:       " for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.\n\nUse this token to access the HTTP API:\n",
			Param:      fmt.Sprintf("%d:%s", newBotId, newBotToken),
			EntityType: mtproto.Predicate_messageEntityCode,
		},
		{
			Text:       "\nKeep your token ",
			Param:      "secure",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       " and ",
			Param:      "store it safely",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text:       ", it can be used by anyone to control your bot.\nFor a description of the Bot API, see this page: ",
			Param:      fmt.Sprintf("https://%s/bots/api", env2.MY_WEB_SITE),
			EntityType: mtproto.Predicate_messageEntityUrl,
		},
	}

	return messageBuilderHelper.MakeMessage(mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		FromId_FLAGPEER: model2.MakePeerUser(botId),
		ToId:            model2.MakePeerUser(toId),
		Date:            int32(time.Now().Unix()),
		Message:         "",
		ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
		Entities:        nil,
	}).To_Message())
}
