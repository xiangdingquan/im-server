package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.setContactSignUpNotification - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.setContactSignUpNotification - error: %v", err)
		return nil, err
	}

	if mtproto.FromBool(request.Silent) {
		s.AccountFacade.SetSettingValue(ctx, md.UserId, "contactSignUpNotification", "false")
	} else {
		s.AccountFacade.SetSettingValue(ctx, md.UserId, "contactSignUpNotification", "true")
	}

	log.Debugf("account.setContactSignUpNotification - reply: {true}")
	return mtproto.BoolTrue, nil
}
