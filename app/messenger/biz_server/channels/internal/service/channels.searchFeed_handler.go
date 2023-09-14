package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsSearchFeed(ctx context.Context, request *mtproto.TLChannelsSearchFeed) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.searchFeed#88325369 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	result := mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages()

	log.Debugf("channels.searchFeed#88325369 - reply %s", result.DebugString())
	return result, nil
}
