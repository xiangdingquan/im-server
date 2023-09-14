package service

import (
	"open.chat/mtproto"
)

func checkRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLAuthCheckedPhone,
		*mtproto.TLAuthLogOut,
		*mtproto.TLAuthSendCode,
		*mtproto.TLAuthResendCode,
		*mtproto.TLAuthSignIn,
		*mtproto.TLAuthSignUp,
		*mtproto.TLAuthExportedAuthorization,
		*mtproto.TLAuthExportAuthorization,
		*mtproto.TLAuthImportAuthorization,
		*mtproto.TLAuthCancelCode,
		*mtproto.TLAuthRequestPasswordRecovery,
		*mtproto.TLAuthCheckPassword,
		*mtproto.TLAuthRecoverPassword,
		*mtproto.TLUploadGetFile,
		*mtproto.TLUploadGetFileHashes,
		*mtproto.TLJsonObject,
		*mtproto.TLAuthExportLoginToken,
		*mtproto.TLAuthAcceptLoginToken,
		*mtproto.TLAuthImportLoginToken:
		return true

	case *mtproto.TLHelpGetConfig,
		*mtproto.TLHelpGetCdnConfig,
		*mtproto.TLHelpGetAppConfig,
		*mtproto.TLHelpGetNearestDc,
		*mtproto.TLAuthBindTempAuthKey:
		return true

	case *mtproto.TLLangpackGetLanguages,
		*mtproto.TLLangpackGetDifference,
		*mtproto.TLLangpackGetLangPack,
		*mtproto.TLLangpackGetStrings:
		return true

	case *mtproto.TLUploadGetWebFile:
		return true

	default:
		return false
	}
}
