package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
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

func (dao *ChannelAdminLogsDAO) SelectByEvent(ctx context.Context, channel_id int32, event int32, date2 int32) (rList []dataobject.ChannelAdminLogsDO, err error) {
	var (
		query = "select id, user_id, channel_id, event, event_data, `query`, date2 from channel_admin_logs where channel_id = ? and event = ? and date2 > ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, event, date2)

	if err != nil {
		log.Errorf("queryx in SelectByEvent(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelAdminLogsDO
	for rows.Next() {
		v := dataobject.ChannelAdminLogsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByEvent(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelAdminLogsDAO) SelectByChannelId(ctx context.Context, channel_id int32, date2 int32) (rList []dataobject.ChannelAdminLogsDO, err error) {
	var (
		query = "select id, user_id, channel_id, event, event_data, `query`, date2 from channel_admin_logs where channel_id = ? and date2 > ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, date2)

	if err != nil {
		log.Errorf("queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelAdminLogsDO
	for rows.Next() {
		v := dataobject.ChannelAdminLogsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByChannelId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
