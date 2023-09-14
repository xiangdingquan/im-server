package service

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("langpack.getDifference - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("langpack.getDifference - error: %v", err)
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

	// check
	version, _, err := s.GetLanguage(ctx, langPack, request.LangCode)
	if err != nil {
		log.Errorf("langpack.getDifference - error: %v", err)
		err = mtproto.ErrLangCodeNotSupported
		return nil, err
	}

	diff := mtproto.MakeTLLangPackDifference(&mtproto.LangPackDifference{
		LangCode:    request.LangCode,
		FromVersion: request.FromVersion,
		Version:     version,
		Strings:     nil,
	}).To_LangPackDifference()

	if request.FromVersion >= version {
		diff.Strings = []*mtproto.LangPackString{}
	} else {
		diff.Strings = s.GetDifference(ctx, langPack, request.LangCode, request.FromVersion)
	}

	log.Debugf("langpack.getDifference - reply: %s", diff.DebugString())
	return diff, nil
}
