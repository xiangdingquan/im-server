package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type MessageDatasDAO struct {
	db *sqlx.DB
}

func NewMessageDatasDAO(db *sqlx.DB) *MessageDatasDAO {
	return &MessageDatasDAO{db}
}

func (dao *MessageDatasDAO) InsertOrGetId(ctx context.Context, do *dataobject.MessageDatasDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_datas(message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date) values (:message_data_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :media_unread, :has_media_unread, :date, :edit_message, :edit_date) on duplicate key update id = last_insert_id(id)"
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

func (dao *MessageDatasDAO) InsertOrGetIdTx(tx *sqlx.Tx, do *dataobject.MessageDatasDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_datas(message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date) values (:message_data_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :media_unread, :has_media_unread, :date, :edit_message, :edit_date) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
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

func (dao *MessageDatasDAO) SelectById(ctx context.Context, id int64) (rValue *dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectById(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessageDatasDAO) SelectMessageListByDataIdList(ctx context.Context, idList []int64) (rList []dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectMessageListByDataIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectMessageListByDataIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageDatasDO
	for rows.Next() {
		v := dataobject.MessageDatasDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectMessageListByDataIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessageDatasDAO) SelectMessageByDataId(ctx context.Context, message_data_id int64) (rValue *dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where message_data_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, message_data_id)

	if err != nil {
		log.Errorf("queryx in SelectMessageByDataId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectMessageByDataId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessageDatasDAO) SelectMessageList(ctx context.Context, dialog_id int64, idList []int32) (rList []dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = ? and dialog_message_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, dialog_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectMessageList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageDatasDO
	for rows.Next() {
		v := dataobject.MessageDatasDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessageDatasDAO) SelectMessage(ctx context.Context, dialog_id int64, dialog_message_id int32) (rValue *dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where dialog_id = ? and dialog_message_id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, dialog_id, dialog_message_id)

	if err != nil {
		log.Errorf("queryx in SelectMessage(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectMessage(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessageDatasDAO) SelectMessageByRandomId(ctx context.Context, sender_user_id int32, random_id int64) (rValue *dataobject.MessageDatasDO, err error) {
	var (
		query = "select id, message_data_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, media_unread, has_media_unread, `date`, edit_message, edit_date from message_datas where sender_user_id = ? and random_id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sender_user_id, random_id)

	if err != nil {
		log.Errorf("queryx in SelectMessageByRandomId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectMessageByRandomId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessageDatasDAO) UpdateFullEditMessage(ctx context.Context, message_type int8, message_data string, edit_message string, edit_date int32, dialog_id int64, dialog_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set message_type = ?, message_data = ?, edit_message = ?, edit_date = ? where dialog_id = ? and dialog_message_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, message_type, message_data, edit_message, edit_date, dialog_id, dialog_message_id)

	if err != nil {
		log.Errorf("exec in UpdateFullEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFullEditMessage(_), error: %v", err)
	}

	return
}

func (dao *MessageDatasDAO) UpdateFullEditMessageTx(tx *sqlx.Tx, message_type int8, message_data string, edit_message string, edit_date int32, dialog_id int64, dialog_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set message_type = ?, message_data = ?, edit_message = ?, edit_date = ? where dialog_id = ? and dialog_message_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, message_type, message_data, edit_message, edit_date, dialog_id, dialog_message_id)

	if err != nil {
		log.Errorf("exec in UpdateFullEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFullEditMessage(_), error: %v", err)
	}

	return
}

func (dao *MessageDatasDAO) UpdateEditMessage(ctx context.Context, edit_message string, edit_date int32, dialog_id int64, dialog_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set edit_message = ?, edit_date = ? where dialog_id = ? and dialog_message_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, edit_message, edit_date, dialog_id, dialog_message_id)

	if err != nil {
		log.Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}

func (dao *MessageDatasDAO) UpdateEditMessageTx(tx *sqlx.Tx, edit_message string, edit_date int32, dialog_id int64, dialog_message_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set edit_message = ?, edit_date = ? where dialog_id = ? and dialog_message_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, edit_message, edit_date, dialog_id, dialog_message_id)

	if err != nil {
		log.Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}

func (dao *MessageDatasDAO) UpdateMediaUnread(ctx context.Context, media_unread int8, message_data_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set media_unread = ? where message_data_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, media_unread, message_data_id)

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

func (dao *MessageDatasDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, media_unread int8, message_data_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_datas set media_unread = ? where message_data_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, media_unread, message_data_id)

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
