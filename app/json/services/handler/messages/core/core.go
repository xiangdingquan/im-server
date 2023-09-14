package core

import "open.chat/app/json/services/handler/messages/dao"

type MessagesCore struct {
	*dao.Dao
}

func New(d *dao.Dao) *MessagesCore {
	if d == nil {
		d = dao.New()
	}
	return &MessagesCore{d}
}
