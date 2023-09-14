package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetAllSecureValues(ctx context.Context, request *mtproto.TLAccountGetAllSecureValues) (*mtproto.Vector_SecureValue, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAllSecureValues - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getAllSecureValues - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.getAllSecureValues - error: %v", err)
	return nil, err
}
