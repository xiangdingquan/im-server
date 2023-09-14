package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/session/internal/server/http"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/interface/session/internal/server/grpc"
	"open.chat/app/interface/session/internal/service"
	"open.chat/app/pkg/env2"
	"open.chat/pkg/grpc_util/server"
)

type Server struct {
	httpSrv        *bm.Engine
	sessionGrpcSrv *server.RPCServer
	pushGrpcSrv    *server.RPCServer
	svc            *service.Service
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	s.svc = service.New()

	var rc struct {
		Session *warden.ServerConfig
		Push    *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	time.Sleep(2)
	s.sessionGrpcSrv = grpc.NewSession(env2.InterfaceSessionSessionId, rc.Session, s.svc)
	s.pushGrpcSrv = grpc.NewPush(env2.InterfaceSessionPushId, rc.Push, s.svc)
	s.httpSrv = http.New(s.svc)

	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.sessionGrpcSrv.Stop()
	s.pushGrpcSrv.Stop()
	s.svc.Close()
}
