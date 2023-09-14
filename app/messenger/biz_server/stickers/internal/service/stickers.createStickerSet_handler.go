package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StickersCreateStickerSet(ctx context.Context, request *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stickers.createStickerSet#9bd86e6a - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not impl StickersCreateStickerSet")
}
