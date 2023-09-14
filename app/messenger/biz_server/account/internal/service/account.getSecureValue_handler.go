package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetSecureValue(ctx context.Context, request *mtproto.TLAccountGetSecureValue) (*mtproto.Vector_SecureValue, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("account.getSecureValue - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getSecureValue - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.getSecureValue - error: %v", err)
	return nil, err
}
