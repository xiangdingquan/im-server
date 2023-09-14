package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/app/bots/botfather/internal/model"
	"open.chat/app/pkg/redis_util"
	"open.chat/pkg/log"
)

const (
	cacheTimeout = 3 * 60
	cachePrefix  = "bot_father"
)

type Redis struct {
	redis *redis_util.Redis
}

func newRedisDao() *Redis {
	return &Redis{
		redis: redis_util.GetSingletonRedis(),
	}
}

func (d *Redis) Close() error {
	return d.redis.Redis.Close()
}

func (d *Redis) Ping(ctx context.Context) (err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func genCacheKey(userId int32) string {
	return fmt.Sprintf("%s_%d", cachePrefix, userId)
}

func (d *Redis) GetBotFatherCommandStates(ctx context.Context, userId int32) (states *model.BotFatherCommandStates, err error) {
	var b []byte
	cacheKey := genCacheKey(userId)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	b, err = redis.Bytes(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
			return
		} else {
			err = nil
		}
	}

	states = model.NewBotFatherCommandStates()
	if b != nil {
		err = json.Unmarshal(b, states)
		if err != nil {
			log.Errorf("error - %v", err)
		}
	}

	return
}

func (d *Redis) PutBotFatherCommandStates(ctx context.Context, userId int32, cData *model.BotFatherCommandStates) (err error) {
	cacheKey := genCacheKey(userId)
	cacheData, _ := json.Marshal(cData)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETEX", cacheKey, cacheTimeout, cacheData); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}

	return
}

func (d *Redis) DeleteBotFatherCommandStates(ctx context.Context, userId int32) (err error) {
	cacheKey := genCacheKey(userId)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("DEL", cacheKey, cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}

	return
}
