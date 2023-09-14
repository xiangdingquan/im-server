package main

import (
	"open.chat/app/interface/gateway/internal/server"
	"open.chat/app/interface/gateway/internal/server/grpc"
	"open.chat/app/pkg/env2"
	"open.chat/pkg/commands"
	rpc_server "open.chat/pkg/grpc_util/server"
)

type Server struct {
	grpcSrv *rpc_server.RPCServer
	server  *server.Server
}

func (s *Server) Initialize() error {
	s.server = server.New()
	s.grpcSrv = grpc.New(env2.InterfaceGatewayId, s.server)
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
