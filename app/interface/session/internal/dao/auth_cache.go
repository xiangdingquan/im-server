package dao

import (
	"context"
	"fmt"
	"time"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

type cacheAuthValue struct {
	UserId        int32
	Layer         int32
	pushSessionId int64
	client        string
	ipAddress     string
	SaltList      []*mtproto.TLFutureSalt
}

func (cv *cacheAuthValue) Size() int {
	return 1
}

func (d *Dao) getCacheValue(authKeyId int64) *cacheAuthValue {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := d.cache.Get(cacheK); !ok {
		cv := new(cacheAuthValue)
		d.cache.Set(cacheK, cv)
		return cv
	} else {
		return v.(*cacheAuthValue)
	}
}

func (d *Dao) GetCacheUserID(ctx context.Context, authKeyId int64) (int32, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.UserId == 0 {
		id, err := d.AuthSessionRpcClient.SessionGetUserId(ctx, &authsessionpb.TLSessionGetUserId{AuthKeyId: authKeyId})
		if err != nil {
			log.Error(err.Error())
			return 0, false
		}
		cv.UserId = id.GetV()
	}
	return cv.UserId, true
}

func (d *Dao) GetCachePushSessionID(ctx context.Context, userId int32, authKeyId int64) (int64, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.pushSessionId == 0 {
		id, err := d.AuthSessionRpcClient.SessionGetPushSessionId(ctx, &authsessionpb.TLSessionGetPushSessionId{
			UserId:    userId,
			AuthKeyId: authKeyId,
			TokenType: 7,
		})
		if err != nil {
			log.Error(err.Error())
			return 0, false
		}
		cv.pushSessionId = id.GetV()
	}

	return cv.pushSessionId, true
}

func (d *Dao) GetCacheApiLayer(ctx context.Context, authKeyId int64) (int32, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.Layer == 0 {
		id, err := d.AuthSessionRpcClient.SessionGetLayer(ctx, &authsessionpb.TLSessionGetLayer{AuthKeyId: authKeyId})
		if err != nil {
			log.Error(err.Error())
			return 0, false
		}
		cv.Layer = id.GetV()
	}

	return cv.Layer, true
}

func (d *Dao) GetCacheClient(ctx context.Context, authKeyId int64) string {
	cv := d.getCacheValue(authKeyId)
	return cv.client
}

func (d *Dao) GetCacheIpAddress(ctx context.Context, authKeyId int64) string {
	cv := d.getCacheValue(authKeyId)
	return cv.ipAddress
}

func (d *Dao) PutCacheApiLayer(ctx context.Context, authKeyId int64, layer int32) {
	cv := d.getCacheValue(authKeyId)
	cv.Layer = layer
}

func (d *Dao) PutCacheClient(ctx context.Context, authKeyId int64, clint string) {
	cv := d.getCacheValue(authKeyId)
	cv.client = clint
}

func (d *Dao) PutCacheIpAddress(ctx context.Context, authKeyId int64, ip string) {
	cv := d.getCacheValue(authKeyId)
	cv.ipAddress = ip
}

func (d *Dao) PutCacheUserId(ctx context.Context, authKeyId int64, userId int32) {
	cv := d.getCacheValue(authKeyId)
	cv.UserId = userId
}

func (d *Dao) PutCachePushSessionId(ctx context.Context, authKeyId, sessionId int64) {
	cv := d.getCacheValue(authKeyId)
	cv.pushSessionId = sessionId
}

func (d *Dao) getFutureSaltList(ctx context.Context, authKeyId int64) ([]*mtproto.TLFutureSalt, bool) {
	var (
		cv   = d.getCacheValue(authKeyId)
		date = int32(time.Now().Unix())
	)

	if len(cv.SaltList) > 0 {
		futureSalts := cv.SaltList
		for i, salt := range futureSalts {
			if salt.Data2.ValidUntil >= date {
				if i > 0 {
					return futureSalts[i-1:], true
				} else {
					return futureSalts[i:], true
				}
			}
		}
	}

	futureSalts, err := d.AuthSessionRpcClient.SessionGetFutureSalts(ctx, &authsessionpb.TLSessionGetFutureSalts{AuthKeyId: authKeyId})
	if err != nil {
		log.Error(err.Error())
		return nil, false
	}

	saltList := futureSalts.GetSalts()
	for i, salt := range saltList {
		if salt.Data2.ValidUntil >= date {
			if i > 0 {
				saltList = saltList[i-1:]
				cv.SaltList = saltList
				return saltList, true
			} else {
				saltList = saltList[i:]
				cv.SaltList = saltList
				return saltList, true
			}
		}
	}

	return nil, false
}

func (d *Dao) PutUploadInitConnection(ctx context.Context, authKeyId int64, layer int32, ip string, initConnection *mtproto.TLInitConnection) error {
	session := &authsessionpb.TLClientSessionInfo{Data2: &authsessionpb.ClientSession{
		AuthKeyId:      authKeyId,
		Ip:             ip,
		Layer:          layer,
		ApiId:          initConnection.GetApiId(),
		DeviceModel:    initConnection.GetDeviceModel(),
		SystemVersion:  initConnection.GetSystemVersion(),
		AppVersion:     initConnection.GetAppVersion(),
		SystemLangCode: initConnection.GetSystemLangCode(),
		LangPack:       initConnection.GetLangPack(),
		LangCode:       initConnection.GetLangCode(),
		Proxy:          "",
		Params:         "",
	}}

	if initConnection.GetProxy() != nil {
		session.Data2.Proxy = hack.String(model.TLObjectToJson(initConnection.Proxy))
	}
	if initConnection.GetParams() != nil {
		session.Data2.Params = hack.String(model.TLObjectToJson(initConnection.Params))
	}

	request := &authsessionpb.TLSessionSetClientSessionInfo{
		Session: session.To_ClientSession(),
	}

	_, err := d.AuthSessionRpcClient.SessionSetClientSessionInfo(ctx, request)

	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func (d *Dao) GetOrFetchNewSalt(ctx context.Context, authKeyId int64) (salt, lastInvalidSalt *mtproto.TLFutureSalt, err error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	if len(cacheSalts) < 2 {
		return nil, nil, fmt.Errorf("get salt error")
	} else {
		if cacheSalts[0].GetValidUntil() >= int32(time.Now().Unix()) {
			return cacheSalts[0], nil, nil
		} else {
			return cacheSalts[1], cacheSalts[0], nil
		}
	}
}

func (d *Dao) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) ([]*mtproto.TLFutureSalt, error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	return cacheSalts, nil
}

func (d *Dao) GetKeyStateData(ctx context.Context, authKeyId int64) (*authsessionpb.KeyStateData, error) {
	return &authsessionpb.KeyStateData{}, nil
}
