package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UsersGetMe(ctx context.Context, request *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.getMe - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("users.getMe - error: %v", err)
		return nil, err
	}

	user, err := s.UserFacade.GetUserByToken(ctx, request.Token)
	if err != nil || user == nil {
		log.Errorf("users.getMe - error: %v", err)
		return nil, err
	} else if user.Id != request.Id {
		err = mtproto.ErrTokenInvalid
		log.Errorf("users.getMe - error: %v", err)
		return nil, err
	}

	log.Debugf("users.getMe - reply: %s", user.DebugString())
	return user, nil
}
