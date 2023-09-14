package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetEmojiURL(ctx context.Context, request *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getEmojiURL - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	return nil, fmt.Errorf("messages.getEmojiURL - not imp MessagesGetEmojiURL")
}
