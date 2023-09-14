package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetFeedSources(ctx context.Context, request *mtproto.TLChannelsGetFeedSources) (*mtproto.Channels_FeedSources, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getFeedSources - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.MakeTLChannelsFeedSourcesNotModified(nil).To_Channels_FeedSources(), nil
}
