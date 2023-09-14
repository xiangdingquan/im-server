package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSaveRecentSticker(ctx context.Context, request *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.saveRecentSticker#392718f8 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.saveRecentSticker#392718f8")

	return mtproto.ToBool(true), nil
}
