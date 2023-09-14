package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.unregisterDevice - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.unregisterDevice - error: %v", err)
		return nil, err
	}

	// Check token format by token_type
	if request.Token == "" {
		err := mtproto.ErrTokenInvalid
		log.Errorf("account.unregisterDevice - error:", err.Error())
		return nil, err
	}

	// check toke_type invalid
	if request.TokenType < model.PushTypeAPNS || request.TokenType > model.PushTypeMaxSize {
		err := mtproto.ErrTokenInvalid
		log.Errorf("account.unregisterDevice - error:", err.Error())
		return nil, err
	}

	// check token
	// 400	TOKEN_INVALID	The provided token is invalid

	if err := s.PushFacade.UnregisterDevice(ctx,
		md.UserId,
		md.AuthId,
		int(request.GetTokenType()),
		request.GetToken(),
		request.GetOtherUids()); err != nil {

		log.Errorf("account.unregisterDevice - error: {%v}", err)
	}

	log.Debugf("account.unregisterDevice - reply: {true}")
	return mtproto.ToBool(true), nil
}
