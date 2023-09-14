package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/wallet/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type WalletRecordDAO struct {
	db *sqlx.DB
}

func NewWalletRecordDAO(db *sqlx.DB) *WalletRecordDAO {
	return &WalletRecordDAO{db}
}

// InsertTx .
func (dao *WalletRecordDAO) InsertTx(tx *sqlx.Tx, do *dataobject.WalletRecordDO) (lastInsertID, rowsAffected int64, err error) {
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

func (dao *WalletRecordDAO) Select(ctx context.Context, id uint32) (rValue *dataobject.WalletRecordDO, err error) {
	var (
		query = "select id, uid, type, amount, related, remarks, `date`, deleted from `wallet_records` where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.WalletRecordDO{}
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

func (dao *WalletRecordDAO) SelectByUser(ctx context.Context, userId int32, date int32, offset int32, limit int32) (rValue []*dataobject.WalletRecordDO, err error) {
	var (
		query = "select id, uid, type, amount, related, remarks, `date`, deleted from `wallet_records` where uid = ? and `date` > ? order by id desc limit ?,?"
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.WalletRecordDO, 0)
	rows, err = dao.db.Query(ctx, query, userId, offset, limit)

	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.WalletRecordDO{}
	for rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUser(_), error: %v", err)
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *WalletRecordDAO) SelectCountByUser(ctx context.Context, userId int32, date int32) int {
	var (
		sql = "select count(id) from `wallet_records` where uid = ? and `date` > ?"
	)

	var count int
	err := dao.db.Get(ctx, &count, sql, userId, date)
	if err != nil {
		log.Errorf("SelectCountByUser - [%s] error: %s", sql, err)
		return 0
	}
	return count
}

func (dao *WalletRecordDAO) SelectByType(ctx context.Context, userId int32, recordType int8, date int32, offset int32, limit int32) (rValue []*dataobject.WalletRecordDO, err error) {
	var (
		query = "select id, uid, type, amount, related, remarks, `date`, deleted from `wallet_records` where uid = ? and type = ? and date > ? order by id desc limit ?,?"
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.WalletRecordDO, 0)
	rows, err = dao.db.Query(ctx, query, userId, recordType, offset, limit)

	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.WalletRecordDO{}
	for rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUser(_), error: %v", err)
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *WalletRecordDAO) SelectCountByType(ctx context.Context, userId int32, recordType int8, date int32) int {
	var (
		sql = "select count(id) from `wallet_records` where uid = ? and type = ? and `date` > ?"
	)

	var count int
	err := dao.db.Get(ctx, &count, sql, userId, recordType, date)
	if err != nil {
		log.Errorf("SelectCountByType - [%s] error: %s", sql, err)
		return 0
	}
	return count
}
