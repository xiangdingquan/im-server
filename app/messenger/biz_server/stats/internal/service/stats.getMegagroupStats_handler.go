package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StatsGetMegagroupStats(ctx context.Context, request *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stats.getMegagroupStats - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("stats.getMegagroupStats - not imp StatsGetMegagroupStats")
}
