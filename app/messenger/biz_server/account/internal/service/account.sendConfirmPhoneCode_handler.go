package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.sendConfirmPhoneCode#1516d7bd - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrMethodNotImpl
	return nil, err
}
