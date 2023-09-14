package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelUnreadMentionsDAO struct {
	db *sqlx.DB
}

func NewChannelUnreadMentionsDAO(db *sqlx.DB) *ChannelUnreadMentionsDAO {
	return &ChannelUnreadMentionsDAO{db}
}

func (dao *ChannelUnreadMentionsDAO) InsertIgnore(ctx context.Context, do *dataobject.ChannelUnreadMentionsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into channel_unread_mentions(user_id, channel_id, mentioned_message_id, deleted) values (:user_id, :channel_id, :mentioned_message_id, 0)"
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

func (dao *ChannelUnreadMentionsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.ChannelUnreadMentionsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into channel_unread_mentions(user_id, channel_id, mentioned_message_id, deleted) values (:user_id, :channel_id, :mentioned_message_id, 0)"
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

func (dao *ChannelUnreadMentionsDAO) Delete(ctx context.Context, user_id int32, channel_id int32, mentioned_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_unread_mentions set deleted = 1 where user_id = ? and channel_id = ? and mentioned_message_id = ? and deleted = 0"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, channel_id, mentioned_message_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ChannelUnreadMentionsDAO) DeleteTx(tx *sqlx.Tx, user_id int32, channel_id int32, mentioned_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_unread_mentions set deleted = 1 where user_id = ? and channel_id = ? and mentioned_message_id = ? and deleted = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, channel_id, mentioned_message_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ChannelUnreadMentionsDAO) DeleteByPeer(ctx context.Context, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_unread_mentions set deleted = 1 where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in DeleteByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteByPeer(_), error: %v", err)
	}

	return
}

func (dao *ChannelUnreadMentionsDAO) DeleteByPeerTx(tx *sqlx.Tx, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_unread_mentions set deleted = 1 where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in DeleteByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteByPeer(_), error: %v", err)
	}

	return
}
