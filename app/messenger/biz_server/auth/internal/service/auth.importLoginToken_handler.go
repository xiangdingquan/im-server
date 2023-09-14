package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthImportLoginToken(ctx context.Context, request *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.importLoginToken - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrAuthTokenInvalid
	log.Errorf("auth.importLoginToken - error: %v", err)

	return nil, err
}
