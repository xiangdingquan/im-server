package service

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("langpack.getLangPack - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("langpack.getLangPack - error: %v", err)
		return nil, err
	}

	langPack := request.LangPack
	if langPack == "" {
		langPack = md.Client
	}

	// 400	LANG_PACK_INVALID	The provided language pack is invalid
	if !model.CheckLangPackInvalid(langPack) {
		err := mtproto.ErrLangPackInvalid
		log.Errorf("langpack.getLangPack - error: %v", err)
		return nil, err
	}

	// check
	version, _, err := s.GetLanguage(ctx, langPack, request.LangCode)
	if err != nil {
		log.Errorf("langpack.getLangPack - error: %v", err)
		err = mtproto.ErrLangCodeNotSupported
		return nil, err
	}

	diff := mtproto.MakeTLLangPackDifference(&mtproto.LangPackDifference{
		LangCode:    request.LangCode,
		FromVersion: 0,
		Version:     version,
		Strings:     s.GetStrings(ctx, langPack, request.LangCode),
	}).To_LangPackDifference()

	log.Debugf("langpack.getLangPack#f2f2330a - reply: %s", diff.DebugString())
	return diff, nil
}
