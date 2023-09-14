package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetPassportConfig(ctx context.Context, request *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getPassportConfig - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getPassportConfig - error: %v", err)
		return nil, err
	}

	err := mtproto.ErrMethodNotImpl

	log.Warnf("account.getPassportConfig - error: %v", err)
	return nil, err
}
