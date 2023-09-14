package mysql_dao

import (
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// WalletDAO .
type UserWalletDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewUserWalletDAO(db *sqlx.DB) *UserWalletDAO {
	return &UserWalletDAO{db}
}

// InsertTx .
func (dao *UserWalletDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserWalletDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `wallets`(uid, address, `date`) VALUES (:uid, :address, :date)"
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
