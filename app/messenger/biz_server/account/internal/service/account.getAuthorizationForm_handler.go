package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetAuthorizationForm(ctx context.Context, request *mtproto.TLAccountGetAuthorizationForm) (*mtproto.Account_AuthorizationForm, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAuthorizationForm - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getAuthorizationForm - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.getAuthorizationForm - error: %v", err)
	return nil, err
}
