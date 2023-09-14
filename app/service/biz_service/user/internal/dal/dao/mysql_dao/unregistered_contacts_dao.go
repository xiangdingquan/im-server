package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UnregisteredContactsDAO struct {
	db *sqlx.DB
}

func NewUnregisteredContactsDAO(db *sqlx.DB) *UnregisteredContactsDAO {
	return &UnregisteredContactsDAO{db}
}

func (dao *UnregisteredContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UnregisteredContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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

func (dao *UnregisteredContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UnregisteredContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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

func (dao *UnregisteredContactsDAO) SelectImportersByPhone(ctx context.Context, phone string) (rList []dataobject.UnregisteredContactsDO, err error) {
	var (
		query = "select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = ? and imported = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		log.Errorf("queryx in SelectImportersByPhone(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UnregisteredContactsDO
	for rows.Next() {
		v := dataobject.UnregisteredContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectImportersByPhone(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UnregisteredContactsDAO) UpdateContactName(ctx context.Context, import_first_name string, import_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, import_first_name, import_last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

func (dao *UnregisteredContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, import_first_name string, import_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, import_first_name, import_last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

func (dao *UnregisteredContactsDAO) DeleteContacts(ctx context.Context, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		log.Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

func (dao *UnregisteredContactsDAO) DeleteContactsTx(tx *sqlx.Tx, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		log.Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

func (dao *UnregisteredContactsDAO) DeleteImportersByPhone(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, phone)

	if err != nil {
		log.Errorf("exec in DeleteImportersByPhone(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteImportersByPhone(_), error: %v", err)
	}

	return
}

func (dao *UnregisteredContactsDAO) DeleteImportersByPhoneTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone)

	if err != nil {
		log.Errorf("exec in DeleteImportersByPhone(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteImportersByPhone(_), error: %v", err)
	}

	return
}
