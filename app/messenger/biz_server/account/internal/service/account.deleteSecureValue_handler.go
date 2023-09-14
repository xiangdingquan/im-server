package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountDeleteSecureValue(ctx context.Context, request *mtproto.TLAccountDeleteSecureValue) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.deleteSecureValue - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.deleteSecureValue - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.deleteSecureValue - error: %v", err)
	return nil, err
}
