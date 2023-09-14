package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsChangeFeedBroadcast2528871E(ctx context.Context, request *mtproto.TLChannelsChangeFeedBroadcast2528871E) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.changeFeedBroadcast2528871E - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolFalse, nil
}
