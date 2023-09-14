package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AppLanguagesDAO struct {
	db *sqlx.DB
}

func NewAppLanguagesDAO(db *sqlx.DB) *AppLanguagesDAO {
	return &AppLanguagesDAO{db}
}

func (dao *AppLanguagesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.AppLanguagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into app_languages(app, lang_code, version, strings_count, translated_count, state) values (:app, :lang_code, :version, :strings_count, :translated_count, 0) on duplicate key update version = values(version), strings_count = values(strings_count), translated_count = values(translated_count), state = 0"
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

func (dao *AppLanguagesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.AppLanguagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into app_languages(app, lang_code, version, strings_count, translated_count, state) values (:app, :lang_code, :version, :strings_count, :translated_count, 0) on duplicate key update version = values(version), strings_count = values(strings_count), translated_count = values(translated_count), state = 0"
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

func (dao *AppLanguagesDAO) SelectLanguageList(ctx context.Context, app string) (rList []dataobject.AppLanguagesDO, err error) {
	var (
		query = "select id, app, lang_code, version, strings_count, translated_count from app_languages where app = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, app)

	if err != nil {
		log.Errorf("queryx in SelectLanguageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AppLanguagesDO
	for rows.Next() {
		v := dataobject.AppLanguagesDO{}
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

func (dao *AppLanguagesDAO) SelectLanguage(ctx context.Context, app string, lang_code string) (rValue *dataobject.AppLanguagesDO, err error) {
	var (
		query = "select id, app, lang_code, version, strings_count, translated_count from app_languages where app = ? and lang_code = ? and state = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, app, lang_code)

	if err != nil {
		log.Errorf("queryx in SelectLanguage(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AppLanguagesDO{}
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
