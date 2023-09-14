package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/app/pkg/redis_util"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	phoneCallTimeout     int64 = 8 * 60 * 60
	cachePhoneCallPrefix       = "phone_call"
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

func genCachePhoneCallKey(callId int64) string {
	return fmt.Sprintf("%s_%d", cachePhoneCallPrefix, callId)
}

func (d *Redis) CreatePhoneCallSession(ctx context.Context,
	video bool,
	adminId int32,
	adminAuthKeyId int64,
	participantId int32,
	randomId int64,
	gAHash []byte,
	protocol *mtproto.PhoneCallProtocol) (*model.PhoneCallSession, error) {

	callSession := &model.PhoneCallSession{
		Video:          video,
		Id:             idgen.GetUUID(),
		AccessHash:     rand.Int63(),
		AdminId:        adminId,
		AdminAuthKeyId: adminAuthKeyId,
		RandomId:       randomId,
		ParticipantId:  participantId,
		AdminProtocol:  protocol,
		GAHash:         gAHash,
		State:          model.CallStateRequested,
		Date:           time.Now().Unix(),
	}

	if err := d.PutPhoneCallSession(ctx, callSession.Id, callSession); err != nil {
		return nil, err
	}

	return callSession, nil
}

func (d *Redis) GetPhoneCallSession(ctx context.Context, callSessionId int64) (callSession *model.PhoneCallSession, err error) {
	cacheKey := genCachePhoneCallKey(callSessionId)

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
		callSession := new(model.PhoneCallSession)
		err = json.Unmarshal(v, callSession)
		return callSession, err
	}
}

func (d *Redis) PutPhoneCallSession(ctx context.Context, callSessionId int64, callSession *model.PhoneCallSession) error {
	cacheKey := genCachePhoneCallKey(callSessionId)

	b, _ := json.Marshal(callSession)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err := conn.Do("SETEX", cacheKey, phoneCallTimeout, b); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
		return err
	}

	return nil
}

func (d *Redis) DeletePhoneCallSession(ctx context.Context, callSessionId int64) error {
	cacheKey := genCachePhoneCallKey(callSessionId)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err := conn.Do("DEL", cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
		return err
	}

	return nil
}
