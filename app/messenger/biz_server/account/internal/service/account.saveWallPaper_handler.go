package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSaveWallPaper(ctx context.Context, request *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.saveWallPaper - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolTrue, nil
}
