package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.faveSticker#b9ffc55b - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.faveSticker#b9ffc55b")

	return mtproto.ToBool(true), nil
}
