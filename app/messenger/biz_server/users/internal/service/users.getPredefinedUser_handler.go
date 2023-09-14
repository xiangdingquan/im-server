package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UsersGetPredefinedUser(ctx context.Context, request *mtproto.TLUsersGetPredefinedUser) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.getPredefinedUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.getPredefinedUser - error: %v", err)
		return nil, err
	}

	user, err := s.UserFacade.GetPredefinedUser(ctx, request.Phone)
	if err != nil {
		log.Errorf("account.getPredefinedUser - error: %v", err)
		err = mtproto.ErrPhoneNumberUnoccupied
		return nil, err
	}

	if user.GetRegisteredUserId().GetValue() != 0 {
		rMap := s.UserFacade.GetLastSeenList(ctx, []int32{user.GetRegisteredUserId().GetValue()})
		user.RegisteredUserId = &types.Int32Value{Value: int32(rMap[user.GetRegisteredUserId().GetValue()])}
	}

	bannedList := s.BannedFacade.GetBannedByPhoneList(ctx, []string{request.Phone})
	user.Banned, _ = bannedList[request.Phone]

	log.Debugf("account.getPredefinedUser - reply: %s", user.DebugString())
	return user, nil
}
