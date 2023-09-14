package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetInactiveChannels(ctx context.Context, request *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getInactiveChannels - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("channels.getInactiveChannels - not imp ChannelsGetInactiveChannels")
}
