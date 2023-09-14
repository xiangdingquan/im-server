package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSetContentSettings(ctx context.Context, request *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.setContentSettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolTrue, nil
}
