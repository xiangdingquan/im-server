package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// stats.getMessageStats#b6e0a3f5 flags:# dark:flags.0?true channel:InputChannel msg_id:int = stats.MessageStats;
func (s *Service) StatsGetMessageStats(ctx context.Context, request *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stats.getMessageStats - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl StatsGetMessageStats logic
	log.Warn("stats.getMessageStats - error: method StatsGetMessageStats not impl")

	return nil, mtproto.ErrMethodNotImpl
}
