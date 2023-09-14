package dao

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/pkg/log"
	"strconv"
)

const (
	passwordFloodPrefix      = "password_flood"
	passwordFloodCountPrefix = "password_flood_count"
)

func passwordFloodKey(uid int32) string {
	return fmt.Sprintf("%s_%d", passwordFloodPrefix, uid)
}

func passwordFloodCountKey(uid int32) string {
	return fmt.Sprintf("%s_%d", passwordFloodCountPrefix, uid)
}

func (d *Redis) IsPasswordFlood(ctx context.Context, uid int32) (bool, error) {
	key := passwordFloodKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	v, err := redis.Bool(conn.Do("EXISTS", key))
	log.Debugf("IsPasswordFlood, redis exists, key: %s, err: %v, v: %v", key, err, v)
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(EXISTS %s) error(%v)", key, err)
		} else {
			err = nil
		}
		return false, err
	}
	return v, nil
}

func (d *Redis) IncPFCount(ctx context.Context, uid int32, floodExpireInSecond, countExpireInSeocnd int32, limit int32) (flooded bool, err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	src := `
local countKey = KEYS[1];
local floodKey = KEYS[2];
local floodExpireInSecond = ARGV[1];
local countExpireInSecond = ARGV[2];
local floodLimit = tonumber(ARGV[3]);
local exists = redis.call("EXISTS", countKey);
if exists and exists == 1 then
	redis.call("INCR", countKey);
else
	redis.call("SETEX", countKey, countExpireInSecond, 1);
end
local count = tonumber(redis.call("GET", KEYS[1]));
if count and count >= floodLimit and floodLimit > 0 then
	redis.call("SETEX", floodKey, floodExpireInSecond, 1);
	return 1;
else
	return 2;
end
`
	script := redis.NewScript(2, src)
	keysAndArgs := []interface{}{
		passwordFloodCountKey(uid), passwordFloodKey(uid),
		strconv.Itoa(int(floodExpireInSecond)), strconv.Itoa(int(countExpireInSeocnd)), strconv.Itoa(int(limit)),
	}

	log.Debugf("IncPFCount countKey: %s, floodKey: %s, floodExpireInSecond: %d, countExpireInSecond: %d, limit: %d", passwordFloodCountKey(uid), passwordFloodKey(uid), floodExpireInSecond, countExpireInSeocnd, limit)

	ret, err := redis.Int64(script.Do(conn, keysAndArgs...))
	if err != nil {
		log.Errorf("IncPFCount countKey: %s, floodKey: %s, floodExpireInSecond: %d, countExpireInSecond: %d, limit: %d, err: %v", passwordFloodCountKey(uid), passwordFloodKey(uid), floodExpireInSecond, countExpireInSeocnd, limit, err)
		return false, err
	}
	log.Debugf("IncPFCount countKey: %s, floodKey: %s, floodExpireInSecond: %d, countExpireInSecond: %d, limit: %d, ret: %d", passwordFloodCountKey(uid), passwordFloodKey(uid), floodExpireInSecond, countExpireInSeocnd, limit, ret)
	return ret == 1, nil
}

func (d *Redis) GetPasswordFloodCount(ctx context.Context, uid int32) (int32, error) {
	key := passwordFloodCountKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	v, err := redis.Int64(conn.Do("GET", key))
	log.Debugf("GetPasswordFloodCount, redis get, key: %s, value: %d, err: %v", key, v, err)
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", key, err)
		} else {
			err = nil
		}
		return 0, err
	}
	return int32(v), nil
}

func (d *Redis) CleanPasswordFloodCount(ctx context.Context, uid int32) error {
	key := passwordFloodCountKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	log.Debugf("CleanPasswordFloodCount, redis del, key: %s, err: %v", key, err)
	if err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", key, err)
	}
	return err
}
