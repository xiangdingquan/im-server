package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountFinishTakeoutSession(ctx context.Context, request *mtproto.TLAccountFinishTakeoutSession) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.finishTakeoutSession - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.finishTakeoutSession - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Errorf("account.finishTakeoutSession - error: %v", err)
	return nil, err
}
