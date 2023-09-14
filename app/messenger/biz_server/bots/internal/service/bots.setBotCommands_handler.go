package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) BotsSetBotCommands(ctx context.Context, request *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("bots.setBotCommands - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if len(request.Commands) == 0 {
		log.Errorf("bots.setBotCommands - commands empty")
		return mtproto.BoolTrue, nil
	}

	err := s.UserFacade.SetBotCommands(ctx, md.UserId, request.Commands)
	if err != nil {
		log.Errorf("bots.setBotCommands - error: %v", err)
		return nil, err
	}

	log.Debugf("bots.setBotCommands - reply: {true}")
	return mtproto.BoolTrue, nil
}
