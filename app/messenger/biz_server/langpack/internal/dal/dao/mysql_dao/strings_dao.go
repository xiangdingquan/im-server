package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type StringsDAO struct {
	db *sqlx.DB
}

func NewStringsDAO(db *sqlx.DB) *StringsDAO {
	return &StringsDAO{db}
}

func (dao *StringsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.StringsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into strings(lang_pack, lang_code, version, key_index, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value) values (:lang_pack, :lang_code, :version, :key_index, :key2, :pluralized, :value, :zero_value, :one_value, :two_value, :few_value, :many_value, :other_value) on duplicate key update version = values(version), value = values(value), zero_value = values(zero_value), one_value = values(one_value), two_value = values(two_value), few_value = values(few_value), many_value = values(many_value), other_value = values(other_value)"
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

func (dao *StringsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.StringsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into strings(lang_pack, lang_code, version, key_index, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value) values (:lang_pack, :lang_code, :version, :key_index, :key2, :pluralized, :value, :zero_value, :one_value, :two_value, :few_value, :many_value, :other_value) on duplicate key update version = values(version), value = values(value), zero_value = values(zero_value), one_value = values(one_value), two_value = values(two_value), few_value = values(few_value), many_value = values(many_value), other_value = values(other_value)"
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

func (dao *StringsDAO) Delete(ctx context.Context, lang_pack string, lang_code string, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update strings set deleted = 1 where lang_pack = ? and lang_code = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, lang_pack, lang_code, key2)

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

func (dao *StringsDAO) DeleteTx(tx *sqlx.Tx, lang_pack string, lang_code string, key2 string) (rowsAffected int64, err error) {
	var (
		query   = "update strings set deleted = 1 where lang_pack = ? and lang_code = ? and key2 = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, lang_pack, lang_code, key2)

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

func (dao *StringsDAO) SelectListByKeyList(ctx context.Context, lang_pack string, lang_code string, keyList []string) (rList []dataobject.StringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from strings where lang_pack = ? and lang_code = ? and key2 in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, lang_pack, lang_code, keyList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByKeyList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByKeyList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.StringsDO
	for rows.Next() {
		v := dataobject.StringsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByKeyList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *StringsDAO) SelectList(ctx context.Context, lang_pack string, lang_code string) (rList []dataobject.StringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from strings where lang_pack = ? and lang_code = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack, lang_code)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.StringsDO
	for rows.Next() {
		v := dataobject.StringsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *StringsDAO) SelectGTVersionList(ctx context.Context, lang_pack string, lang_code string, version int32) (rList []dataobject.StringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from strings where lang_pack = ? and lang_code = ? and version > ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack, lang_code, version)

	if err != nil {
		log.Errorf("queryx in SelectGTVersionList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.StringsDO
	for rows.Next() {
		v := dataobject.StringsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectGTVersionList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
