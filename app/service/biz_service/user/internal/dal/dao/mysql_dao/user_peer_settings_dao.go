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

type UserPeerSettingsDAO struct {
	db *sqlx.DB
}

func NewUserPeerSettingsDAO(db *sqlx.DB) *UserPeerSettingsDAO {
	return &UserPeerSettingsDAO{db}
}

func (dao *UserPeerSettingsDAO) InsertIgnore(ctx context.Context, do *dataobject.UserPeerSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)"
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

func (dao *UserPeerSettingsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.UserPeerSettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_settings(user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance) values (:user_id, :peer_type, :peer_id, :hide, :report_spam, :add_contact, :block_contact, :share_contact, :need_contacts_exception, :report_geo, :autoarchived, :geo_distance)"
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

func (dao *UserPeerSettingsDAO) Select(ctx context.Context, user_id int32, peer_type int8, peer_id int32) (rValue *dataobject.UserPeerSettingsDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, hide, report_spam, add_contact, block_contact, share_contact, need_contacts_exception, report_geo, autoarchived, geo_distance from user_peer_settings where user_id = ? and peer_type = ? and peer_id = ? and hide = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPeerSettingsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserPeerSettingsDAO) Update(ctx context.Context, cMap map[string]interface{}, user_id int32, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_peer_settings set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

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

func (dao *UserPeerSettingsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, user_id int32, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update user_peer_settings set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, user_id)
	aValues = append(aValues, peer_type)
	aValues = append(aValues, peer_id)

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

func (dao *UserPeerSettingsDAO) Delete(ctx context.Context, user_id int32, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_type, peer_id)

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

func (dao *UserPeerSettingsDAO) DeleteTx(tx *sqlx.Tx, user_id int32, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_settings set hide = 1 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_type, peer_id)

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
