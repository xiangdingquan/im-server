package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/auth_session/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthKeysDAO struct {
	db *sqlx.DB
}

func NewAuthKeysDAO(db *sqlx.DB) *AuthKeysDAO {
	return &AuthKeysDAO{db}
}

func (dao *AuthKeysDAO) Insert(ctx context.Context, do *dataobject.AuthKeysDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *AuthKeysDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthKeysDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *AuthKeysDAO) SelectByAuthKeyId(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthKeysDO, err error) {
	var (
		query = "select auth_key_id, body from auth_keys where auth_key_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id)

	if err != nil {
		log.Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthKeysDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}
