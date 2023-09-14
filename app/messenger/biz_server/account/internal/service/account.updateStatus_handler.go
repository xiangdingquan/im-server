package service

import (
	"context"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/pkg/util"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateStatus - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// var status *mtproto.UserStatus
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 401	SESSION_PASSWORD_NEEDED	2FA is enabled, use a password to login
	//
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.updateStatus - error: %v", err)
		return nil, err
	}

	var (
		offline     = mtproto.FromBool(request.GetOffline())
		now         = time.Now().Unix()
		pushUpdates = mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateUserStatus(&mtproto.Update{
				UserId: md.UserId,
				Status: nil,
			}).To_Update(),
			Date: int32(now) - 1,
		}).To_Updates()
	)
	// online
	s.UserFacade.UpdateUserStatus(ctx, md.UserId, now)
	if !offline {
		pushUpdates.Update.Status = mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
			Expires: int32(now) + 300,
		}).To_UserStatus()
	} else {
		pushUpdates.Update.Status = mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
			WasOnline: int32(now),
		}).To_UserStatus()
	}

	log.Debugf("account.updateStatus - reply: {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		// log.Debugf("ready push to other contacts...")
		// push to other contacts.
		contactIdList := s.UserFacade.GetContactUserIdList(context.Background(), md.UserId)
		blockedIdList := s.UserFacade.CheckBlockUserList(context.Background(), md.UserId, contactIdList)

		// push updateUserStatus规则
		for _, id := range contactIdList {
			if md.UserId == id {
				// why??
				continue
			}

			if blocked, _ := util.Contains(id, blockedIdList); blocked {
				continue
			}

			// log.Debugf("check blocked...")
			blocked := s.UserFacade.IsBlockedByUser(context.Background(), md.UserId, id)
			if blocked {
				continue
			}

			sync_client.PushUpdates(context.Background(), id, pushUpdates)
		}

	}).(*mtproto.Bool), nil
}
