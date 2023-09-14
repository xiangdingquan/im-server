package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PopularContactsDAO struct {
	db *sqlx.DB
}

func NewPopularContactsDAO(db *sqlx.DB) *PopularContactsDAO {
	return &PopularContactsDAO{db}
}

func (dao *PopularContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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

func (dao *PopularContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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

func (dao *PopularContactsDAO) IncreaseImporters(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, phone)

	if err != nil {
		log.Errorf("exec in IncreaseImporters(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseImporters(_), error: %v", err)
	}

	return
}

func (dao *PopularContactsDAO) IncreaseImportersTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone)

	if err != nil {
		log.Errorf("exec in IncreaseImporters(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseImporters(_), error: %v", err)
	}

	return
}

func (dao *PopularContactsDAO) IncreaseImportersList(ctx context.Context, phoneList []string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		log.Errorf("sqlx.In in IncreaseImportersList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in IncreaseImportersList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseImportersList(_), error: %v", err)
	}

	return
}

func (dao *PopularContactsDAO) IncreaseImportersListTx(tx *sqlx.Tx, phoneList []string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		log.Errorf("sqlx.In in IncreaseImportersList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in IncreaseImportersList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncreaseImportersList(_), error: %v", err)
	}

	return
}

func (dao *PopularContactsDAO) SelectImporters(ctx context.Context, phone string) (rValue *dataobject.PopularContactsDO, err error) {
	var (
		query = "select phone, importers from popular_contacts where phone = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		log.Errorf("queryx in SelectImporters(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PopularContactsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectImporters(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *PopularContactsDAO) SelectImportersList(ctx context.Context, phoneList []string) (rList []dataobject.PopularContactsDO, err error) {
	var (
		query = "select phone, importers from popular_contacts where phone in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		log.Errorf("sqlx.In in SelectImportersList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectImportersList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PopularContactsDO
	for rows.Next() {
		v := dataobject.PopularContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectImportersList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
