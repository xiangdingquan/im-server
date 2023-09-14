package core

import (
	"context"

	"open.chat/app/service/biz_service/account/internal/dal/dataobject"
)

func (m *AccountCore) SetSettingValue(ctx context.Context, userId int32, key, value string) error {
	_, _, err := m.UserSettingsDAO.InsertOrUpdate(ctx, &dataobject.UserSettingsDO{
		UserId: userId,
		Key2:   key,
		Value:  value,
	})

	return err
}

func (m *AccountCore) GetSettingValue(ctx context.Context, userId int32, key string) (v string, ok bool) {
	if do, err := m.UserSettingsDAO.SelectByKey(ctx, userId, key); do != nil {
		v = do.Value
		ok = true
	} else if err == nil {
		v = ""
		ok = false
	}
	return
}
