package mysql_dao

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type LangPackStringsDAO struct {
	db *sqlx.DB
}

func NewLangPackStringsDAO(db *sqlx.DB) *LangPackStringsDAO {
	return &LangPackStringsDAO{db}
}

func (dao *LangPackStringsDAO) SelectListByKeyList(ctx context.Context, lang_pack string, lang_code string, keyList []string) (rList []dataobject.LangPackStringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from lang_pack_strings where lang_pack = ? and lang_code = ? and key2 in (?) and deleted = 0"
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

	var values []dataobject.LangPackStringsDO
	for rows.Next() {
		v := dataobject.LangPackStringsDO{}
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

func (dao *LangPackStringsDAO) SelectList(ctx context.Context, lang_pack string, lang_code string) (rList []dataobject.LangPackStringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from lang_pack_strings where lang_pack = ? and lang_code = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack, lang_code)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.LangPackStringsDO
	for rows.Next() {
		v := dataobject.LangPackStringsDO{}
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

func (dao *LangPackStringsDAO) SelectGTVersionList(ctx context.Context, lang_pack string, lang_code string, version int32) (rList []dataobject.LangPackStringsDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, key2, pluralized, value, zero_value, one_value, two_value, few_value, many_value, other_value from lang_pack_strings where lang_pack = ? and lang_code = ? and version > ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack, lang_code, version)

	if err != nil {
		log.Errorf("queryx in SelectGTVersionList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.LangPackStringsDO
	for rows.Next() {
		v := dataobject.LangPackStringsDO{}
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
