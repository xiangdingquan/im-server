package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpAcceptTermsOfService(ctx context.Context, request *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.acceptTermsOfService - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.acceptTermsOfService - error: %v", err)
		return nil, err
	}

	reply := mtproto.ToBool(true)

	log.Debugf("help.acceptTermsOfService - reply: %s", reply.DebugString())
	return reply, nil
}
