package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editInlineBotMessage#b0e08243 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not impl MessagesEditInlineBotMessage")
}
