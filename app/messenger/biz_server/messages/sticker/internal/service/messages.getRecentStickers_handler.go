package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getRecentStickers#5ea192c9 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := &mtproto.TLMessagesRecentStickers{Data2: &mtproto.Messages_RecentStickers{
		Hash:     request.Hash,
		Stickers: []*mtproto.Document{},
	}}

	log.Debugf("messages.getRecentStickers#5ea192c9 - reply: %s", stickers.DebugString())
	return stickers.To_Messages_RecentStickers(), nil
}
