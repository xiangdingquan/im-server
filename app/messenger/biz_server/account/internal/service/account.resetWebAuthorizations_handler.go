package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountResetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountResetWebAuthorizations) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.resetWebAuthorizations#682d2594 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := mtproto.ToBool(true)

	log.Debugf("account.resetWebAuthorizations#682d2594 - reply: {true}")
	return reply, nil
}
