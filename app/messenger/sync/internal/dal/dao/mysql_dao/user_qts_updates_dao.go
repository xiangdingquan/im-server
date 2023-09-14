package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/sync/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserQtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewUserQtsUpdatesDAO(db *sqlx.DB) *UserQtsUpdatesDAO {
	return &UserQtsUpdatesDAO{db}
}

func (dao *UserQtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.UserQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)"
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

func (dao *UserQtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)"
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

func (dao *UserQtsUpdatesDAO) SelectLastQts(ctx context.Context, user_id int32) (rValue *dataobject.UserQtsUpdatesDO, err error) {
	var (
		query = "select qts from user_qts_updates where user_id = ? order by qts desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastQts(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserQtsUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastQts(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserQtsUpdatesDAO) SelectByGtQts(ctx context.Context, user_id int32, qts int32) (rList []dataobject.UserQtsUpdatesDO, err error) {
	var (
		query = "select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = ? and qts > ? order by qts asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, qts)

	if err != nil {
		log.Errorf("queryx in SelectByGtQts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserQtsUpdatesDO
	for rows.Next() {
		v := dataobject.UserQtsUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtQts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
