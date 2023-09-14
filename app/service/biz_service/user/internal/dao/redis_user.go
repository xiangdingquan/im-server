package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"strconv"

	"open.chat/app/infra/databus/pkg/cache/redis"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	userKeyPrefix = "user"
)

func getUserCacheKey(userId int32) string {
	return fmt.Sprintf("%s_%d", userKeyPrefix, userId)
}

func (d *Redis) PutCacheUser2(ctx context.Context, user *model.ImmutableUser) (err error) {
	var (
		args = []interface{}{getUserCacheKey(user.ID())}
		b    []byte
	)

	// "user"
	b, _ = json.Marshal(user.User)
	args = append(args, "user2", b)

	// bot
	if user.IsBot() {
		b, _ = json.Marshal(user.Bot)
		args = append(args, "bot2", b)
	}

	// photo
	b, _ = json.Marshal(user.ProfilePhoto())
	args = append(args, "photo2", b)

	// last_seen_at
	args = append(args, "was_online2", user.LastSeenAt())

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("HMSET", args...)
	if err != nil {
		log.Error("conn.Set(HMSET) error(%v)", err)
	}
	return
}

func (d *Redis) GetCacheUser2(ctx context.Context, userId int32) (user *model.ImmutableUser, err error) {
	var (
		key    = getUserCacheKey(userId)
		values [][]byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	values, err = redis.ByteSlices(conn.Do("HMGET", key, "user2", "photo2", "was_online2"))
	if err != nil {
		log.Error("conn.Do(HMGET) error(%v)", err)
		return
	}
	if len(values[0]) == 0 {
		return
	}

	userData := &model.UserData{}
	err = json.Unmarshal(values[0], userData)
	if err != nil {
		return
	}

	if len(values[1]) > 0 {
		photo := &mtproto.UserProfilePhoto{}
		if json.Unmarshal(values[1], photo) == nil {
			user.User.ProfilePhoto = photo
		}
	}
	if user.User.ProfilePhoto == nil {
		user.User.ProfilePhoto = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}
	var wasOnline int64
	if len(values[2]) > 0 {
		wasOnline, _ = strconv.ParseInt(hack.String(values[2]), 10, 64)
	}
	user.User.LastSeenAt = int32(wasOnline)

	return
}
