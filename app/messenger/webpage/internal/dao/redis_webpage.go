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
	webPagePrefix = "webpage_"
)

func genWebPageCodeKey(url string) string {
	return fmt.Sprintf("%s_%s", webPagePrefix, url)
}

func (d *Redis) GetCacheWebPage(ctx context.Context, url string) (webPage *mtproto.WebPage, err error) {
	var (
		key       = genWebPageCodeKey(url)
		cacheData []byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	cacheData, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", key, err)
		} else {
			err = nil
		}
		return
	}

	webPage = new(mtproto.WebPage)
	if err = json.Unmarshal(cacheData, webPage); err != nil {
		log.Errorf("json.Unmarshal error: %v", err)
	}

	return
}

func (d *Redis) PutCacheWebPage(ctx context.Context, url string, webPage *mtproto.WebPage, expiredIn int) (err error) {
	var (
		key       = genWebPageCodeKey(url)
		cacheData []byte
	)

	if cacheData, err = json.Marshal(webPage); err != nil {
		log.Errorf("json.Marshal error: %v", err)
		return
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("SETEX", key, expiredIn, cacheData); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", key, err)
	}

	return
}

func (d *Redis) DeleteCacheWebPage(ctx context.Context, url string) (err error) {
	key := genWebPageCodeKey(url)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	if _, err = conn.Do("DEL", key); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", key, err)
	}

	return
}
