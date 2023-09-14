package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdatePredefinedCode(ctx context.Context, request *mtproto.TLAccountUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updatePredefinedCode - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.updatePredefinedCode - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.updatePredefinedCode - error: %v", err)
		return nil, err
	}

	r, err := s.UserFacade.UpdatePredefinedCode(ctx,
		request.Phone,
		request.GetCode())

	if err != nil {
		log.Errorf("account.updatePredefinedCode - error: %v", err)
		return nil, err
	}

	log.Debugf("account.updatePredefinedCode - reply: %s", r.DebugString())
	return r, nil
}
