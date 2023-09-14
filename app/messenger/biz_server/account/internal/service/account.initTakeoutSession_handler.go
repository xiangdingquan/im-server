package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountInitTakeoutSession(ctx context.Context, request *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.initTakeoutSession - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.initTakeoutSession - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Errorf("account.initTakeoutSession - error: %v", err)
	return nil, err
}
