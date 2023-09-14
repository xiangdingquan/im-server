package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getFeaturedStickers#2dacca4f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := &mtproto.TLMessagesFeaturedStickers{Data2: &mtproto.Messages_FeaturedStickers{
		Hash:   request.Hash,
		Sets:   []*mtproto.StickerSetCovered{},
		Unread: []int64{},
	}}

	log.Debugf("messages.getFeaturedStickers#2dacca4f - reply: %s", stickers.DebugString())
	return stickers.To_Messages_FeaturedStickers(), nil
}
