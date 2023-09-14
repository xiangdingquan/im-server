package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdatePredefinedUsername(ctx context.Context, request *mtproto.TLAccountUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updatePredefinedUsername - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.updatePredefinedUsername - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.updatePredefinedUsername - error: %v", err)
		return nil, err
	}

	r, err := s.UserFacade.UpdatePredefinedUsername(ctx,
		request.Phone,
		request.GetUsername())

	if err != nil {
		log.Errorf("account.updatePredefinedUsername - error: %v", err)
		return nil, err
	}

	if r.GetRegisteredUserId() == nil {
		user, _ := s.UserFacade.GetUserSelfByPhoneNumber(ctx, request.GetPhone())
		if user != nil {
			r.RegisteredUserId = &types.Int32Value{Value: user.GetId()}
		}
	}

	if r.GetRegisteredUserId() != nil {
		md.UserId = r.GetRegisteredUserId().GetValue()
		if _, err = s.updateUsername(ctx, md, request.GetUsername()); err != nil {
			log.Errorf("account.updatePredefinedUsername - error: %v", err)
			return nil, err
		}
	}

	log.Debugf("account.updatePredefinedUsername - reply: %s", r.DebugString())
	return r, nil
}
