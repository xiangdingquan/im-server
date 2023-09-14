package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
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
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)"
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
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)"
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

func (dao *ChannelPtsUpdatesDAO) SelectByGtPts(ctx context.Context, channel_id int32, pts int32, limit int32) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id = ? and pts > ? order by pts asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, pts, limit)

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

func (dao *ChannelPtsUpdatesDAO) SelectByGtDate2(ctx context.Context, idList []int32, date2 int32) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id in (?) and date2 > ? order by date2 asc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList, date2)
	if err != nil {
		log.Errorf("sqlx.In in SelectByGtDate2(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByGtDate2(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtDate2(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
