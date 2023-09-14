package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetSearchCounters(ctx context.Context, request *mtproto.TLMessagesGetSearchCounters) (*mtproto.Vector_Messages_SearchCounter, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getSearchCounters - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getSearchCounters - error: %v", err)
		return nil, err
	}

	return nil, fmt.Errorf("messages.getSearchCounters - not imp MessagesGetSearchCounters")
}
