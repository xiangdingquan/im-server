package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/gogo/protobuf/proto"
	"open.chat/app/pkg/redis_util"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
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

const (
	duplicateMessageId   = "duplicate_message_id"
	duplicateMessageData = "duplicate_message_data"
	expireTimeout        = 60 // 60s
)

func makeDuplicateMessageKey(prefix string, senderUserId int32, clientRandomId int64) string {
	return fmt.Sprintf("%s_%d_%d", prefix, senderUserId, clientRandomId)
}

func (d *Redis) HasDuplicateMessage(ctx context.Context, senderUserId int32, clientRandomId int64) (bool, error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	k := makeDuplicateMessageKey(duplicateMessageId, senderUserId, clientRandomId)
	seq, err := redis.Int64(conn.Do("INCR", k))
	if err != nil {
		log.Errorf("checkDuplicateMessage - INCR {%s}, error: {%v}", k, err)
		return false, err
	}

	if _, err = conn.Do("EXPIRE", k, expireTimeout); err != nil {
		log.Errorf("expire DuplicateMessage - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
		return false, err
	}

	return seq > 1, nil
}

func (d *Redis) PutDuplicateMessage(ctx context.Context, senderUserId int32, clientRandomId int64, upd *mtproto.Updates) error {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	cacheData, _ := proto.Marshal(upd)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err := conn.Do("SET", k, cacheData, "EX", expireTimeout); err != nil {
		log.Errorf("putDuplicateMessage - SET {%s, %s, %d}, error: %s", k, cacheData, expireTimeout, err)
		return err
	}
	return nil
}

func (d *Redis) GetDuplicateMessage(ctx context.Context, senderUserId int32, clientRandomId int64) (*mtproto.Updates, error) {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	conn := d.redis.Redis.Get(ctx)
	var upd *mtproto.Updates

	defer conn.Close()
	if cacheData, err := redis.Bytes(conn.Do("GET", k)); err != nil {
		if err.Error() == "redigo: nil returned" {
			return nil, nil
		}

		log.Errorf("getDuplicateMessage - GET {%s}, error: %s", k, err)
		return nil, err
	} else {
		upd = &mtproto.Updates{}
		proto.Unmarshal(cacheData, upd)
	}

	return upd, nil
}

const (
	setName  string = "setCountDownMessage"
	ssetName string = "ssetCountDownMessage"
)

func (m *Redis) AddCountDownMsg(ctx context.Context, uId, peerType, peerId, msgId int32, startime int64) (success bool) {
	var (
		err error
		n   = 2
	)
	msg := &model.CountDownMessage{
		UId:      uId,
		PeerType: peerType,
		PeerId:   peerId,
		MsgId:    msgId,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Error("json.Marshal(%s) error(%v)", msg, err)
		return
	}

	member_name := fmt.Sprintf("count_down_msg_%d%d%d%d", msg.UId, msg.PeerType, msg.PeerId, msg.MsgId)

	conn := m.redis.Redis.Get(ctx)
	defer conn.Close()

	if err = conn.Send("ZADD", ssetName, startime, member_name); err != nil {
		log.Error("conn.Send(ZADD %s %d %s) error(%v)", ssetName, startime, member_name, err)
		return false
	}

	if err = conn.Send("HSET", setName, member_name, string(b)); err != nil {
		log.Error("conn.Send(HGETALL %s %s) error(%v)", setName, member_name, err)
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
	return true
}

func (m *Redis) GetCountDownMsg(ctx context.Context, nowtime int64) (messages []*model.CountDownMessage) {
	messages = make([]*model.CountDownMessage, 0)
	conn := m.redis.Redis.Get(ctx)
	defer conn.Close()

	var (
		err          error
		member_names []string
	)

	if member_names, err = redis.Strings(conn.Do("ZRANGEBYSCORE", ssetName, 0, nowtime)); err != nil {
		log.Error("conn.Do() error(%v)", err)
		return
	}

	if len(member_names) == 0 {
		return
	}

	args := []interface{}{setName}
	for _, v := range member_names {
		args = append(args, v)
	}

	var msgDatas [][]byte
	if msgDatas, err = redis.ByteSlices(conn.Do("HMGET", args...)); err != nil {
		log.Error("conn.Do() error(%v)", err)
		return
	}

	for _, data := range msgDatas {
		recordMsg := &model.CountDownMessage{}
		err = json.Unmarshal(data, recordMsg)
		if err != nil {
			log.Error("json.Marshal(%s) error(%v)", string(data), err)
			return
		}
		messages = append(messages, recordMsg)
	}
	return
}

func (m *Redis) DelCountDownMsgs(ctx context.Context, messages []*model.CountDownMessage) (success bool) {
	conn := m.redis.Redis.Get(ctx)
	defer conn.Close()

	var (
		n            = 2
		err          error
		args         []interface{}
		member_names []string = make([]string, len(messages))
	)

	for _, msg := range messages {
		args = append(args, fmt.Sprintf("count_down_msg_%d%d%d%d", msg.UId, msg.PeerType, msg.PeerId, msg.MsgId))
	}

	if err = conn.Send("HDEL", append([]interface{}{setName}, args...)...); err != nil {
		log.Error("conn.Send(HDEL %s %v) error(%v)", ssetName, member_names, err)
		return
	}

	if err = conn.Send("ZREM", append([]interface{}{ssetName}, args...)...); err != nil {
		log.Error("conn.Send(ZREM %s %v) error(%v)", ssetName, member_names, err)
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
	return true
}
