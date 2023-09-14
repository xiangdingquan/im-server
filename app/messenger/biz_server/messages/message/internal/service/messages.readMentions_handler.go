package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReadMentions(ctx context.Context, request *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("MessagesReadMentions - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Debugf("messages.readMentions#f0189d3 - reply: {}")
	return nil, fmt.Errorf("not impl MessagesReadMentions")
}
