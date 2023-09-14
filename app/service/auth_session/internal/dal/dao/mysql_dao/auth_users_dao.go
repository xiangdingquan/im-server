package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/auth_session/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

func (dao *AuthUsersDAO) InsertOrUpdates(ctx context.Context, do *dataobject.AuthUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, platform, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :platform, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdates(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdates(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdates(%v)_error: %v", do, err)
	}

	return
}

func (dao *AuthUsersDAO) InsertOrUpdatesTx(tx *sqlx.Tx, do *dataobject.AuthUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, platform, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :platform, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdates(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdates(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdates(%v)_error: %v", do, err)
	}

	return
}

func (dao *AuthUsersDAO) Select(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthUsersDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, hash, platform, date_created, date_actived from auth_users where auth_key_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *AuthUsersDAO) SelectAuthKeyIds(ctx context.Context, user_id int32) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, hash, platform, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectAuthKeyIds(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthUsersDO
	for rows.Next() {
		v := dataobject.AuthUsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAuthKeyIds(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *AuthUsersDAO) DeleteByHashList(ctx context.Context, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteByHashList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in DeleteByHashList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteByHashList(_), error: %v", err)
	}

	return
}

func (dao *AuthUsersDAO) DeleteByHashListTx(ctx context.Context, tx *sqlx.Tx, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteByHashList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in DeleteByHashList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteByHashList(_), error: %v", err)
	}

	return
}

func (dao *AuthUsersDAO) SelectListByUserId(ctx context.Context, user_id int32) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, hash, platform, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectListByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthUsersDO
	for rows.Next() {
		v := dataobject.AuthUsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByUserId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *AuthUsersDAO) Delete(ctx context.Context, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *AuthUsersDAO) DeleteTx(ctx context.Context, tx *sqlx.Tx, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *AuthUsersDAO) DeleteUser(ctx context.Context, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		log.Errorf("exec in DeleteUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteUser(_), error: %v", err)
	}

	return
}

func (dao *AuthUsersDAO) DeleteUserTx(tx *sqlx.Tx, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

	if err != nil {
		log.Errorf("exec in DeleteUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteUser(_), error: %v", err)
	}

	return
}
