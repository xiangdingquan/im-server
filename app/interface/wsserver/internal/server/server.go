package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/wsserver/internal/server/http"
	"open.chat/app/interface/wsserver/internal/service"
)

type Server struct {
	httpSrv *bm.Engine
	svc     *service.Service
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	s.svc = service.New()
	s.httpSrv = http.New(s.svc)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.svc.Close()
}
