package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetSplitRanges(ctx context.Context, request *mtproto.TLMessagesGetSplitRanges) (*mtproto.Vector_MessageRange, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("MessagesGetSplitRanges - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not impl MessagesGetSplitRanges")
}
