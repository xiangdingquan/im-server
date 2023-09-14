package mysql_dao

import (
	"database/sql"

	"context"

	"open.chat/app/messenger/biz_server/auth/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthOpLogsDAO struct {
	db *sqlx.DB
}

func NewAuthOpLogsDAO(db *sqlx.DB) *AuthOpLogsDAO {
	return &AuthOpLogsDAO{db}
}

func (dao *AuthOpLogsDAO) Insert(ctx context.Context, do *dataobject.AuthOpLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"
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

func (dao *AuthOpLogsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthOpLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"
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
