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
	duplicateBlogId   = "duplicate_blog_id"
	duplicateBlogData = "duplicate_blog_data"
	expireTimeout     = 60
)

func makeDuplicateBlogKey(prefix string, senderUserId int32, clientRandomId int64) string {
	return fmt.Sprintf("%s_%d_%d", prefix, senderUserId, clientRandomId)
}

func (d *Redis) HasDuplicateBlog(ctx context.Context, senderUserId int32, clientRandomId int64) (bool, error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	k := makeDuplicateBlogKey(duplicateBlogId, senderUserId, clientRandomId)
	seq, err := redis.Int64(conn.Do("INCR", k))
	if err != nil {
		log.Errorf("checkDuplicateBlog - INCR {%s}, error: {%v}", k, err)
		return false, err
	}

	if _, err = conn.Do("EXPIRE", k, expireTimeout); err != nil {
		log.Errorf("expire DuplicateBlog - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
		return false, err
	}

	return seq > 1, nil
}

func (d *Redis) PutDuplicateBlog(ctx context.Context, senderUserId int32, clientRandomId int64, blog *mtproto.MicroBlog) error {
	k := makeDuplicateBlogKey(duplicateBlogData, senderUserId, clientRandomId)
	cacheData, _ := proto.Marshal(blog)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err := conn.Do("SET", k, cacheData, "EX", expireTimeout); err != nil {
		log.Errorf("putDuplicateBlog - SET {%s, %s, %d}, error: %s", k, cacheData, expireTimeout, err)
		return err
	}

	return nil
}

func (d *Redis) GetDuplicateBlog(ctx context.Context, senderUserId int32, clientRandomId int64) (*mtproto.MicroBlog, error) {
	k := makeDuplicateBlogKey(duplicateBlogData, senderUserId, clientRandomId)

	conn := d.redis.Redis.Get(ctx)
	var blog *mtproto.MicroBlog

	defer conn.Close()
	if cacheData, err := redis.Bytes(conn.Do("GET", k)); err != nil {
		if err.Error() == "redigo: nil returned" {
			return nil, nil
		}

		log.Errorf("getDuplicateBlog - GET {%s}, error: %s", k, err)
		return nil, err
	} else {
		blog = &mtproto.MicroBlog{}
		proto.Unmarshal(cacheData, blog)
	}

	return blog, nil
}
