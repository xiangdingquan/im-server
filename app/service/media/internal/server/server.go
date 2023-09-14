package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/service/media/internal/server/http"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/media/internal/server/grpc"
	"open.chat/app/service/media/internal/service"
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
	s.grpcSrv = grpc.New(env2.ServiceMediaId, s.svc)
	s.httpSrv = http.New(s.svc)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.svc.Close()
}
