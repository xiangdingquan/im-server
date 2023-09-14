package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getFavedStickers#21ce0b0e - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := mtproto.TLMessagesFavedStickers{Data2: &mtproto.Messages_FavedStickers{
		Hash:     request.Hash,
		Packs:    []*mtproto.StickerPack{},
		Stickers: []*mtproto.Document{},
	}}

	log.Debugf("messages.getFavedStickers#21ce0b0e - reply: %s", stickers.DebugString())
	return stickers.To_Messages_FavedStickers(), nil
}
