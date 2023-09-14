package mysql_dao

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type LanguagesDAO struct {
	db *sqlx.DB
}

func NewLanguagesDAO(db *sqlx.DB) *LanguagesDAO {
	return &LanguagesDAO{db}
}

func (dao *LanguagesDAO) SelectAllLangCode(ctx context.Context) (rList []string, err error) {
	var query = "select lang_code from languages"
	err = dao.db.Select(ctx, &rList, query)

	if err != nil {
		log.Errorf("select in SelectAllLangCode(_), error: %v", err)
	}

	return
}

func (dao *LanguagesDAO) SelectLanguageList(ctx context.Context, codeList []string) (rList []dataobject.LanguagesDO, err error) {
	var (
		query = "select id, lang_code, base_lang_code, link, official, rtl, beta, name, native_name, plural_code, translations_url from languages where state = 0 and lang_code in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, codeList)
	if err != nil {
		log.Errorf("sqlx.In in SelectLanguageList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectLanguageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.LanguagesDO
	for rows.Next() {
		v := dataobject.LanguagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectLanguageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *LanguagesDAO) SelectLanguage(ctx context.Context, lang_code string) (rValue *dataobject.LanguagesDO, err error) {
	var (
		query = "select id, lang_code, base_lang_code, link, official, rtl, beta, name, native_name, plural_code, translations_url from languages where lang_code = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_code)

	if err != nil {
		log.Errorf("queryx in SelectLanguage(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.LanguagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLanguage(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}
