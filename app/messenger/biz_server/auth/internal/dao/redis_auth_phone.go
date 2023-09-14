package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	phoneNumberTimeout     = 300 // salt timeout
	cachePhoneNumberPrefix = "phone_codes"
)

func genCachePhoneNumberKey(authKeyId int64) string {
	return fmt.Sprintf("%s_%d", cachePhoneNumberPrefix, authKeyId)
}

func (d *Redis) GetCachePhoneNumber(ctx context.Context, authKeyId int64) (string, error) {
	cacheKey := genCachePhoneNumberKey(authKeyId)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	PhoneNumber := ""
	v, err := redis.Bytes(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
	} else {
		log.Debugf(hack.String(v))
		PhoneNumber = string(v)
	}
	return PhoneNumber, err
}

func (d *Redis) PutCachePhoneNumber(ctx context.Context, authKeyId int64, phoneNumber string) (err error) {
	cacheKey := genCachePhoneNumberKey(authKeyId)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETEX", cacheKey, phoneNumberTimeout, phoneNumber); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Redis) DeleteCachePhoneNumber(ctx context.Context, authKeyId int64) (err error) {
	cacheKey := genCachePhoneNumberKey(authKeyId)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("DEL", cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}
	return
}
