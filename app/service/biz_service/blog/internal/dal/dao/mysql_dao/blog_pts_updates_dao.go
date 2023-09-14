package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewBlogPtsUpdatesDAO(db *sqlx.DB) *BlogPtsUpdatesDAO {
	return &BlogPtsUpdatesDAO{db}
}

func (dao *BlogPtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.BlogPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_pts_updates(user_id, pts, pts_count, update_type, update_data, `date`) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date)"
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

func (dao *BlogPtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_pts_updates(user_id, pts, pts_count, update_type, update_data, `date`) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date)"
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

func (dao *BlogPtsUpdatesDAO) SelectLastPts(ctx context.Context, user_id int32) (rValue *dataobject.BlogPtsUpdatesDO, err error) {
	var (
		query = "select pts from blog_pts_updates where user_id = ? order by pts desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BlogPtsUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastPts(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogPtsUpdatesDAO) SelectByGtPts(ctx context.Context, user_id int32, pts int32, limit int32) (rList []dataobject.BlogPtsUpdatesDO, err error) {
	var (
		query = "select user_id, pts, pts_count, update_type, update_data from blog_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, pts, limit)

	if err != nil {
		log.Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BlogPtsUpdatesDO
	for rows.Next() {
		v := dataobject.BlogPtsUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtPts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
