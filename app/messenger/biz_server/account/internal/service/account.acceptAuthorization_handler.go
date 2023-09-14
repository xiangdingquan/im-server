package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountAcceptAuthorization(ctx context.Context, request *mtproto.TLAccountAcceptAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.acceptAuthorization - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.acceptAuthorization - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.acceptAuthorization - error: %v", err)
	return nil, err
}
