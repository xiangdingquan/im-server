package core

import (
	"open.chat/app/json/services/handler/wallet/dao"
)

// WalletCore .
type WalletCore struct {
	*dao.Dao
}

// New .
func New(d *dao.Dao) *WalletCore {
	if d == nil {
		d = dao.New()
	}
	return &WalletCore{d}
}
