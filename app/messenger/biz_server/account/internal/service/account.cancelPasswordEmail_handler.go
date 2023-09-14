package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountCancelPasswordEmail(ctx context.Context, request *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.cancelPasswordEmail - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.cancelPasswordEmail - error: %v", err)
		return nil, err
	}

	log.Debugf("account.cancelPasswordEmail - reply: {false}")
	return mtproto.BoolFalse, nil
}
