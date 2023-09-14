package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSendVerifyEmailCode(ctx context.Context, request *mtproto.TLAccountSendVerifyEmailCode) (*mtproto.Account_SentEmailCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.sendVerifyEmailCode - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.sendVerifyEmailCode - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.sendVerifyEmailCode - error: %v", err)
	return nil, err
}
