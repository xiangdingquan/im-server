package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelMessageVisiblesDAO struct {
	db *sqlx.DB
}

func NewChannelMessageVisiblesDAO(db *sqlx.DB) *ChannelMessageVisiblesDAO {
	return &ChannelMessageVisiblesDAO{db}
}

func (dao *ChannelMessageVisiblesDAO) InsertBulk(ctx context.Context, doList []*dataobject.ChannelMessageVisiblesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_message_visibles(user_id, channel_id, message_id) values (:user_id, :channel_id, :message_id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetId(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetId(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetId(%v)_error: %v", doList, err)
	}

	return
}

func (dao *ChannelMessageVisiblesDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ChannelMessageVisiblesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_message_visibles(user_id, channel_id, message_id) values (:user_id, :channel_id, :message_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetIdTx(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetIdTx(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetIdTx(%v)_error: %v", doList, err)
	}

	return
}
