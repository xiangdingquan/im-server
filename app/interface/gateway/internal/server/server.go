package server

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"

	"open.chat/app/interface/gateway/internal/dao"
	"open.chat/app/interface/gateway/internal/server/http"
	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/ketama"
	"open.chat/pkg/log"
	"open.chat/pkg/net2"
	"open.chat/pkg/time2"
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

type Server struct {
	c              *Config
	server         *net2.TcpServer2
	handshake      *handshake
	sessionServers map[string]*sessionClient
	ketama         *ketama.Ketama
	authSessionMgr *authSessionManager
	dao            *dao.Dao
	timer          *time2.Timer
	httpSrv        *bm.Engine
}

func New() *Server {
	var (
		ac  = &Config{}
		err error
		s   = new(Server)
	)

	if err := paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	log.Info("config: %#v", ac)

	s.c = ac
	s.authSessionMgr = NewAuthSessionManager()

	s.sessionServers = make(map[string]*sessionClient)
	s.ketama = ketama.NewKetama(10, nil)

	s.dao = dao.New(1024*1024, ac.WardenClient)
	s.timer = time2.NewTimer(1024)

	keyFingerprint, err := strconv.ParseUint(ac.KeyFingerprint, 10, 64)
	if err != nil {
		panic(err)
	}
	s.handshake, err = newHandshake(s.c.KeyFile, keyFingerprint,
		func(ctx context.Context, keyInfo *authsessionpb.AuthKeyInfo, salt *mtproto.FutureSalt) error {
			return s.dao.PutAuthKey(ctx, keyInfo, salt)
		})

	if err != nil {
		panic(err)
	}

	s.server, err = net2.NewTcpServer2(s.c.Server, s.c.MaxProc, s)
	if err != nil {
		panic(err)
	}
	s.server.Serve()
	s.httpSrv = http.New()

	s.watchSession()

	return s
}

func (s *Server) Close() {
	s.server.Stop()
}

func (s *Server) Ping(ctx context.Context) (err error) {
	return nil
}

func (s *Server) watchSession() {
	dis, err := etcd.New(&clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	resolver := dis.Build(env2.InterfaceSessionSessionId)

	event := resolver.Watch()
	select {
	case _, ok := <-event:
		if !ok {
			panic("watchComet init failed")
		}
		if ins, ok := resolver.Fetch(context.Background()); ok {
			if err := s.newAddress(ins); err != nil {
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
				if err := s.newAddress(ins); err != nil {
					log.Error("watchComet newAddress(%+v) error(%+v)", ins, err)
					continue
				}
				log.Info("watchComet change newAddress:%+v", ins)
			}
		}
	}()
}

func (s *Server) newAddress(insInfo *naming.InstancesInfo) error {
	var (
		addAddrs    []string
		removeAddrs []string
	)

	ins := insInfo.Instances[env.Zone]
	if len(ins) == 0 {
		return fmt.Errorf("watchSession instance is empty")
	}
	sessions := map[string]*sessionClient{}
	for _, in := range ins {
		if old, ok := s.sessionServers[in.Hostname]; ok {
			sessions[in.Hostname] = old
			continue
		}
		c, err := newSession(in, s.c)
		if err != nil {
			log.Error("watchSession NewSession(%+v) error(%v)", in, err)
			return err
		}
		sessions[in.Hostname] = c
		addAddrs = append(addAddrs, in.Hostname)
		log.Info("watchSession AddSession grpc:%+v", in)
	}
	for key, old := range s.sessionServers {
		if _, ok := sessions[key]; !ok {
			old.cancel()
			removeAddrs = append(removeAddrs, key)
			log.Info("watchSession DelSession:%s", key)
		}
	}
	s.sessionServers = sessions
	if len(addAddrs) > 0 {
		s.ketama.Add(addAddrs...)
	}
	if len(removeAddrs) > 0 {
		s.ketama.Remove(removeAddrs...)
	}

	return nil
}

func (s *Server) GetSessionClient(key string) (c *sessionClient, err error) {
	if addr, ok := s.ketama.Get(key); !ok {
		err = fmt.Errorf("not found ketama addr by key: %s", key)
	} else {
		sessionServers := s.sessionServers
		if c2, ok := sessionServers[addr]; !ok {
			err = fmt.Errorf("not found ketama addr by (key: %s, addr: %s)", key, addr)
		} else {
			c = c2
		}
	}
	return
}
