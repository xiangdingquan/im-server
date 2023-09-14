package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	xtime "github.com/go-kratos/kratos/pkg/time"

	status_facade "open.chat/app/service/status/facade"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

const (
	onlineKeyPrefix  = "online"    //
	userKeyIdsPrefix = "user_keys" //
)

// ////////////////////////////////////////////////////////////////////////
type redisStatusClient struct {
	redis       *redis.Pool
	redisExpire int32
}

func New() status_facade.StatusFacade {
	var (
		rc struct {
			Status       *redis.Config
			StatusExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("status.toml").UnmarshalTOML(&rc))

	return &redisStatusClient{
		redis:       redis.NewPool(rc.Status),
		redisExpire: int32(time.Duration(rc.StatusExpire) / time.Second),
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getUserKey(id int32) string {
	return fmt.Sprintf("%s_%d", userKeyIdsPrefix, id)
}

func getAuthKeyIdKey(id int64) string {
	return fmt.Sprintf("%s_%d", onlineKeyPrefix, id)
}

func (c *redisStatusClient) AddOnline(ctx context.Context, userId int32, authKeyId int64, serverId string) (err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	var n = 2
	if err = conn.Send("HSET", getUserKey(userId), authKeyId, serverId); err != nil {
		log.Error("conn.Send(HSET %d, %d, %s) error(%v)", userId, authKeyId, serverId, err)
		return
	}
	if err = conn.Send("EXPIRE", getUserKey(userId), c.redisExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%d,%s) error(%v)", userId, authKeyId, serverId, err)
		return
	}

	n += 2
	if err = conn.Send("SET", getAuthKeyIdKey(authKeyId), serverId); err != nil {
		log.Error("conn.Send(SET %d,%d,%s) error(%v)", userId, authKeyId, serverId, err)
		return
	}
	if err = conn.Send("EXPIRE", getAuthKeyIdKey(authKeyId), c.redisExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%d,%s) error(%v)", userId, authKeyId, serverId, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

func (c *redisStatusClient) DelOnline(ctx context.Context, userId int32, authKeyId int64) (err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	n := 1
	if err = conn.Send("HDEL", getUserKey(userId), authKeyId); err != nil {
		log.Error("conn.Send(HDEL %d,%d) error(%v)", userId, authKeyId, err)
		return
	}
	n++
	if err = conn.Send("DEL", getAuthKeyIdKey(authKeyId)); err != nil {
		log.Error("conn.Send(HDEL %d,%d) error(%v)", userId, authKeyId, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if _, err = redis.Bool(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}

	return
}

func (c *redisStatusClient) ExpireOnline(ctx context.Context, userId int32, authKeyId int64) (has bool, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	var n = 1
	if err = conn.Send("EXPIRE", getUserKey(userId), c.redisExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%s) error(%v)", userId, authKeyId, err)
		return
	}
	n++
	if err = conn.Send("EXPIRE", getAuthKeyIdKey(authKeyId), c.redisExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%s) error(%v)", userId, authKeyId, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if has, err = redis.Bool(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

func (c *redisStatusClient) GetOnlineListByKeyIdList(ctx context.Context, authKeyIds []int64) (res []string, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	var args []interface{}
	for _, key := range authKeyIds {
		args = append(args, getAuthKeyIdKey(key))
	}
	if res, err = redis.Strings(conn.Do("MGET", args...)); err != nil {
		log.Error("conn.Do(MGET %v) error(%v)", args, err)
	}
	return
}

func (c *redisStatusClient) GetOnlineByKeyId(ctx context.Context, authKeyId int64) (res string, err error) {
	var ress []string

	ress, err = c.GetOnlineListByKeyIdList(ctx, []int64{authKeyId})
	if len(ress) == 1 {
		res = ress[0]
	}

	return
}

func (c *redisStatusClient) GetOnlineListExcludeKeyId(ctx context.Context, userId int32, authKeyId int64) (res map[int64]string, err error) {
	res, _, err = c.GetOnlineMapByUserList(ctx, []int32{userId})

	if err != nil {
		delete(res, authKeyId)
	}

	return
}

func (c *redisStatusClient) GetOnlineListByUser(ctx context.Context, userId int32) (res map[int64]string, err error) {
	res, _, err = c.GetOnlineMapByUserList(ctx, []int32{userId})
	return
}

func (c *redisStatusClient) GetOnlineMapByUserList(ctx context.Context, userIdList []int32) (ress map[int64]string, onUserList []int32, err error) {
	conn := c.redis.Get(ctx)
	defer conn.Close()

	ress = make(map[int64]string)
	for _, userId := range userIdList {
		if err = conn.Send("HGETALL", getUserKey(userId)); err != nil {
			log.Error("conn.Do(HGETALL %d) error(%v)", userId, err)
			return
		}
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for idx := 0; idx < len(userIdList); idx++ {
		var (
			res map[string]string
		)
		if res, err = redis.StringMap(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
		if len(res) > 0 {
			onUserList = append(onUserList, userIdList[idx])
		}
		for k, v := range res {
			authKeyId, _ := util.StringToInt64(k)
			ress[authKeyId] = v
		}
	}
	return
}

func init() {
	status_facade.Register("redis", New)
}
