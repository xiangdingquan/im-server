package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/util"
)

func (s *Service) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.exportAuthorization#e5bfffcd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	exported := mtproto.MakeTLAuthExportedAuthorization(&mtproto.Auth_ExportedAuthorization{
		Id:    md.UserId,
		Bytes: hack.Bytes(util.Int32ToString(request.DcId)),
	})

	log.Debugf("auth.exportAuthorization#e5bfffcd - reply: %s", logger.JsonDebugData(exported))
	return exported.To_Auth_ExportedAuthorization(), nil
}
