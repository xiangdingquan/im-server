package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type MessageBoxesDAO struct {
	db *sqlx.DB
}

func NewMessageBoxesDAO(db *sqlx.DB) *MessageBoxesDAO {
	return &MessageBoxesDAO{db}
}

func (dao *MessageBoxesDAO) Insert(ctx context.Context, do *dataobject.MessageBoxesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into message_boxes(user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, message_filter_type, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :message_data_id, :pts, :pts_count, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :message_filter_type, :date2)"
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

func (dao *MessageBoxesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.MessageBoxesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into message_boxes(user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, message_filter_type, date2) values (:user_id, :user_message_box_id, :dialog_id, :dialog_message_id, :message_data_id, :pts, :pts_count, :message_box_type, :reply_to_msg_id, :mentioned, :media_unread, :message_filter_type, :date2)"
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

func (dao *MessageBoxesDAO) SelectByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and deleted = 0 and user_message_box_id in (?) order by user_message_box_id desc"
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

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectByMessageId(ctx context.Context, user_id int32, user_message_box_id int32) (rValue *dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectByMessageDataIdList(ctx context.Context, idList []int64) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where deleted = 0 and message_data_id in (?) order by user_message_box_id desc"
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

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectByMessageDataId(ctx context.Context, user_id int32, message_data_id int64) (rValue *dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and message_data_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_data_id)

	if err != nil {
		log.Errorf("queryx in SelectByMessageDataId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectBackwardByOffsetLimit(ctx context.Context, user_id int32, dialog_id int64, user_message_box_id int32, limit int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and dialog_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectForwardByPeerOffsetLimit(ctx context.Context, user_id int32, dialog_id int64, user_message_box_id int32, limit int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and dialog_id = ? and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id, user_message_box_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectForwardByPeerOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectPeerMessageId(ctx context.Context, peerId int32, user_id int32, user_message_box_id int32) (rValue *dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_message_box_id, message_box_type from message_boxes where user_id = ? and deleted = 0 and message_data_id = (select message_data_id from message_boxes where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peerId, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectPeerMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectPeerDialogMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and dialog_message_id in (select dialog_message_id from message_boxes where user_id = ? and user_message_box_id in (?)) and deleted = 0"
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

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectDialogMessageListByMessageId(ctx context.Context, user_id int32, user_message_box_id int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where dialog_message_id = (select dialog_message_id from message_boxes where user_id = ? and user_message_box_id = ?) and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectDialogMessageListByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectPeerDialogMessageListByMessageId(ctx context.Context, user_id int32, user_message_box_id int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ?) and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("queryx in SelectPeerDialogMessageListByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPeerDialogMessageListByMessageId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessageBoxesDAO) SelectLastTwoMessageId(ctx context.Context, user_id int32) (rValue *dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_message_box_id from message_boxes where user_id = ? and deleted = 0 order by user_message_box_id desc limit 2"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastTwoMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastTwoMessageId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *MessageBoxesDAO) SelectDialogsByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and user_message_box_id in (?) and deleted = 0"
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

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) DeleteMessagesByMessageIdList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
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

func (dao *MessageBoxesDAO) DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
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

func (dao *MessageBoxesDAO) SelectDialogMessageIdList(ctx context.Context, user_id int32, dialog_id int64) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_message_box_id, date2 from message_boxes where user_id = ? and dialog_id = ? and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id)

	if err != nil {
		log.Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) SelectPeerMessageList(ctx context.Context, user_id int32, message_data_id int64) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id != ? and message_data_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_data_id)

	if err != nil {
		log.Errorf("queryx in SelectPeerMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
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

func (dao *MessageBoxesDAO) UpdateMediaUnread(ctx context.Context, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set media_unread = 0 where user_id = ? and user_message_box_id = ?"
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

func (dao *MessageBoxesDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set media_unread = 0 where user_id = ? and user_message_box_id = ?"
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

func (dao *MessageBoxesDAO) UpdateMessageDataId(ctx context.Context, dialog_message_id int32, message_data_id int64, message_box_type int8, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set dialog_message_id = ?, message_data_id = ?, message_box_type = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, dialog_message_id, message_data_id, message_box_type, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("exec in UpdateMessageDataId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMessageDataId(_), error: %v", err)
	}

	return
}

func (dao *MessageBoxesDAO) UpdateMessageDataIdTx(tx *sqlx.Tx, dialog_message_id int32, message_data_id int64, message_box_type int8, user_id int32, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update message_boxes set dialog_message_id = ?, message_data_id = ?, message_box_type = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, dialog_message_id, message_data_id, message_box_type, user_id, user_message_box_id)

	if err != nil {
		log.Errorf("exec in UpdateMessageDataId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMessageDataId(_), error: %v", err)
	}

	return
}

func (dao *MessageBoxesDAO) SelectByFilterType(ctx context.Context, user_id int32, message_filter_type int8, user_message_box_id int32, limit int32) (rList []dataobject.MessageBoxesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id, dialog_message_id, message_data_id, pts, pts_count, message_box_type, reply_to_msg_id, mentioned, media_unread, date2 from message_boxes where user_id = ? and deleted = 0 and message_filter_type = ? and user_message_box_id < ? order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_filter_type, user_message_box_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectByFilterType(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByFilterType(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
