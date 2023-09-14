package core

import (
	"open.chat/app/service/biz_service/account/internal/dao"
)

type AccountCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *AccountCore {
	return &AccountCore{Dao: dao}
}
