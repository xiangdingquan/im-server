package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/app/pkg/redis_util"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	rpcMetadataTimeout          int64 = 15 // salt timeout
	cacheRpcMetadataPrefix            = "rpc_metadata"
	cacheInlineBotResultsPrefix       = "inline_bot_results"
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

func genCacheRpcMetadataKey(queryId int64) string {
	return fmt.Sprintf("%s_%d", cacheRpcMetadataPrefix, queryId)
}

func genCacheInlineBotResultsKey(queryId int64) string {
	return fmt.Sprintf("%s_%d", cacheInlineBotResultsPrefix, queryId)
}

func (d *Redis) GetCacheRpcMetadata(ctx context.Context, queryId int64) (*grpc_util.RpcMetadata, error) {
	cacheKey := genCacheRpcMetadataKey(queryId)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	v, err := redis.Bytes(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
		return nil, err
	} else {
		log.Debugf(hack.String(v))
		md := new(grpc_util.RpcMetadata)
		err = json.Unmarshal(v, md)
		log.Debugf("cache md: %v", md)
		return md, err
	}
}

func (d *Redis) PutCacheRpcMetadata(ctx context.Context, md *grpc_util.RpcMetadata) (queryId int64, err error) {
	queryId = idgen.GetUUID()
	cacheKey := genCacheRpcMetadataKey(queryId)
	b, _ := json.Marshal(md)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("SETEX", cacheKey, rpcMetadataTimeout, b); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Redis) GetCacheInlineBotResults(ctx context.Context, queryId int64, id string) (botInlineResult *model.BotInlineIdResult, err error) {
	var (
		cacheKey = genCacheInlineBotResultsKey(queryId)
		v        []byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	v, err = redis.Bytes(conn.Do("HGET", cacheKey, id))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
	} else {
		log.Debugf(hack.String(v))
		botInlineResult = new(model.BotInlineIdResult)
		err = json.Unmarshal(v, botInlineResult)
		log.Debugf("cache md: %v", botInlineResult)
	}

	return
}

func (d *Redis) PutCacheInlineBotResults(ctx context.Context, botId int32, botInlineResults []*mtproto.BotInlineResult, cacheTime int32) (queryId int64, err error) {
	queryId = idgen.GetUUID()
	cacheKey := genCacheInlineBotResultsKey(queryId)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	for _, botInlineResult := range botInlineResults {
		b, _ := json.Marshal(&model.BotInlineIdResult{
			BotId:           botId,
			BotInlineResult: botInlineResult,
		})

		if _, err = conn.Do("HSET", cacheKey, botInlineResult.Id, b); err != nil {
			log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
		}
	}

	return
}
