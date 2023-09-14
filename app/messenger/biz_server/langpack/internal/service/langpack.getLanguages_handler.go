package service

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("langpack.getLanguages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("langpack.getLanguages - error: %v", err)
		return nil, err
	}

	langPack := request.LangPack
	if langPack == "" {
		langPack = md.Client
	}

	// 400	LANG_PACK_INVALID	The provided language pack is invalid
	if !model.CheckLangPackInvalid(langPack) {
		err := mtproto.ErrLangPackInvalid
		log.Errorf("langpack.getLanguages - error: %v", err)
		return nil, err
	}

	languages := new(mtproto.Vector_LangPackLanguage)
	languages.Datas = s.GetLanguages(ctx, langPack)

	log.Debugf("langpack.getLanguages - reply: %s", languages.DebugString())
	return languages, nil
}
