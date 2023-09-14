package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	privacyKeyPrefix = "privacy"
)

func getPrivacyCacheKey(userId int32) string {
	return fmt.Sprintf("%s_%d", privacyKeyPrefix, userId)
}

func (d *Redis) GetPrivacy(ctx context.Context, userId int32, keyType int) (reply []*mtproto.PrivacyRule, err error) {
	var (
		key = getPrivacyCacheKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	b, err := redis.Bytes(conn.Do("HGET", key, keyType))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGET %s %v) error(%v)", key, keyType, err)
		} else {
			err = nil
		}
		return
	}

	reply = make([]*mtproto.PrivacyRule, 0)
	if err = json.Unmarshal(b, &reply); err != nil {
		log.Error("getPrivacy json.Unmarshal(%s) error(%v)", b, err)
	}

	return
}

func (d *Redis) SetPrivacy(ctx context.Context, userId int32, keyType int, rules []*mtproto.PrivacyRule) (err error) {
	var (
		key = getPrivacyCacheKey(userId)
		b   []byte
	)

	b, err = json.Marshal(rules)
	if err != nil {
		log.Errorf("setPeerNotifySettings - json.Marshal(%s) error(%v)", b, err)
		return
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("HSET", key, keyType, b)
	if err != nil {
		log.Error("conn.Set(HSET) error(%v)", err)
	}

	return
}

func (d *Redis) SetPrivacyList(ctx context.Context, userId int32, privacyList map[int][]*mtproto.PrivacyRule) (err error) {
	var (
		key = getPrivacyCacheKey(userId)
		b   []byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	for k, v := range privacyList {
		b, err = json.Marshal(v)
		if err != nil {
			log.Errorf("SetPrivacyList - json.Marshal(%d, %s) error(%v)", k, v, err)
			return
		}
		if err = conn.Send("HSET", key, k, b); err != nil {
			log.Error("conn.Send(HSET %s,%v) error(%v)", key, k, err)
			return
		}
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < len(privacyList); i++ {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}
