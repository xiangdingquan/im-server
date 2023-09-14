package service

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("langpack.getStrings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("langpack.getStrings - error: %v", err)
		return nil, err
	}

	langPack := request.LangPack
	if langPack == "" {
		langPack = md.Client
	}

	// 400	LANG_PACK_INVALID	The provided language pack is invalid
	if !model.CheckLangPackInvalid(langPack) {
		err := mtproto.ErrLangPackInvalid
		log.Errorf("langpack.getStrings - error: %v", err)
		return nil, err
	}

	_, _, err := s.GetLanguage(ctx, langPack, request.LangCode)
	if err != nil {
		log.Errorf("langpack.getStrings - error: %v", err)
		err = mtproto.ErrLangCodeNotSupported
		return nil, err
	}

	langPackStrings := new(mtproto.Vector_LangPackString)
	langPackStrings.Datas = s.GetStringListByIdList(ctx, langPack, request.LangCode, request.Keys)

	log.Debugf("langpack.getStrings - reply: %s", langPackStrings.DebugString())
	return langPackStrings, nil
}
