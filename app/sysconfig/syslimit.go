package sysconfig

import (
	"context"
	"fmt"

	"open.chat/app/infra/databus/pkg/cache/redis"
	"open.chat/pkg/log"
)

const (
	limitTimePrefix = "limit_time"
)

func limitTimeKey(uid int32, key string) string {
	return fmt.Sprintf("%s_%d_%s", limitTimePrefix, uid, key)
}

func KeyIsExist(ctx context.Context, uid int32, key string, longtime uint32) bool {
	cacheKey := limitTimeKey(uid, key)
	conn := getGConfigs().Redis.Redis.Get(ctx)
	defer conn.Close()
	exists, _ := redis.Bool(conn.Do("EXISTS", cacheKey))
	if !exists {
		if _, err := conn.Do("SETEX", cacheKey, longtime, 1); err != nil {
			log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
		}
		return false
	}
	return true
}
