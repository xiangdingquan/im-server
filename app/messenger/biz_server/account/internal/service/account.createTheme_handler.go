package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountCreateTheme(ctx context.Context, request *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.createTheme - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrMethodNotImpl
	return nil, err
}
