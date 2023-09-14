package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.registerDevice - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.registerDevice - error: %v", err)
		return nil, err
	}

	// Check token format by token_type
	if request.Token == "" {
		err := mtproto.ErrTokenInvalid
		log.Errorf("account.registerDevice - error:", err.Error())
		return nil, err
	}

	// check toke_type invalid
	if request.TokenType < model.PushTypeAPNS || request.TokenType > model.PushTypeMaxSize {
		err := mtproto.ErrTokenInvalid
		log.Errorf("account.registerDevice - error:", err.Error())
		return nil, err
	}

	// check token
	// 400	TOKEN_INVALID	The provided token is invalid

	appSandbox := mtproto.FromBool(request.GetAppSandbox())
	if err := s.PushFacade.RegisterDevice(ctx,
		md.UserId,
		md.AuthId,
		int(request.GetTokenType()),
		request.GetToken(),
		request.GetNoMuted(),
		appSandbox,
		request.GetSecret(),
		request.GetOtherUids()); err != nil {

		log.Errorf("account.registerDevice#5cbea590 - error: {%v}", err)
	}

	log.Debugf("account.registerDevice#5cbea590 - reply: {true}")
	return mtproto.ToBool(true), nil
}
