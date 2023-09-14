package core

import (
	"open.chat/app/json/services/handler/call/dao"
)

// AVCallCore .
type AVCallCore struct {
	*dao.Dao
}

// New .
func New(d *dao.Dao) *AVCallCore {
	if d == nil {
		d = dao.New()
	}
	return &AVCallCore{d}
}
