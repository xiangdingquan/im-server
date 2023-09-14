package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/account/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ThemesDAO struct {
	db *sqlx.DB
}

func NewThemesDAO(db *sqlx.DB) *ThemesDAO {
	return &ThemesDAO{db}
}

func (dao *ThemesDAO) Insert(ctx context.Context, do *dataobject.ThemesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into themes(theme_id, access_hash, slug, title) values (:theme_id, :access_hash, :slug, :title)"
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

func (dao *ThemesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ThemesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into themes(theme_id, access_hash, slug, title) values (:theme_id, :access_hash, :slug, :title)"
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

func (dao *ThemesDAO) SelectBySlug(ctx context.Context, slug string) (rValue *dataobject.ThemesDO, err error) {
	var (
		query = "select id, theme_id, access_hash, creator, default2, slug, title, settings, installs_count from themes where slug = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, slug)

	if err != nil {
		log.Errorf("queryx in SelectBySlug(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ThemesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBySlug(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ThemesDAO) SelectByThemeId(ctx context.Context, theme_id int64) (rValue *dataobject.ThemesDO, err error) {
	var (
		query = "select id, theme_id, access_hash, creator, default2, slug, title, settings, installs_count from themes where theme_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, theme_id)

	if err != nil {
		log.Errorf("queryx in SelectByThemeId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ThemesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByThemeId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ThemesDAO) AddInstallsCount(ctx context.Context, theme_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update themes set installs_count = installs_count + 1 where theme_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, theme_id)

	if err != nil {
		log.Errorf("exec in AddInstallsCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in AddInstallsCount(_), error: %v", err)
	}

	return
}

func (dao *ThemesDAO) AddInstallsCountTx(tx *sqlx.Tx, theme_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update themes set installs_count = installs_count + 1 where theme_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, theme_id)

	if err != nil {
		log.Errorf("exec in AddInstallsCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in AddInstallsCount(_), error: %v", err)
	}

	return
}
