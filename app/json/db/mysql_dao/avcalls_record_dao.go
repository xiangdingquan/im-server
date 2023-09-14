package mysqldao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// AvcallsRecordsDAO .
type AvcallsRecordsDAO struct {
	db *sqlx.DB
}

// NewAvcallsRecordsDAO .
func NewAvcallsRecordsDAO(db *sqlx.DB) *AvcallsRecordsDAO {
	return &AvcallsRecordsDAO{db}
}

// Insert .
func (dao *AvcallsRecordsDAO) Insert(ctx context.Context, do *dbo.AvcallRecordDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `avcall_records`(call_id, user_id, is_read, enter_at) VALUES (:call_id, :user_id, :is_read, :enter_at)"
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
func (dao *AvcallsRecordsDAO) InsertTx(tx *sqlx.Tx, do *dbo.AvcallRecordDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `avcall_records`(call_id, user_id, is_read, enter_at) VALUES (:call_id, :user_id, :is_read, :enter_at)"
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

// Select .
func (dao *AvcallsRecordsDAO) Select(ctx context.Context, callID, userID uint32) (rValue *dbo.AvcallRecordDO, err error) {
	var (
		query = "SELECT id, call_id, user_id, is_read, call_time, enter_at, leave_at FROM `avcall_records` WHERE call_id = ? AND user_id = ? AND deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, callID, userID)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.AvcallRecordDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

// Selects .
func (dao *AvcallsRecordsDAO) Selects(ctx context.Context, callID uint32) (rList []dbo.AvcallRecordDO, err error) {
	var (
		query = "SELECT id, call_id, user_id, is_read, call_time, enter_at, leave_at FROM `avcall_records` WHERE call_id = ? AND deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, callID)

	if err != nil {
		log.Errorf("queryx in Selects(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dbo.AvcallRecordDO
	for rows.Next() {
		v := dbo.AvcallRecordDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in Selects(_), error: %v", err)
		}
		values = append(values, v)
	}
	rList = values
	return
}

// UpdateWithID .
func (dao *AvcallsRecordsDAO) UpdateWithID(ctx context.Context, cMap map[string]interface{}, rID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, rID)

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
func (dao *AvcallsRecordsDAO) UpdateWithIDTx(tx *sqlx.Tx, cMap map[string]interface{}, rID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, rID)

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

// UpdateWithCallAndUser .
func (dao *AvcallsRecordsDAO) UpdateWithCallAndUser(ctx context.Context, cMap map[string]interface{}, callID uint32, userID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE call_id = ? AND user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID, userID)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithCallAndUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithCallAndUser(_), error: %v", err)
	}

	return
}

// UpdateWithCallAndUserTx .
func (dao *AvcallsRecordsDAO) UpdateWithCallAndUserTx(tx *sqlx.Tx, cMap map[string]interface{}, callID uint32, userID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE call_id = ? AND user_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID, userID)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithCallAndUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithCallAndUser(_), error: %v", err)
	}

	return
}

// UpdateWithCallID .
func (dao *AvcallsRecordsDAO) UpdateWithCallID(ctx context.Context, cMap map[string]interface{}, callID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE call_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithCallID(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithCallID(_), error: %v", err)
	}

	return
}

// UpdateWithCallIDTx .
func (dao *AvcallsRecordsDAO) UpdateWithCallIDTx(tx *sqlx.Tx, cMap map[string]interface{}, callID uint32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("UPDATE `avcall_records` SET %s WHERE call_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, callID)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateWithCallIDTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateWithCallIDTx(_), error: %v", err)
	}

	return
}

// UpdateRead .
func (dao *AvcallsRecordsDAO) UpdateRead(ctx context.Context, rID uint32) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE `avcall_records` SET is_read = 1 WHERE user_id = ? AND is_read = 0 AND deleted = 0"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, rID)

	if err != nil {
		log.Errorf("exec in UpdateRead(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateRead(_), error: %v", err)
	}

	return
}

// UpdateReadTx .
func (dao *AvcallsRecordsDAO) UpdateReadTx(tx *sqlx.Tx, rID uint32) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE `avcall_records` SET is_read = 1 WHERE user_id = ? AND is_read = 0 AND deleted = 0"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, rID)

	if err != nil {
		log.Errorf("exec in UpdateReadTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadTx(_), error: %v", err)
	}

	return
}
