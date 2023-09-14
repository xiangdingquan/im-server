package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getArchivedStickers#57f17692 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := &mtproto.TLMessagesArchivedStickers{Data2: &mtproto.Messages_ArchivedStickers{
		Count: 0,
		Sets:  []*mtproto.StickerSetCovered{},
	}}

	log.Debugf("messages.getArchivedStickers#57f17692 - reply: %s", stickers.DebugString())
	return stickers.To_Messages_ArchivedStickers(), nil
}
