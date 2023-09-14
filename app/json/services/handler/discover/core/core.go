package core

import (
	"open.chat/app/json/services/handler/discover/dao"
)

// DiscoverCore .
type DiscoverCore struct {
	*dao.Dao
}

// New .
func New(d *dao.Dao) *DiscoverCore {
	if d == nil {
		d = dao.New()
	}
	return &DiscoverCore{d}
}
