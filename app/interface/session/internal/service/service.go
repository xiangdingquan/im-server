package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"

	"open.chat/app/interface/session/internal/dao"
	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
	"open.chat/pkg/log"
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

type Service struct {
	ac              *Config
	mu              sync.RWMutex
	sessionsManager map[int64]*authSessions
	eGateServers    map[string]*Gateway
	reqCache        *RequestManager
	*dao.Dao
}

func New() *Service {
	var (
		ac  = &Config{}
		err error
		s   = new(Service)
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s.ac = ac
	s.Dao = dao.New(ac.WardenClient)
	s.sessionsManager = make(map[int64]*authSessions)
	s.eGateServers = make(map[string]*Gateway)
	s.reqCache = NewRequestManager()

	s.watchGateway()
	return s
}

func (s *Service) Close() error {
	for _, c := range s.eGateServers {
		if err := c.Close(); err != nil {
			log.Error("c.Close() error(%v)", err)
		}
	}

	s.Dao.Close()
	return nil
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

func (s *Service) newAddress(insMap map[string][]*naming.Instance) error {
	ins := insMap[env.Zone]
	if len(ins) == 0 {
		return fmt.Errorf("watchComet instance is empty")
	}
	eGates := map[string]*Gateway{}
	options := gatewayOptions{
		RoutineSize: s.ac.Routine.Size,
		RoutineChan: s.ac.Routine.Chan,
	}

	for _, in := range ins {
		if old, ok := s.eGateServers[in.Hostname]; ok {
			eGates[in.Hostname] = old
			continue
		}
		c, err := NewGateway(in, s.ac, options)
		if err != nil {
			log.Errorf("watchComet NewComet(%+v) error(%v)", in, err)
			return err
		}
		eGates[in.Hostname] = c
		log.Infof("watchComet AddComet grpc:%+v", in)
	}
	for key, old := range s.eGateServers {
		if _, ok := eGates[key]; !ok {
			_ = old
			log.Infof("watchComet DelComet:%s", key)
		}
	}
	s.eGateServers = eGates
	return nil
}

func (s *Service) watchGateway() {
	dis, err := etcd.New(&clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	resolver := dis.Build(env2.InterfaceGatewayId)

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

func (s *Service) SendDataToGateway(ctx context.Context, gatewayId string, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	if c, ok := s.eGateServers[gatewayId]; ok {
		return c.SendDataToGate(ctx, authKeyId, sessionId, SerializeToBuffer2(salt, sessionId, msg))
	} else {
		log.Errorf("not found k: %s, %v", gatewayId, s.eGateServers)
		return false, fmt.Errorf("not found k: %s", gatewayId)
	}
}

func (s *Service) SendHttpDataToGateway(ctx context.Context, ch chan interface{}, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	select {
	case ch <- SerializeToBuffer2(salt, sessionId, msg):
		close(ch)
		return true, nil
	default:
		log.Errorf("Default fail !!!!! ch closed")
		return false, fmt.Errorf("ch closed")
	}
}

func (s *Service) DeleteByAuthKeyId(authKeyId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok := s.sessionsManager[authKeyId]; ok {
		sessList.Stop()
		delete(s.sessionsManager, authKeyId)
	}
}
