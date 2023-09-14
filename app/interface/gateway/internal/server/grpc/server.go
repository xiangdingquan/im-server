package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/interface/gateway/egatepb"
	"open.chat/app/interface/gateway/internal/server"
	server2 "open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *server.Server) *server2.RPCServer {
	return server2.NewRpcServer(appId, func(s *grpc.Server) {
		egatepb.RegisterEGateServer(s, svc)
	})
}
