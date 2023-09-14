package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountGetWebAuthorizations) (*mtproto.Account_WebAuthorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getWebAuthorizations#182e6d6f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// account.webAuthorizations#ed56c9fc authorizations:Vector<WebAuthorization> users:Vector<User> = account.WebAuthorizations;
	reply := mtproto.MakeTLAccountWebAuthorizations(&mtproto.Account_WebAuthorizations{
		Authorizations: []*mtproto.WebAuthorization{},
		Users:          []*mtproto.User{},
	}).To_Account_WebAuthorizations()

	log.Debugf("account.getWebAuthorizations#182e6d6f - reply: %s", reply.DebugString())
	return reply, nil
}
