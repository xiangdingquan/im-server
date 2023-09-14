package mysqldao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// WalletDAO .
type WalletDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewWalletDAO(db *sqlx.DB) *WalletDAO {
	return &WalletDAO{db}
}

// InsertTx .
func (dao *WalletDAO) InsertTx(tx *sqlx.Tx, do *dbo.WalletDO) (lastInsertID, rowsAffected int64, err error) {
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

// SelectByID .
func (dao *WalletDAO) SelectByID(ctx context.Context, id uint32) (rValue *dbo.WalletDO, err error) {
	var (
		query = "SELECT id, uid, address, balance, password, `date` FROM `wallets` WHERE id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectByID(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.WalletDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByID(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

// SelectByUid .
func (dao *WalletDAO) SelectByUid(ctx context.Context, uid uint32) (rValue *dbo.WalletDO, err error) {
	var (
		query = "SELECT id, uid, address, balance, password, `date` FROM `wallets` WHERE uid = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, uid)

	if err != nil {
		log.Errorf("queryx in SelectByUid(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.WalletDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUid(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

// SelectByAddress .
func (dao *WalletDAO) SelectByAddress(ctx context.Context, address string) (rValue *dbo.WalletDO, err error) {
	var (
		query = "SELECT id, uid, address, balance, password, `date` FROM `wallets` WHERE address = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, address)

	if err != nil {
		log.Errorf("queryx in SelectByAddress(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.WalletDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByAddress(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

// UpdatePassword .
func (dao *WalletDAO) UpdatePassword(ctx context.Context, walletID uint32, password string) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE `wallets` SET password = ? WHERE id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, password, walletID)

	if err != nil {
		log.Errorf("exec in UpdateWithID(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithID(_), error: %v", err)
	}

	return
}

// UpdateWithID .
func (dao *WalletDAO) UpdateWithID(ctx context.Context, cMap map[string]interface{}, walletID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, walletID)
	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithID(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithID(_), error: %v", err)
	}

	return
}

// UpdateWithIDTx .
func (dao *WalletDAO) UpdateWithIDTx(tx *sqlx.Tx, cMap map[string]interface{}, walletID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, walletID)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithIDTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithIDTx(_), error: %v", err)
	}

	return
}

// UpdateWithUid .
func (dao *WalletDAO) UpdateWithUid(ctx context.Context, cMap map[string]interface{}, uID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE uid = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, uID)
	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithUid(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithUid(_), error: %v", err)
	}

	return
}

// UpdateWithUidTx .
func (dao *WalletDAO) UpdateWithUidTx(tx *sqlx.Tx, cMap map[string]interface{}, uID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE uid = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, uID)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithUidTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithUidTx(_), error: %v", err)
	}

	return
}

// UpdateWithAddress .
func (dao *WalletDAO) UpdateWithAddress(ctx context.Context, cMap map[string]interface{}, address string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE address = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, address)
	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithAddress(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithAddress(_), error: %v", err)
	}

	return
}

// UpdateWithAddressTx .
func (dao *WalletDAO) UpdateWithAddressTx(tx *sqlx.Tx, cMap map[string]interface{}, address string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `wallets` SET %s WHERE address = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, address)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithAddressTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithAddressTx(_), error: %v", err)
	}

	return
}

// IncreaseBalance 增加余额.
func (dao *WalletDAO) IncreaseBalance(ctx context.Context, uid uint32, money float64) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `wallets` SET balance = balance + ? WHERE uid = ?"
		r     sql.Result
	)
	if money < 0.0 {
		money = -money
		query = "UPDATE `wallets` SET balance = balance - ? WHERE uid = ? AND balance >= ?"
		r, err = dao.db.Exec(ctx, query, money, uid, money)
	} else {
		r, err = dao.db.Exec(ctx, query, money, uid)
	}

	if err != nil {
		log.Errorf("Exec in IncreaseBalance()_error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseBalance()_error: %v", err)
		return
	}

	return
}

// IncreaseBalanceTx 增加余额.
func (dao *WalletDAO) IncreaseBalanceTx(tx *sqlx.Tx, uid uint32, money float64) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `wallets` SET balance = balance + ? WHERE uid = ?"
		r     sql.Result
	)
	if money < 0.0 {
		m := -money
		query = "UPDATE `wallets` SET balance = balance - ? WHERE uid = ? AND balance >= ?"
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
