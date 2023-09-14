package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/interface/botway/internal/service"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		sessionpb.RegisterRPCPushServer(s, svc)
	})
}
