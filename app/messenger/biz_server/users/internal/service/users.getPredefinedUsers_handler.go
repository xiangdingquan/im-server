package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UsersGetPredefinedUsers(ctx context.Context, request *mtproto.TLUsersGetPredefinedUsers) (*mtproto.Vector_PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.getPredefinedUsers - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.getPredefinedUsers - error: %v", err)
		return nil, err
	}

	predefinedUser := &mtproto.Vector_PredefinedUser{
		Datas: s.UserFacade.GetPredefinedUserList(ctx),
	}

	// set last_seen_at
	if len(predefinedUser.Datas) > 0 {
		idList := make([]int32, 0, len(predefinedUser.Datas))
		pList := make([]string, 0, len(predefinedUser.Datas))
		for _, u := range predefinedUser.Datas {
			pList = append(pList, u.Phone)
			if u.GetRegisteredUserId().GetValue() > 0 {
				idList = append(idList, u.GetRegisteredUserId().GetValue())
			}
		}
		if len(idList) > 0 {
			rMap := s.UserFacade.GetLastSeenList(ctx, idList)
			for id, lastSeenAt := range rMap {
				for i := 0; i < len(predefinedUser.Datas); i++ {
					if predefinedUser.Datas[i].RegisteredUserId.GetValue() == id {
						predefinedUser.Datas[i].LastSeenAt = &types.Int32Value{Value: int32(lastSeenAt)}
						break
					}
				}
			}
		}

		bannedList := s.BannedFacade.GetBannedByPhoneList(ctx, pList)
		for i := 0; i < len(predefinedUser.Datas); i++ {
			predefinedUser.Datas[i].Banned = bannedList[predefinedUser.Datas[i].Phone]
		}
	}

	log.Debugf("account.getPredefinedUsers - reply: %s", predefinedUser.DebugString())
	return predefinedUser, nil
}
