package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/bots/botfather/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

func (dao *UsersDAO) Insert(ctx context.Context, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, username, phone, country_code, is_bot) values (2, :access_hash, :secret_key_id, :first_name, :username, :phone, :country_code, 1)"
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

func (dao *UsersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, username, phone, country_code, is_bot) values (2, :access_hash, :secret_key_id, :first_name, :username, :phone, :country_code, 1)"
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

func (dao *UsersDAO) SelectListByCreator(ctx context.Context, ownerUserId int32) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, username, phone, country_code, is_bot from users where id in (select bot_id from bots where creator_user_id = ?)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, ownerUserId)

	if err != nil {
		log.Errorf("queryx in SelectListByCreator(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByCreator(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
