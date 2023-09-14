package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime/debug"

	"open.chat/app/messenger/msg/msgpb"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/bots/botfather/internal/dao"
	"open.chat/app/bots/botfather/internal/model"
	"open.chat/app/bots/botpb"
	"open.chat/app/infra/databus/pkg/queue/databus"
	msg_facade "open.chat/app/messenger/msg/facade"
	_ "open.chat/app/service/biz_service/message/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	_ "open.chat/app/service/biz_service/username/facade"
	username_facade "open.chat/app/service/biz_service/username/facade"
	model2 "open.chat/model"
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
	conf     *Config
	consumer *databus.Databus

	*dao.Dao
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
	s.consumer = databus.New(ac.Databus)
	s.Dao = dao.New(ac.RPC)

	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	go s.consume()

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

// botCallback
func (s *Service) getUsername() username_facade.UsernameFacade {
	return s.UsernameFacade
}

func (s *Service) getDao() *dao.Dao {
	return s.Dao
}

func (s *Service) getUser() user_client.UserFacade {
	return s.UserFacade
}

func (s *Service) getMessage() message_facade.MessageFacade {
	return s.MessageFacade
}

func (s *Service) consume() {
	msgs := s.consumer.Messages()
	for {
		msg, ok := <-msgs
		if !ok {
			log.Warnf("[job] consumer has been closed")
			return
		}
		if msg.Topic != s.conf.Databus.Topic {
			log.Errorf("unknown message:%v", msg)
			continue
		}

		s.onBotsData(context.Background(), msg.Key, msg.Value)

		msg.Commit()
	}
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
	switch r.Updates.PredicateName {
	case mtproto.Predicate_updatesTooLong:
	case mtproto.Predicate_updateShortMessage:
	case mtproto.Predicate_updateShortChatMessage:
	case mtproto.Predicate_updateShort:
	case mtproto.Predicate_updatesCombined:
	case mtproto.Predicate_updates:
		for _, update := range r.Updates.Updates {
			switch update.PredicateName {
			case mtproto.Predicate_updateNewMessage:
				s.onBotFatherCommand(ctx, update.Message_MESSAGE)
			default:
			}
		}
	case mtproto.Predicate_updateShortSentMessage:
	default:
	}
	return nil
}

func (s *Service) onBotFatherCommand(ctx context.Context, m *mtproto.Message) error {
	fromUserId := m.GetFromId_FLAGPEER().GetUserId()
	states, err := s.Dao.GetBotFatherCommandStates(ctx, fromUserId)
	if err != nil {
		return err
	}

	var (
		botMessage *mtproto.Message
		r2         int
	)

	command := model2.GetBotCommandByMessage(m)
	if command == nil {
		if states.MainCmd == "" {
			botMessage = makeBotHelpMessage(model.BotFatherID, fromUserId)
		} else {
			cmdH := NewCommandHandler(states.MainCmd, s)
			if cmdH != nil {
				botMessage, r2 = cmdH.onDoNextCall(ctx, fromUserId, states, m)
			}
		}
	} else {
		cmdH := NewCommandHandler(command.CommandName, s)
		if cmdH != nil {
			botMessage, r2 = cmdH.onDoMainCmd(ctx, fromUserId, states, command.Params)
		}
	}

	switch r2 & 0xffff {
	case OpDelete:
		s.Dao.DeleteBotFatherCommandStates(ctx, fromUserId)
	case OpSave:
		s.Dao.PutBotFatherCommandStates(ctx, fromUserId, states)
	case OpNone:
	default:
	}

	if botMessage == nil {
		botMessage = makeBotHelpMessage(model.BotFatherID, fromUserId)
	}

	if r2>>16 == sendMessage {
		_, err = s.MsgFacade.SendMessage(ctx,
			model.BotFatherID,
			0,
			model2.MakeUserPeerUtil(fromUserId),
			&msgpb.OutboxMessage{
				NoWebpage:    true,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      botMessage,
				ScheduleDate: nil,
			})
	} else {
		s.MsgFacade.EditMessage(ctx,
			model.BotFatherID,
			0,
			model2.MakeUserPeerUtil(fromUserId),
			&msgpb.OutboxMessage{
				NoWebpage:    false,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      botMessage,
				ScheduleDate: nil,
			})
	}
	return err
}
