package core

import (
	"context"

	"open.chat/app/json/db/dbo"
	"open.chat/app/json/services/http/account/dao"
	"open.chat/pkg/database/sqlx"
)

// BannedIpCore .
type BannedIpCore struct {
	*dao.Dao
}

// New .
func New(d *dao.Dao) *BannedIpCore {
	if d == nil {
		d = dao.New()
	}
	return &BannedIpCore{d}
}

func (d *BannedIpCore) IsBannedIp(ctx context.Context, ipAddr string) bool {
	ipdo, err := d.BannedIpDAO.Select(ctx, ipAddr)
	if ipdo == nil || err != nil {
		return false
	}
	return true
}

func (d *BannedIpCore) AddBannedIp(ctx context.Context, ipAddr string) bool {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		bdo := &dbo.BannedIp{
			IpAddr: ipAddr,
		}
		bid, _, err := d.BannedIpDAO.InsertTx(tx, bdo)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = (uint32)(bid)
	})
	return tR.Err == nil
}
