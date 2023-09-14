package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/messenger/webpage/internal/service"
	"open.chat/app/messenger/webpage/webpagepb"
	"open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		webpagepb.RegisterRPCWebPageServer(s, svc)
	})
}
