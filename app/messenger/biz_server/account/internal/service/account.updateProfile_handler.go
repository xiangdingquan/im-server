package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	sync_client "open.chat/app/messenger/sync/client"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateProfile#78515775 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	user, err := s.updateProfile(ctx, md, request.GetFirstName(), request.GetLastName(), request.GetAbout())
	if err == nil {
		log.Debugf("account.updateProfile#78515775 - reply: {%v}", user)
	}

	return user, err
}

func (s *Service) updateProfile(ctx context.Context, md *grpc_util.RpcMetadata, firstName, lastName, about *types.StringValue) (*mtproto.User, error) {
	user, _ := s.UserFacade.GetUserById(ctx, md.UserId, md.UserId)

	var isUpdateAbout = about != nil
	if isUpdateAbout {
		//// about长度<70并且可以为emtpy
		if len(about.GetValue()) > 70 {
			// return error
		}

		_, _ = s.UserFacade.UpdateAbout(ctx, md.UserId, about.GetValue())
	} else {
		if firstName == nil || firstName.Value == "" {
			err := mtproto.ErrFirstNameInvalid
			log.Errorf("bad request - %v", err)
			return nil, err
		}

		s.UserFacade.UpdateFirstAndLastName(ctx, md.UserId, firstName.GetValue(), lastName.GetValue())
		user.FirstName = firstName
		user.LastName = lastName

		go func() {
			// sync to other sessions
			updateUserName := mtproto.MakeTLUpdateUserName(&mtproto.Update{
				UserId:    md.UserId,
				FirstName: firstName.GetValue(),
				LastName:  lastName.GetValue(),
			}).To_Update()

			sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, model.MakeUpdatesByUpdates(updateUserName))
		}()
	}

	return user, nil
}
