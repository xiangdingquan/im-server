package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserNotifySettingsDAO struct {
	db *sqlx.DB
}

func NewUserNotifySettingsDAO(db *sqlx.DB) *UserNotifySettingsDAO {
	return &UserNotifySettingsDAO{db}
}

func (dao *UserNotifySettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
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

func (dao *UserNotifySettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
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

func (dao *UserNotifySettingsDAO) SelectAll(ctx context.Context, user_id int32) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserNotifySettingsDO
	for rows.Next() {
		v := dataobject.UserNotifySettingsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAll(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserNotifySettingsDAO) Select(ctx context.Context, user_id int32, peer_type int8, peer_id int32) (rValue *dataobject.UserNotifySettingsDO, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserNotifySettingsDO{}
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

func (dao *UserNotifySettingsDAO) DeleteAll(ctx context.Context, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		log.Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}

func (dao *UserNotifySettingsDAO) DeleteAllTx(tx *sqlx.Tx, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

	if err != nil {
		log.Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}
