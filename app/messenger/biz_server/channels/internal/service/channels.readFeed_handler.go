package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsReadFeed(ctx context.Context, request *mtproto.TLChannelsReadFeed) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.readFeed - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	return model.MakeEmptyUpdates(), nil
}
