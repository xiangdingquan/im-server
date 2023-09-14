package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/push/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type DevicesDAO struct {
	db *sqlx.DB
}

func NewDevicesDAO(db *sqlx.DB) *DevicesDAO {
	return &DevicesDAO{db}
}

func (dao *DevicesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DevicesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into devices(auth_key_id, user_id, token_type, token, no_muted, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :no_muted, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), no_muted = values(no_muted), secret = values(secret), other_uids = values(other_uids), state = 0"
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

func (dao *DevicesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DevicesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into devices(auth_key_id, user_id, token_type, token, no_muted, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :no_muted, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), no_muted = values(no_muted), secret = values(secret), other_uids = values(other_uids), state = 0"
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

func (dao *DevicesDAO) Select(ctx context.Context, auth_key_id int64, user_id int32, token_type int8) (rValue *dataobject.DevicesDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, token_type, token, no_muted, locked_period, app_sandbox, secret, other_uids from devices where auth_key_id = ? and user_id = ? and token_type = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id, user_id, token_type)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.DevicesDO{}
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

func (dao *DevicesDAO) SelectListByUser(ctx context.Context, user_id int32) (rList []dataobject.DevicesDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, token_type, token, no_muted, locked_period, app_sandbox, secret, other_uids from devices where user_id = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectListByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DevicesDO
	for rows.Next() {
		v := dataobject.DevicesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByUser(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *DevicesDAO) SelectUserFcmList(ctx context.Context, user_id int32) (rList []dataobject.DevicesDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, token_type, token, no_muted, locked_period, app_sandbox, secret, other_uids from devices where user_id = ? and token_type = 2 and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectUserFcmList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DevicesDO
	for rows.Next() {
		v := dataobject.DevicesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserFcmList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *DevicesDAO) SelectListById(ctx context.Context, token_type int8, token string) (rList []dataobject.DevicesDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, token_type, token, no_muted, locked_period from devices where token_type = ? and token = ? and state = 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, token_type, token)

	if err != nil {
		log.Errorf("queryx in SelectListById(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DevicesDO
	for rows.Next() {
		v := dataobject.DevicesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListById(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *DevicesDAO) UpdateState(ctx context.Context, state int8, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where auth_key_id = ? and user_id = ?" // and token_type
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateStateTx(tx *sqlx.Tx, state int8, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where auth_key_id = ? and user_id = ?" // and token_type
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateStateById(ctx context.Context, state int8, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, id)

	if err != nil {
		log.Errorf("exec in UpdateStateById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateById(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateStateByIdTx(tx *sqlx.Tx, state int8, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, id)

	if err != nil {
		log.Errorf("exec in UpdateStateById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateById(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateStateByToken(ctx context.Context, state int8, token_type int8, token string) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where token_type = ? and token = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, token_type, token)

	if err != nil {
		log.Errorf("exec in UpdateStateByToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByToken(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateStateByTokenTx(tx *sqlx.Tx, state int8, token_type int8, token string) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where token_type = ? and token = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, token_type, token)

	if err != nil {
		log.Errorf("exec in UpdateStateByToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateStateByToken(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateLockedPeriod(ctx context.Context, locked_period int32, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update devices set locked_period = ? where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, locked_period, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLockedPeriod(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLockedPeriod(_), error: %v", err)
	}

	return
}

func (dao *DevicesDAO) UpdateLockedPeriodTx(tx *sqlx.Tx, locked_period int32, auth_key_id int64, user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update devices set locked_period = ? where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, locked_period, auth_key_id, user_id)

	if err != nil {
		log.Errorf("exec in UpdateLockedPeriod(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLockedPeriod(_), error: %v", err)
	}

	return
}
