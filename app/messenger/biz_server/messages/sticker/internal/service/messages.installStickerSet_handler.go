package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.installStickerSet#c78fe460 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.installStickerSet#c78fe460")

	installResult := mtproto.MakeTLMessagesStickerSetInstallResultSuccess(nil)
	return installResult.To_Messages_StickerSetInstallResult(), nil
}
