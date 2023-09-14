package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getMaskStickers#65b8c79f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := &mtproto.TLMessagesAllStickers{Data2: &mtproto.Messages_AllStickers{
		Hash: request.Hash,
		Sets: []*mtproto.StickerSet{},
	}}
	log.Debugf("messages.getMaskStickers#65b8c79f - reply: %s", stickers.DebugString())
	return stickers.To_Messages_AllStickers(), nil
}
