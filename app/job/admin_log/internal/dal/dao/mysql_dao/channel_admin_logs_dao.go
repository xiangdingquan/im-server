package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/job/admin_log/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelAdminLogsDAO struct {
	db *sqlx.DB
}

func NewChannelAdminLogsDAO(db *sqlx.DB) *ChannelAdminLogsDAO {
	return &ChannelAdminLogsDAO{db}
}

func (dao *ChannelAdminLogsDAO) Insert(ctx context.Context, do *dataobject.ChannelAdminLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_admin_logs(user_id, channel_id, event, event_data, `query`, date2) values (:user_id, :channel_id, :event, :event_data, :query, :date2)"
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

func (dao *ChannelAdminLogsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelAdminLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_admin_logs(user_id, channel_id, event, event_data, `query`, date2) values (:user_id, :channel_id, :event, :event_data, :query, :date2)"
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
