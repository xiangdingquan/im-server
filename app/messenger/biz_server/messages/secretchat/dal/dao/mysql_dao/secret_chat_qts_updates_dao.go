package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type SecretChatQtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewSecretChatQtsUpdatesDAO(db *sqlx.DB) *SecretChatQtsUpdatesDAO {
	return &SecretChatQtsUpdatesDAO{db}
}

func (dao *SecretChatQtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.SecretChatQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chat_qts_updates(user_id, auth_key_id, chat_id, qts, chat_message_id, date2) values (:user_id, :auth_key_id, :chat_id, :qts, :chat_message_id, :date2)"
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

func (dao *SecretChatQtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.SecretChatQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chat_qts_updates(user_id, auth_key_id, chat_id, qts, chat_message_id, date2) values (:user_id, :auth_key_id, :chat_id, :qts, :chat_message_id, :date2)"
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

func (dao *SecretChatQtsUpdatesDAO) SelectGtQts(ctx context.Context, auth_key_id int64, qts int32) (rList []dataobject.SecretChatQtsUpdatesDO, err error) {
	var (
		query = "select user_id, auth_key_id, chat_id, qts, chat_message_id, date2 from secret_chat_qts_updates where auth_key_id = ? and qts > ? order by qts asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id, qts)

	if err != nil {
		log.Errorf("queryx in SelectGtQts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.SecretChatQtsUpdatesDO
	for rows.Next() {
		v := dataobject.SecretChatQtsUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectGtQts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
