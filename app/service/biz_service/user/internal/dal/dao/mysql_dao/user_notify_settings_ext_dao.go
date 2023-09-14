package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

func (dao *UserNotifySettingsDAO) InsertOrUpdateExt(ctx context.Context, userId, peerType, peerId int32, cMap map[string]interface{}) (lastInsertId, rowsAffected int64, err error) {
	var (
		s1 []string
		s2 []string
		s3 []string
		r  sql.Result
	)

	for k := range cMap {
		s1 = append(s1, k)
		s2 = append(s2, ":"+k)
		s3 = append(s3, fmt.Sprintf("%s = :%s", k, k))
	}

	cMap["user_id"] = userId
	cMap["peer_type"] = peerType
	cMap["peer_id"] = peerId

	query := `
		insert into user_notify_settings
			(user_id, peer_type, peer_id, %s) 
		values 
			(:user_id, :peer_type, :peer_id, %s)
		on duplicate key update %s, deleted = 0`

	ss := fmt.Sprintf(query, strings.Join(s1, ","), strings.Join(s2, ","), strings.Join(s3, ", "))
	r, err = dao.db.NamedExec(ctx, ss, cMap)
	if err != nil {
		log.Errorf("namedExec (%s) in InsertOrUpdate(%v), error: %v", query, cMap, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
	}

	return
}

func (dao *UserNotifySettingsDAO) InsertOrUpdateExtTx(tx *sqlx.Tx, userId int32, peerType int8, peerId int32, cMap map[string]interface{}) (lastInsertId, rowsAffected int64, err error) {
	var (
		s1 []string
		s2 []string
		s3 []string
		r  sql.Result
	)

	for k := range cMap {
		s1 = append(s1, k)
		s2 = append(s2, ":"+k)
		s3 = append(s3, fmt.Sprintf("%s = (%s)", k, k))
	}

	cMap["user_id"] = userId
	cMap["peer_type"] = peerType
	cMap["peer_id"] = peerId

	query := `
		insert into user_notify_settings
			(user_id, peer_type, peer_id, %s) 
		values 
			(:user_id, :peer_type, :peer_id, %s)
		on duplicate key update %s, deleted = 0`

	r, err = tx.NamedExec(fmt.Sprintf(query, strings.Join(s1, ","), strings.Join(s2, ","), strings.Join(s3, ", ")), cMap)
	if err != nil {
		log.Errorf("namedExec (%s) in InsertOrUpdate(%v), error: %v", query, cMap, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
	}

	return
}

func (dao *UserNotifySettingsDAO) SelectList(ctx context.Context, userId int32, userIdList, chatIdList, channelIdList []int32) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		rows *sqlx.Rows
		qVs  []string
		args []interface{}
		a    []interface{}
	)

	if len(userIdList) == 0 && len(chatIdList) == 0 && len(channelIdList) == 0 {
		log.Errorf("idList empty")
		return
	}

	query := `
	select 
		id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound 
	from 
		user_notify_settings 
	where 
		user_id = ? AND deleted = 0 AND (%s)`

	args = append(args, userId)
	if len(userIdList) > 0 {
		qVs = append(qVs, "(peer_type = 2 AND peer_id IN (?)) ")
		args = append(args, userIdList)
	}
	if len(chatIdList) > 0 {
		qVs = append(qVs, "(peer_type = 3 AND peer_id IN (?)) ")
		args = append(args, chatIdList)
	}
	if len(channelIdList) > 0 {
		qVs = append(qVs, "(peer_type = 4 AND peer_id IN (?)) ")
		args = append(args, channelIdList)
	}
	query = fmt.Sprintf(query, strings.Join(qVs, " OR "))

	query, a, err = sqlx.In(query, args...)
	if err != nil {
		log.Errorf("sqlx.In (%s) in SelectNotifySettingsList(_), error: %v", query, err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectNotifySettingsList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserNotifySettingsDO
	for rows.Next() {
		v := dataobject.UserNotifySettingsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectNotifySettingsList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
