package core

import (
	"context"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
)

func (m *UserCore) SetAccountDaysTTL(ctx context.Context, userId int32, ttl int32) (err error) {
	_, err = m.UsersDAO.UpdateAccountDaysTTL(ctx, ttl, userId)
	return
}

func (m *UserCore) GetAccountDaysTTL(ctx context.Context, userId int32) (ttl int32, err error) {
	var userDO *dataobject.UsersDO
	userDO, err = m.UsersDAO.SelectAccountDaysTTL(ctx, userId)
	if userDO != nil {
		ttl = userDO.AccountDaysTtl
	}
	return
}
