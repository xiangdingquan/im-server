package mysqldao

import (
	"context"
	"database/sql"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type MessagesDao struct {
	db *sqlx.DB
}

func NewMessagesDao(db *sqlx.DB) *MessagesDao {
	return &MessagesDao{db: db}
}

func (dao *MessagesDao) InsertBatchSendMessage(ctx context.Context, do *dbo.BatchSendMessageDo) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `batch_send_messages` (uid, message, to_users) VALUES (:uid, :message, :to_users)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertBatchSendMessage(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertID in InsertBatchSendMessage(%v), error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBatchSendMessage(%v), error: %v", do, err)
		return
	}

	return
}

func (dao *MessagesDao) SelectBatchSendMessage(ctx context.Context, uid, fromId, limit int32) (rValue []*dbo.BatchSendMessageDo, err error) {
	var (
		query = "SELECT id, message, to_users FROM batch_send_messages where uid=? and id<? and deleted=0 order by id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, uid, fromId, limit)

	if err != nil {
		log.Errorf("query in SelectBatchSendMesage(%d), error: %v", uid, err)
		return
	}

	defer rows.Close()

	rValue = make([]*dbo.BatchSendMessageDo, 0)
	for rows.Next() {
		do := &dbo.BatchSendMessageDo{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBatchSendMessage(%d), error: %v", uid, err)
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *MessagesDao) DeleteBatchSendMessageDelete(ctx context.Context, uid uint32) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE batch_send_messages set deleted=1 where uid=?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, uid)

	if err != nil {
		log.Errorf("exec in DeleteBatchSendMessageDelete(%d), error: %v", uid, err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteBatchSendMessageDelete(%d), error: %v", uid, err)
	}

	return
}

func (dao *MessagesDao) SelectDialogId(ctx context.Context, userId, peerId int32) (dialogId int64, err error) {
	var (
		query = "SELECT dialog_id FROM messages where user_id=? and peer_type=2 and peer_id=? limit 1"
	)

	err = dao.db.Get(ctx, &dialogId, query, userId, peerId)
	if err != nil {
		log.Errorf("get in SelectDialogId(%d, %d), error: %v", userId, peerId, err)
	}
	//log.Debugf("SelectDialogId(%d, %d), dialogId:%d", userId, peerId, dialogId)
	return
}

func (dao *MessagesDao) SelectUserMessages(ctx context.Context, userId int32, messageIds []int32) (rList []*dbo.MessagesDo, err error) {
	var (
		query = "SELECT user_message_box_id,dialog_id,dialog_message_id FROM messages WHERE user_id=? AND user_message_box_id in (?)"
		rows  *sqlx.Rows
	)

	query, args, err := sqlx.In(query, userId, messageIds)
	if err != nil {
		log.Error("sqlx.In in SelectUserMessages(%d, _), error: %v", userId, err)
		return
	}

	rows, err = dao.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("queryx in SelectUserMessages(%d, _), error: %v", userId, err)
		return
	}

	defer rows.Close()

	var values []*dbo.MessagesDo
	for rows.Next() {
		v := &dbo.MessagesDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserMessages(%d, _), error: %v", userId, err)
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *MessagesDao) SelectDialogMessages(ctx context.Context, userId int32, dialogId int64, messageIds []int32) (rList []*dbo.MessagesDo, err error) {
	var (
		query = "SELECT user_message_box_id,dialog_id,dialog_message_id FROM messages WHERE user_id=? AND dialog_id=? AND dialog_message_id in (?)"
		rows  *sqlx.Rows
	)

	query, args, err := sqlx.In(query, userId, dialogId, messageIds)
	if err != nil {
		log.Error("sqlx.In in SelectUserMessages(%d, _), error: %v", userId, err)
		return
	}

	rows, err = dao.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("queryx in SelectUserMessages(%d, _), error: %v", userId, err)
		return
	}

	defer rows.Close()

	var values []*dbo.MessagesDo
	for rows.Next() {
		v := &dbo.MessagesDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserMessages(%d, _), error: %v", userId, err)
		}
		values = append(values, v)
	}
	rList = values

	return
}

// to be deleted
func (dao MessagesDao) SelectLowerBound(ctx context.Context, userId, messageBoxId int32) (lowerBound int32, err error) {
	var (
		query = "SELECT max(id) FROM messages WHERE user_id=? AND user_message_box_id<=?"
	)

	err = dao.db.Get(ctx, &lowerBound, query, userId, messageBoxId)
	if err != nil {
		log.Errorf("get in SelectLowerBound(%d, %d), error: %v", userId, messageBoxId, err)
	}
	return
}

func (dao MessagesDao) SelectUpperBound(ctx context.Context, userId, messageBoxId int32) (upperBound int32, err error) {
	var (
		query = "SELECT min(id) FROM messages WHERE user_id=? AND user_message_box_id>=?"
	)

	err = dao.db.Get(ctx, &upperBound, query, userId, messageBoxId)
	if err != nil {
		log.Errorf("get in SelectUpperBound(%d, %d), error: %v", userId, messageBoxId, err)
	}
	return
}

func (dao MessagesDao) SelectLowerBoundById(ctx context.Context, userId, id int32) (lowerBound int32, err error) {
	var (
		query = "SELECT max(id) FROM messages WHERE user_id=? AND id<=?"
	)

	err = dao.db.Get(ctx, &lowerBound, query, userId, id)
	if err != nil {
		log.Errorf("get in SelectLowerBound(%d, %d), error: %v", userId, id, err)
	}
	return
}

func (dao MessagesDao) SelectUserSentMessages(ctx context.Context, userId, lower, upper int32) (rList []*dbo.MessagesDo, err error) {
	var (
		query = "SELECT id,user_message_box_id,dialog_id,dialog_message_id FROM messages WHERE user_id=? AND id BETWEEN ? and ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, userId, lower, upper)

	if err != nil {
		log.Errorf("queryx in SelectUserSentMessages(%d, %d, %d), error: %v", userId, lower, upper, err)
		return
	}

	var values []*dbo.MessagesDo
	for rows.Next() {
		v := &dbo.MessagesDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserSentMessages(%d, %d, %d), error: %v", userId, lower, upper, err)
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao MessagesDao) SelectUserReceivedMessages(ctx context.Context, userId, lower, upper int32) (rList []*dbo.MessagesDo, err error) {
	var (
		query = "SELECT id,user_message_box_id,dialog_id,dialog_message_id FROM messages WHERE peer_type=2 AND peer_id=? AND id BETWEEN ? and ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, userId, lower, upper)

	if err != nil {
		log.Errorf("queryx in SelectUserSentMessages(%d, %d, %d), error: %v", userId, lower, upper, err)
		return
	}

	var values []*dbo.MessagesDo
	for rows.Next() {
		v := &dbo.MessagesDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserSentMessages(%d, %d, %d), error: %v", userId, lower, upper, err)
		}
		values = append(values, v)
	}
	rList = values

	return
}
