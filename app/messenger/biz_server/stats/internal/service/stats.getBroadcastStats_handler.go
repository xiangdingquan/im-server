package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StatsGetBroadcastStats(ctx context.Context, request *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stats.getBroadcastStats - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("stats.getBroadcastStats - not imp StatsGetBroadcastStats")
}
