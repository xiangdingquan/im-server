package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsGetBankCardData(ctx context.Context, request *mtproto.TLPaymentsGetBankCardData) (*mtproto.Payments_BankCardData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.getBankCardData - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("payments.getBankCardData - not imp PaymentsGetBankCardData")
}
