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

type ChannelParticipantsDAO struct {
	db *sqlx.DB
}

func NewChannelParticipantsDAO(db *sqlx.DB) *ChannelParticipantsDAO {
	return &ChannelParticipantsDAO{db}
}

func (dao *ChannelParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, is_creator, draft_message_data, inviter_user_id, state, date2, available_min_id, available_min_pts) values (:channel_id, :user_id, :is_creator, '', :inviter_user_id, :state, :date2, :available_min_id, :available_min_pts)"
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

func (dao *ChannelParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, is_creator, draft_message_data, inviter_user_id, state, date2, available_min_id, available_min_pts) values (:channel_id, :user_id, :is_creator, '', :inviter_user_id, :state, :date2, :available_min_id, :available_min_pts)"
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

func (dao *ChannelParticipantsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, is_creator, draft_message_data, inviter_user_id, migrated_from_max_id, state, date2, available_min_id, available_min_pts) values (:channel_id, :user_id, :is_creator, '', :inviter_user_id, :migrated_from_max_id, :state, :date2, :available_min_id, :available_min_pts)"
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

func (dao *ChannelParticipantsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, is_creator, draft_message_data, inviter_user_id, migrated_from_max_id, state, date2, available_min_id, available_min_pts) values (:channel_id, :user_id, :is_creator, '', :inviter_user_id, :migrated_from_max_id, :state, :date2, :available_min_id, :available_min_pts)"
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

func (dao *ChannelParticipantsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, inviter_user_id, state, date2, draft_message_data, available_min_id, available_min_pts, kicked_by, banned_rights, banned_until_date) values (:channel_id, :user_id, :inviter_user_id, :state, :date2, '', :available_min_id, :available_min_pts, :kicked_by, :banned_rights, :banned_until_date) on duplicate key update inviter_user_id = values(inviter_user_id), available_min_id = values(available_min_id), available_min_pts = values(available_min_pts), state = values(state), date2 = values(date2), draft_type = 0, draft_message_data = '', promoted_by = 0, admin_rights = 0, kicked_by = values(kicked_by), banned_rights = values(banned_rights), banned_until_date = values(banned_until_date), is_pinned = 0"
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

func (dao *ChannelParticipantsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ChannelParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, inviter_user_id, state, date2, draft_message_data, available_min_id, available_min_pts, kicked_by, banned_rights, banned_until_date) values (:channel_id, :user_id, :inviter_user_id, :state, :date2, '', :available_min_id, :available_min_pts, :kicked_by, :banned_rights, :banned_until_date) on duplicate key update inviter_user_id = values(inviter_user_id), available_min_id = values(available_min_id), available_min_pts = values(available_min_pts), state = values(state), date2 = values(date2), draft_type = 0, draft_message_data = '', promoted_by = 0, admin_rights = 0, kicked_by = values(kicked_by), banned_rights = values(banned_rights), banned_until_date = values(banned_until_date), is_pinned = 0"
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

func (dao *ChannelParticipantsDAO) SelectByChannelId(ctx context.Context, channel_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id)

	if err != nil {
		log.Errorf("queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByChannelId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectByUserIdList(ctx context.Context, channel_id int32, idList []int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ? and user_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, channel_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByUserIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByUserIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByUserIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectByUserId(ctx context.Context, channel_id int32, user_id int32) (rValue *dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ? and user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, user_id)

	if err != nil {
		log.Errorf("queryx in SelectByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelParticipantsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectAdminAndBannedList(ctx context.Context, channel_id int32, banned_until_date int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ? and (state = 0 and (is_creator = 1 or admin_rights > 0) or (banned_rights > 0 and banned_until_date <= ?))"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, banned_until_date)

	if err != nil {
		log.Errorf("queryx in SelectAdminAndBannedList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAdminAndBannedList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectAdminList(ctx context.Context, channel_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ? and state = 0 and (is_creator = 1 or admin_rights > 0)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id)

	if err != nil {
		log.Errorf("queryx in SelectAdminList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAdminList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectBannedList(ctx context.Context, channel_id int32, banned_until_date int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where channel_id = ? and banned_rights > 0 and banned_until_date <= ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, banned_until_date)

	if err != nil {
		log.Errorf("queryx in SelectBannedList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectBannedList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) Update(ctx context.Context, cMap map[string]interface{}, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channel_participants set %s where channel_id = ? and user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, channel_id)
	aValues = append(aValues, user_id)

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

func (dao *ChannelParticipantsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channel_participants set %s where channel_id = ? and user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, channel_id)
	aValues = append(aValues, user_id)

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

func (dao *ChannelParticipantsDAO) UpdateLeave(ctx context.Context, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set state = 1, date2 = ?, promoted_by = 0, admin_rights = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLeave(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLeave(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateLeaveTx(tx *sqlx.Tx, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set state = 1, date2 = ?, promoted_by = 0, admin_rights = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLeave(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLeave(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateBannedRights(ctx context.Context, kicked_by int32, banned_rights int32, banned_until_date int32, state int8, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set kicked_by = ?, banned_rights = ?, banned_until_date = ?, state = ?, date2 = ?, promoted_by = 0, admin_rights = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, kicked_by, banned_rights, banned_until_date, state, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateBannedRightsTx(tx *sqlx.Tx, kicked_by int32, banned_rights int32, banned_until_date int32, state int8, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set kicked_by = ?, banned_rights = ?, banned_until_date = ?, state = ?, date2 = ?, promoted_by = 0, admin_rights = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, kicked_by, banned_rights, banned_until_date, state, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateAdminRights(ctx context.Context, promoted_by int32, admin_rights int32, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set promoted_by = ?, admin_rights = ?, date2 = ?, state = 0, kicked_by = 0, banned_rights = 0, banned_until_date = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, promoted_by, admin_rights, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateAdminRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateAdminRightsTx(tx *sqlx.Tx, promoted_by int32, admin_rights int32, date2 int32, channel_id int32, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set promoted_by = ?, admin_rights = ?, date2 = ?, state = 0, kicked_by = 0, banned_rights = 0, banned_until_date = 0 where channel_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, promoted_by, admin_rights, date2, channel_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateAdminRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateReadInboxMaxId(ctx context.Context, read_inbox_max_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set read_inbox_max_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_inbox_max_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateReadInboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadInboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateReadInboxMaxIdTx(tx *sqlx.Tx, read_inbox_max_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set read_inbox_max_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_inbox_max_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateReadInboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadInboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateReadOutboxMaxId(ctx context.Context, read_inbox_max_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set read_outbox_max_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_inbox_max_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateReadOutboxMaxIdTx(tx *sqlx.Tx, read_inbox_max_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set read_outbox_max_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_inbox_max_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectListByUserId(ctx context.Context, user_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where user_id = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectListByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByUserId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectListByChannelIdList(ctx context.Context, user_id int32, idList []int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where user_id = ? and (state = 0 or state = 2) and channel_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByChannelIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByChannelIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByChannelIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) UpdatePinned(ctx context.Context, is_pinned int8, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set is_pinned = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, is_pinned, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) UpdatePinnedTx(tx *sqlx.Tx, is_pinned int8, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set is_pinned = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, is_pinned, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) UpdateMarkDialogUnread(ctx context.Context, unread_mark int8, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set unread_mark = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, unread_mark, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) UpdateMarkDialogUnreadTx(tx *sqlx.Tx, unread_mark int8, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set unread_mark = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, unread_mark, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) SelectMarkDialogUnreadList(ctx context.Context, user_id int32) (rList []int32, err error) {
	var query = "select channel_id from channel_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectMyAdminPublicList(ctx context.Context, user_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where user_id = ? and state = 0 and (is_creator = 1 or admin_rights > 0)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectMyAdminPublicList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectMyAdminPublicList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectLeftList(ctx context.Context, user_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, date2 from channel_participants where user_id = ? and state > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectLeftList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectLeftList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) UpdateAvailableMinId(ctx context.Context, available_min_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set available_min_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, available_min_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateAvailableMinId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAvailableMinId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) UpdateAvailableMinIdTx(tx *sqlx.Tx, available_min_id int32, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set available_min_id = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, available_min_id, user_id, channel_id)

	if err != nil {
		log.Errorf("exec in UpdateAvailableMinId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAvailableMinId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectAvailableMinId(ctx context.Context, user_id int32, channel_id int32) (rValue int32, err error) {
	var query = "select available_min_id from channel_participants where user_id = ? and channel_id = ?"
	err = dao.db.Get(ctx, &rValue, query, user_id, channel_id)

	if err != nil {
		log.Errorf("get in SelectAvailableMinId(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectAvailableMinPts(ctx context.Context, user_id int32, channel_id int32) (rValue int32, err error) {
	var query = "select available_min_pts from channel_participants where user_id = ? and channel_id = ?"
	err = dao.db.Get(ctx, &rValue, query, user_id, channel_id)

	if err != nil {
		log.Errorf("get in SelectAvailableMinPts(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SaveDraft(ctx context.Context, draft_type int8, draft_message_data string, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set draft_type = ?, draft_message_data = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, draft_type, draft_message_data, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) SaveDraftTx(tx *sqlx.Tx, draft_type int8, draft_message_data string, user_id int32, channel_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set draft_type = ?, draft_message_data = ? where user_id = ? and channel_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, draft_type, draft_message_data, user_id, channel_id)

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

func (dao *ChannelParticipantsDAO) SelectAllDrafts(ctx context.Context, user_id int32) (rList []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select user_id, channel_id, draft_type, draft_message_data from channel_participants where user_id = ? and draft_type > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectAllDrafts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsDO{}
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

func (dao *ChannelParticipantsDAO) ClearAllDrafts(ctx context.Context, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set draft_type = 0, draft_message_data = '' where user_id = ? and draft_type = 2"
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

func (dao *ChannelParticipantsDAO) ClearAllDraftsTx(tx *sqlx.Tx, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set draft_type = 0, draft_message_data = '' where user_id = ? and draft_type = 2"
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

func (dao *ChannelParticipantsDAO) UpdateUnPinnedList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set is_pinned = 0, order_pinned = 0 where user_id = ? and folder_id = 0 and is_pinned = 1 and channel_id not in (?)"
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

func (dao *ChannelParticipantsDAO) UpdateUnPinnedListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set is_pinned = 0, order_pinned = 0 where user_id = ? and folder_id = 0 and is_pinned = 1 and channel_id not in (?)"
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

func (dao *ChannelParticipantsDAO) UpdateFolderUnPinnedList(ctx context.Context, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set folder_pinned = 0, folder_order_pinned = 0 where user_id = ? and folder_id = 1 and is_pinned = 1 and channel_id not in (?)"
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

func (dao *ChannelParticipantsDAO) UpdateFolderUnPinnedListTx(tx *sqlx.Tx, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set folder_pinned = 0, folder_order_pinned = 0 where user_id = ? and folder_id = 1 and is_pinned = 1 and channel_id not in (?)"
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

func (dao *ChannelParticipantsDAO) UpdateFolderId(ctx context.Context, folder_id int32, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set folder_id = ? where user_id = ? and channel_id in (?)"
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

func (dao *ChannelParticipantsDAO) UpdateFolderIdTx(tx *sqlx.Tx, folder_id int32, user_id int32, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update channel_participants set folder_id = ? where user_id = ? and channel_id in (?)"
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
