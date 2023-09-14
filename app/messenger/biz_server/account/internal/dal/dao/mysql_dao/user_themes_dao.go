package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserThemesDAO struct {
	db *sqlx.DB
}

func NewUserThemesDAO(db *sqlx.DB) *UserThemesDAO {
	return &UserThemesDAO{db}
}

func (dao *UserThemesDAO) InsertIgnore(ctx context.Context, do *dataobject.UserThemesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_themes(user_id, theme_id, format) values (:user_id, :theme_id, :format)"
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

func (dao *UserThemesDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.UserThemesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into user_themes(user_id, theme_id, format) values (:user_id, :theme_id, :format)"
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

func (dao *UserThemesDAO) SelectByUserIdAndFormat(ctx context.Context, user_id int32, format string) (rList []dataobject.UserThemesDO, err error) {
	var (
		query = "select id, user_id, theme_id, format from user_themes where user_id = ? and format = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, format)

	if err != nil {
		log.Errorf("queryx in SelectByUserIdAndFormat(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserThemesDO
	for rows.Next() {
		v := dataobject.UserThemesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByUserIdAndFormat(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
