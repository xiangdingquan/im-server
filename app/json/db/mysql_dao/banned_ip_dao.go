package mysqldao

import (
	"context"
	"database/sql"

	"open.chat/app/json/db/dbo"
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

// InsertTx .
func (dao *BannedIpDAO) InsertTx(tx *sqlx.Tx, do *dbo.BannedIp) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `banned_ips`(ip_addr) VALUES (:ip_addr)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(%v)_error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v)_error: %v", do, err)
	}

	return
}

// Select .
func (dao *BannedIpDAO) Select(ctx context.Context, ip string) (rValue *dbo.BannedIp, err error) {
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

	do := &dbo.BannedIp{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
		} else {
			do.CreateAt = (uint32)(do.CreateTime.Unix())
			rValue = do
		}
	}

	return
}
