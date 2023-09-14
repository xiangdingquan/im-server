package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.importAuthorization#e3ef9613 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	authorization := mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		TmpSessions: &types.Int32Value{Value: request.GetId()},
		User:        mtproto.MakeTLUserEmpty(nil).To_User(),
	})

	log.Debugf("auth.importAuthorization#e3ef9613- reply: %s", logger.JsonDebugData(authorization))
	return authorization.To_Auth_Authorization(), nil
}
