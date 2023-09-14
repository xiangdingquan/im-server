package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetEmojiKeywordsLanguages(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getEmojiKeywordsLanguages - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	emojiLanguageList := &mtproto.Vector_EmojiLanguage{
		Datas: []*mtproto.EmojiLanguage{
			mtproto.MakeTLEmojiLanguage(&mtproto.EmojiLanguage{
				PredicateName: mtproto.Predicate_emojiLanguage,
				LangCode:      "en",
			}).To_EmojiLanguage(),
			mtproto.MakeTLEmojiLanguage(&mtproto.EmojiLanguage{
				PredicateName: mtproto.Predicate_emojiLanguage,
				LangCode:      "zh-hans",
			}).To_EmojiLanguage(),
			mtproto.MakeTLEmojiLanguage(&mtproto.EmojiLanguage{
				PredicateName: mtproto.Predicate_emojiLanguage,
				LangCode:      "zh-hant",
			}).To_EmojiLanguage(),
		},
	}

	log.Debugf("messages.getEmojiKeywordsLanguages#4e9963b2 - reply: %v", emojiLanguageList)
	return emojiLanguageList, nil
}
