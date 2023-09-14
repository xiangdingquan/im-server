package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthRequestPasswordRecovery(ctx context.Context, request *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.requestPasswordRecovery#d897bc66 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	passwordRecovery, err := s.AccountFacade.RequestPasswordRecovery(ctx, md.UserId)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("auth.requestPasswordRecovery#d897bc66 - reply: %s\n", passwordRecovery.DebugString())
	return passwordRecovery, nil
}
