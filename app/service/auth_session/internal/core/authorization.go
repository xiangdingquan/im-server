package core

import (
	"context"
	"net"

	"open.chat/app/service/auth_session/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (m *AuthSessionCore) getCountryAndRegionByIp(ip string) (string, string) {
	r, err := m.MMDB.City(net.ParseIP(ip))
	if err != nil {
		log.Errorf("getCountryAndRegionByIp - error: %v", err)
		return "", ""
	}

	return r.City.Names["en"] + ", " + r.Country.Names["en"], r.Country.IsoCode
}

func getAppNameByAppId(appId int32) string {
	return "tdesktop"
}

func (m *AuthSessionCore) GetAuthorization(ctx context.Context, authKeyId int64) (*mtproto.Authorization, error) {
	authUsersDO, err := m.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
	if err != nil || authUsersDO == nil {
		return nil, err
	}

	country, region := m.getCountryAndRegionByIp(authUsersDO.ClientIp)
	return mtproto.MakeTLAuthorization(&mtproto.Authorization{
		Current:       false,
		OfficialApp:   true,
		Hash:          0,
		DeviceModel:   authUsersDO.DeviceModel,
		Platform:      "",
		SystemVersion: authUsersDO.SystemVersion,
		ApiId:         authUsersDO.ApiId,
		AppName:       authUsersDO.LangPack,
		AppVersion:    authUsersDO.AppVersion,
		DateCreated:   0,
		DateActive:    0,
		Ip:            authUsersDO.ClientIp,
		Country:       country,
		Region:        region,
	}).To_Authorization(), nil

}

func (m *AuthSessionCore) GetAuthorizations(ctx context.Context, userId int32, excludeAuthKeyId int64) (authorizations []*mtproto.Authorization) {
	authUsersDOList, _ := m.AuthUsersDAO.SelectAuthKeyIds(ctx, userId)
	authorizations = make([]*mtproto.Authorization, 0, len(authUsersDOList))
	idList := make([]int64, 0, len(authUsersDOList))

	for i := 0; i < len(authUsersDOList); i++ {
		idList = append(idList, authUsersDOList[i].AuthKeyId)
	}
	if len(idList) == 0 {
		return
	}

	getAuthUsersDO := func(authKeyId int64) *dataobject.AuthUsersDO {
		for i := 0; i < len(authUsersDOList); i++ {
			if authKeyId == authUsersDOList[i].AuthKeyId {
				return &authUsersDOList[i]
			}
		}
		return nil
	}

	myIdx := -1
	authsDOList, _ := m.AuthsDAO.SelectSessions(ctx, idList)
	for i := 0; i < len(authsDOList); i++ {
		authUsersDO := getAuthUsersDO(authsDOList[i].AuthKeyId)
		if excludeAuthKeyId == authsDOList[i].AuthKeyId {
			myIdx = i
			continue
		}

		country, region := m.getCountryAndRegionByIp(authsDOList[i].ClientIp)
		authorization := mtproto.MakeTLAuthorization(&mtproto.Authorization{
			OfficialApp:   true,
			Hash:          authUsersDO.Hash,
			DeviceModel:   authsDOList[i].DeviceModel,
			Platform:      authUsersDO.Platform,
			SystemVersion: authsDOList[i].SystemVersion,
			ApiId:         authsDOList[i].ApiId,
			AppName:       authsDOList[i].LangPack,
			AppVersion:    authsDOList[i].AppVersion,
			DateCreated:   authUsersDO.DateCreated,
			DateActive:    authUsersDO.DateActived,
			Ip:            authsDOList[i].ClientIp,
			Country:       country,
			Region:        region,
		}).To_Authorization()

		log.Debugf("%d - %s", i, authorization.DebugString())
		authorizations = append(authorizations, authorization)
	}

	if myIdx != -1 {
		log.Debugf("excludeAuthKeyId - %d", excludeAuthKeyId)
		authUsersDO := getAuthUsersDO(excludeAuthKeyId)
		country, region := m.getCountryAndRegionByIp(authsDOList[myIdx].ClientIp)
		authorizations = append([]*mtproto.Authorization{mtproto.MakeTLAuthorization(&mtproto.Authorization{
			Current:       true,
			OfficialApp:   true,
			Hash:          0,
			DeviceModel:   authsDOList[myIdx].DeviceModel,
			Platform:      authUsersDO.Platform,
			SystemVersion: authsDOList[myIdx].SystemVersion,
			ApiId:         authsDOList[myIdx].ApiId,
			AppName:       authsDOList[myIdx].LangPack,
			AppVersion:    authsDOList[myIdx].AppVersion,
			DateCreated:   authUsersDO.DateCreated,
			DateActive:    authUsersDO.DateActived,
			Ip:            authsDOList[myIdx].ClientIp,
			Country:       country,
			Region:        region,
		}).To_Authorization()}, authorizations...)
	}

	return
}

func (m *AuthSessionCore) ResetAuthorization(ctx context.Context, userId int32, authKeyId, hash int64) []int64 {
	doList, _ := m.AuthUsersDAO.SelectListByUserId(ctx, userId)
	if len(doList) == 0 {
		return []int64{}
	}

	var (
		keyIdList []int64
		idList    []int32
	)
	if hash == 0 {
		for i := 0; i < len(doList); i++ {
			if doList[i].AuthKeyId != authKeyId {
				idList = append(idList, doList[i].Id)
				keyIdList = append(keyIdList, doList[i].AuthKeyId)
			}
		}
	} else {
		for i := 0; i < len(doList); i++ {
			if doList[i].Hash == hash && doList[i].AuthKeyId != authKeyId {
				idList = append(idList, doList[i].Id)
				keyIdList = append(keyIdList, doList[i].AuthKeyId)
			}
		}
	}

	if len(idList) > 0 {
		m.AuthUsersDAO.DeleteByHashList(ctx, idList)
	} else {
		keyIdList = []int64{}
	}
	return keyIdList
}
