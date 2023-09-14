package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelMessagesDeleteDAO struct {
	db *sqlx.DB
}

func NewChannelMessagesDeleteDAO(db *sqlx.DB) *ChannelMessagesDeleteDAO {
	return &ChannelMessagesDeleteDAO{db}
}

func (dao *ChannelMessagesDeleteDAO) InsertOrGetId(ctx context.Context, do *dataobject.ChannelMessagesDeleteDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_messages_delete(user_id, channel_id, message_id) values (:user_id, :channel_id, :message_id) on duplicate key update id = last_insert_id(id), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetId(%v)_error: %v", do, err)
	}

	return
}

func (dao *ChannelMessagesDeleteDAO) InsertOrGetIdTx(tx *sqlx.Tx, do *dataobject.ChannelMessagesDeleteDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_messages_delete(user_id, channel_id, message_id) values (:user_id, :channel_id, :message_id) on duplicate key update id = last_insert_id(id), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetIdTx(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetIdTx(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetIdTx(%v)_error: %v", do, err)
	}

	return
}
