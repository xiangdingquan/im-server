package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsSetFeedBroadcasts7E91B8F2(ctx context.Context, request *mtproto.TLChannelsSetFeedBroadcasts7E91B8F2) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.setFeedBroadcasts7E91B8F2 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolFalse, nil
}
