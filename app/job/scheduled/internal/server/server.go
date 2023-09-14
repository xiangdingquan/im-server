package server

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/job/scheduled/internal/server/http"
	"open.chat/app/job/scheduled/internal/service"
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
	s.svc.RunLoop()
}

func (s *Server) Destroy() {
	s.svc.Close()
}
