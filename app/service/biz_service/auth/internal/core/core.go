package core

import (
	"context"
	"github.com/pkg/errors"
	"open.chat/app/service/biz_service/auth/internal/dao"
	"open.chat/pkg/log"
	"strings"
)

type AuthCore struct {
	*dao.Dao
}

const (
	PlayformInvalid = 0
	PlatformIOS     = 1
	PlatformAndroid = 2
	PlatformDesktop = 3
	PlatformWeb     = 4
)

func New(dao *dao.Dao) *AuthCore {
	return &AuthCore{Dao: dao}
}

func (m *AuthCore) GetPlatform(ctx context.Context, authKeyId int64) (int32, error) {
	do, err := m.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return PlayformInvalid, err
	}

	if do == nil {
		log.Errorf("GetPlatform, auth not found by key:%d", authKeyId)
		return PlayformInvalid, errors.New("auth not found")
	}

	isIOS := func(deviceModel string) bool {
		lower := strings.ToLower(deviceModel)
		l := []string{"iphone", "ios", "ipad"}
		for _, s := range l {
			if strings.Contains(lower, s) {
				return true
			}
		}
		return false
	}

	isDesktop := func(systemVersion string) bool {
		lower := strings.ToLower(systemVersion)
		l := []string{"windows", "macos"}
		for _, s := range l {
			if strings.Contains(lower, s) {
				return true
			}
		}
		return false
	}

	if do.ApiId == 81 {
		return PlatformWeb, nil
	} else if isIOS(do.DeviceModel) {
		return PlatformIOS, nil
	} else if isDesktop(do.SystemVersion) {
		return PlatformDesktop, nil
	} else {
		return PlatformAndroid, nil
	}
}
