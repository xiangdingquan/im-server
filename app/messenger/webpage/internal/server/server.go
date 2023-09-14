package server

import (
	"open.chat/app/messenger/webpage/internal/server/grpc"
	"open.chat/app/messenger/webpage/internal/service"
	"open.chat/app/pkg/env2"
	"open.chat/pkg/grpc_util/server"
)

type Server struct {
	grpcSrv *server.RPCServer
	svc     *service.Service
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	s.svc = service.New()
	s.grpcSrv = grpc.New(env2.MessengerWebPageId, s.svc)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.svc.Close()
}
