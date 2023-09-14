package dao

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/model"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"strconv"
)

const (
	cacheAuthKeyPrefix = "auth_keys"
)

func genCacheAuthKeyKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyPrefix, id)
}

func (d *Redis) PutAuthKey(ctx context.Context, keyId int64, keyData *model.AuthKeyData, expiredIn int32) (err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	var (
		key = genCacheAuthKeyKey(keyId)

		args = []interface{}{
			key,
			"auth_key_type", keyData.AuthKeyType,
			"auth_key_id", keyData.AuthKeyId,
			"auth_key", keyData.AuthKey,
			"perm_auth_key_id", keyData.PermAuthKeyId,
			"temp_auth_key_id", keyData.TempAuthKeyId,
			"media_temp_auth_key_id", keyData.MediaTempAuthKeyId,
		}
	)

	if err = conn.Send("HMSET", args...); err != nil {
		log.Error("conn.Send(HMSET %s,%v) error(%v)", key, args, err)
		return
	}

	if expiredIn > 0 {
		if err = conn.Send("EXPIRE", key, expiredIn); err != nil {
			log.Error("conn.Send(EXPIRE %d,%d) error(%v)", key, expiredIn, err)
			return
		}
	}

	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	if _, err = conn.Receive(); err != nil {
		log.Error("conn.Receive() error(%v)", err)
		return
	}
	if expiredIn > 0 {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

func (d *Redis) UnsafeBindKeyId(ctx context.Context, keyId int64, bindType int, bindKeyId int64) (err error) {
	var (
		key = genCacheAuthKeyKey(keyId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	switch bindType {
	case model.AuthKeyTypePerm:
		if _, err = conn.Do("HSET", key, "perm_auth_key_id", bindKeyId); err != nil {
			log.Errorf("conn.Do(HSET %s,perm_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	case model.AuthKeyTypeTemp:
		if _, err = conn.Do("HSET", key, "temp_auth_key_id", bindKeyId); err != nil {
			log.Errorf("conn.Do(HSET %s,temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	case model.AuthKeyTypeMediaTemp:
		if _, err = conn.Do("HSET", key, "media_temp_auth_key_id", bindKeyId); err != nil {
			log.Errorf("conn.Do(HSET %s,media_temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	default:
		return
	}

	return
}

func (d *Redis) GetAuthKey(ctx context.Context, keyId int64) (keyData *model.AuthKeyData, err error) {
	var (
		key    = genCacheAuthKeyKey(keyId)
		values [][]byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	values, err = redis.ByteSlices(conn.Do("HGETALL", key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGETALL %s) error(%v)", key, err)
		} else {
			err = nil
		}
		return
	} else if len(values) == 0 {
		err = fmt.Errorf("invalid auth_key")
		return
	}

	keyData = &model.AuthKeyData{}
	for i := 0; i < len(values); i = i + 2 {
		switch hack.String(values[i]) {
		case "auth_key_type":
			keyData.AuthKeyType, _ = strconv.Atoi(hack.String(values[i+1]))
		case "auth_key_id":
			keyData.AuthKeyId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "auth_key":
			keyData.AuthKey = values[i+1]
		case "perm_auth_key_id":
			keyData.PermAuthKeyId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "temp_auth_key_id":
			keyData.TempAuthKeyId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "media_temp_auth_key_id":
			keyData.MediaTempAuthKeyId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		}
	}

	return
}
