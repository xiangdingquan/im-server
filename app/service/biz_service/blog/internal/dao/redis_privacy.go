package dao

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/pkg/log"
	"strconv"
)

const privacyKeyPrefix = "blog_privacy"

func getPrivacyCacheKey(userId int32) string {
	return fmt.Sprintf("%s_%d", privacyKeyPrefix, userId)
}

func (d *Redis) SetPrivacy(ctx context.Context, userId int32, key int8, rules string) (err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("HSET", cKey, key, rules)
	if err != nil {
		log.Error("conn.Do(HSET %s %d %s) error(%v)", cKey, key, rules, err)
	}

	return
}

func (d *Redis) SetPrivacyList(ctx context.Context, userId int32, m map[int8]string) (err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	args := []interface{}{cKey}
	for k, v := range m {
		args = append(args, strconv.Itoa(int(k)))
		args = append(args, v)
	}
	_, err = conn.Do("HMSET", args...)

	if err != nil {
		log.Errorf("conn.Do(HMSET %s), error(%v)", cKey, err)
	}

	return
}

func (d *Redis) GetPrivacy(ctx context.Context, userId int32, key int8) (rules string, err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	rules, err = redis.String(conn.Do("HGET", cKey, key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGET %s %d), error(%v)", cKey, key, err)
		} else {
			err = nil
		}
	}

	return
}

func (d *Redis) GetUserPrivacy(ctx context.Context, userId int32) (out map[int8]string, err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	m, err := redis.StringMap(conn.Do("HGETALL", cKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGETALL %s), error(%v)", cKey, err)
		} else {
			err = nil
		}
		return
	} else if len(m) == 0 {
		return
	}

	out = make(map[int8]string)
	for k, v := range m {
		var i int64
		i, err = strconv.ParseInt(k, 10, 8)
		if err != nil {
			log.Errorf("parseInt in GetUserPricacy(%d), s: %s, errror: %v", userId, k, err)
			return
		}
		out[int8(i)] = v
	}

	return
}

func (d *Redis) ExpirePrivacy(ctx context.Context, userId int32, seconds int32) (err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("EXPIRE", cKey, seconds)
	if err != nil {
		log.Errorf("conn.Do(EXPIRE %s %d), error(%v)", cKey, seconds, err)
	}

	return
}

func (d *Redis) IsPrivacyExists(ctx context.Context, userId int32) (isExists bool, err error) {
	var (
		cKey = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	isExists, err = redis.Bool(conn.Do("EXISTS", cKey))
	if err != nil {
		log.Errorf("conn.Do(EXISTS %s), error(%v)", cKey, err)
	}

	return
}
