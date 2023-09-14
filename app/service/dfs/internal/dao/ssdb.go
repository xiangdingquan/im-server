package dao

import (
	"context"
	"fmt"
	"io"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/pkg/log"
)

const (
	_fileKeyPrefix = "file_%d_%d"
)

func getFileKey(ownerId, fileId int64) string {
	return fmt.Sprintf(_fileKeyPrefix, ownerId, fileId)
}

func (d *Dao) WriteFilePartData(ctx context.Context, ownerId, fileId int64, filePart int32, bytes []byte) (err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	k := getFileKey(ownerId, fileId)

	var n = 2
	if err = conn.Send("HSET", k, filePart, bytes); err != nil {
		log.Error("conn.Send(HSET %d, %d, %s) error(%v)", ownerId, fileId, filePart, err)
		return
	}
	log.Debugf("conn.Send(HSET %s, %d)", k, filePart)
	if err = conn.Send("EXPIRE", k, d.redis.RedisExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%d,%s) error(%v)", ownerId, fileId, filePart, err)
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

func (d *Dao) GetFileParts(ctx context.Context, ownerId, fileId int64) (fileParts int32, err error) {
	var (
		k    = getFileKey(ownerId, fileId)
		hLen int
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if hLen, err = redis.Int(conn.Do("HLEN", k)); err == nil {
		fileParts = int32(hLen)
	}

	return
}

func (d *Dao) ReadFile(ctx context.Context, ownerId, fileId int64, parts int32) (partLength int32, bytes []byte, err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	k := getFileKey(ownerId, fileId)

	for i := int32(0); i < parts; i++ {
		var b []byte
		b, err = redis.Bytes(conn.Do("HGET", k, i))
		if err != nil {
			log.Error("conn.Send(HGET %s, %d) error(%v)", k, i, err)
			return 0, nil, err
		}
		if bytes == nil {
			bytes = make([]byte, 0, len(b)*int(parts))
		}
		if i == 0 {
			partLength = int32(len(b))
		}
		bytes = append(bytes, b...)
	}

	return
}

func (d *Dao) ReadFileCB(ctx context.Context, ownerId, fileId int64, parts int32, cb func(part int32, bytes []byte) error) (err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	k := getFileKey(ownerId, fileId)

	for i := int32(0); i < parts; i++ {
		var b []byte
		b, err = redis.Bytes(conn.Do("HGET", k, i))
		if err != nil {
			log.Error("conn.Do(HGET %s, %d) error(%v)", k, i, err)
			return
		}
		log.Debugf("conn.Do(HGET %s, %d), len: %d", k, i, len(b))
		if err = cb(i, b); err != nil {
			return
		}
	}

	return
}

func (d *Dao) ReadOffsetLimit(ctx context.Context, ownerId, fileId int64, offset, limit int32) (bytes []byte, err error) {
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	//k := getFileKey(ownerId, fileId)
	//
	//for i := int32(0); i < parts; i++ {
	//	var b []byte
	//	b, err = redis.Bytes(conn.Do("HGET", k, i))
	//	if err != nil {
	//		log.Error("conn.Send(HGET %s, %d) error(%v)", k, i, err)
	//		return nil, err
	//	}
	//	if bytes == nil {
	//		bytes = make([]byte, 0, len(b)*int(parts))
	//	}
	//	bytes = append(bytes, b...)
	//}

	return
}

func (d *Dao) OpenFile(ctx context.Context, ownerId, fileId int64, parts int32) (r io.Reader, err error) {
	return NewSSDBReader(d.redis, ownerId, fileId, parts), nil
}
