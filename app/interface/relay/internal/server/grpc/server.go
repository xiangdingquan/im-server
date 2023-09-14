package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/interface/relay/internal/server"
	"open.chat/app/interface/relay/relaypb"
	server2 "open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *server.Server) *server2.RPCServer {
	return server2.NewRpcServer(appId, func(s *grpc.Server) {
		relaypb.RegisterRPCRelayServer(s, svc)
	})
}
