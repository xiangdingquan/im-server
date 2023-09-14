package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.validateRequestedInfo#770a8e74 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl PaymentsValidateRequestedInfo")
}
