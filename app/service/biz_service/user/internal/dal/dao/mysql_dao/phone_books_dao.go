package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PhoneBooksDAO struct {
	db *sqlx.DB
}

func NewPhoneBooksDAO(db *sqlx.DB) *PhoneBooksDAO {
	return &PhoneBooksDAO{db}
}

func (dao *PhoneBooksDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
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

func (dao *PhoneBooksDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
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
