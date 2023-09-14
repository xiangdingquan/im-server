package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"open.chat/model"
	"open.chat/pkg/log"
	"strconv"
)

const (
	locationKey                 = "location_for_nearby"
	locationExpirationZSetKey   = "zset_location_expires_for_nearby"
	locationExpirationHashesKey = "hashes_location_expires_for_nearby"
)

func (d *Redis) GeoAdd(ctx context.Context, lat, long float64, uid int32) (bool, error) {
	log.Debugf("GeoAdd, lat:%f, long:%f, uid:%d", lat, long, uid)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	member := strconv.Itoa(int(uid))
	seq, err := redis.Int64(conn.Do("GEOADD", locationKey, long, lat, member))
	if err != nil {
		log.Errorf("geoadd of user(%d), error: {%v}", uid, err)
		return false, err
	}

	return seq == 1, nil
}

func (d *Redis) GeoDel(ctx context.Context, uid int32) (bool, error) {
	log.Debugf("GeoDel, uid:%d", uid)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	src := `
		redis.call('ZREM', KEYS[1], ARGV[1]);
		redis.call('ZREM', KEYS[2], ARGV[1]);
		redis.call('HDEL', KEYS[3], ARGV[1]);
		return 1;
	`
	script := redis.NewScript(3, src)
	keyAndArgs := []string{
		locationKey, locationExpirationZSetKey, locationExpirationHashesKey,
		strconv.Itoa(int(uid)),
	}

	seq, err := redis.Int64(script.Do(conn, keyAndArgs))
	if err != nil {
		log.Errorf("del user(%d) from nearby location, error: {%v}", uid, err)
		return false, err
	}

	return seq == 1, nil
}

func (d *Redis) GeoRadius(ctx context.Context, lat, long float64, radius int32, limit int32) ([]*model.NearByUser, error) {
	log.Debugf("GeoRadius, lat:%f, long:%f, radius:%d, limit:%d", lat, long, radius, limit)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	//count := fmt.Sprintf("Count %d ANY", limit)
	//values, err := redis.Values(conn.Do("GEORADIUS", locationKey, long, lat, radius, "m", "WITHDIST", "COUNT", limit, "ANY"))
	values, err := redis.Values(conn.Do("GEORADIUS", locationKey, long, lat, radius, "m", "WITHDIST", "COUNT", limit))
	if err != nil {
		log.Errorf("GeoRadius, get value from redis failed, lat:%f, long:%f, radius:%d, limit:%d, error:%v", lat, long, radius, limit, err)
		return nil, err
	}

	out := make([]*model.NearByUser, 0, len(values))
	for _, v := range values {
		infos, err := redis.Values(v, nil)
		if err != nil {
			log.Errorf("GeoRadius, parse near by information failed, v:%v, err:%v", v, err)
			continue
		}
		if len(infos) != 2 {
			log.Errorf("GeoRadius, invalid len of near by information")
			continue
		}
		idString, err := redis.String(infos[0], nil)
		if err != nil {
			log.Errorf("GeoRadius, get id from near by information failed, v:%v, err:%v", infos[0], err)
			continue
		}
		distString, err := redis.String(infos[1], nil)

		n, err := model.NearByUserFromString(idString, distString)
		if err != nil {
			log.Errorf("GeoRadius, new near by user failed, err:%v", err)
			continue
		}
		out = append(out, n)
	}

	return out, nil
}

func (d *Redis) SetGeoExpiration(ctx context.Context, uid int32, expiration int32) (bool, error) {
	log.Debugf("SetGeoExpiration, uid:%d, expiration:%d", uid, expiration)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	src := `
		redis.call('ZADD', KEYS[1], ARGV[1], ARGV[2]);
		redis.call('HSET', KEYS[2], ARGV[2], ARGV[1]);
		return 1;
	`
	script := redis.NewScript(2, src)
	keyAndArgs := []string{
		locationExpirationZSetKey, locationExpirationHashesKey,
		strconv.Itoa(int(expiration)), strconv.Itoa(int(uid)),
	}
	seq, err := redis.Int64(script.Do(conn, keyAndArgs))
	if err != nil {
		log.Errorf("SetGeoExpiration uid:%d expires:%d, error: {%v}", uid, expiration, err)
		return false, err
	}

	return seq == 1, nil
}

func (d *Redis) DelGeoExpiration(ctx context.Context, uid int32) (bool, error) {
	log.Debugf("DelGeoExpiration, uid:%d", uid)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	src := `
		redis.call('ZREM', KEYS[1], ARGV[1]);
		redis.call('HDEL', KEYS[2], ARGV[1]);
		return 1;
	`
	script := redis.NewScript(2, src)
	keyAndArgs := []string{
		locationExpirationZSetKey, locationExpirationHashesKey,
		strconv.Itoa(int(uid)),
	}
	seq, err := redis.Int64(script.Do(conn, keyAndArgs))
	if err != nil {
		log.Errorf("DelGeoExpiration uid:%d error: {%v}", uid, err)
		return false, err
	}

	return seq == 1, nil
}

func (d *Redis) GetExpiredLocation(ctx context.Context, expireTime int32) ([]int32, error) {
	log.Debugf("GetExpiredLocation, expireTime:%d", expireTime)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	members, err := redis.ByteSlices(conn.Do("ZRANGE", locationExpirationZSetKey, 0, expireTime, "BYSCORE"))
	if err != nil {
		log.Errorf("GetExpiredLocation, expireTime:{%d}, error: %v", expireTime, err)
		return nil, err
	}

	ids := make([]int32, len(members))
	for _, member := range members {
		u, err := strconv.Atoi(string(member))
		if err != nil {
			log.Errorf("GetExpiredLocation, convert member(%s) failed, error: %v", member, err)
			continue
		}
		ids = append(ids, int32(u))
	}

	return ids, nil
}

func (d *Redis) GetLocationExpirations(ctx context.Context, uidList []int32) (map[int32]int32, error) {
	log.Debugf("GetLocationExpirations, uidList.count:%d", len(uidList))

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	args := []interface{}{locationExpirationHashesKey}
	for _, u := range uidList {
		args = append(args, strconv.Itoa(int(u)))
	}
	l, err := redis.Values(conn.Do("HMGET", args...))
	if err != nil {
		log.Errorf("GetLocationExpiration, call redis failed, uid_count:%d, error:%v", len(uidList), err)
		return nil, err
	}

	m := make(map[int32]int32, len(l))
	for i, v := range l {
		s, err := redis.String(v, nil)
		if err != nil {
			log.Errorf("GetLocationExpiration, error:%v", err)
			continue
		}
		expiration, err := strconv.Atoi(s)
		if err != nil {
			log.Errorf("GetLocationExpiration, invalid expiration, s:%s, error:%v", s, err)
			continue
		}
		m[uidList[i]] = int32(expiration)
	}

	return m, nil
}
