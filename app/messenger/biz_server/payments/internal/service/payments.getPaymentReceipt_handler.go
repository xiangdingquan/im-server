package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsGetPaymentReceipt(ctx context.Context, request *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.getPaymentReceipt#a092a980 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl PaymentsGetPaymentReceipt")
}
