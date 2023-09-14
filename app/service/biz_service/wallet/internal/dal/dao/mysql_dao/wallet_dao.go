package mysql_dao

import (
	"context"
	"database/sql"
	"errors"

	"open.chat/app/service/biz_service/wallet/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type WalletDAO struct {
	db *sqlx.DB
}

func NewWalletDAO(db *sqlx.DB) *WalletDAO {
	return &WalletDAO{db}
}

// InsertTx .
func (dao *WalletDAO) InsertTx(tx *sqlx.Tx, do *dataobject.WalletDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `wallets`(uid, address, password, `date`) VALUES (:uid, :address, :password, :date)"
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

func (dao *WalletDAO) Select(ctx context.Context, id uint32) (rValue *dataobject.WalletDO, err error) {
	var (
		query = "select id, uid, address, balance, password, `date`, deleted from `wallets` where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.WalletDO{}
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

func (dao *WalletDAO) SelectByUser(ctx context.Context, userId int32) (rValue *dataobject.WalletDO, err error) {
	var (
		query = "select id, uid, address, balance, password, `date`, deleted from `wallets` where uid = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.WalletDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUser(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

// IncreaseBalanceTx 增加余额.
func (dao *WalletDAO) IncreaseBalanceTx(tx *sqlx.Tx, uid int32, money float64) (rowsAffected int64, err error) {
	var (
		query string = "update `wallets` set balance = balance + ? where uid = ?"
		r     sql.Result
	)
	if money < 0.0 {
		m := -money
		query = "update `wallets` set balance = balance - ? where uid = ? and balance >= ?"
		r, err = tx.Exec(query, m, uid, m)
	} else {
		r, err = tx.Exec(query, money, uid)
	}

	if err != nil {
		log.Errorf("exec in IncreaseBalanceTx(_), error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseBalanceTx(_), error: %v", err)
	}

	if money < 0.0 && rowsAffected == 0 {
		err = errors.New("the balance is too little")
	}

	return
}
