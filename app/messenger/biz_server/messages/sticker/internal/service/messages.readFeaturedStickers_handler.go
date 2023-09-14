package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("MessagesReadFeaturedStickers - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.readFeaturedStickers#5b11812")

	return mtproto.ToBool(true), nil
}
