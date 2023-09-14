package core

import (
	"open.chat/app/service/biz_service/poll/internal/dao"
)

type PollCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *PollCore {
	return &PollCore{Dao: dao}
}
