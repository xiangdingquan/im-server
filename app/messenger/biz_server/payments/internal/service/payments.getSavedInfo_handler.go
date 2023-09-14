package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PaymentsGetSavedInfo(ctx context.Context, request *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("payments.getSavedInfo#227d824b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl PaymentsGetSavedInfo")
}
