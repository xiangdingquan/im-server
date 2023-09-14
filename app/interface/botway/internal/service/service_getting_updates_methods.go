package service

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"open.chat/app/interface/botway/botapi"
	status_client "open.chat/app/service/status/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	chatPrivate    = "private"
	chatGroup      = "group"
	chatSuperGroup = "supergroup"
	chatChannel    = "channel"
)

func (s *Service) toBotMessage(ctx context.Context, m *mtproto.Message) (botMsg *botapi.Message, err error) {
	var (
		tmpUser *botapi.User
	)

	log.Debugf("toBotMessage - %s", m.DebugString())
	botMsg = new(botapi.Message)
	botMsg.MessageId = m.Id
	botMsg.Text = m.Message
	if m.GetFromId_FLAGPEER() != nil {
		tmpUser, err = s.dao.GetUser(ctx, m.GetFromId_FLAGPEER().GetUserId())
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		botMsg.From = tmpUser
	}
	if len(m.Entities) > 0 {
		botMsg.Entities = make([]*botapi.MessageEntity, 0, len(m.Entities))
		for _, e := range m.Entities {
			switch e.PredicateName {
			case mtproto.Predicate_messageEntityMention:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "mention",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityHashtag:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "hashtag",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityBotCommand:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "bot_command",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityUrl:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "url",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityEmail:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "email",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityBold:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "bold",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityItalic:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "italic",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityCode:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "code",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityPre:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "pre",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityTextUrl:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "text_link",
					Offset: e.Offset,
					Length: e.Length,
					Url:    e.Url,
				})
			case mtproto.Predicate_messageEntityMentionName:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "text_mention",
					Offset: e.Offset,
					Length: e.Length,
					User:   model.GetFirstValue(s.dao.GetUser(ctx, e.UserId_INT32)).(*botapi.User),
				})
			case mtproto.Predicate_messageEntityPhone:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "phone_number",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityCashtag:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "cashtag",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityUnderline:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "underline",
					Offset: e.Offset,
					Length: e.Length,
				})
			case mtproto.Predicate_messageEntityStrike:
				botMsg.Entities = append(botMsg.Entities, &botapi.MessageEntity{
					Type:   "strikethrough",
					Offset: e.Offset,
					Length: e.Length,
				})

			default:
			}
		}
	}

	if m.ToId != nil {
		switch m.ToId.PredicateName {
		case mtproto.Predicate_peerUser:
			tmpUser, err = s.dao.GetUser(ctx, m.ToId.UserId)
			if err != nil {
				log.Errorf("%v", err)
				return
			}
			botMsg.Chat = &botapi.Chat{
				Id:        botapi.ChatIdPrivate.ToChatId(m.ToId.UserId),
				Type:      chatPrivate,
				Username:  tmpUser.Username,
				FirstName: tmpUser.FirstName,
				LastName:  tmpUser.LastName,
			}
		case mtproto.Predicate_peerChat:
		case mtproto.Predicate_peerChannel:
		}
	}
	return
}

func (s *Service) toBotCallbackQuery(ctx context.Context, m *mtproto.TLUpdateBotCallbackQuery) (cbQuery *botapi.CallbackQuery, err error) {
	cbQuery = new(botapi.CallbackQuery)
	cbQuery.Id = strconv.Itoa(int(m.GetQueryId()))
	cbQuery.From = nil
	cbQuery.Message = nil
	cbQuery.InlineMessageId = ""
	cbQuery.ChatInstance = ""
	cbQuery.Data = hack.String(m.GetData_FLAGBYTES())
	cbQuery.GameShortName = ""
	return cbQuery, nil
}

func (s *Service) getBotUpdates(ctx context.Context, botId int32, req *botapi.GetUpdates2) ([]*botapi.Update, error) {
	doList, err := s.dao.BotUpdatesDAO.SelectByGtUpdateId(ctx, botId, int32(req.Offset), int32(req.Limit))
	if err != nil {
		log.Errorf("getSession error: %v", err)
		return nil, err
	}

	var res []*botapi.Update
	for i := 0; i < len(doList); i++ {
		update := &mtproto.Update{}
		err := json.Unmarshal([]byte(doList[i].UpdateData), update)
		if err != nil {
			log.Errorf("unmarshal pts's update(%d)error: %v", doList[i].Id, err)
			continue
		}
		if update.Message_MESSAGE != nil {
			update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
		}
		switch update.PredicateName {
		case mtproto.Predicate_updateNewMessage:
			if !req.CheckAllowUpdate("message") {
				continue
			}
			newMessage := update.To_UpdateNewMessage()
			botMsg, _ := s.toBotMessage(ctx, newMessage.GetMessage_MESSAGE())
			res = append(res, &botapi.Update{
				UpdateId: doList[i].UpdateId,
				Message:  botMsg,
			})
		case mtproto.Predicate_updateEditMessage:
			if !req.CheckAllowUpdate("edited_message") {
				continue
			}
			editMessage := update.To_UpdateEditMessage()
			botMsg, _ := s.toBotMessage(ctx, editMessage.GetMessage_MESSAGE())
			res = append(res, &botapi.Update{
				UpdateId: doList[i].UpdateId,
				Message:  botMsg,
			})
		case mtproto.Predicate_updateBotCallbackQuery:
			if !req.CheckAllowUpdate("callback_query") {
				continue
			}
			cbQuery, _ := s.toBotCallbackQuery(ctx, update.To_UpdateBotCallbackQuery())
			res = append(res, &botapi.Update{
				UpdateId:      doList[i].UpdateId,
				CallbackQuery: cbQuery,
			})
		default:
			continue
		}
	}

	return res, nil
}

func (s *Service) GetUpdates(ctx context.Context, token string, req *botapi.GetUpdates2) ([]*botapi.Update, error) {
	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	botSession := s.sessions.Get(auth.AuthKeyId())
	if botSession == nil {
		botSession = newCacheBotSession(auth.UserId(), auth.AuthKeyId(), auth.Layer())
		s.sessions.Put(botSession)
	}

	if botSession.TrySetLastTime(60) {
		status_client.AddOnline(ctx, botSession.userId, botSession.authKeyId, env.Hostname)
	}

	botUpdates, err := s.getBotUpdates(ctx, botSession.userId, req)
	if err != nil {
		return nil, err
	} else if botUpdates == nil {
		respChan := make(chan mtproto.TLObject, 2)
		s.sessions.reqCache.cache(botSession.authKeyId, respChan)
		timer := time.NewTimer(time.Second * time.Duration(req.Timeout))
		select {
		case meUpdates := <-respChan:
			if meUpdates != nil {
				if updates, ok := meUpdates.(*mtproto.Updates); ok {
					log.Debugf("pushUpdates: %s", updates.DebugString())
					botUpdates, _ = DecodeToBotApi(updates).([]*botapi.Update)
				}
			}
		case <-timer.C:
			s.sessions.reqCache.dispose(botSession.authKeyId)
		}

		if botUpdates == nil {
			botUpdates = make([]*botapi.Update, 0)
		}
		log.Debugf("botUpdates: %v", botUpdates)
	} else {
		log.Debugf("botUpdates: %v", botUpdates)
	}

	return botUpdates, nil
}

func (s *Service) SetWebhook(ctx context.Context, token string, req *botapi.SetWebhook2) (bool, error) {
	log.Warnf("setWebhook - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) DeleteWebhook(ctx context.Context, token string, req *botapi.DeleteWebhook2) (bool, error) {
	log.Warnf("deleteWebhook - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) GetWebhookInfo(ctx context.Context, token string, req *botapi.GetWebhookInfo2) (*botapi.WebhookInfo, error) {
	log.Warnf("getWebhookInfo - method not impl")

	r := &botapi.WebhookInfo{
		Url:                  "",
		HasCustomCertificate: false,
		PendingUpdateCount:   0,
	}

	log.Debugf("getWebhookInfo: %#v", r)
	return r, nil
}
