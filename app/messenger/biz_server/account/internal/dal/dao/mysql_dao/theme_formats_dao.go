package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ThemeFormatsDAO struct {
	db *sqlx.DB
}

func NewThemeFormatsDAO(db *sqlx.DB) *ThemeFormatsDAO {
	return &ThemeFormatsDAO{db}
}

func (dao *ThemeFormatsDAO) Insert(ctx context.Context, do *dataobject.ThemeFormatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into theme_formats(theme_id, format, document_id) values (:theme_id, :format, :document_id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
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

func (dao *ThemeFormatsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ThemeFormatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into theme_formats(theme_id, format, document_id) values (:theme_id, :format, :document_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
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

func (dao *ThemeFormatsDAO) SelectByThemeIdAndFormat(ctx context.Context, theme_id int64, format string) (rValue *dataobject.ThemeFormatsDO, err error) {
	var (
		query = "select id, theme_id, format, document_id from theme_formats where theme_id = ? and format = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, theme_id, format)

	if err != nil {
		log.Errorf("queryx in SelectByThemeIdAndFormat(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ThemeFormatsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByThemeIdAndFormat(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ThemeFormatsDAO) SelectByThemeId(ctx context.Context, theme_id int64) (rList []dataobject.ThemeFormatsDO, err error) {
	var (
		query = "select id, theme_id, format, document_id from theme_formats where theme_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, theme_id)

	if err != nil {
		log.Errorf("queryx in SelectByThemeId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ThemeFormatsDO
	for rows.Next() {
		v := dataobject.ThemeFormatsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByThemeId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
