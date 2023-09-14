package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updatePasswordSettings#fa7c4b86 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Password == nil || request.NewSettings == nil {
		log.Errorf("password and new_settings is nil - bad request")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	err := s.AccountFacade.UpdatePasswordSetting(ctx, md.UserId, request.Password, request.NewSettings)

	if err != nil {
		log.Errorf("account.updatePasswordSettings#fa7c4b86 - error: %v", err)
		log.Errorf("account.updatePasswordSettings#fa7c4b86 - error: %v", err)
		return nil, err
	}

	log.Debugf("account.getPassword#548a30f5 - reply: {true}")
	return mtproto.ToBool(true), nil
}
