package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	phoneCodeTimeout = 90 // salt timeout
	cachePhonePrefix = "phone_codes"
)

func genCachePhoneCodeKey(authKeyId int64, phoneNumber string) string {
	return fmt.Sprintf("%s_%d_%s", cachePhonePrefix, authKeyId, phoneNumber)
}

func (d *Redis) GetCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (*model.PhoneCodeTransaction, error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)
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
		codeData := &model.PhoneCodeTransaction{}
		err = json.Unmarshal(v, codeData)
		log.Debugf("codeData: %v", codeData)
		return codeData, err
	}
}

func (d *Redis) PutCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string, codeData *model.PhoneCodeTransaction) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)
	b, _ := json.Marshal(codeData)

	//	b, err = json.Marshal(salts)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("SETEX", cacheKey, phoneCodeTimeout, b); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Redis) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("DEL", cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}

	return
}
