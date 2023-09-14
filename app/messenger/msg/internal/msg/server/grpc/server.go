package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/messenger/msg/internal/msg/service"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		msgpb.RegisterRPCMsgServer(s, svc)
	})
}
