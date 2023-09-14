package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (s *Service) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.checkUsername#2714d86c - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	// Check username format
	// You can choose a username on Telegram.
	// If you do, other people will be able to find
	// you by this username and contact you
	// without knowing your phone number.
	//
	// You can use a-z, 0-9 and underscores.
	// Minimum length is 5 characters.";
	//
	if !model.CheckUsernameInvalid(request.Username) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
		log.Errorf("account.checkUsername#2714d86c - format error: %v", err)
		return nil, err
	}
	if len(request.Username) < model.MinUsernameLen ||
		!util.IsAlNumString(request.Username) ||
		util.IsNumber(request.Username[0]) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
		log.Errorf("account.checkUsername#2714d86c - format error: %v", err)
		return nil, err
	} else {
		existed, err := s.UsernameFacade.CheckAccountUsername(ctx, md.UserId, request.GetUsername())
		if err != nil {
			return nil, err
		}
		if existed == model.UsernameExistedNotMe {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
			log.Errorf("account.checkUsername#2714d86c - exists username: %v", err)
			return mtproto.ToBool(false), nil
		}
	}

	log.Debugf("account.checkUsername#2714d86c - reply: {true}")
	return mtproto.ToBool(true), nil
}
