package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.uninstallStickerSet#f96e55de - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.uninstallStickerSet#f96e55de")

	return mtproto.ToBool(true), nil
}
