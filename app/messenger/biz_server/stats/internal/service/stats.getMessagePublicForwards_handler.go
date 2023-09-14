package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// stats.getMessagePublicForwards#5630281b channel:InputChannel msg_id:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *Service) StatsGetMessagePublicForwards(ctx context.Context, request *mtproto.TLStatsGetMessagePublicForwards) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stats.getMessagePublicForwards - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl StatsGetMessagePublicForwards logic
	log.Warn("stats.getMessagePublicForwards - error: method StatsGetMessagePublicForwards not impl")

	return nil, mtproto.ErrMethodNotImpl
}
