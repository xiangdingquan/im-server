package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserSettingsDAO struct {
	db *sqlx.DB
}

func NewUserSettingsDAO(db *sqlx.DB) *UserSettingsDAO {
	return &UserSettingsDAO{db}
}

func (dao *UserSettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
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

func (dao *UserSettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_settings(user_id, key2, value) values (:user_id, :key2, :value) on duplicate key update value = values(value)"
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

func (dao *UserSettingsDAO) SelectByKey(ctx context.Context, user_id int32, key2 string) (rValue *dataobject.UserSettingsDO, err error) {
	var (
		query = "select id, user_id, key2, value from user_settings where user_id = ? and key2 = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, key2)

	if err != nil {
		log.Errorf("queryx in SelectByKey(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserSettingsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByKey(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserSettingsDAO) Update(ctx context.Context, value string, user_id int32, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, value, user_id, key2)

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

func (dao *UserSettingsDAO) UpdateTx(tx *sqlx.Tx, value string, user_id int32, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update user_settings set value = ?, deleted = 0 where user_id = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, value, user_id, key2)

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
