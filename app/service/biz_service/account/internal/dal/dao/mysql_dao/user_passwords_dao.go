package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserPasswordsDAO struct {
	db *sqlx.DB
}

func NewUserPasswordsDAO(db *sqlx.DB) *UserPasswordsDAO {
	return &UserPasswordsDAO{db}
}

func (dao *UserPasswordsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPasswordsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_passwords(user_id, new_algo_salt1, v, srp_id, srp_b, B, hint, email, has_recovery, state) values (:user_id, :new_algo_salt1, :v, :srp_id, :srp_b, :B, :hint, :email, :has_recovery, :state) on duplicate key update id = last_insert_id(id), new_algo_salt1 = values(new_algo_salt1), v = values(v), srp_id = values(srp_id), srp_b = values(srp_b), B = values(B), hint = values(hint), email = values(email), has_recovery = values(has_recovery), state = values(state)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *UserPasswordsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPasswordsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_passwords(user_id, new_algo_salt1, v, srp_id, srp_b, B, hint, email, has_recovery, state) values (:user_id, :new_algo_salt1, :v, :srp_id, :srp_b, :B, :hint, :email, :has_recovery, :state) on duplicate key update id = last_insert_id(id), new_algo_salt1 = values(new_algo_salt1), v = values(v), srp_id = values(srp_id), srp_b = values(srp_b), B = values(B), hint = values(hint), email = values(email), has_recovery = values(has_recovery), state = values(state)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *UserPasswordsDAO) SelectByUserId(ctx context.Context, user_id int32) (rValue *dataobject.UserPasswordsDO, err error) {
	var (
		query = "select id, user_id, new_algo_salt1, v, srp_id, srp_b, B, hint, email, has_recovery, code, code_expired, attempts, state from user_passwords where user_id = ? limit 1"
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

func (dao *UserPasswordsDAO) Update(ctx context.Context, cMap map[string]interface{}, user_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_passwords set %s where user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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

func (dao *UserPasswordsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_passwords set %s where user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)

	rResult, err = tx.Exec(query, aValues...)

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
