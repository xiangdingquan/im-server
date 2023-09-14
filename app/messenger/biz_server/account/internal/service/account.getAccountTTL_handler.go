package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (reply *mtproto.AccountDaysTTL, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAccountTTL - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.getAccountTTL - error: %v", err)
		return
	}

	var days int32
	if days, err = s.UserFacade.GetAccountDaysTTL(ctx, md.UserId); err != nil {
		log.Errorf("account.getAccountTTL - error: %v", err)
		return
	}

	reply = mtproto.MakeTLAccountDaysTTL(&mtproto.AccountDaysTTL{
		Days: days,
	}).To_AccountDaysTTL()

	log.Infof("account.getAccountTTL - reply: %s\n", reply.DebugString())
	return
}
