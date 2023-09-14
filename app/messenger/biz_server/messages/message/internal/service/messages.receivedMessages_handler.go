package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesReceivedMessages(ctx context.Context, request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.receivedMessages#5a954c0 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	receivedM := &mtproto.Vector_ReceivedNotifyMessage{
		Datas: []*mtproto.ReceivedNotifyMessage{},
	}

	log.Debugf("messages.receivedMessages#5a954c0 - reply: %s", logger.JsonDebugData(receivedM))

	return receivedM, nil
}
