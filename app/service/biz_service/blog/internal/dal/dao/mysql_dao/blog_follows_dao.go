package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogFollowsDAO struct {
	db *sqlx.DB
}

func NewBlogFollowsDAO(db *sqlx.DB) *BlogFollowsDAO {
	return &BlogFollowsDAO{db}
}

func (dao *BlogFollowsDAO) Insert(ctx context.Context, do *dataobject.BlogFollowsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_follows(user_id, target_uid, `date`) values (:user_id, :target_uid, :date)"
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

func (dao *BlogFollowsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogFollowsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_follows(user_id, target_uid, `date`) values (:user_id, :target_uid, :date)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v)_error: %v", do, err)
	}

	return
}

func (dao *BlogFollowsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogFollowsDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_follows set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)
	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *BlogFollowsDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_follows set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTx(_), error: %v", err)
	}

	return
}

func (dao *BlogFollowsDAO) SelectAllByUser(ctx context.Context, userId int32) (rValue []*dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in SelectAllByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogFollowsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectAllByUser(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogFollowsDAO) SelectAllByTargetUser(ctx context.Context, userId int32) (rValue []*dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where target_uid = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in SelectAllByTargetUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogFollowsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectAllByTargetUser(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogFollowsDAO) SelectByUser(ctx context.Context, userId int32, offset int32, limit int32) (rValue []*dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where user_id = ? and deleted = 0 order by id desc limit ?, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId, offset, limit)

	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogFollowsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUser(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogFollowsDAO) SelectByTargetUser(ctx context.Context, userId int32, offset int32, limit int32) (rValue []*dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where target_uid = ? and deleted = 0 order by id desc limit ?, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId, offset, limit)

	if err != nil {
		log.Errorf("queryx in SelectByTargetUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogFollowsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByTargetUser(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogFollowsDAO) SelectByUserAndTargetUser(ctx context.Context, userId int32, target_uid int32) (rValue *dataobject.BlogFollowsDO, err error) {
	var (
		query = "select id, user_id, target_uid, `date`, `deleted` from blog_follows where user_id = ? and target_uid = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId, target_uid)

	if err != nil {
		log.Errorf("queryx in SelectByUserAndTargetUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogFollowsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserAndTargetUser(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogFollowsDAO) UpdateFollowTx(tx *sqlx.Tx, fromUserId int32, id int32, followed bool) (rowsAffected int64, err error) {
	var (
		query   = "update `blog_follows` set deleted = ? where id = ? and user_id = ?"
		rResult sql.Result
	)
	deleted := !followed
	rResult, err = tx.Exec(query, deleted, id, fromUserId)

	if err != nil {
		log.Errorf("exec in UpdateFollowTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFollowTx(_), error: %v", err)
	}

	return
}
