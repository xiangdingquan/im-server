package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) BizInvokeBizDataRaw(ctx context.Context, request *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("biz.invokeBizDataRaw - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("biz.invokeBizDataRaw - not imp BizInvokeBizDataRaw")
}
