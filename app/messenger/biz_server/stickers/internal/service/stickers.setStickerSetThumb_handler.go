package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StickersSetStickerSetThumb(ctx context.Context, request *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stickers.setStickerSetThumb - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("stickers.setStickerSetThumb - not imp StickersSetStickerSetThumb")
}
