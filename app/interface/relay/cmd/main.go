package main

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/interface/relay/internal/server"
	"open.chat/app/interface/relay/internal/server/grpc"
	"open.chat/app/interface/relay/internal/server/http"
	"open.chat/app/pkg/env2"
	"open.chat/pkg/commands"
	rpc_server "open.chat/pkg/grpc_util/server"
)

type Server struct {
	httpSrv *bm.Engine
	grpcSrv *rpc_server.RPCServer
	server  *server.Server
}

func (s *Server) Initialize() error {
	s.server = server.New()
	s.grpcSrv = grpc.New(env2.InterfaceRelayId, s.server)
	s.httpSrv = http.New(s.server)
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.server.Close()
}

func main() {
	commands.Run(new(Server))
}
