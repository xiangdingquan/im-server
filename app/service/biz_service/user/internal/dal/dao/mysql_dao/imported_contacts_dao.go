package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ImportedContactsDAO struct {
	db *sqlx.DB
}

func NewImportedContactsDAO(db *sqlx.DB) *ImportedContactsDAO {
	return &ImportedContactsDAO{db}
}

func (dao *ImportedContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ImportedContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
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

func (dao *ImportedContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ImportedContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
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

func (dao *ImportedContactsDAO) SelectList(ctx context.Context, user_id int32) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}
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

func (dao *ImportedContactsDAO) SelectListByImportedList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByImportedList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByImportedList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByImportedList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ImportedContactsDAO) SelectAllList(ctx context.Context, user_id int32) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectAllList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ImportedContactsDO
	for rows.Next() {
		v := dataobject.ImportedContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAllList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ImportedContactsDAO) Delete(ctx context.Context, user_id int32, imported_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, imported_user_id)

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

func (dao *ImportedContactsDAO) DeleteTx(tx *sqlx.Tx, user_id int32, imported_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, imported_user_id)

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
