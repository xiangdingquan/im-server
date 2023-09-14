package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsGetPaymentForm(ctx context.Context, request *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.getPaymentForm#99f09745 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	return nil, fmt.Errorf("not impl PaymentsGetPaymentForm")
}
