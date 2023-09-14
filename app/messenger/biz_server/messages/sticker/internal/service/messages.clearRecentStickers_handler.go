package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("MessagesClearRecentStickers - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.clearRecentStickers#8999602d")

	return mtproto.ToBool(true), nil
}
