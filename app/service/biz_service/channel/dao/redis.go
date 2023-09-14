package dao

import (
	"context"
	"open.chat/app/pkg/redis_util"
	"open.chat/pkg/log"
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
