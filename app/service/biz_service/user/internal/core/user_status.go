package core

import (
	"context"
	"time"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/mtproto"
)

var (
	userStatusEmpty = mtproto.MakeTLUserStatusEmpty(nil).To_UserStatus()
)

func makeUserStatusOnline() *mtproto.UserStatus {
	now := time.Now().Unix()
	status := &mtproto.UserStatus{
		PredicateName: mtproto.Predicate_userStatusOnline,
		Constructor:   mtproto.CRC32_userStatusOnline,
		Expires:       int32(now + 60),
	}
	return status
}

func makeUserStatus(do *dataobject.UserPresencesDO, showStatus bool) *mtproto.UserStatus {
	now := time.Now().Unix()

	if showStatus {
		if now <= do.LastSeenAt+60 {
			status := mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
				Expires: int32(do.LastSeenAt + 60),
			})
			return status.To_UserStatus()
		} else {
			status := mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
				WasOnline: int32(do.LastSeenAt),
			})
			return status.To_UserStatus()
		}
	} else {
		if now-do.LastSeenAt >= 60*60*24*30 {
			return nil
		} else if now-do.LastSeenAt >= 60*60*24*7 {
			return mtproto.MakeTLUserStatusLastMonth(nil).To_UserStatus()
		} else if now-do.LastSeenAt >= 60*60*24*3 {
			return mtproto.MakeTLUserStatusLastWeek(nil).To_UserStatus()
		} else {
			return mtproto.MakeTLUserStatusRecently(nil).To_UserStatus()
		}
	}
}

func (m *UserCore) GetUserStatus(ctx context.Context, selfId int32, presencesDO *dataobject.UserPresencesDO, isContact, isBlocked bool) *mtproto.UserStatus {
	if isBlocked {
		return nil
	}

	if presencesDO == nil {
		return nil
	}

	showStatus := false
	if env2.PredefinedUser {
		do, _ := m.UsersDAO.SelectById(ctx, selfId)
		if do != nil {
			if do.Verified == 1 {
				showStatus = true
			}
		}
	} else {
		showStatus = m.CheckLastSeenOnline(ctx, presencesDO.UserId, selfId, isContact)
	}
	return makeUserStatus(presencesDO, showStatus)
}

func (m *UserCore) GetUserStatus2(ctx context.Context, selfId, userId int32, isContact, isBlocked bool) *mtproto.UserStatus {
	if isBlocked {
		return nil
	}

	do, _ := m.UserPresencesDAO.Select(ctx, userId)
	return m.GetUserStatus(ctx, selfId, do, isContact, isBlocked)
}

func (m *UserCore) UpdateUserStatus(ctx context.Context, userId int32, lastSeenAt int64) error {
	_, err := m.UserPresencesDAO.UpdateLastSeenAt(ctx, lastSeenAt, userId)
	return err
}

func (m *UserCore) GetLastSeenList(ctx context.Context, id []int32) map[int32]int64 {
	if len(id) == 0 {
		return make(map[int32]int64)
	}

	doList, _ := m.UserPresencesDAO.SelectList(ctx, id)

	r := make(map[int32]int64, len(doList))
	for i := 0; i < len(doList); i++ {
		r[doList[i].UserId] = doList[i].LastSeenAt
	}

	return r
}
