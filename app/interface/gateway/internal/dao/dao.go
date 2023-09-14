package dao

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/app/service/auth_session/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/cache"
	"open.chat/pkg/log"
)

type Dao struct {
	cache  *cache.LRUCache
	client authsessionpb.RPCSessionClient
}

func New(cap int64, c *warden.ClientConfig) *Dao {
	cli, err := authsession_client.New(c)
	if err != nil {
		panic(err)
	}

	return &Dao{
		client: cli,
		cache:  cache.NewLRUCache(cap),
	}
}

func (d *Dao) GetAuthKey(ctx context.Context, authKeyId int64) (*model.AuthKeyData, error) {
	var (
		cacheK = strconv.Itoa(int(authKeyId))
		value  *model.AuthKeyData
	)

	if v, ok := d.cache.Get(cacheK); !ok {
		keyInfo, err := d.client.SessionQueryAuthKey(ctx, &authsessionpb.TLSessionQueryAuthKey{AuthKeyId: authKeyId})
		if err != nil {
			log.Error("queryAuthKey error: auth_key_id:%d, %s", authKeyId, err.Error())
			return nil, err
		}

		value = &model.AuthKeyData{
			AuthKeyId:          authKeyId,
			AuthKey:            keyInfo.AuthKey,
			AuthKeyType:        int(keyInfo.AuthKeyType),
			PermAuthKeyId:      keyInfo.PermAuthKeyId,
			TempAuthKeyId:      keyInfo.TempAuthKeyId,
			MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
		}
		d.cache.Set(cacheK, value)
	} else {
		value = v.(*model.AuthKeyData)
	}

	return value, nil
}

func (d *Dao) PutAuthKey(ctx context.Context, keyInfo *authsessionpb.AuthKeyInfo, salt *mtproto.FutureSalt) error {
	r, err := d.client.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey:    keyInfo,
		FutureSalt: salt,
	})
	if err != nil || !mtproto.FromBool(r) {
		log.Errorf("saveAuthKeyInfo error: auth_key_id:%d, err: %v", keyInfo.AuthKeyId, err)
		return err
	}

	var (
		cacheK = strconv.Itoa(int(keyInfo.AuthKeyId))
	)
	d.cache.Set(cacheK, &model.AuthKeyData{
		AuthKeyId:          keyInfo.AuthKeyId,
		AuthKey:            keyInfo.AuthKey,
		AuthKeyType:        int(keyInfo.AuthKeyType),
		PermAuthKeyId:      keyInfo.PermAuthKeyId,
		TempAuthKeyId:      keyInfo.TempAuthKeyId,
		MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId})
	return nil
}
