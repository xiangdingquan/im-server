package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.checkPassword - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("auth.checkPassword - error: %v", err)
		return nil, err
	}

	ok, err := s.AccountFacade.CheckPassword(ctx, md.UserId, request.Password)
	if err != nil {
		log.Error(err.Error())
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		return nil, err
	}

	if !ok {
		log.Errorf("auth.checkPassword, password not match")
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		return nil, err
	}

	user, _ := s.UserFacade.GetUserSelf(ctx, md.UserId)
	authAuthorization := mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user,
	}).To_Auth_Authorization()

	log.Debugf("auth.checkPassword - reply: %s\n", authAuthorization.DebugString())
	return authAuthorization, nil
}
