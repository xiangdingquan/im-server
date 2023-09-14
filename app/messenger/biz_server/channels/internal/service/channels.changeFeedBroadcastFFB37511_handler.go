package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsChangeFeedBroadcastFFB37511(ctx context.Context, request *mtproto.TLChannelsChangeFeedBroadcastFFB37511) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.changeFeedBroadcastFFB37511 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return model.MakeEmptyUpdates(), nil
}
