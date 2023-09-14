package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StatsLoadAsyncGraph(ctx context.Context, request *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stats.loadAsyncGraph - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("stats.loadAsyncGraph - not imp StatsLoadAsyncGraph")
}
