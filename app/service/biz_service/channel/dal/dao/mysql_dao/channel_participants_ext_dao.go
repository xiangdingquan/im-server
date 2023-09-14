package mysql_dao

import (
	"context"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

func (dao *ChannelParticipantsDAO) SelectByGTOffsetDate2(ctx context.Context, userId, date2 int32) (rList []int32, err error) {
	var query = `
           SELECT
               channels.id
           FROM
               channels, channel_participants
           WHERE
               channel_participants.user_id = ? AND channels.date2 > ? AND channels.id = channel_participants.channel_id AND channels.deleted = 0`
	err = dao.db.Select(ctx, &rList, query, userId, date2)

	if err != nil {
		log.Errorf("select in SelectByGTOffsetDate2(_), error: %v", err)
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectDialogs(ctx context.Context, userId int32, folderId int32) (rList []dataobject.ChannelParticipantsExtDO, err error) {
	var (
		query = `
			SELECT 
				channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, rank, state, channels.date as date, channels.date2 as date2, channels.top_message as top_message, channels.pts as pts
			FROM 
				channel_participants, channels
			WHERE 
				channel_participants.user_id = ? AND channel_participants.folder_id = ? AND (channel_participants.state = 0 OR channel_participants.state = 2) AND channels.id = channel_participants.channel_id`

		rows *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId, folderId)

	if err != nil {
		log.Errorf("queryx in SelectDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsExtDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsExtDO{}
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

func (dao *ChannelParticipantsDAO) SelectExcludePinnedDialogs(ctx context.Context, userId int32, pinned string, folderId int32) (rList []dataobject.ChannelParticipantsExtDO, err error) {
	var (
		query = `
			SELECT 
				channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, rank, state, channels.date as date, channels.date2 as date2, channels.top_message as top_message, channels.pts as pts
			FROM 
				channel_participants, channels
           WHERE
				channel_participants.user_id = ? AND ? = 0 AND channel_participants.folder_id = ? AND (channel_participants.state = 0 OR channel_participants.state = 2) AND channels.id = channel_participants.channel_id`
		rows *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId, pinned, folderId)

	if err != nil {
		log.Errorf("queryx in SelectExcludePinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsExtDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsExtDO{}
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

func (dao *ChannelParticipantsDAO) SelectPinnedDialogs(ctx context.Context, user_id int32, folder_id int32) (rList []dataobject.ChannelParticipantsExtDO, err error) {
	var (
		query = `
			SELECT 
				channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, rank, state, channels.date as date, channels.date2 as date2, channels.top_message as top_message, channels.pts as pts
			FROM 
				channel_participants, channels
           WHERE
				channel_participants.user_id = ? AND channel_participants.folder_id = ? AND channel_participants.is_pinned = 1 AND (channel_participants.state = 0 OR channel_participants.state = 2) AND channels.id = channel_participants.channel_id`

		rows *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, folder_id)

	if err != nil {
		log.Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelParticipantsExtDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsExtDO{}
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

func (dao *ChannelParticipantsDAO) SelectExtListByChannelIdList(ctx context.Context, userId int32, idList []int32) (rList []dataobject.ChannelParticipantsExtDO, err error) {
	var (
		query = `
			SELECT 
				channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, rank, state, channels.date as date, channels.date2 as date2, channels.top_message as top_message, channels.pts as pts
			FROM 
				channel_participants, channels
           WHERE
				channel_participants.user_id = ? AND (channel_participants.state = 0 OR channel_participants.state = 2) and channel_participants.channel_id IN (?) AND channels.id = channel_participants.channel_id`

		a    []interface{}
		rows *sqlx.Rows
	)
	query, a, err = sqlx.In(query, userId, idList)
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

	var values []dataobject.ChannelParticipantsExtDO
	for rows.Next() {
		v := dataobject.ChannelParticipantsExtDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByUserIdChannelIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelParticipantsDAO) SelectExtByChannelIdUserId(ctx context.Context, channelId int32, userId int32) (rValue *dataobject.ChannelParticipantsExtDO, err error) {
	var (
		query = `
			SELECT 
				channel_id, user_id, is_creator, is_pinned, order_pinned, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, folder_order_pinned, inviter_user_id, promoted_by, admin_rights, hidden_prehistory, hidden_prehistory_message_id, kicked_by, banned_rights, banned_until_date, migrated_from_max_id, available_min_id, available_min_pts, state, channels.date as date, channels.date2 as date2, channels.top_message as top_message, channels.pts as pts
			FROM 
				channel_participants, channels
           WHERE
				channel_participants.channel_id = ? AND channel_participants.user_id = ?`

		rows *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channelId, userId)

	if err != nil {
		log.Errorf("queryx in SelectByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelParticipantsExtDO{}
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

func (dao *ChannelParticipantsDAO) SelectUnreadCountByChannelIdUserId(ctx context.Context, channelId int32, userId int32) (rValue int32) {
	var (
		query = `SELECT
					COUNT(id) as c
				FROM
					channel_messages
				WHERE
					channel_id = ? AND
					channel_message_id > (SELECT read_inbox_max_id FROM channel_participants WHERE channel_id = ? AND user_id = ?)`
		err error
	)

	err = dao.db.Get(ctx, &rValue, query, channelId, channelId, userId)

	if err != nil {
		log.Errorf("queryx in SelectUnreadCountByChannelIdUserId(_), error: %v", err)
		return
	}

	return
}

func (dao *ChannelParticipantsDAO) SelectUnreadMentionsCount(ctx context.Context, channelId int32, userId int32, msgId int32) (rValue int32) {
	var (
		query = `SELECT
					COUNT(id) AS c
				FROM
					channel_unread_mentions
				WHERE
					user_id = ? AND channel_id = ? AND mentioned_message_id > ? AND deleted = 0`
		err error
	)

	err = dao.db.Get(ctx, &rValue, query, userId, channelId, msgId)

	if err != nil {
		log.Errorf("queryx in SelectUnreadMentionCount(_), error: %v", err)
		return
	}

	return
}
