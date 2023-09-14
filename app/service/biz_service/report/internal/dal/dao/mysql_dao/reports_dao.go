package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/report/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ReportsDAO struct {
	db *sqlx.DB
}

func NewReportsDAO(db *sqlx.DB) *ReportsDAO {
	return &ReportsDAO{db}
}

func (dao *ReportsDAO) Insert(ctx context.Context, do *dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :message_sender_user_id, :message_id, :reason, :text)"
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

func (dao *ReportsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :message_sender_user_id, :message_id, :reason, :text)"
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

func (dao *ReportsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :message_sender_user_id, :message_id, :reason, :text)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *ReportsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ReportsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into reports(user_id, report_type, peer_type, peer_id, message_sender_user_id, message_id, reason, `text`) values (:user_id, :report_type, :peer_type, :peer_id, :message_sender_user_id, :message_id, :reason, :text)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}
