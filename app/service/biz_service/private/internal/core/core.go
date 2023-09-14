package core

import (
	"open.chat/app/service/biz_service/private/internal/dao"
)

type PrivateCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *PrivateCore {
	return &PrivateCore{
		Dao: dao,
	}
}
