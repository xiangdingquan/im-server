package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/messenger/biz_server/internal/server/grpc"
	"open.chat/app/messenger/biz_server/internal/server/http"
	"open.chat/app/messenger/biz_server/internal/service"
	"open.chat/pkg/grpc_util/server"
)

// ///////////////////////////////////////////////////////////////////////////
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
	s.grpcSrv = grpc.New()
	s.httpSrv = http.New(s.svc)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
