package model

import (
	"context"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

type LangType int32

const (
	LocalizationDefault LangType = iota
	LocalizationEN
	LocalizationCN
)

type LocalizationWords map[LangType]string

// default is english
func isEnglish(code string) bool {
	return code == "en" || code == ""
}

func isChinese(code string) bool {
	return code == "zh" || code == "cn" || code == "zh-hans-raw"
}

func getLangCode(ctx context.Context, client authsessionpb.RPCSessionClient) string {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	mtLangCode, err := client.SessionGetLangCode(ctx, &authsessionpb.TLSessionGetLangCode{
		AuthKeyId: md.GetAuthId(),
	})
	if err != nil {
		return "en"
	}
	return mtLangCode.GetV()
}

func GetLangType(ctx context.Context, client authsessionpb.RPCSessionClient) LangType {
	code := getLangCode(ctx, client)
	if isEnglish(code) {
		return LocalizationEN
	}
	if isChinese(code) {
		return LocalizationCN
	}
	return LocalizationDefault
}

func Localize(ctx context.Context, client authsessionpb.RPCSessionClient, words LocalizationWords) string {
	t := GetLangType(ctx, client)
	log.Debugf("Localize words with lang type: %d", t)
	return words[t]
}
