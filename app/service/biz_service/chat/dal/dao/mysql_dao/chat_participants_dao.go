package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChatParticipantsDAO struct {
	db *sqlx.DB
}

func NewChatParticipantsDAO(db *sqlx.DB) *ChatParticipantsDAO {
	return &ChatParticipantsDAO{db}
}

func (dao *ChatParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2)"
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

func (dao *ChatParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2)"
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

func (dao *ChatParticipantsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *ChatParticipantsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *ChatParticipantsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *ChatParticipantsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_type, draft_message_data, date2) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, 0, '', :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectList(ctx context.Context, chat_id int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, chat_id)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
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

func (dao *ChatParticipantsDAO) SelectByParticipantId(ctx context.Context, chat_id int32, user_id int32) (rValue *dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where chat_id = ? and user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, chat_id, user_id)

	if err != nil {
		log.Errorf("queryx in SelectByParticipantId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatParticipantsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByParticipantId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChatParticipantsDAO) Update(ctx context.Context, participant_type int8, inviter_user_id int32, invited_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_type, inviter_user_id, invited_at, id)

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

func (dao *ChatParticipantsDAO) UpdateTx(tx *sqlx.Tx, participant_type int8, inviter_user_id int32, invited_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_type, inviter_user_id, invited_at, id)

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

func (dao *ChatParticipantsDAO) UpdateKicked(ctx context.Context, kicked_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 2, kicked_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, kicked_at, id)

	if err != nil {
		log.Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateKickedTx(tx *sqlx.Tx, kicked_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 2, kicked_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, kicked_at, id)

	if err != nil {
		log.Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateLeft(ctx context.Context, left_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 1, left_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, left_at, id)

	if err != nil {
		log.Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateLeftTx(tx *sqlx.Tx, left_at int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 1, left_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, left_at, id)

	if err != nil {
		log.Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdatePinnedMsgId(ctx context.Context, pinned_msg_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, pinned_msg_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdatePinnedMsgIdTx(tx *sqlx.Tx, pinned_msg_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned_msg_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateUnPinnedMsgId(ctx context.Context, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = 0 where chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateUnPinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnPinnedMsgId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateUnPinnedMsgIdTx(tx *sqlx.Tx, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = 0 where chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateUnPinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnPinnedMsgId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateParticipantType(ctx context.Context, participant_type int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_type, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateParticipantTypeTx(tx *sqlx.Tx, participant_type int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_type, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateOutboxDialog(ctx context.Context, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = 0, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateOutboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = 0, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateUnreadByPeer(ctx context.Context, read_inbox_max_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_inbox_max_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateUnreadByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnreadByPeer(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateUnreadByPeerTx(tx *sqlx.Tx, read_inbox_max_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_inbox_max_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateUnreadByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnreadByPeer(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateReadOutboxMaxIdByPeer(ctx context.Context, read_outbox_max_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set read_outbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_outbox_max_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateReadOutboxMaxIdByPeerTx(tx *sqlx.Tx, read_outbox_max_id int32, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set read_outbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_outbox_max_id, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectListByChatIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and chat_id in (?) order by top_message desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByChatIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByChatIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByChatIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) UpdatePinned(ctx context.Context, is_pinned int8, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, is_pinned, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdatePinnedTx(tx *sqlx.Tx, is_pinned int8, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, is_pinned, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = unread_count + 1, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = unread_count + 1, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateMarkDialogUnread(ctx context.Context, unread_mark int8, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_mark = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, unread_mark, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateMarkDialogUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMarkDialogUnread(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateMarkDialogUnreadTx(tx *sqlx.Tx, unread_mark int8, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_mark = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, unread_mark, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in UpdateMarkDialogUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMarkDialogUnread(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectMarkDialogUnreadList(ctx context.Context, user_id int32) (rList []int32, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, chat_id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SaveDraft(ctx context.Context, draft_type int8, draft_message_data string, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = ?, draft_message_data = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, draft_type, draft_message_data, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SaveDraftTx(tx *sqlx.Tx, draft_type int8, draft_message_data string, user_id int32, chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = ?, draft_message_data = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, draft_type, draft_message_data, user_id, chat_id)

	if err != nil {
		log.Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectAllDrafts(ctx context.Context, user_id int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = ? and draft_type > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectAllDrafts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAllDrafts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) ClearAllDrafts(ctx context.Context, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 0, draft_message_data = '' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		log.Errorf("exec in ClearAllDrafts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearAllDrafts(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) ClearAllDraftsTx(tx *sqlx.Tx, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 0, draft_message_data = '' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

	if err != nil {
		log.Errorf("exec in ClearAllDrafts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearAllDrafts(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectDialogs(ctx context.Context, user_id int32, folder_id int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and folder_id = ? and (state = 0 or state = 2)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, folder_id)

	if err != nil {
		log.Errorf("queryx in SelectDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) SelectExcludePinnedDialogs(ctx context.Context, user_id int32, pinned string, folder_id int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and ? = 0 and folder_id = ? and (state = 0 or state = 2)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, pinned, folder_id)

	if err != nil {
		log.Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectExcludePinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) SelectByOffsetId(ctx context.Context, user_id int32, top_message int32, limit int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and (state = 0 or state = 2) and top_message < ? order by top_message desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, top_message, limit)

	if err != nil {
		log.Errorf("queryx in SelectByOffsetId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByOffsetId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) SelectExcludePinnedByOffsetId(ctx context.Context, user_id int32, top_message int32, limit int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and is_pinned = 0 and (state = 0 or state = 2) and top_message > 0 and top_message < ? order by top_message desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, top_message, limit)

	if err != nil {
		log.Errorf("queryx in SelectExcludePinnedByOffsetId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectExcludePinnedByOffsetId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) SelectPinnedDialogs(ctx context.Context, user_id int32, folder_id int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, chat_id, user_id, participant_type, is_pinned, order_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, invited_at, kicked_at, left_at, state, date2 from chat_participants where user_id = ? and folder_id = ? and is_pinned = 1 and (state = 0 or state = 2) order by top_message desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, folder_id)

	if err != nil {
		log.Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPinnedDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) UpdateUnPinnedList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = 0, order_pinned = 0 where user_id = ? and folder_id = 0 and is_pinned = 1 and chat_id not in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateUnPinnedList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateUnPinnedList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnPinnedList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateUnPinnedListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = 0, order_pinned = 0 where user_id = ? and folder_id = 0 and is_pinned = 1 and chat_id not in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateUnPinnedList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateUnPinnedList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnPinnedList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateFolderUnPinnedList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set folder_pinned = 0, folder_order_pinned = 0 where user_id = ? and folder_id = 1 and is_pinned = 1 and chat_id not in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateFolderUnPinnedList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateFolderUnPinnedList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFolderUnPinnedList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateFolderUnPinnedListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set folder_pinned = 0, folder_order_pinned = 0 where user_id = ? and folder_id = 1 and is_pinned = 1 and chat_id not in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateFolderUnPinnedList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateFolderUnPinnedList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFolderUnPinnedList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateFolderId(ctx context.Context, folder_id int32, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set folder_id = ? where user_id = ? and chat_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, folder_id, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateFolderId(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in UpdateFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFolderId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) UpdateFolderIdTx(tx *sqlx.Tx, folder_id int32, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set folder_id = ? where user_id = ? and chat_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, folder_id, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in UpdateFolderId(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in UpdateFolderId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFolderId(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectUsersChatIdList(ctx context.Context, idList []int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (?) and chats.id = chat_participants.chat_id and chats.deactivated = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectUsersChatIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectUsersChatIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatParticipantsDO
	for rows.Next() {
		v := dataobject.ChatParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUsersChatIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChatParticipantsDAO) SelectMyAdminList(ctx context.Context, user_id int32) (rList []int32, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and participant_type = 1 and state = 0"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectMyAdminList(_), error: %v", err)
	}

	return
}

func (dao *ChatParticipantsDAO) SelectMyAllList(ctx context.Context, user_id int32) (rList []int32, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and state = 0"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectMyAllList(_), error: %v", err)
	}

	return
}
