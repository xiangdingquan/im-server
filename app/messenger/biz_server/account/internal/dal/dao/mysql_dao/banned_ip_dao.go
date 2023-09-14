package mysql_dao

import (
	"context"

	"open.chat/app/messenger/biz_server/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// BannedIpDAO .
type BannedIpDAO struct {
	db *sqlx.DB
}

// NewBannedIpDAO .
func NewBannedIpDAO(db *sqlx.DB) *BannedIpDAO {
	return &BannedIpDAO{db}
}

// Select .
func (dao *BannedIpDAO) Select(ctx context.Context, ip string) (rValue *dataobject.BannedIp, err error) {
	var (
		query = "SELECT id, ip_addr created_at FROM `banned_ips` WHERE ip_addr = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, ip)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BannedIp{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}
