package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountConfirmPasswordEmail(ctx context.Context, request *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.confirmPasswordEmail#8fdf1920 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolFalse, nil
}
