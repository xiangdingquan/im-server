package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/app/service/auth_session/internal/service"
	"open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		authsessionpb.RegisterRPCSessionServer(s, svc)
	})
}
