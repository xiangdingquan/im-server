package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getEmojiKeywordsDifference#1508b6af - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	emojiKeywordsDifference := mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
		LangCode: request.LangCode,
		Version:  0,
		Keywords: []*mtproto.EmojiKeyword{},
	})

	log.Debugf("messages.getEmojiKeywordsDifference#1508b6af - reply: %s", emojiKeywordsDifference.DebugString())
	return emojiKeywordsDifference.To_EmojiKeywordsDifference(), nil
}
