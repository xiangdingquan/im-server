package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type SecretChatMessagesDAO struct {
	db *sqlx.DB
}

func NewSecretChatMessagesDAO(db *sqlx.DB) *SecretChatMessagesDAO {
	return &SecretChatMessagesDAO{db}
}

func (dao *SecretChatMessagesDAO) Insert(ctx context.Context, do *dataobject.SecretChatMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chat_messages(sender_user_id, chat_id, random_id, peer_id, message_type, message_data, date2) values (:sender_user_id, :chat_id, :random_id, :peer_id, :message_type, :message_data, :date2)"
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

func (dao *SecretChatMessagesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.SecretChatMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chat_messages(sender_user_id, chat_id, random_id, peer_id, message_type, message_data, date2) values (:sender_user_id, :chat_id, :random_id, :peer_id, :message_type, :message_data, :date2)"
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

func (dao *SecretChatMessagesDAO) SelectList(ctx context.Context, idList []int64) (rList []dataobject.SecretChatMessagesDO, err error) {
	var (
		query = "select id, sender_user_id, chat_id, random_id, peer_id, message_type, message_data, date2 from secret_chat_messages where id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.SecretChatMessagesDO
	for rows.Next() {
		v := dataobject.SecretChatMessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
