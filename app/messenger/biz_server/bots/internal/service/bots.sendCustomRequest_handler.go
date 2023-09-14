package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) BotsSendCustomRequest(ctx context.Context, request *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("bots.sendCustomRequest#aa2769ed - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	reply, err := s.Service.Handle(ctx, md, request.GetCustomMethod(), request.GetParams())

	log.Debugf("bots.sendCustomRequest#aa2769ed - reply: {%s}", logger.JsonDebugData(reply))
	return reply, err
}
