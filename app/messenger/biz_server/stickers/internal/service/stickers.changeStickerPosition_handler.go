package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) StickersChangeStickerPosition(ctx context.Context, request *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("stickers.changeStickerPosition#ffb6d4ca - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not impl StickersChangeStickerPosition")
}
