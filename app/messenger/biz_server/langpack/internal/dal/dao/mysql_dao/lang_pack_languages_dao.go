package mysql_dao

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type LangPackLanguagesDAO struct {
	db *sqlx.DB
}

func NewLangPackLanguagesDAO(db *sqlx.DB) *LangPackLanguagesDAO {
	return &LangPackLanguagesDAO{db}
}

func (dao *LangPackLanguagesDAO) SelectByLangPack(ctx context.Context, lang_pack string) (rList []dataobject.LangPackLanguagesDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, base_lang_code, official, rtl, beta, name, native_name, plural_code, strings_count, translated_count, translations_url from lang_pack_languages where lang_pack = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack)

	if err != nil {
		log.Errorf("queryx in SelectByLangPack(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.LangPackLanguagesDO
	for rows.Next() {
		v := dataobject.LangPackLanguagesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByLangPack(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *LangPackLanguagesDAO) SelectByLangPackCode(ctx context.Context, lang_pack string, lang_code string) (rValue *dataobject.LangPackLanguagesDO, err error) {
	var (
		query = "select lang_pack, lang_code, version, base_lang_code, official, rtl, beta, name, native_name, plural_code, strings_count, translated_count, translations_url from lang_pack_languages where lang_pack = ? and lang_code = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, lang_pack, lang_code)

	if err != nil {
		log.Errorf("queryx in SelectByLangPackCode(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.LangPackLanguagesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByLangPackCode(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *LangPackLanguagesDAO) SelectVersion(ctx context.Context, lang_pack string, lang_code string) (rValue int32, err error) {
	var query = "select version from lang_pack_languages where lang_pack = ? and lang_code = ? and state = 0"
	err = dao.db.Get(ctx, &rValue, query, lang_pack, lang_code)

	if err != nil {
		log.Errorf("get in SelectVersion(_), error: %v", err)
	}

	return
}
