package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (s *Service) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateUsername#3e0bdd7c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !md.IsAdmin && !sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysPermitModifyUserName, false, 0) {
		user, _ := s.UserFacade.GetUserById(ctx, md.UserId, md.UserId)
		if len(user.GetUsername().GetValue()) > 0 {
			log.Errorf("account.updateUsername not update name")
			return nil, mtproto.ErrNotModifyUserName
		}
	}

	user, err := s.updateUsername(ctx, md, request.GetUsername())
	if err == nil {
		log.Debugf("account.updateUsername#3e0bdd7c - reply: %s", user.DebugString())
		return user, nil
	}

	return user, err
}

func (s *Service) updateUsername(ctx context.Context, md *grpc_util.RpcMetadata, username string) (*mtproto.User, error) {
	username2 := username
	if username2 != "" {
		if len(username) < model.MinUsernameLen ||
			!util.IsAlNumString(username) ||
			util.IsNumber(username[0]) {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
			log.Errorf("account.updateUsername#3e0bdd7c - format error: %v", err)
			return nil, err
		}

		ok, err := s.UsernameFacade.UpdateUsername(ctx, model.PEER_USER, md.UserId, username2)
		log.Debugf("ok: %v, err: %v", ok, err)
		if err != nil {
			log.Errorf("account.updateUsername#3e0bdd7c - format error: %v", err)
			return nil, err
		} else {
			if !ok {
				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
				log.Errorf("account.updateUsername#3e0bdd7c - format error: %v", err)
				return nil, err
			}
		}
	} else {
		// delete username
		_, err := s.UsernameFacade.DeleteUsername(ctx, username)
		if err != nil {
			log.Errorf("account.updateUsername#3e0bdd7c - format error: %v", err)
			return nil, err
		}
	}

	_, err := s.UserFacade.UpdateUsername(ctx, md.UserId, username2)
	if err != nil {
		log.Errorf("account.updateUsername#3e0bdd7c - format error: %v", err)
		return nil, err
	}

	user, _ := s.UserFacade.GetUserById(ctx, md.UserId, md.UserId)
	// 要考虑到数据库主从同步问题
	user.Username = &types.StringValue{
		Value: username,
	}

	go func() {
		sync_client.SyncUpdatesNotMe(context.Background(),
			md.UserId,
			md.AuthId,
			model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserName(&mtproto.Update{
				UserId:    md.UserId,
				FirstName: user.GetFirstName().GetValue(),
				LastName:  user.GetLastName().GetValue(),
				Username:  username,
			}).To_Update()))
	}()

	return user, nil
}
