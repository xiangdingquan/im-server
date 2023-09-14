package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesSetInlineGameScore(ctx context.Context, request *mtproto.TLMessagesSetInlineGameScore) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.setInlineGameScore#15ad9f64 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl MessagesSetInlineGameScore")
}
