package core

import (
	"open.chat/app/service/biz_service/user/internal/dao"
)

type UserCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *UserCore {
	return &UserCore{Dao: dao}
}
