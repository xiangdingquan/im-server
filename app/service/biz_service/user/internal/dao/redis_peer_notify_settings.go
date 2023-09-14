package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/redis"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	peerNotifySettingsKeyPrefix = "peer_notify_settings"
)

func getPeerNotifySettingsKey(userId int32) string {
	return fmt.Sprintf("%s_%d", peerNotifySettingsKeyPrefix, userId)
}

func (d *Redis) GetPeerNotifySettingsByPeerList(ctx context.Context, userId int32, peerListMap map[int32][]int32) (reply map[int64]*mtproto.PeerNotifySettings, err error) {
	var (
		values [][]byte
		args   = []interface{}{getPeerNotifySettingsKey(userId)}
		idx    = 0
	)

	for k, v := range peerListMap {
		for _, id := range v {
			args = append(args, int64(k)<<32|int64(id))
		}
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	log.Debugf("get settings - %v", args)
	if values, err = redis.ByteSlices(conn.Do("HMGET", args...)); err != nil {
		log.Error("conn.Do(HMGET %v) error(%v)", args, err)
		if err == redis.ErrNil {
			return
		}
		return
	}
	log.Debugf("get settings - %d", len(values))

	reply = make(map[int64]*mtproto.PeerNotifySettings)
	for k, v := range peerListMap {
		for _, id := range v {
			if len(values[idx]) == 0 {
				reply[int64(k)<<32|int64(id)] = nil
			} else {
				settings := mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
				if err2 := json.Unmarshal(values[idx], settings); err2 != nil {
					log.Error("getPeerNotifySettings json.Unmarshal(%s) error(%v)", values[idx], err2)
					reply[int64(k)<<32|int64(id)] = nil
				} else {
					reply[int64(k)<<32|int64(id)] = settings
				}
			}
			idx++
		}
	}

	return
}

func (d *Redis) GetPeerNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil) (reply *mtproto.PeerNotifySettings, err error) {
	var (
		key = getPeerNotifySettingsKey(userId)
		b   []byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	b, err = redis.Bytes(conn.Do("HGET", key, int64(peer.PeerType)<<32|int64(peer.PeerId)))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(HGET %s %v) error(%v)", key, peer, err)
		} else {
			err = nil
		}
		return
	}

	reply = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	if err = json.Unmarshal(b, reply); err != nil {
		log.Error("getPeerNotifySettings json.Unmarshal(%s) error(%v)", b, err)
		reply = nil
	}

	return
}

func (d *Redis) SetPeerNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil, settings *mtproto.PeerNotifySettings) (err error) {
	var (
		key = getPeerNotifySettingsKey(userId)
		b   []byte
	)

	b, err = json.Marshal(settings)
	if err != nil {
		log.Errorf("setPeerNotifySettings - json.Marshal(%s) error(%v)", settings, err)
		return
	}

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("HSET", key, int64(peer.PeerType)<<32|int64(peer.PeerId), b)
	if err != nil {
		log.Error("conn.Set(HSET) error(%v)", err)
	}

	return
}

func (d *Redis) SetPeerNotifySettingsList(ctx context.Context, userId int32, settings map[int64]*mtproto.PeerNotifySettings) (err error) {
	var (
		key = getPeerNotifySettingsKey(userId)
		b   []byte
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	for k, v := range settings {
		b, err = json.Marshal(v)
		if err != nil {
			log.Errorf("SetPeerNotifySettingsList - json.Marshal(%d, %s) error(%v)", k, v, err)
			return
		}
		if err = conn.Send("HSET", key, k, b); err != nil {
			log.Error("conn.Send(HSET %s,%v) error(%v)", key, k, err)
			return
		}
	}
	return
}

func (d *Redis) DelAllPeerNotifySettings(ctx context.Context, userId int32) (err error) {
	var (
		key = getPeerNotifySettingsKey(userId)
	)

	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	if err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}

	return
}
