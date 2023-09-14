package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetRecentLocations(ctx context.Context, request *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getRecentLocations#bbc45b09 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	messages := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	})

	log.Debugf("messages.getRecentLocations#bbc45b09 - reply: %s", messages.DebugString())
	return messages.To_Messages_Messages(), nil
}
