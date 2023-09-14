package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountResetWebAuthorization(ctx context.Context, request *mtproto.TLAccountResetWebAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.resetWebAuthorization#2d01b9ef - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := mtproto.ToBool(true)

	log.Debugf("account.resetWebAuthorization#2d01b9ef - reply: {true}")
	return reply, nil
}
