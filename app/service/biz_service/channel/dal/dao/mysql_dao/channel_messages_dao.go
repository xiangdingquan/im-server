package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelMessagesDAO struct {
	db *sqlx.DB
}

func NewChannelMessagesDAO(db *sqlx.DB) *ChannelMessagesDAO {
	return &ChannelMessagesDAO{db}
}

func (dao *ChannelMessagesDAO) SelectByMessageIdList(ctx context.Context, user_id int32, channel_id int32, idList []int32) (rList []dataobject.ChannelMessagesDO, err error) {
	var (
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages cm where channel_id = ? and deleted = 0 and channel_message_id in (?) and (has_remove = 0 or not exists (select message_id from channel_messages_delete where user_id = ? and channel_id = cm.channel_id and message_id = cm.channel_message_id and deleted = 0)) and (has_dm = 0 or exists (select message_id from channel_message_visibles where user_id = ? and channel_id = cm.channel_id and message_id = cm.channel_message_id)) order by channel_message_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if user_id > 0 {
		query, a, err = sqlx.In(query, channel_id, idList, user_id, user_id)
		if err != nil {
			log.Errorf("sqlx.In in SelectByMessageIdList(_), error: %v", err)
			return
		}
	} else {
		query = "select id, channel_id, channel_message_id, sender_user_id, random_id, message_data_id, message_type, message_data, message, media_type, media_unread, has_media_unread, edit_message, edit_date, ttl_seconds, has_remove, has_dm, views, `date` from channel_messages where channel_id = ? and deleted = 0 and channel_message_id in (?) order by channel_message_id desc"
		query, a, err = sqlx.In(query, channel_id, idList)
		if err != nil {
			log.Errorf("sqlx.In in SelectByMessageIdList(_), error: %v", err)
			return
		}
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
