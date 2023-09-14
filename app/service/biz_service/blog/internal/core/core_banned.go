package core

import (
	"context"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"time"
)

type BlogBannedUser struct {
	UserId  int32 `json:"userId"`
	BanFrom int32 `json:"banFrom"`
	BanTo   int32 `json:"banTo"`
}

func (m *BlogCore) AddBannedUser(ctx context.Context, userId int32, banFrom, banTo int32) error {
	do := &dataobject.BlogBannedUsersDo{
		UserId:  userId,
		BanFrom: banFrom,
		BanTo:   banTo,
	}
	return m.BlogBannedUsersDAO.InsertOrUpdate(ctx, do)
}

func (m *BlogCore) GetBannedUsers(ctx context.Context, offset, limit int32) (map[int32][]int32, error) {
	l, err := m.BlogBannedUsersDAO.Select(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	out := make(map[int32][]int32)
	for _, v := range l {
		out[v.UserId] = []int32{v.BanFrom, v.BanTo}
	}
	return out, nil
}

func (m *BlogCore) IsBannedUser(ctx context.Context, userId int32) (bool, error) {
	doList, err := m.BlogBannedUsersDAO.SelectByUsers(ctx, []int32{userId})
	if err != nil {
		return true, err
	}

	var isBanned bool
	if len(doList) == 0 {
		isBanned = false
	} else {
		now := int32(time.Now().Unix())
		for _, do := range doList {
			if do.UserId == userId {
				isBanned = do.BanFrom < now && now < do.BanTo
				break
			}
		}
		isBanned = false
	}

	return isBanned, nil
}

func (m *BlogCore) GetBannedByUserIds(ctx context.Context, uidList []int32) (map[int32][]int32, error) {
	l, err := m.BlogBannedUsersDAO.SelectByUsers(ctx, uidList)
	if err != nil {
		return nil, err
	}

	out := make(map[int32][]int32)
	for _, v := range l {
		out[v.UserId] = []int32{v.BanFrom, v.BanTo}
	}
	return out, nil
}
