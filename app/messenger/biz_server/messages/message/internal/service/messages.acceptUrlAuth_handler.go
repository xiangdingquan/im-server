package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesAcceptUrlAuth(ctx context.Context, request *mtproto.TLMessagesAcceptUrlAuth) (*mtproto.UrlAuthResult, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.acceptUrlAuth - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("messages.acceptUrlAuth - not imp MessagesAcceptUrlAuth")
}
