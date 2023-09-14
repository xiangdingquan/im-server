package core

import (
	"open.chat/app/messenger/biz_server/phone/internal/dao"
)

type PhoneCallCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *PhoneCallCore {
	return &PhoneCallCore{dao}
}
