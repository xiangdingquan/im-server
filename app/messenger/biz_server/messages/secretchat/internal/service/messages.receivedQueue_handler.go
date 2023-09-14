package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesReceivedQueue(ctx context.Context, request *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.receivedQueue#55a5bb66 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	receivedQ := &mtproto.Vector_Long{
		Datas: []int64{},
	}

	log.Debugf("messages.receivedQueue#55a5bb66 - reply: %s", logger.JsonDebugData(receivedQ))
	return receivedQ, nil
}
