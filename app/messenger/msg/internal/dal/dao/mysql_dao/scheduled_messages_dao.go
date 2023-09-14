package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ScheduledMessagesDAO struct {
	db *sqlx.DB
}

func NewScheduledMessagesDAO(db *sqlx.DB) *ScheduledMessagesDAO {
	return &ScheduledMessagesDAO{db}
}

func (dao *ScheduledMessagesDAO) InsertOrReturnId(ctx context.Context, do *dataobject.ScheduledMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into scheduled_messages(user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state) values (:user_id, :user_message_box_id, :peer_type, :peer_id, :dialog_id, :random_id, :message_type, :message_data_type, :message_data, :message_box_type, :scheduled_date, :date2, 0) on duplicate key update id = last_insert_id(id)"
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

func (dao *ScheduledMessagesDAO) InsertOrReturnIdTx(tx *sqlx.Tx, do *dataobject.ScheduledMessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into scheduled_messages(user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state) values (:user_id, :user_message_box_id, :peer_type, :peer_id, :dialog_id, :random_id, :message_type, :message_data_type, :message_data, :message_box_type, :scheduled_date, :date2, 0) on duplicate key update id = last_insert_id(id)"
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

func (dao *ScheduledMessagesDAO) SelectById(ctx context.Context, id int64) (rValue *dataobject.ScheduledMessagesDO, err error) {
	var (
		query = "select id, user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state from scheduled_messages where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ScheduledMessagesDO{}
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

func (dao *ScheduledMessagesDAO) SelectByMessageIdList(ctx context.Context, user_id int32, peer_type int8, peer_id int32, idList []int32) (rList []dataobject.ScheduledMessagesDO, err error) {
	var (
		query = "select id, user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state from scheduled_messages where user_id = ? and state = 0 and peer_type = ? and peer_id = ? and user_message_box_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, peer_type, peer_id, idList)
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

	var values []dataobject.ScheduledMessagesDO
	for rows.Next() {
		v := dataobject.ScheduledMessagesDO{}
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

func (dao *ScheduledMessagesDAO) SelectHistory(ctx context.Context, user_id int32, peer_type int8, peer_id int32) (rList []dataobject.ScheduledMessagesDO, err error) {
	var (
		query = "select id, user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state from scheduled_messages where user_id = ? and state = 0 and peer_type = ? and peer_id = ? order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		log.Errorf("queryx in SelectHistory(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ScheduledMessagesDO
	for rows.Next() {
		v := dataobject.ScheduledMessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectHistory(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ScheduledMessagesDAO) SelectScheduled(ctx context.Context, scheduled_date int32) (rList []dataobject.ScheduledMessagesDO, err error) {
	var (
		query = "select id, user_id, user_message_box_id, peer_type, peer_id, dialog_id, random_id, message_type, message_data_type, message_data, message_box_type, scheduled_date, date2, state from scheduled_messages where state = 0 and scheduled_date > ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, scheduled_date)

	if err != nil {
		log.Errorf("queryx in SelectScheduled(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ScheduledMessagesDO
	for rows.Next() {
		v := dataobject.ScheduledMessagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectScheduled(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ScheduledMessagesDAO) UpdateStateByMessageId(ctx context.Context, state int8, user_id int32, peer_type int8, peer_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update scheduled_messages set state = ? where user_id = ? and state = 0 and peer_type = ? and peer_id = ? and user_message_box_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, state, user_id, peer_type, peer_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateStateByMessageId(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateStateByMessageId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByMessageId(_), error: %v", err)
	}

	return
}

func (dao *ScheduledMessagesDAO) UpdateStateByMessageIdTx(tx *sqlx.Tx, state int8, user_id int32, peer_type int8, peer_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update scheduled_messages set state = ? where user_id = ? and state = 0 and peer_type = ? and peer_id = ? and user_message_box_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, state, user_id, peer_type, peer_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateStateByMessageId(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateStateByMessageId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByMessageId(_), error: %v", err)
	}

	return
}

func (dao *ScheduledMessagesDAO) UpdateStateByIdList(ctx context.Context, state int8, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update scheduled_messages set state = ? where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, state, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateStateByIdList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateStateByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByIdList(_), error: %v", err)
	}

	return
}

func (dao *ScheduledMessagesDAO) UpdateStateByIdListTx(tx *sqlx.Tx, state int8, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update scheduled_messages set state = ? where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, state, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateStateByIdList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateStateByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByIdList(_), error: %v", err)
	}

	return
}
