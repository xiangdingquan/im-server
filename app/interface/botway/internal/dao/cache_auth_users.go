package dao

import (
	"context"
	"open.chat/pkg/log"
)

type cacheAuthUser struct {
	userId    int32
	authKeyId int64
	layer     int32
}

func newCacheAuthUser() *cacheAuthUser {
	return &cacheAuthUser{
		userId:    0,
		authKeyId: 0,
		layer:     110,
	}
}

func (cv *cacheAuthUser) Size() int {
	return 1
}

func (cv *cacheAuthUser) UserId() int32 {
	return cv.userId
}

func (cv *cacheAuthUser) AuthKeyId() int64 {
	return cv.authKeyId
}

func (cv *cacheAuthUser) Layer() int32 {
	return cv.layer
}

func (d *Dao) GetCacheAuthUser(ctx context.Context, botId, token string) (cv *cacheAuthUser, err error) {
	if v, ok := d.cache.Get(token); !ok {
		cv = newCacheAuthUser()
		err = d.Mysql.DB.QueryRow(ctx, "SELECT id, secret_key_id FROM users WHERE id = ?", botId).Scan(&cv.userId, &cv.authKeyId)
		if err != nil {
			log.Error("d.GetAuthUsersById.Query error(%v)", err)
			return
		} else {
			d.cache.Set(token, cv)
		}
	} else {
		cv = v.(*cacheAuthUser)
	}
	return
}
