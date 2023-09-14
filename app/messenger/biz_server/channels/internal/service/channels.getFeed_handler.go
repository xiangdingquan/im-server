package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetFeed(ctx context.Context, request *mtproto.TLChannelsGetFeed) (*mtproto.Messages_FeedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getFeed - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.MakeTLMessagesFeedMessagesNotModified(nil).To_Messages_FeedMessages(), nil
}
