package mysqldao

import (
	"context"
	"database/sql"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// WalletRecordDAO .
type WalletRecordDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewWalletRecordDAO(db *sqlx.DB) *WalletRecordDAO {
	return &WalletRecordDAO{db}
}

// InsertTx .
func (dao *WalletRecordDAO) InsertTx(tx *sqlx.Tx, do *dbo.WalletRecordDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `wallet_records`(uid, type, amount, related, remarks, `date`) VALUES (:uid, :type, :amount, :related, :remarks, :date)"
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

// SelectsByUid .
func (dao *WalletRecordDAO) SelectsByUid(ctx context.Context, uid, count, page uint32) (rList []dbo.WalletRecordDO, err error) {
	if count == 0 {
		count = 20
	}
	if page == 0 {
		page = 1
	}
	var startPs = count * (page - 1)
	var (
		query = "SELECT id, uid, type, amount, related, remarks, `date`, deleted FROM `wallet_records` WHERE uid = ? AND deleted = 0 ORDER BY id DESC LIMIT ?,?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, uid, startPs, count)

	if err != nil {
		log.Errorf("queryx in SelectsByUid(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		v := dbo.WalletRecordDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectsByUid(_), error: %v", err)
		}
		rList = append(rList, v)
	}

	return
}
