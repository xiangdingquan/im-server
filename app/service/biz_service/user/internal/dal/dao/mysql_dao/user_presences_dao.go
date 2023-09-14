package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserPresencesDAO struct {
	db *sqlx.DB
}

func NewUserPresencesDAO(db *sqlx.DB) *UserPresencesDAO {
	return &UserPresencesDAO{db}
}

func (dao *UserPresencesDAO) Insert(ctx context.Context, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at) values (:user_id, :last_seen_at)"
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

func (dao *UserPresencesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at) values (:user_id, :last_seen_at)"
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

func (dao *UserPresencesDAO) Select(ctx context.Context, user_id int32) (rValue *dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at from user_presences where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPresencesDO{}
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

func (dao *UserPresencesDAO) SelectList(ctx context.Context, idList []int32) (rList []dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at from user_presences where user_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPresencesDO
	for rows.Next() {
		v := dataobject.UserPresencesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserPresencesDAO) UpdateLastSeenAt(ctx context.Context, last_seen_at int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, last_seen_at, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}

func (dao *UserPresencesDAO) UpdateLastSeenAtTx(tx *sqlx.Tx, last_seen_at int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, last_seen_at, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}
