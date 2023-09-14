package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelMessagesDAO struct {
	db *sqlx.DB
}

func NewChannelMessagesDAO(db *sqlx.DB) *ChannelMessagesDAO {
	return &ChannelMessagesDAO{db}
}

func (dao *ChannelMessagesDAO) InsertOrGetId(ctx context.Context, do *dataobject.ChannelMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_messages(channel_id, channel_message_id, sender_user_id, random_id, pts, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_dm, views, `date`) values (:channel_id, :channel_message_id, :sender_user_id, :random_id, :pts, :message_data_id, :message_type, :message_data, :message, :media_type, :media_unread, :has_media_unread, :edit_message, :edit_date, :ttl_seconds, :has_dm, :views, :date) on duplicate key update id = last_insert_id(id), deleted = 0"
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

func (dao *ChannelMessagesDAO) InsertOrGetIdTx(tx *sqlx.Tx, do *dataobject.ChannelMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_messages(channel_id, channel_message_id, sender_user_id, random_id, pts, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_dm, views, `date`) values (:channel_id, :channel_message_id, :sender_user_id, :random_id, :pts, :message_data_id, :message_type, :message_data, :message, :media_type, :media_unread, :has_media_unread, :edit_message, :edit_date, :ttl_seconds, :has_dm, :views, :date) on duplicate key update id = last_insert_id(id), deleted = 0"
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

func (dao *ChannelMessagesDAO) SelectById(ctx context.Context, id int64) (rValue *dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelMessagesDO{}
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

func (dao *ChannelMessagesDAO) SelectByMessageIdList(ctx context.Context, channel_id int32, idList []int32) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where channel_id = ? and deleted = 0 and channel_message_id in (?) order by channel_message_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, channel_id, idList)
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

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}

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

func (dao *ChannelMessagesDAO) SelectByMessageDataIdList(ctx context.Context, idList []int64) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where deleted = 0 and message_data_id in (?) order by `date` desc"
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

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}
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

func (dao *ChannelMessagesDAO) SelectByMessageId(ctx context.Context, channel_id int32, channel_message_id int32) (rValue *dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where channel_id = ? and channel_message_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, channel_message_id)

	if err != nil {
		log.Errorf("queryx in SelectByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelMessagesDO{}
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

func (dao *ChannelMessagesDAO) SelectMessagesViews(ctx context.Context, channel_id int32, idList []int32) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select channel_message_id, views from channel_messages where channel_id = ? and channel_message_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectMessagesViews(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectMessagesViews(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectMessagesViews(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelMessagesDAO) UpdateMessagesViews(ctx context.Context, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set views = views + 1 where channel_id = ? and channel_message_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateMessagesViews(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateMessagesViews(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMessagesViews(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) UpdateMessagesViewsTx(tx *sqlx.Tx, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set views = views + 1 where channel_id = ? and channel_message_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateMessagesViews(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateMessagesViews(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMessagesViews(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) SelectByRandomId(ctx context.Context, sender_user_id int32, random_id int64) (rValue *dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where sender_user_id = ? and random_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sender_user_id, random_id)

	if err != nil {
		log.Errorf("queryx in SelectByRandomId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelMessagesDO{}
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

func (dao *ChannelMessagesDAO) SelectBackwardByOffsetLimit(ctx context.Context, channel_id int32, channel_message_id int32, available_min_id int32, limit int32) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where channel_id = ? and channel_message_id < ? and channel_message_id > ? and deleted = 0 order by channel_message_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, channel_message_id, available_min_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}
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

func (dao *ChannelMessagesDAO) SelectForwardByOffsetLimit(ctx context.Context, channel_id int32, channel_message_id int32, available_min_id int32, limit int32) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where channel_id = ? and channel_message_id >= ? and channel_message_id > ? and deleted = 0 order by channel_message_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, channel_message_id, available_min_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectForwardByOffsetLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelMessagesDO
	for rows.Next() {
		v := dataobject.ChannelMessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectForwardByOffsetLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelMessagesDAO) DeleteMessages(ctx context.Context, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set deleted = 1 where channel_id = ? and channel_message_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteMessages(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in DeleteMessages(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessages(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) DeleteMessagesTx(tx *sqlx.Tx, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set deleted = 1 where channel_id = ? and channel_message_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in DeleteMessages(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in DeleteMessages(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessages(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) SelectLastMessageNotIdList(ctx context.Context, channel_id int32, idList []int32) (rValue int32, err error) {
	var (
		query = "select channel_message_id from channel_messages where channel_id = ? and deleted = 0 and channel_message_id not in (?) order by channel_message_id desc limit 1"
		a     []interface{}
	)
	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectLastMessageNotIdList(_), error: %v", err)
		return
	}

	err = dao.db.Get(ctx, &rValue, query, a...)

	if err != nil {
		log.Errorf("get in SelectLastMessageNotIdList(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) DeleteMessagesBySenderUserId(ctx context.Context, channel_id int32, sender_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set deleted = 1 where channel_id = ? and sender_user_id = ? and deleted = 0"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, channel_id, sender_user_id)

	if err != nil {
		log.Errorf("exec in DeleteMessagesBySenderUserId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessagesBySenderUserId(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) DeleteMessagesBySenderUserIdTx(tx *sqlx.Tx, channel_id int32, sender_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set deleted = 1 where channel_id = ? and sender_user_id = ? and deleted = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, channel_id, sender_user_id)

	if err != nil {
		log.Errorf("exec in DeleteMessagesBySenderUserId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteMessagesBySenderUserId(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) SelectLastMessageNotSenderUserId(ctx context.Context, channel_id int32, sender_user_id int32) (rValue int32, err error) {
	var query = "select channel_message_id from channel_messages where channel_id = ? and deleted = 0 and (sender_user_id > ? or sender_user_id < ?) order by channel_message_id desc limit 1"
	err = dao.db.Get(ctx, &rValue, query, channel_id, sender_user_id, sender_user_id)

	if err != nil {
		log.Errorf("get in SelectLastMessageNotSenderUserId(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) SelectMessageIdListBySenderUserId(ctx context.Context, channel_id int32, sender_user_id int32) (rList []int32, err error) {
	var query = "select channel_message_id from channel_messages where channel_id = ? and deleted = 0 and sender_user_id = ? order by channel_message_id desc"
	err = dao.db.Select(ctx, &rList, query, channel_id, sender_user_id)

	if err != nil {
		log.Errorf("select in SelectMessageIdListBySenderUserId(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) Update(ctx context.Context, cMap map[string]interface{}, channel_id int32, channel_message_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channel_messages set %s where channel_id = ? and channel_message_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, channel_id)
	aValues = append(aValues, channel_message_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, channel_id int32, channel_message_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channel_messages set %s where channel_id = ? and channel_message_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, channel_id)
	aValues = append(aValues, channel_message_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) RemoveMessages(ctx context.Context, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set has_remove = 1 where channel_id = ? and channel_message_id in (?) and has_remove = 0"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in RemoveMessages(_), error: %v", err)
		return
	}

	rResult, err = dao.db.Exec(ctx, query, a...)
	if err != nil {
		log.Errorf("exec in RemoveMessages(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in RemoveMessages(_), error: %v", err)
	}

	return
}

func (dao *ChannelMessagesDAO) RemoveMessagesTx(tx *sqlx.Tx, channel_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_messages set has_remove = 1 where channel_id = ? and channel_message_id in (?) and has_remove = 0"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in RemoveMessagesTx(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)
	if err != nil {
		log.Errorf("exec in RemoveMessagesTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in RemoveMessagesTx(_), error: %v", err)
	}

	return
}
