package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ConversationsDAO struct {
	db *sqlx.DB
}

func NewConversationsDAO(db *sqlx.DB) *ConversationsDAO {
	return &ConversationsDAO{db}
}

func (dao *ConversationsDAO) InsertIgnore(ctx context.Context, do *dataobject.ConversationsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into conversations(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, :draft_message_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

func (dao *ConversationsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.ConversationsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into conversations(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, :draft_message_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

func (dao *ConversationsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ConversationsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into conversations(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, '', :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)"
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

func (dao *ConversationsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ConversationsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into conversations(user_id, peer_type, peer_id, top_message, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :top_message, :unread_count, :unread_mentions_count, '', :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)"
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

func (dao *ConversationsDAO) SelectPinnedDialogs(ctx context.Context, user_id int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and is_pinned = 1 and deleted = 0 order by top_message desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
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

func (dao *ConversationsDAO) SelectByPeer(ctx context.Context, user_id int32, peer_id int32) (rValue *dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_id)

	if err != nil {
		log.Errorf("queryx in SelectByPeer(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ConversationsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByPeer(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ConversationsDAO) SelectUserDialogs(ctx context.Context, user_id int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectUserDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserDialogs(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ConversationsDAO) SelectByOffsetId(ctx context.Context, user_id int32, top_message int32, limit int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and deleted = 0 and top_message > 0 and top_message < ? order by top_message desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, top_message, limit)

	if err != nil {
		log.Errorf("queryx in SelectByOffsetId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
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

func (dao *ConversationsDAO) SelectExcludePinnedByOffsetId(ctx context.Context, user_id int32, top_message int32, limit int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and is_pinned = 0 and deleted = 0 and top_message > 0 and top_message < ? order by top_message desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, top_message, limit)

	if err != nil {
		log.Errorf("queryx in SelectExcludePinnedByOffsetId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
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

func (dao *ConversationsDAO) SelectByPeerType(ctx context.Context, user_id int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and peer_type = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectByPeerType(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByPeerType(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ConversationsDAO) SelectExcludePinnedByPeerType(ctx context.Context, user_id int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and peer_type = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectExcludePinnedByPeerType(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectExcludePinnedByPeerType(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ConversationsDAO) SelectListByPeerList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, date2 from conversations where user_id = ? and peer_type = ? and deleted = 0 and peer_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByPeerList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByPeerList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByPeerList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ConversationsDAO) UpdateUnreadByPeer(ctx context.Context, read_inbox_max_id int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_inbox_max_id, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdateUnreadByPeerTx(tx *sqlx.Tx, read_inbox_max_id int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_inbox_max_id, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdateReadOutboxMaxIdByPeer(ctx context.Context, read_outbox_max_id int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set read_outbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_outbox_max_id, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdateReadOutboxMaxIdByPeerTx(tx *sqlx.Tx, read_outbox_max_id int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set read_outbox_max_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_outbox_max_id, user_id, peer_id)

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

func (dao *ConversationsDAO) SaveDraft(ctx context.Context, draft_message_data string, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set draft_type = 2, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, draft_message_data, user_id, peer_id)

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

func (dao *ConversationsDAO) SaveDraftTx(tx *sqlx.Tx, draft_message_data string, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set draft_type = 2, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, draft_message_data, user_id, peer_id)

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

func (dao *ConversationsDAO) SelectDraftList(ctx context.Context, user_id int32) (rList []dataobject.ConversationsDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, draft_type, draft_message_data from conversations where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectDraftList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ConversationsDO
	for rows.Next() {
		v := dataobject.ConversationsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectDraftList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ConversationsDAO) ClearDraft(ctx context.Context, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set draft_type = 0, draft_message_data = '' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in ClearDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearDraft(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) ClearDraftTx(tx *sqlx.Tx, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set draft_type = 0, draft_message_data = '' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in ClearDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearDraft(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdatePinned(ctx context.Context, is_pinned int8, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set is_pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, is_pinned, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdatePinnedTx(tx *sqlx.Tx, is_pinned int8, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set is_pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, is_pinned, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdateDialog(ctx context.Context, top_message int32, unreadCount int32, unreadMentionCount int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set top_message = ?, unread_count = unread_count + ?, unread_mentions_count = unread_mentions_count + ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, top_message, unreadCount, unreadMentionCount, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDialog(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdateDialogTx(tx *sqlx.Tx, top_message int32, unreadCount int32, unreadMentionCount int32, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set top_message = ?, unread_count = unread_count + ?, unread_mentions_count = unread_mentions_count + ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, top_message, unreadCount, unreadMentionCount, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDialog(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdateUnreadMentionCount(ctx context.Context, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_mentions_count = unread_mentions_count - 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateUnreadMentionCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnreadMentionCount(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdateUnreadMentionCountTx(tx *sqlx.Tx, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_mentions_count = unread_mentions_count - 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateUnreadMentionCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUnreadMentionCount(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) Delete(ctx context.Context, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "delete from conversations where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) DeleteTx(tx *sqlx.Tx, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "delete from conversations where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdateOutboxDialog(ctx context.Context, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set unread_count = 0, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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

func (dao *ConversationsDAO) UpdateOutboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set unread_count = 0, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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

func (dao *ConversationsDAO) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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

func (dao *ConversationsDAO) UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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

func (dao *ConversationsDAO) UpdateMarkDialogUnread(ctx context.Context, unread_mark int8, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_mark = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, unread_mark, user_id, peer_id)

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

func (dao *ConversationsDAO) UpdateMarkDialogUnreadTx(tx *sqlx.Tx, unread_mark int8, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update conversations set unread_mark = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, unread_mark, user_id, peer_id)

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

func (dao *ConversationsDAO) SelectMarkDialogUnreadList(ctx context.Context, user_id int32) (rList []int32, err error) {
	var query = "select peer_id from conversations where user_id = ? and unread_mark = 1"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
	}

	return
}

func (dao *ConversationsDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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

func (dao *ConversationsDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update conversations set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_id)

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
