package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/sync/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewChannelPtsUpdatesDAO(db *sqlx.DB) *ChannelPtsUpdatesDAO {
	return &ChannelPtsUpdatesDAO{db}
}

func (dao *ChannelPtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.ChannelPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)"
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

func (dao *ChannelPtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)"
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

func (dao *ChannelPtsUpdatesDAO) SelectLastPts(ctx context.Context, channel_id int32) (rValue *dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select pts from channel_pts_updates where channel_id = ? order by pts desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id)

	if err != nil {
		log.Errorf("queryx in SelectLastPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelPtsUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastPts(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelPtsUpdatesDAO) SelectByGtPts(ctx context.Context, channel_id int32, pts int32) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, update_data from channel_pts_updates where channel_id = ? and pts > ? order by pts asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, pts)

	if err != nil {
		log.Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtPts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
