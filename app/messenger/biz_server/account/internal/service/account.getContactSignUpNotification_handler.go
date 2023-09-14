package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getContactSignUpNotification - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getContactSignUpNotification - error: %v", err)
		return nil, err
	}

	// "ContactSignUpNotification"
	rValue := s.AccountFacade.GetSettingValueBool(ctx, md.UserId, "contactSignUpNotification", true)

	log.Debugf("account.getContactSignUpNotification - reply: {%v}", rValue)
	return mtproto.ToBool(rValue), nil
}
