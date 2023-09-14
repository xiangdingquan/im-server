package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/service/media/internal/service"
	"open.chat/app/service/media/mediapb"
	"open.chat/pkg/grpc_util/server"
)

// New new a grpc server.
func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		mediapb.RegisterRPCNbfsServer(s, svc)
	})
}
