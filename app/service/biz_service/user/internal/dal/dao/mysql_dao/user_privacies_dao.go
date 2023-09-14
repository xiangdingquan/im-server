package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserPrivaciesDAO struct {
	db *sqlx.DB
}

func NewUserPrivaciesDAO(db *sqlx.DB) *UserPrivaciesDAO {
	return &UserPrivaciesDAO{db}
}

func (dao *UserPrivaciesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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

func (dao *UserPrivaciesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules) on duplicate key update rules = values(rules)"
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

func (dao *UserPrivaciesDAO) InsertBulk(ctx context.Context, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *UserPrivaciesDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.UserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_privacies(user_id, key_type, rules) values (:user_id, :key_type, :rules)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *UserPrivaciesDAO) SelectPrivacy(ctx context.Context, user_id int32, key_type int8) (rValue *dataobject.UserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, key_type)

	if err != nil {
		log.Errorf("queryx in SelectPrivacy(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPrivaciesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectPrivacy(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserPrivaciesDAO) SelectPrivacyList(ctx context.Context, user_id int32, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ? and key_type in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, keyList)
	if err != nil {
		log.Errorf("sqlx.In in SelectPrivacyList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectPrivacyList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPrivaciesDO
	for rows.Next() {
		v := dataobject.UserPrivaciesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPrivacyList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserPrivaciesDAO) SelectUsersPrivacyList(ctx context.Context, idList []int32, keyList []int32) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id in (?) and key_type in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList, keyList)
	if err != nil {
		log.Errorf("sqlx.In in SelectUsersPrivacyList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectUsersPrivacyList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPrivaciesDO
	for rows.Next() {
		v := dataobject.UserPrivaciesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUsersPrivacyList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserPrivaciesDAO) SelectPrivacyAll(ctx context.Context, user_id int32) (rList []dataobject.UserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from user_privacies where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectPrivacyAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPrivaciesDO
	for rows.Next() {
		v := dataobject.UserPrivaciesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPrivacyAll(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
