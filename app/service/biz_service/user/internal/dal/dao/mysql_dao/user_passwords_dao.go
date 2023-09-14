package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserPasswordsDAO struct {
	db *sqlx.DB
}

func NewUserPasswordsDAO(db *sqlx.DB) *UserPasswordsDAO {
	return &UserPasswordsDAO{db}
}

func (dao *UserPasswordsDAO) Insert(ctx context.Context, do *dataobject.UserPasswordsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_passwords(user_id, server_salt, hash, salt, hint, email, state) values (:user_id, :server_salt, '', '', '', '', 0)"
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

func (dao *UserPasswordsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserPasswordsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_passwords(user_id, server_salt, hash, salt, hint, email, state) values (:user_id, :server_salt, '', '', '', '', 0)"
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

func (dao *UserPasswordsDAO) SelectByUserId(ctx context.Context, user_id int32) (rValue *dataobject.UserPasswordsDO, err error) {
	var (
		query = "select user_id, server_salt, hash, salt, hint, email, state from user_passwords where user_id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPasswordsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserPasswordsDAO) SelectCode(ctx context.Context, user_id int32) (rValue *dataobject.UserPasswordsDO, err error) {
	var (
		query = "select code, code_expired, attempts from user_passwords where user_id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectCode(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPasswordsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectCode(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserPasswordsDAO) Update(ctx context.Context, hint string, email string, state int8, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_passwords set salt = ?, hash = ?, hint = ?, email = ?, state = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, hint, email, state, user_id)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *UserPasswordsDAO) UpdateTx(tx *sqlx.Tx, hint string, email string, state int8, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_passwords set salt = ?, hash = ?, hint = ?, email = ?, state = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, hint, email, state, user_id)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
