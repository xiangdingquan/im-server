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

type BlogLikesDAO struct {
	db *sqlx.DB
}

func NewBlogLikesDAO(db *sqlx.DB) *BlogLikesDAO {
	return &BlogLikesDAO{db}
}

func (dao *BlogLikesDAO) InsertOrGetId(tx *sqlx.Tx, do *dataobject.BlogLikesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_likes(user_id, type, blog_id, comment_id, `date`) values (:user_id, :type, :blog_id, :comment_id, :date) on duplicate key update id = last_insert_id(id), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetId(%v)_error: %v", do, err)
	}

	return
}

func (dao *BlogLikesDAO) Insert(ctx context.Context, do *dataobject.BlogLikesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_likes(user_id, type, blog_id, comment_id, `date`) values (:user_id, :type, :blog_id, :comment_id, :date)"
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

func (dao *BlogLikesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogLikesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_likes(user_id, type, blog_id, comment_id, `date`) values (:user_id, :type, :blog_id, :comment_id, :date)"
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

func (dao *BlogLikesDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogLikesDO{}
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

func (dao *BlogLikesDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_likes set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogLikesDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_likes set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogLikesDAO) SelectByBlogId(ctx context.Context, user_id, blogId int32) (rValue *dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where user_id = ? AND type = 0 AND blog_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, blogId)

	if err != nil {
		log.Errorf("queryx in SelectByBlogId(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByBlogId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogLikesDAO) SelectByCommentId(ctx context.Context, user_id, commentId int32) (rValue *dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where user_id = ? AND type = 1 AND comment_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, commentId)

	if err != nil {
		log.Errorf("queryx in SelectByCommentId(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByCommentId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogLikesDAO) SelectListByBlogId(ctx context.Context, blogId, offsetId, limit int32) (rValue []*dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where id > ? and type = 0 and blog_id = ? and deleted = 0 ORDER BY id DESC limit 0, ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, offsetId, blogId, limit)
	if err != nil {
		log.Errorf("queryx in SelectListByBlogId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogLikesDO, 0)
	for rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectListByBlogId(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogLikesDAO) SelectListByCommentId(ctx context.Context, commentId, offsetId, limit int32) (rValue []*dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where id > ? and type = 1 and comment_id = ? and deleted = 0 ORDER BY id DESC limit 0, ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, offsetId, commentId, limit)
	if err != nil {
		log.Errorf("queryx in SelectListByCommentId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogLikesDO, 0)
	for rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectListByCommentId(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogLikesDAO) SelectByUserAndBlogIds(ctx context.Context, user_id int32, blogids []int32) (rValue []*dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where user_id = ? AND type = 0 AND blog_id IN (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogLikesDO, 0)
	if len(blogids) == 0 {
		return rValue, nil
	}

	query, a, err = sqlx.In(query, user_id, blogids)
	if err != nil {
		log.Errorf("sqlx.In in SelectByUserAndBlogIds(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectByUserAndBlogIds(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserAndBlogIds(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogLikesDAO) SelectByUserAndCommentIds(ctx context.Context, user_id int32, commentIds []int32) (rValue []*dataobject.BlogLikesDO, err error) {
	var (
		query = "select id, user_id, type, blog_id, comment_id, `date`, `deleted` from blog_likes where user_id = ? AND type = 1 AND comment_id IN (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogLikesDO, 0)
	if len(commentIds) == 0 {
		return rValue, nil
	}
	query, a, err = sqlx.In(query, user_id, commentIds)
	if err != nil {
		log.Errorf("sqlx.In in SelectByUserAndCommentIds(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectByUserAndCommentIds(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogLikesDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserAndCommentIds(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogLikesDAO) UpdateLikeTx(tx *sqlx.Tx, id int32, liked bool) (rowsAffected int64, err error) {
	var (
		query   = "update blog_likes set deleted = ? where id = ?"
		rResult sql.Result
	)
	deleted := !liked
	rResult, err = tx.Exec(query, deleted, id)

	if err != nil {
		log.Errorf("exec in UpdateLikeTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLikeTx(_), error: %v", err)
	}

	return
}
