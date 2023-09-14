package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.recoverPassword#4ea56e92 - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	var (
		err error = nil
	)

	if request.Code == "" {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
		log.Error(err.Error())
		return nil, err
	}

	err = s.AccountFacade.RecoverPassword(ctx, md.UserId, request.Code)
	if err != nil {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
		log.Error(err.Error())
		return nil, err
	}

	user, _ := s.UserFacade.GetUserSelf(ctx, md.UserId)
	authAuthorization := mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user,
	})

	log.Debugf("auth.recoverPassword#4ea56e92 - reply: %s\n", authAuthorization.DebugString())
	return authAuthorization.To_Auth_Authorization(), nil
}
