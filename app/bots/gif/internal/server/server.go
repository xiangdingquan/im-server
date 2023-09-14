package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/bots/gif/internal/server/grpc"
	"open.chat/app/bots/gif/internal/server/http"
	"open.chat/app/bots/gif/internal/service"
	"open.chat/app/pkg/env2"
	"open.chat/pkg/grpc_util/server"
)

type Server struct {
	httpSrv *bm.Engine
	grpcSrv *server.RPCServer
	svc     *service.Service
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	s.svc = service.New()
	s.grpcSrv = grpc.New(env2.BotsGifId, s.svc)
	s.httpSrv = http.New(s.svc)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.svc.Close()
}
