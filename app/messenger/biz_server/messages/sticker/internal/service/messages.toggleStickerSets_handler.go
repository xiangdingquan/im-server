package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesToggleStickerSets(ctx context.Context, request *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.toggleStickerSets - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolFalse, nil
}
