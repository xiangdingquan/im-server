package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/bots/botpb"
	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/interface/session/sessionpb"
	push_client "open.chat/app/messenger/push/client"
	"open.chat/app/messenger/sync/internal/dao"
	"open.chat/app/messenger/sync/syncpb"
	"open.chat/app/pkg/databus_util"
	"open.chat/app/pkg/env2"
	_ "open.chat/app/service/biz_service/channel/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	_ "open.chat/app/service/biz_service/chat/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

var (
	endpoints string
)

func init() {
	endpoints = os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		panic(fmt.Errorf("invalid etcd config endpoints:%+v", endpoints))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Service struct {
	*dao.Dao
	*push_client.PushClient
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	user_client.UserFacade
	conf *Config

	databus        *databus.Databus
	databusHandler *databus_util.DatabusHandler

	sessionServers   map[string]*Session
	botsClient       *botpb.BotsClient
	bingClient       *botpb.BotsClient
	picClient        *botpb.BotsClient
	gifClient        *botpb.BotsClient
	foursquareClient *botpb.BotsClient
}

func New() *Service {
	var (
		ac  = &Config{}
		err error
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s := &Service{
		conf:           ac,
		databus:        databus.New(ac.Databus),
		databusHandler: databus_util.NewDatabusHandler(),
		sessionServers: make(map[string]*Session),
		Dao:            dao.New(),
		PushClient:     push_client.New(),
		botsClient:     botpb.New("bots"),
		gifClient:      botpb.New("gif"),
	}

	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)

	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	s.watchSession()

	s.databusHandler.GoWatch(s.databus, func(msg *databus.Message) error {
		log.Debugf("recv {key: %s, value: %s", msg.Key, string(msg.Value))

		// var err error
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("handle panic: %s", debug.Stack())
			}
		}()

		if msg.Topic != s.conf.Databus.Topic {
			log.Error("unknown message:%v", msg)
			return nil
		}

		return s.onSyncData(context.Background(), msg.Key, msg.Value)
	})

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

// Close close the resources.
func (s *Service) Close() error {
	s.databusHandler.Close()
	for _, c := range s.sessionServers {
		if err := c.Close(); err != nil {
			log.Error("c.Close() error(%v)", err)
		}
	}
	return nil
}

func (s *Service) watchSession() {
	dis, err := etcd.New(&clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	resolver := dis.Build(env2.InterfaceSessionPushId)

	event := resolver.Watch()
	select {
	case _, ok := <-event:
		if !ok {
			panic("watchComet init failed")
		}
		if ins, ok := resolver.Fetch(context.Background()); ok {
			if err := s.newAddress(ins.Instances); err != nil {
				log.Error("watchComet newAddress(%+v) error(%+v)", ins, err)
			} else {
				log.Info("watchComet init newAddress:%+v", ins)
			}
		}
	case <-time.After(10 * time.Second):
		log.Error("watchComet init instances timeout")
	}
	go func() {
		for {
			if _, ok := <-event; !ok {
				log.Info("watchComet exit")
				return
			}
			ins, ok := resolver.Fetch(context.Background())
			if ok {
				if err := s.newAddress(ins.Instances); err != nil {
					log.Error("watchComet newAddress(%+v) error(%+v)", ins, err)
					continue
				}
				log.Info("watchComet change newAddress:%+v", ins)
			}
		}
	}()
}

func (s *Service) newAddress(insMap map[string][]*naming.Instance) error {
	ins := insMap[env.Zone]
	if len(ins) == 0 {
		return fmt.Errorf("watchComet instance is empty")
	}
	sessions := map[string]*Session{}
	options := SessionOptions{
		RoutineSize: s.conf.Routine.Size,
		RoutineChan: s.conf.Routine.Chan,
	}
	for _, data := range ins {
		if old, ok := s.sessionServers[data.Hostname]; ok {
			sessions[data.Hostname] = old
			continue
		}
		c, err := NewSession(data, s.conf, options)
		if err != nil {
			log.Error("watchComet NewComet(%+v) error(%v)", data, err)
			return err
		}
		sessions[data.Hostname] = c
		log.Info("watchComet AddComet grpc:%+v", data)
	}
	for key, old := range s.sessionServers {
		if _, ok := sessions[key]; !ok {
			old.cancel()
			log.Info("watchComet DelComet:%s", key)
		}
	}
	s.sessionServers = sessions
	return nil
}

func (s *Service) onSyncData(ctx context.Context, key string, value []byte) error {
	switch key {
	case proto.MessageName((*syncpb.TLSyncSyncUpdates)(nil)):
		r := new(syncpb.TLSyncSyncUpdates)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onSyncUpdates(ctx, r)
	case proto.MessageName((*syncpb.TLSyncPushUpdates)(nil)):
		r := new(syncpb.TLSyncPushUpdates)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onPushUpdates(ctx, r)
	case proto.MessageName((*syncpb.TLSyncBroadcastUpdates)(nil)):
		r := new(syncpb.TLSyncBroadcastUpdates)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onBroadcastUpdates(ctx, r)
	case proto.MessageName((*syncpb.TLSyncPushRpcResult)(nil)):
		r := new(syncpb.TLSyncPushRpcResult)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.PushRpcResultToSession(ctx, r.ServerId, &sessionpb.PushRpcResultData{
			AuthKeyId:      r.AuthKeyId,
			SessionId:      r.SessionId,
			ClientReqMsgId: r.ReqMsgId,
			RpcResultData:  r.Result,
		})
	default:
		err := fmt.Errorf("invalid key: %s", key)
		log.Error(err.Error())
		return err
	}
}

func (s *Service) PushUpdatesToSession(ctx context.Context, serverId string, msg *sessionpb.PushUpdatesData) (err error) {
	if c, ok := s.sessionServers[serverId]; ok {
		log.Info("push updates to serverId: (%s, %s)", serverId, logger.JsonDebugData(msg))
		return c.PushUpdates(ctx, msg)
	} else {
		log.Errorf("not found k: %s, %v", serverId, s.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
}

func (s *Service) PushSessionUpdatesToSession(ctx context.Context, serverId string, msg *sessionpb.PushSessionUpdatesData) (err error) {
	if c, ok := s.sessionServers[serverId]; ok {
		log.Info("push session updates to serverId: (%s, %s)", serverId, logger.JsonDebugData(msg))
		return c.PushSessionUpdates(ctx, msg)
	} else {
		log.Errorf("not found k: %s, %v", serverId, s.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
}

func (s *Service) PushRpcResultToSession(ctx context.Context, serverId string, msg *sessionpb.PushRpcResultData) (err error) {
	if c, ok := s.sessionServers[serverId]; ok {
		log.Debugf("push rpc result to serverId: (%s, %s)", serverId, logger.JsonDebugData(msg))
		return c.PushRpcResult(ctx, msg)
	} else {
		log.Errorf("not found k: %s, %v", serverId, s.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
}
