package logic

import (
	"context"

	"github.com/gogo/protobuf/proto"

	"open.chat/app/messenger/biz_server/auth/internal/core"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

const (
	opTypeUnknown    = 0
	opTypeSendCode   = 1
	opTypeSignUp     = 2
	opTypeSignIn     = 3
	opTypeLogout     = 4
	opTypeResendCode = 5
	opTypeCancelCode = 6
)

func GetActionType(request proto.Message) int {
	switch request.(type) {
	case *mtproto.TLAuthSendCode:
		return opTypeSendCode
	case *mtproto.TLAuthResendCode:
		return opTypeResendCode
	case *mtproto.TLAuthSignIn:
		return opTypeSignIn
	case *mtproto.TLAuthSignUp:
		return opTypeSignUp
	case *mtproto.TLAuthLogOut:
		return opTypeLogout
	case *mtproto.TLAuthCancelCode:
		return opTypeCancelCode
	}
	return opTypeUnknown
}

// async
func DoLogAuthAction(core *core.AuthCore, md *grpc_util.RpcMetadata, phoneNumber string, actionType int, log string) {
	go func(authKeyId, msgId int64, clientIp string, phoneNumber string, actionType int, log string) {
		core.LogAuthAction(context.Background(), authKeyId, msgId, clientIp, phoneNumber, actionType, log)
	}(md.AuthId, md.ClientMsgId, md.ClientAddr, phoneNumber, actionType, log)
}
