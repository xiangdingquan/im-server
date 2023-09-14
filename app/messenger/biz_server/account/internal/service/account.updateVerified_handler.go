package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdateVerified(ctx context.Context, request *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateVerified - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	user, err := s.updateVerified(ctx, md, request.Id, request.Verified)
	if err == nil {
		log.Debugf("account.updateVerified - reply: %s", user.DebugString())
	}

	return user, err
}

func (s *Service) updateVerified(ctx context.Context, md *grpc_util.RpcMetadata, id int32, verified bool) (*mtproto.User, error) {
	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.updateVerified - error: %v", err)
		return nil, err
	}

	user, err := s.UserFacade.GetUserSelf(ctx, id)
	if err != nil {
		log.Errorf("account.updateVerified - error: %v", err)
		return nil, err
	}

	_, err = s.UserFacade.UpdateVerified(ctx, user.Id, verified)
	if err != nil {
		log.Errorf("account.updateVerified - error: %v", err)
		return nil, err
	}

	user.Verified = verified

	go func() {
		sync_client.PushUpdates(context.Background(),
			user.Id,
			model.MakeUpdatesByUpdatesUsersChats([]*mtproto.User{user}, nil))
	}()

	return user, nil
}
