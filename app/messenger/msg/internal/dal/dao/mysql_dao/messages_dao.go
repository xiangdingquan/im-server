package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type MessagesDAO struct {
	db *sqlx.DB
}

func NewMessagesDAO(db *sqlx.DB) *MessagesDAO {
	return &MessagesDAO{db}
}

func (dao *MessagesDAO) InsertOrReturnId(ctx context.Context, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :message_data_id, :message_data_type, :message, :pts, :pts_count, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :has_media_unread, :ttl_seconds, :date2) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

func (dao *MessagesDAO) InsertOrReturnIdTx(tx *sqlx.Tx, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :message_data_id, :message_data_type, :message, :pts, :pts_count, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :has_media_unread, :ttl_seconds, :date2) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

func (dao *MessagesDAO) SelectByRandomId(ctx context.Context, sender_user_id int32, random_id int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where sender_user_id = ? and random_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sender_user_id, random_id)

	if err != nil {
		log.Errorf("queryx in SelectByRandomId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByRandomId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessagesDAO) SelectByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and deleted = 0 and user_message_box_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectByMessageId(ctx context.Context, user_id int32, user_message_box_id int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByMessageId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessagesDAO) SelectByMessageDataIdList(ctx context.Context, idList []int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where deleted = 0 and message_data_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByMessageDataIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByMessageDataIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectByMessageDataId(ctx context.Context, user_id int32, message_data_id int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and message_data_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_data_id)

	if err != nil {
		log.Errorf("queryx in SelectByMessageDataId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByMessageDataId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessagesDAO) SelectBackwardByOffsetLimit(ctx context.Context, user_id int32, dialog_id int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and dialog_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectBackwardByOffsetLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectForwardByPeerOffsetLimit(ctx context.Context, user_id int32, dialog_id int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and dialog_id = ? and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectForwardByPeerOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectForwardByPeerOffsetLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectPeerMessageId(ctx context.Context, peerId int32, user_id int32, user_message_box_id int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and deleted = 0 and message_data_id = (select message_data_id from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peerId, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectPeerMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectPeerMessageId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessagesDAO) SelectPeerDialogMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id != ? and message_data_id in (select message_data_id from messages where user_id = ? and user_message_box_id in (?)) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPeerDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectDialogMessageListByMessageId(ctx context.Context, user_id int32, user_message_box_id int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ?) and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectDialogMessageListByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDialogMessageListByMessageId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectDialogLastMessageList(ctx context.Context, user_id int32, dialog_id int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and dialog_id = ? and deleted = 0 order by dialog_message_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDialogLastMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectDialogsByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDialogsByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) DeleteMessagesByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

func (dao *MessagesDAO) DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

func (dao *MessagesDAO) SelectDialogMessageIdList(ctx context.Context, user_id int32, dialog_id int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id = ? and dialog_id = ? and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id)

	if err != nil {
		log.Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) SelectPeerMessageList(ctx context.Context, user_id int32, message_data_id int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, message_data_id, message_data_type, message, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, has_media_unread, ttl_seconds, date2 from messages where user_id != ? and message_data_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_data_id)

	if err != nil {
		log.Errorf("queryx in SelectPeerMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPeerMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDAO) UpdateMediaUnread(ctx context.Context, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, user_message_box_id)

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

func (dao *MessagesDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, user_message_box_id)

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

func (dao *MessagesDAO) UpdateMentionedAndMediaUnread(ctx context.Context, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

func (dao *MessagesDAO) UpdateMentionedAndMediaUnreadTx(tx *sqlx.Tx, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

func (dao *MessagesDAO) UpdateEditMessage(ctx context.Context, message_type int8, message_data string, message string, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_type = ?, message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, message_type, message_data, message, user_id, user_message_box_id)

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

func (dao *MessagesDAO) UpdateEditMessageTx(tx *sqlx.Tx, message_type int8, message_data string, message string, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_type = ?, message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, message_type, message_data, message, user_id, user_message_box_id)

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
