package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/sync/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthSeqUpdatesDAO struct {
	db *sqlx.DB
}

func NewAuthSeqUpdatesDAO(db *sqlx.DB) *AuthSeqUpdatesDAO {
	return &AuthSeqUpdatesDAO{db}
}

func (dao *AuthSeqUpdatesDAO) Insert(ctx context.Context, do *dataobject.AuthSeqUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
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

func (dao *AuthSeqUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthSeqUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
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

func (dao *AuthSeqUpdatesDAO) SelectLastSeq(ctx context.Context, auth_id int64, user_id int32) (rValue *dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastSeq(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthSeqUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastSeq(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *AuthSeqUpdatesDAO) SelectByGtSeq(ctx context.Context, auth_id int64, user_id int32, seq int32) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, seq)

	if err != nil {
		log.Errorf("queryx in SelectByGtSeq(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtSeq(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *AuthSeqUpdatesDAO) SelectByGtDate(ctx context.Context, auth_id int64, user_id int32, date2 int32) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, date2)

	if err != nil {
		log.Errorf("queryx in SelectByGtDate(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtDate(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
