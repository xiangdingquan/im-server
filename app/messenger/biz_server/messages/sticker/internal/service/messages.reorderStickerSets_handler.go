package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.reorderStickerSets#78337739 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.reorderStickerSets#78337739")

	return mtproto.ToBool(true), nil
}
