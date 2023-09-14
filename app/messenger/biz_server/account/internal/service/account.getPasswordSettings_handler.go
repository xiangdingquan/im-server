package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getPasswordSettings#bc8d11bb - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	settings, err := s.AccountFacade.GetPasswordSetting(ctx, md.UserId, request.GetPassword())
	if err != nil {
		log.Errorf("account.getPassword#548a30f5 - error: %v", err)
		return nil, err
	}

	log.Debugf("account.getPasswordSettings#bc8d11bb - reply: %s", settings.DebugString())
	return settings, nil
}
