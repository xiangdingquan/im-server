package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/app/service/auth_session/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	cacheSaltPrefix = "salts"
)

func genCacheSaltKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheSaltPrefix, id)
}

func (d *Redis) PutSalts(ctx context.Context, keyId int64, salts []*mtproto.TLFutureSalt) (err error) {
	var (
		b   []byte
		key = genCacheSaltKey(keyId)
	)
	b, err = json.Marshal(salts)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("SETEX", key, int64(len(salts)*model.SALT_TIMEOUT), b); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", key, err)
	}
	return
}

func (d *Redis) GetSalts(ctx context.Context, keyId int64) (salts []*mtproto.TLFutureSalt, err error) {
	var (
		key = genCacheSaltKey(keyId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	b, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", key, err)
		} else {
			err = nil
		}
		return
	}

	salts = make([]*mtproto.TLFutureSalt, 0, 32)
	if err = json.Unmarshal(b, &salts); err != nil {
		log.Error("getSalts json.Unmarshal(%s) error(%v)", b, err)
	}
	return
}
