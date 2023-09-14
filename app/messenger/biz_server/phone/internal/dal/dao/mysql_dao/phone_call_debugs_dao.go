package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/phone/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PhoneCallDebugsDAO struct {
	db *sqlx.DB
}

func NewPhoneCallDebugsDAO(db *sqlx.DB) *PhoneCallDebugsDAO {
	return &PhoneCallDebugsDAO{db}
}

func (dao *PhoneCallDebugsDAO) InsertIgnore(ctx context.Context, do *dataobject.PhoneCallDebugsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into phone_call_debugs(call_id, participant_id, participant_auth_key_id, debug_data) values (:call_id, :participant_id, :participant_auth_key_id, :debug_data)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

func (dao *PhoneCallDebugsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.PhoneCallDebugsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into phone_call_debugs(call_id, participant_id, participant_auth_key_id, debug_data) values (:call_id, :participant_id, :participant_auth_key_id, :debug_data)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}
