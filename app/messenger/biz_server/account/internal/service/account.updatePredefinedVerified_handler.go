package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdatePredefinedVerified(ctx context.Context, request *mtproto.TLAccountUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updatePredefinedVerified - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.updatePredefinedVerified - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.updatePredefinedVerified - error: %v", err)
		return nil, err
	}

	r, err := s.UserFacade.UpdatePredefinedVerified(ctx,
		request.Phone,
		request.GetVerified())

	if err != nil {
		log.Errorf("account.updatePredefinedVerified - error: %v", err)
		return nil, err
	}

	if r.GetRegisteredUserId() == nil {
		user, _ := s.UserFacade.GetUserSelfByPhoneNumber(ctx, request.GetPhone())
		if user != nil {
			r.RegisteredUserId = &types.Int32Value{Value: user.GetId()}
		}
	}
	if r.GetRegisteredUserId() != nil {
		if _, err = s.updateVerified(ctx, md, r.GetRegisteredUserId().GetValue(), request.Verified); err != nil {
			return nil, err
		}
	}

	log.Debugf("account.updatePredefinedVerified - reply: %s", r.DebugString())
	return r, nil
}
