package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetAttachedStickers(ctx context.Context, request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getAttachedStickers#cc5b67cc - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("messages.getAttachedStickers#cc5b67cc")

	return &mtproto.Vector_StickerSetCovered{Datas: []*mtproto.StickerSetCovered{}}, nil
}
