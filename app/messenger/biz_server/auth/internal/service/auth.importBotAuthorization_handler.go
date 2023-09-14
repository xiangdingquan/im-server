package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AuthImportBotAuthorization(ctx context.Context, request *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.importBotAuthorization#67a3ff2c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl AuthImportBotAuthorization")
}
