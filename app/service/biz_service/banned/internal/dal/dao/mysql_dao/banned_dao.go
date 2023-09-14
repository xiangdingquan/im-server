package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/banned/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BannedDAO struct {
	db *sqlx.DB
}

func NewBannedDAO(db *sqlx.DB) *BannedDAO {
	return &BannedDAO{db}
}

func (dao *BannedDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BannedDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)"
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

func (dao *BannedDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BannedDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)"
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

func (dao *BannedDAO) CheckBannedByPhone(ctx context.Context, phone string) (rValue *dataobject.BannedDO, err error) {
	var (
		query = "select id from banned where phone = ? and state > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		log.Errorf("queryx in CheckBannedByPhone(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BannedDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in CheckBannedByPhone(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BannedDAO) SelectPhoneList(ctx context.Context, pList []string) (rList []string, err error) {
	var (
		query = "select phone from banned where state > 0 and phone in (?)"
		a     []interface{}
	)
	query, a, err = sqlx.In(query, pList)
	if err != nil {
		log.Errorf("sqlx.In in SelectPhoneList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		log.Errorf("select in SelectPhoneList(_), error: %v", err)
	}

	return
}

func (dao *BannedDAO) Update(ctx context.Context, expires int64, log2 string, state int8, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update banned set expires = ?, state = 0, log = ?, state = ? where phone = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, expires, log2, state, phone)

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

func (dao *BannedDAO) UpdateTx(tx *sqlx.Tx, expires int64, log2 string, state int8, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update banned set expires = ?, state = 0, log = ?, state = ? where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, expires, log2, state, phone)

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
