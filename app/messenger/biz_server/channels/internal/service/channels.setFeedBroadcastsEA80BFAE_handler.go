package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsSetFeedBroadcastsEA80BFAE(ctx context.Context, request *mtproto.TLChannelsSetFeedBroadcastsEA80BFAE) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.setFeedBroadcastsEA80BFAE - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return model.MakeEmptyUpdates(), nil
}
