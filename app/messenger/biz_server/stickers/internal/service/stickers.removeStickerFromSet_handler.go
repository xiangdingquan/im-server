package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stickers.removeStickerFromSet#f7760f51 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not impl StickersRemoveStickerFromSet")
}
