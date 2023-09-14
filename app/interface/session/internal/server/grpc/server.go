package grpc

import (
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"

	"open.chat/app/interface/session/internal/service"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/pkg/grpc_util/server"
)

func NewSession(appId string, config *warden.ServerConfig, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer2(appId, config, func(s *grpc.Server) {
		sessionpb.RegisterRPCSessionServer(s, svc)
	})
}

func NewPush(appId string, config *warden.ServerConfig, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer2(appId, config, func(s *grpc.Server) {
		sessionpb.RegisterRPCPushServer(s, svc)
	})
}
