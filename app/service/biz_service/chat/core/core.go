package core

import (
	"open.chat/app/service/biz_service/chat/dao"
)

type ChatCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *ChatCore {
	return &ChatCore{
		Dao: dao,
	}
}
