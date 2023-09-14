package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	cacheQRCodePrefix = "qr_codes"
)

func genQRLoginCodeKey(authKeyId int64) string {
	return fmt.Sprintf("%s_%d", cacheQRCodePrefix, authKeyId)
}

func (d *Redis) GetCacheQRLoginCode(ctx context.Context, keyId int64) (code *model.QRCodeTransaction, err error) {
	var (
		key    = genQRLoginCodeKey(keyId)
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
		return
	}

	code = new(model.QRCodeTransaction)
	for i := 0; i < len(values); i = i + 2 {
		switch hack.String(values[i]) {
		case "auth_key_id":
			code.AuthKeyId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "session_id":
			code.SessionId, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "server_id":
			code.ServerId = hack.String(values[i+1])
		case "api_id":
			v, _ := strconv.ParseInt(hack.String(values[i+1]), 10, 64)
			code.ApiId = int32(v)
		case "api_hash":
			code.ApiHash = hack.String(values[i+1])
		case "code_hash":
			code.CodeHash = hack.String(values[i+1])
		case "expire_at":
			code.ExpireAt, _ = strconv.ParseInt(hack.String(values[i+1]), 10, 64)
		case "user_id":
			v, _ := strconv.ParseInt(hack.String(values[i+1]), 10, 64)
			code.UserId = int32(v)
		case "state":
			v, _ := strconv.ParseInt(hack.String(values[i+1]), 10, 64)
			code.State = int(v)
		}
	}

	return
}

func (d *Redis) PutCacheQRLoginCode(ctx context.Context, keyId int64, qrCode *model.QRCodeTransaction, expiredIn int) (err error) {
	var (
		key = genQRLoginCodeKey(keyId)

		args = []interface{}{
			key,
			"auth_key_id", qrCode.AuthKeyId,
			"session_id", qrCode.SessionId,
			"server_id", qrCode.ServerId,
			"api_id", qrCode.ApiId,
			"api_hash", qrCode.ApiHash,
			"code_hash", qrCode.CodeHash,
			"expire_at", qrCode.ExpireAt,
			"state", qrCode.State,
			"user_id", qrCode.UserId,
		}
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if err = conn.Send("HMSET", args...); err != nil {
		log.Error("conn.Send(HMSET %s,%v) error(%v)", key, args, err)
		return
	}

	if expiredIn > 0 {
		if err = conn.Send("EXPIRE", key, expiredIn+2); err != nil {
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

func (d *Redis) UpdateCacheQRLoginCode(ctx context.Context, keyId int64, values map[string]interface{}) (err error) {
	var (
		key  = genQRLoginCodeKey(keyId)
		args = []interface{}{key}
	)

	for k, v := range values {
		args = append(args, []interface{}{k, v}...)
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("HMSET", args...); err != nil {
		log.Errorf("conn.HSET(%s) error(%v)", key, err)
	}

	return
}

func (d *Redis) DeleteCacheQRLoginCode(ctx context.Context, authKeyId int64) (err error) {
	key := genQRLoginCodeKey(authKeyId)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("DEL", key); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", key, err)
	}

	return
}
