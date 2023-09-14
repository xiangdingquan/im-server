package mysqldao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// AvcallsDAO .
type AvcallsDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewAvcallsDAO(db *sqlx.DB) *AvcallsDAO {
	return &AvcallsDAO{db}
}

// Insert .
func (dao *AvcallsDAO) Insert(ctx context.Context, do *dbo.AvcallDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `avcalls`(channel_name, chat_id, owner_uid, member_uids, is_video, is_meet) VALUES (:channel_name, :chat_id, :owner_uid, :member_uids, :is_video, :is_meet)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
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

// InsertTx .
func (dao *AvcallsDAO) InsertTx(tx *sqlx.Tx, do *dbo.AvcallDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `avcalls`(channel_name, chat_id, owner_uid, member_uids, is_video, is_meet) VALUES (:channel_name, :chat_id, :owner_uid, :member_uids, :is_video, :is_meet)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(%v)_error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v)_error: %v", do, err)
	}

	return
}

// SelectByCallID .
func (dao *AvcallsDAO) SelectByCallID(ctx context.Context, callID uint32) (rValue *dbo.AvcallDO, err error) {
	var (
		query = "SELECT id, channel_name, chat_id, owner_uid, IFNULL(`member_uids`,'[]') member_uids, start_at, is_video, is_meet, close_at FROM `avcalls` WHERE id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, callID)

	if err != nil {
		log.Errorf("queryx in SelectByCallID(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.AvcallDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByCallID(_), error: %v", err)
		} else {
			rValue = do
		}
		do.CreateAt = (uint32)(do.CreateTime.Unix())
		json.Unmarshal([]byte(do.MemberInfo), &do.Members)
	}

	return
}

// SelectByChannel .
func (dao *AvcallsDAO) SelectByChannel(ctx context.Context, channelName string) (rValue *dbo.AvcallDO, err error) {
	var (
		query = "SELECT id, channel_name, chat_id, owner_uid, IFNULL(`member_uids`,'[]') member_uids, start_at, is_video, is_meet, close_at FROM `avcalls` WHERE channel_name = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channelName)

	if err != nil {
		log.Errorf("queryx in SelectByChannel(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.AvcallDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByChannel(_), error: %v", err)
		} else {
			rValue = do
		}
		do.CreateAt = (uint32)(do.CreateTime.Unix())
		json.Unmarshal([]byte(do.MemberInfo), &do.Members)
	}

	return
}

// UpdateWithID .
func (dao *AvcallsDAO) UpdateWithID(ctx context.Context, cMap map[string]interface{}, callID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcalls` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithID(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithID(_), error: %v", err)
	}

	return
}

// UpdateWithIDTx .
func (dao *AvcallsDAO) UpdateWithIDTx(tx *sqlx.Tx, cMap map[string]interface{}, callID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcalls` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithIDTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithIDTx(_), error: %v", err)
	}

	return
}
