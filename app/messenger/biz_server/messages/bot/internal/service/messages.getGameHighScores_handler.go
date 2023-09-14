package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getGameHighScores#e822649d - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("not impl MessagesGetGameHighScores")
}
