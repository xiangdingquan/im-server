package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdatePredefinedProfile(ctx context.Context, request *mtproto.TLAccountUpdatePredefinedProfile) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updatePredefinedProfile - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.updatePredefinedProfile - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.updatePredefinedProfile - error: %v", err)
		return nil, err
	}

	r, err := s.UserFacade.UpdatePredefinedFirstAndLastName(ctx,
		request.Phone,
		request.GetFirstName().GetValue(),
		request.GetLastName().GetValue())
	if err != nil {
		log.Errorf("account.updatePredefinedProfile - error: %v", err)
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
		if _, err = s.updateProfile(ctx, md, request.GetFirstName(), request.GetLastName(), nil); err != nil {
			return nil, err
		}
	}

	log.Debugf("account.updatePredefinedProfile - reply: %s", r.DebugString())
	return r, nil
}
