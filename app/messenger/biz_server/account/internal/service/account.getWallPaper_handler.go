package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetWallPaper(ctx context.Context, request *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getWallPaper - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrMethodNotImpl
	return nil, err
}
