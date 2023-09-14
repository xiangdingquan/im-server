package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsClearSavedInfo(ctx context.Context, request *mtproto.TLPaymentsClearSavedInfo) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.clearSavedInfo#d83d70c1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	log.Debugf("payments.clearSavedInfo#d83d70c1 - metadata {true}")
	return mtproto.ToBool(true), nil
}
