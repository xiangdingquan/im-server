package core

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/app/service/auth_session/internal/dal/dataobject"
	"open.chat/app/service/auth_session/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

type AuthSessionCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *AuthSessionCore {
	return &AuthSessionCore{dao}
}

func (m *AuthSessionCore) QueryAuthKey(ctx context.Context, authKeyId int64) (*authsessionpb.AuthKeyInfo, error) {
	var keyInfo *authsessionpb.AuthKeyInfo

	cacheKeyData, err := m.Dao.GetAuthKey(ctx, authKeyId)
	if err != nil {
		log.Errorf("queryAuthKey - error: %v", err)
		return nil, err
	} else if cacheKeyData != nil {
		keyInfo = &authsessionpb.AuthKeyInfo{
			AuthKeyId:          cacheKeyData.AuthKeyId,
			AuthKey:            cacheKeyData.AuthKey,
			AuthKeyType:        int32(cacheKeyData.AuthKeyType),
			PermAuthKeyId:      cacheKeyData.PermAuthKeyId,
			TempAuthKeyId:      cacheKeyData.TempAuthKeyId,
			MediaTempAuthKeyId: cacheKeyData.MediaTempAuthKeyId,
		}
	} else {
		do, _ := m.AuthKeysDAO.SelectByAuthKeyId(ctx, authKeyId)
		if do == nil {
			err := fmt.Errorf("not find key - keyId = %d", authKeyId)
			return nil, err
		}
		authKey, err := base64.RawStdEncoding.DecodeString(do.Body)
		if err != nil {
			log.Errorf("read keyData error - keyId = %d, %v", authKeyId, err)
			return nil, err
		}
		keyInfo = &authsessionpb.AuthKeyInfo{
			AuthKeyId:          authKeyId,
			AuthKey:            authKey,
			AuthKeyType:        model.AuthKeyTypePerm,
			PermAuthKeyId:      authKeyId,
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		}
	}
	return keyInfo, nil
}

func (m *AuthSessionCore) InsertAuthKey(ctx context.Context, authKey *model.AuthKeyData, salt *mtproto.TLFutureSalt) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("storage auth_key error: auth_key_id = %v", authKey)
		}
	}()

	do := &dataobject.AuthKeysDO{
		AuthKeyId: authKey.AuthKeyId,
		Body:      base64.RawStdEncoding.EncodeToString(authKey.AuthKey),
	}

	lastInsertId, _, _ := m.AuthKeysDAO.Insert(ctx, do)
	do.Id = int32(lastInsertId)

	if salt != nil {
		err2 := m.putSaltCache(ctx, authKey.AuthKeyId, salt)
		if err2 != nil {
			log.Errorf("put cache error: ", err2)
		}

		m.Dao.PutAuthKey(ctx, authKey.AuthKeyId, authKey, 0)
	}
	return nil
}

func (m *AuthSessionCore) GetApiLayer(ctx context.Context, authKeyId int64) int32 {
	layer, err := m.AuthsDAO.SelectLayer(ctx, authKeyId)
	if err != nil {
		log.Errorf("not find layer - keyId = %d", authKeyId)
		return 0
	}
	return layer
}

func (m *AuthSessionCore) GetLangCode(ctx context.Context, authKeyId int64) string {
	langCode, systemLangCode, err := m.AuthsDAO.SelectLangCode(ctx, authKeyId)
	if err != nil {
		log.Errorf("not find lang_code - keyId = %d", authKeyId)
		return "en"
	}
	if langCode == "" {
		langCode = systemLangCode
	}
	return langCode
}

func (m *AuthSessionCore) GetAuthKeyUserId(ctx context.Context, authKeyId int64) int32 {
	do, _ := m.AuthUsersDAO.Select(ctx, authKeyId)
	if do == nil {
		log.Errorf("not find user - keyId = %d", authKeyId)
		return 0
	}
	return do.UserId
}

func (m *AuthSessionCore) GetPushSessionId(ctx context.Context, userId int32, authKeyId int64, tokenType int32) int64 {
	do, _ := m.DevicesDAO.Select(ctx, authKeyId, userId, int8(tokenType))
	if do == nil {
		log.Errorf("not find token - keyId = %d", authKeyId)
		return 0
	}
	sessionId, _ := util.StringToUint64(do.Token)
	return int64(sessionId)
}

func (m *AuthSessionCore) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int32) int64 {
	now := int32(time.Now().Unix())
	authUsersDO := &dataobject.AuthUsersDO{
		AuthKeyId:   authKeyId,
		UserId:      userId,
		Hash:        rand.Int63(),
		DateCreated: now,
		DateActived: now,
	}
	m.AuthUsersDAO.InsertOrUpdates(ctx, authUsersDO)
	return authUsersDO.Hash
}

func (m *AuthSessionCore) UnbindAuthUser(ctx context.Context, authKeyId int64, userId int32) bool {
	if authKeyId == 0 {
		m.AuthUsersDAO.DeleteUser(ctx, userId)
	} else {
		m.AuthUsersDAO.Delete(ctx, authKeyId, userId)
	}
	return true
}

func (m *AuthSessionCore) SetClientSessionInfo(ctx context.Context, session *authsessionpb.ClientSession) bool {
	do := &dataobject.AuthsDO{
		AuthKeyId:      session.GetAuthKeyId(),
		Layer:          session.GetLayer(),
		ApiId:          session.GetApiId(),
		DeviceModel:    session.GetDeviceModel(),
		SystemVersion:  session.GetSystemVersion(),
		AppVersion:     session.GetAppVersion(),
		SystemLangCode: session.GetSystemLangCode(),
		LangPack:       session.GetLangPack(),
		LangCode:       session.GetLangCode(),
		ClientIp:       session.GetIp(),
		Proxy:          session.GetProxy(),
		Params:         session.GetParams(),
	}
	m.AuthsDAO.InsertOrUpdate(ctx, do)
	return true
}

func (m *AuthSessionCore) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*mtproto.TLFutureSalts, error) {
	pSalts, err := m.getOrNotInsertSaltList(ctx, authKeyId, num)
	if err != nil {
		return nil, err
	}
	salts := &mtproto.TLFutureSalts{Data2: &mtproto.FutureSalts{
		ReqMsgId: 0,
		Now:      0,
		Salts:    pSalts,
	}}
	return salts, nil
}

func (m *AuthSessionCore) GetPermAuthKeyId(ctx context.Context, authKeyId int64) int64 {
	if k, err := m.Dao.GetAuthKey(ctx, authKeyId); err != nil || k == nil {
		return 0
	} else {
		return k.PermAuthKeyId
	}
}
