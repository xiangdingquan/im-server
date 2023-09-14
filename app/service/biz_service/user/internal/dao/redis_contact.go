package dao

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"fmt"

	"open.chat/model"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

const (
	contactKeyPrefix = "contact"
)

func getContactCacheKey(userId int32) string {
	return fmt.Sprintf("%s_%d", contactKeyPrefix, userId)
}

func (d *Redis) GetContactList(ctx context.Context, userId int32, contactIdList ...int32) (reply map[int32]*model.Contact, err error) {
	var (
		values [][]byte
		args   = []interface{}{getContactCacheKey(userId), versionField}
	)

	for _, id := range contactIdList {
		args = append(args, id)
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	values, err = redis.ByteSlices(conn.Do("HMGET", args...))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HMGET %s) error(%v)", args, err)
		} else {
			err = nil
		}
		return
	}

	reply2 := make(map[int32]*model.Contact, len(values))
	for i, v := range values {
		if i == 0 {
			if len(v) == 0 {
				return
			}
		} else {
			if len(v) == 0 {
				continue
			} else {
				c := &model.Contact{}
				err = json.Unmarshal(v, c)
				if err != nil {
					err = nil
					return
				}
				reply2[contactIdList[i-1]] = c
			}
		}
	}
	reply = reply2
	return
}

func (d *Redis) SetContactList(ctx context.Context, userId int32, contactList ...*model.Contact) (keyMiss bool, err error) {
	var (
		key  = getContactCacheKey(userId)
		b    []byte
		args = []interface{}{key}
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	for _, c := range contactList {
		b, err = json.Marshal(c)
		if err != nil {
			log.Errorf("SetContactList - json.Marshal(%s) error(%v)", hack.String(b), err)
			return
		}
		args = append(args, c.ContactUserId, b)
	}

	if err = conn.Send("HINCRBY", key, versionField, 1); err != nil {
		log.Error("conn.Send(HINCRBY %s,%v) error(%v)", key, versionField, err)
		return
	}

	if len(args) > 1 {
		if err = conn.Send("HMSET", args...); err != nil {
			log.Error("conn.Send(HMSET %s,%v) error(%v)", key, args, err)
			return
		}
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}

	var version int64
	if version, err = redis.Int64(conn.Receive()); err != nil {
		log.Error("conn.Receive() error(%v)", err)
		return
	}
	if version == 1 {
		keyMiss = true
	}

	if len(args) > 1 {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}
