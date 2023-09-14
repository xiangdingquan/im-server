package dao

import (
	"context"
	"fmt"
	"io"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/app/pkg/redis_util"
	"open.chat/pkg/log"
)

type SSDBReader struct {
	ownerId int64
	fileId  int64
	parts   int32
	buf     []byte
	partIdx int32
	redis   *redis_util.Redis
}

func NewSSDBReader(r *redis_util.Redis, ownerId, fileId int64, parts int32) *SSDBReader {
	return &SSDBReader{
		ownerId: ownerId,
		fileId:  fileId,
		parts:   parts,
		partIdx: 0,
		redis:   r,
	}
}

func (r *SSDBReader) Read(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		return 0, fmt.Errorf("len(p) == 0")
	}

	if len(r.buf) == 0 && r.partIdx >= r.parts {
		return 0, io.EOF
	}
	if len(r.buf) >= n {
		copy(p, r.buf[:n])
		r.buf = r.buf[n:]
	} else {
		l := len(r.buf)
		copy(p, r.buf)
		r.buf = r.buf[:0]

		var b []byte
		for i := r.partIdx; i < r.parts; i++ {
			b, err = readFile(context.Background(), r.redis, r.ownerId, r.fileId, r.partIdx)
			if err != nil {
				return 0, err
			}
			r.partIdx += 1
			if len(b) >= n-l {
				copy(p[l:], b[:n-l])
				r.buf = b[n-l:]
				return n, nil
			} else {
				copy(p[l:], b)
				l += len(b)
			}
		}

		n = l
	}
	return
}

func readFile(ctx context.Context, r *redis_util.Redis, ownerId, fileId int64, filePart int32) ([]byte, error) {
	conn := r.Redis.Get(ctx)
	defer conn.Close()

	var (
		err error
		k   = getFileKey(ownerId, fileId)
	)

	var b []byte
	b, err = redis.Bytes(conn.Do("HGET", k, filePart))
	if err != nil {
		log.Error("conn.Send(HGET %s, %d) error(%v)", k, filePart, err)
		return nil, err
	}

	return b, nil
}
