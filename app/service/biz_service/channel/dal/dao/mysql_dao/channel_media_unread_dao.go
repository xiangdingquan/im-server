package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelMediaUnreadDAO struct {
	db *sqlx.DB
}

func NewChannelMediaUnreadDAO(db *sqlx.DB) *ChannelMediaUnreadDAO {
	return &ChannelMediaUnreadDAO{db}
}

func (dao *ChannelMediaUnreadDAO) Insert(ctx context.Context, do *dataobject.ChannelMediaUnreadDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_media_unread(user_id, channel_id, channel_message_id, media_unread) values (:user_id, :channel_id, :channel_message_id, :media_unread)"
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

func (dao *ChannelMediaUnreadDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelMediaUnreadDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_media_unread(user_id, channel_id, channel_message_id, media_unread) values (:user_id, :channel_id, :channel_message_id, :media_unread)"
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

func (dao *ChannelMediaUnreadDAO) SelectMediaUnread(ctx context.Context, user_id int32, channel_id int32, channel_message_id int32) (rValue *dataobject.ChannelMediaUnreadDO, err error) {
	var (
		query = "select media_unread from channel_media_unread where user_id = ? and channel_id = ? and channel_message_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, channel_id, channel_message_id)

	if err != nil {
		log.Errorf("queryx in SelectMediaUnread(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelMediaUnreadDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectMediaUnread(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelMediaUnreadDAO) UpdateMediaUnread(ctx context.Context, user_id int32, channel_id int32, channel_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_media_unread set media_unread = 0 where user_id = ? and channel_id = ? and channel_message_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, channel_id, channel_message_id)

	if err != nil {
		log.Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}

func (dao *ChannelMediaUnreadDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, user_id int32, channel_id int32, channel_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_media_unread set media_unread = 0 where user_id = ? and channel_id = ? and channel_message_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, channel_id, channel_message_id)

	if err != nil {
		log.Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}
