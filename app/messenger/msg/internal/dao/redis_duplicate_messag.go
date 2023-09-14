package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/gogo/protobuf/proto"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	duplicateMessageId   = "duplicate_message_id"
	duplicateMessageData = "duplicate_message_data"
	expireTimeout        = 60
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
