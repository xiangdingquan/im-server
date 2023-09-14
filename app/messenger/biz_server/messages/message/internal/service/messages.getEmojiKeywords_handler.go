package service

import (
	"golang.org/x/net/context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetEmojiKeywords(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getEmojiKeywords#35a0e062 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	emojiKeywordsDifference := mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
		LangCode:    request.LangCode,
		FromVersion: 0,
		Keywords:    []*mtproto.EmojiKeyword{},
	}).To_EmojiKeywordsDifference()

	switch request.LangCode {
	case "en":
		emojiKeywordsDifference.Version = 5923
		emojiKeywordsDifference.Keywords = []*mtproto.EmojiKeyword{
			mtproto.MakeTLEmojiKeyword(&mtproto.EmojiKeyword{
				Keyword:   "lion",
				Emoticons: []string{"🦁"},
			}).To_EmojiKeyword(),
			mtproto.MakeTLEmojiKeyword(&mtproto.EmojiKeyword{
				Keyword:   "bitcoin",
				Emoticons: []string{"💸", "💵", "💲", "💰"},
			}).To_EmojiKeyword(),
			mtproto.MakeTLEmojiKeyword(&mtproto.EmojiKeyword{
				Keyword:   "lion",
				Emoticons: []string{"🇦🇮"},
			}).To_EmojiKeyword(),
		}
	case "zh-hans":
		emojiKeywordsDifference.Version = 0
	case "zh-hant":
		emojiKeywordsDifference.Version = 0
	}

	log.Debugf("messages.getEmojiKeywordsLanguages#4e9963b2 - reply: %s", emojiKeywordsDifference.DebugString())
	return emojiKeywordsDifference, nil
}
