package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/bots/botpb"
	"open.chat/app/bots/gif/internal/dao"
	"open.chat/app/infra/databus/pkg/queue/databus"
	msg_facade "open.chat/app/messenger/msg/facade"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/pkg/databus_util"
	_ "open.chat/app/service/biz_service/message/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	_ "open.chat/app/service/biz_service/username/facade"
	username_facade "open.chat/app/service/biz_service/username/facade"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type Config struct {
	Log       *log.Config
	HTTP      *bm.ServerConfig
	RPC       *warden.ClientConfig
	Databus   *databus.Config
	Discovery *discovery.Config
}

var (
	_service *Service
)

type Service struct {
	conf *Config

	*dao.Dao

	databus        *databus.Databus
	databusHandler *databus_util.DatabusHandler

	msg_facade.MsgFacade
	username_facade.UsernameFacade
	user_client.UserFacade
	message_facade.MessageFacade
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		err error
		ac  = &Config{}
		s   = new(Service)
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s.conf = ac
	s.Dao = dao.New()

	s.databus = databus.New(ac.Databus)
	s.databusHandler = databus_util.NewDatabusHandler()

	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	media_client.New()

	s.databusHandler.GoWatch(s.databus, func(msg *databus.Message) error {
		log.Debugf("recv {key: %s, value: %s", msg.Key, string(msg.Value))

		defer func() {
			if r := recover(); r != nil {
				log.Errorf("handle panic: %s", debug.Stack())
			}
		}()

		if msg.Topic != s.conf.Databus.Topic {
			log.Errorf("unknown message:%v", msg)
			return nil
		}

		s.onBotsData(context.Background(), msg.Key, msg.Value)
		return nil
	})

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.Dao.Ping(ctx)
}

// Close close the resources.
func (s *Service) Close() (err error) {
	s.Dao.Close()
	return
}

func (s *Service) getUsername() username_facade.UsernameFacade {
	return s.UsernameFacade
}

func (s *Service) getDao() *dao.Dao {
	return s.Dao
}

func (s *Service) getUser() user_client.UserFacade {
	return s.UserFacade
}

func (s *Service) onBotsData(ctx context.Context, key string, value []byte) error {
	log.Debugf("recv {key: %s, value: %s", key, string(value))
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("handle panic: %s", debug.Stack())
		}
	}()

	switch key {
	case proto.MessageName((*botpb.BotUpdates)(nil)):
		r := new(botpb.BotUpdates)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onBotUpdates(ctx, r)
	default:
		err := fmt.Errorf("invalid key: %s", key)
		log.Error(err.Error())
		return err
	}
}

func (s *Service) onBotUpdates(ctx context.Context, r *botpb.BotUpdates) error {
	model.VisitUpdates(r.BotId, r.Updates, map[string]model.UpdateVisitedFunc{
		mtproto.Predicate_updateNewMessage: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			var (
				m          = update.Message_MESSAGE
				fromUserId = m.GetFromId_FLAGPEER().GetUserId()
			)

			gifHelpMessage := mtproto.MakeTLMessage(&mtproto.Message{
				Out:             true,
				FromId_FLAGPEER: model.MakePeerUser(model.BotGifId),
				ToId:            model.MakePeerUser(fromUserId),
				Date:            int32(time.Now().Unix()),
				Message:         "This bot can help you find and share GIFs. It works automatically, no need to add it anywhere. Simply open any of your chats and type @gif something in the message field. Then tap on a result to send.\n\nFor example, try typing @gif happy dog here.",
				ReplyMarkup:     mtproto.MakeTLReplyKeyboardHide(&mtproto.ReplyMarkup{Selective: false}).To_ReplyMarkup(),
				Entities: []*mtproto.MessageEntity{
					mtproto.MakeTLMessageEntityCode(&mtproto.MessageEntity{
						Offset: 134,
						Length: 14,
					}).To_MessageEntity(),
					mtproto.MakeTLMessageEntityCode(&mtproto.MessageEntity{
						Offset: 226,
						Length: 14,
					}).To_MessageEntity(),
				},
			}).To_Message()
			if _, err := s.MsgFacade.SendMessage(ctx,
				model.BotGifId,
				0,
				model.MakeUserPeerUtil(fromUserId),
				&msgpb.OutboxMessage{
					NoWebpage:    true,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      gifHelpMessage,
					ScheduleDate: nil,
				}); err != nil {
			}
		},
	})
	return nil
}
