package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	xtime "github.com/go-kratos/kratos/pkg/time"

	id_facade "open.chat/app/service/idgen/facade"
	"open.chat/app/service/idgen/facade/dao"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

type RedisSeqClient struct {
	redis       *redis.Pool
	redisExpire int32
	*dao.Dao
}

func New() id_facade.SeqIDGen {
	var (
		rc struct {
			Seqsvr       *redis.Config
			SeqsvrExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("seqsvr.toml").UnmarshalTOML(&rc))

	return &RedisSeqClient{
		redis:       redis.NewPool(rc.Seqsvr),
		redisExpire: int32(time.Duration(rc.SeqsvrExpire) / time.Second),
		Dao:         dao.New(),
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *RedisSeqClient) GetCurrentSeqID(ctx context.Context, key string) (seq int64, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	seq, err = redis.Int64(conn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("redis_seq_client.GetCurrentSeqID - GET {%s}, error: {%v}", key, err)
		} else {
			err = nil
		}
	}
	return
}

func (c *RedisSeqClient) GetNextSeqID(ctx context.Context, key string) (seq int64, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()
	seq, err = redis.Int64(conn.Do("INCR", key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("redis_seq_client.GetNextSeqID - INCR {%s}, error: {%v}", key, err)
		} else {
			err = nil
		}
	}

	return
}

func (c *RedisSeqClient) GetNextNSeqID(ctx context.Context, key string, n int) (seq int64, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()
	seq, err = redis.Int64(conn.Do("INCRBY", key, n))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("redis_seq_client.GetNextNSeqID - INCR {%s}, error: {%v}", key, err)
		} else {
			err = nil
		}
	}

	return
}

func (c *RedisSeqClient) SetCurrentSeqID(ctx context.Context, key string, v int64) (err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("SET", key, v)
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("redis_seq_client.SetSeqID - SET {%s, %d}, error: {%v}", key, v, err)
		} else {
			err = nil
		}
	}

	return
}

func (c *RedisSeqClient) GetNextPhoneNumber(ctx context.Context, prefix string) (string, error) {
	if len(prefix) != 5 {
		return "", errors.New("prefix error")
	}
	tR := sqlx.TxWrapper(ctx, c.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		seq, err := c.Dao.GetNextSeqTx(tx, prefix)
		result.Data = seq
		result.Err = err
	})
	if tR.Err != nil {
		log.Errorf("GetNextPhoneNumber(GET %s) error(%v)", prefix, tR.Err)
		return "", tR.Err
	}
	number := util.Int64ToString(tR.Data.(int64))
	max := 8 - len(number)
	for i := 0; i < max; i++ {
		number = "0" + number
	}
	return prefix + number, nil
}

func init() {
	id_facade.SeqIDGenRegister("redis", New)
}
