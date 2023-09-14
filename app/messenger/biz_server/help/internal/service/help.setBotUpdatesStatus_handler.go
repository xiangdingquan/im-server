package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpSetBotUpdatesStatus(ctx context.Context, request *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.setBotUpdatesStatus - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 1. for bots only
	// 2. handle: pending_updates_count:int message:string

	log.Debugf("help.setBotUpdatesStatus - reply: {true}")
	return mtproto.ToBool(true), nil
}
