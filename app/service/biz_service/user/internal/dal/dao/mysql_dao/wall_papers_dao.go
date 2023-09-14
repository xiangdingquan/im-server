package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type WallPapersDAO struct {
	db *sqlx.DB
}

func NewWallPapersDAO(db *sqlx.DB) *WallPapersDAO {
	return &WallPapersDAO{db}
}

func (dao *WallPapersDAO) Insert(ctx context.Context, do *dataobject.WallPapersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into wall_papers(type, title, color, bg_color, photo_id) values (:type, :title, :color, :bg_color, :photo_id)"
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

func (dao *WallPapersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.WallPapersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into wall_papers(type, title, color, bg_color, photo_id) values (:type, :title, :color, :bg_color, :photo_id)"
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

func (dao *WallPapersDAO) SelectAll(ctx context.Context) (rList []dataobject.WallPapersDO, err error) {
	var (
		query = "select id, type, title, color, bg_color, photo_id from wall_papers where deleted_at = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.WallPapersDO
	for rows.Next() {
		v := dataobject.WallPapersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAll(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
