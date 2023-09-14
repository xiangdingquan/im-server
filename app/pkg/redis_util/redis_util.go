package redis_util

import (
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	xtime "github.com/go-kratos/kratos/pkg/time"
)

type Redis struct {
	Redis       *redis.Pool
	RedisExpire int32
}

var _self *Redis

func GetSingletonRedis() *Redis {
	if _self == nil {
		var (
			rc struct {
				Redis       *redis.Config
				RedisExpire xtime.Duration
			}
		)

		checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
		_self = &Redis{
			Redis:       redis.NewPool(rc.Redis),
			RedisExpire: int32(time.Duration(rc.RedisExpire) / time.Second),
		}
	}

	return _self
}

var _ssdb *Redis

func GetSingletonSsdb() *Redis {
	if _ssdb == nil {
		var (
			rc struct {
				Redis       *redis.Config
				RedisExpire xtime.Duration
			}
		)

		checkErr(paladin.Get("ssdb.toml").UnmarshalTOML(&rc))
		_ssdb = &Redis{
			Redis:       redis.NewPool(rc.Redis),
			RedisExpire: int32(time.Duration(rc.RedisExpire) / time.Second),
		}
	}

	return _ssdb
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
