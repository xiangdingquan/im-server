package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsSendPaymentForm(ctx context.Context, request *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.sendPaymentForm#2b8879b3 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl PaymentsSendPaymentForm")
}
