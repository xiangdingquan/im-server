package mysql_dao

import (
	"context"
	"database/sql"
	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type SecretChatsCloseRequestsDAO struct {
	db *sqlx.DB
}

func NewSecretChatsCloseRequestsDAO(db *sqlx.DB) *SecretChatsCloseRequestsDAO {
	return &SecretChatsCloseRequestsDAO{db}
}

func (dao *SecretChatsCloseRequestsDAO) Insert(ctx context.Context, do *dataobject.SecretChatCloseRequestsDo) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chat_close_requests(secret_chat_id, from_uid, to_uid) values (:secret_chat_id, :from_uid, :to_uid)"
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

func (dao *SecretChatsCloseRequestsDAO) SelectByUser(ctx context.Context, toUID int32) (rList []*dataobject.SecretChatCloseRequestsDo, err error) {
	var (
		query = "select secret_chat_id from secret_chat_close_requests where to_uid=? and closed=0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, toUID)

	if err != nil {
		log.Errorf("queryx in Select(%d), error: %v", toUID, err)
		return
	}

	defer rows.Close()

	rList = make([]*dataobject.SecretChatCloseRequestsDo, 0)
	for rows.Next() {
		do := &dataobject.SecretChatCloseRequestsDo{}
		err = rows.StructScan(do)
		if err != nil {
			rList = nil
			log.Errorf("structScan in Select(%d), error: %v", toUID, err)
			return
		} else {
			rList = append(rList, do)
		}
	}

	return
}

func (dao *SecretChatsCloseRequestsDAO) SelectByUserAndChat(ctx context.Context, toUID, secretChatId int32) (rValue *dataobject.SecretChatCloseRequestsDo, err error) {
	var (
		query = "select secret_chat_id from secret_chat_close_requests where to_uid=? and secret_chat_id=? and closed=0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, toUID, secretChatId)

	if err != nil {
		log.Errorf("queryx in Select(%d, %d), error: %v", toUID, secretChatId, err)
		return
	}

	defer rows.Close()

	do := &dataobject.SecretChatCloseRequestsDo{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(%d, %d), error: %v", toUID, secretChatId, err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *SecretChatsCloseRequestsDAO) UpdateClosed(ctx context.Context, toUID, secretChatId int32) (rowsAffected int64, err error) {
	var (
		query   = "update secret_chat_close_requests set closed=1 where to_uid=? and secret_chat_id=? and closed=0"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, toUID, secretChatId)

	if err != nil {
		log.Errorf("exec in UpdateClosed(%d, %d), error: %v", toUID, secretChatId, err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateClosed(%d, %d), error: %v", toUID, secretChatId, err)
	}

	return
}
