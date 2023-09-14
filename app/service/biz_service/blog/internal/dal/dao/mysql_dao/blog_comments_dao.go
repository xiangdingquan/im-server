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

type BlogCommentsDAO struct {
	db *sqlx.DB
}

func NewBlogCommentsDAO(db *sqlx.DB) *BlogCommentsDAO {
	return &BlogCommentsDAO{db}
}

func (dao *BlogCommentsDAO) Insert(ctx context.Context, do *dataobject.BlogCommentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_comments(user_id, type, text, blog_id, comment_id, reply_id, `date`) values (:user_id, :type, :text, :blog_id, :comment_id, :reply_id, :date)"
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

func (dao *BlogCommentsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogCommentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_comments(user_id, type, text, blog_id, comment_id, reply_id, `date`) values (:user_id, :type, :text, :blog_id, :comment_id, :reply_id, :date)"
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

func (dao *BlogCommentsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogCommentsDO, err error) {
	var (
		query = "select id, user_id, type, text, blog_id, comment_id, reply_id, `date`, `deleted` from blog_comments where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogCommentsDO{}
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

func (dao *BlogCommentsDAO) SelectList(ctx context.Context, ids []int32) (rValue []*dataobject.BlogCommentsDO, err error) {
	var (
		query = "select id, user_id, type, text, blog_id, comment_id, reply_id, `date`, `deleted` from blog_comments where id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogCommentsDO, 0)
	if len(ids) == 0 {
		return rValue, nil
	}
	query, a, err = sqlx.In(query, ids)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogCommentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogCommentsDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_comments set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogCommentsDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_comments set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogCommentsDAO) SelectListByBlogId(ctx context.Context, blogId, offsetId, limit int32) (rValue []*dataobject.BlogCommentsDO, err error) {
	var (
		query = "select id, user_id, type, text, blog_id, comment_id, reply_id, `date`, `deleted` from blog_comments where id < ? and type = 0 and blog_id = ? and deleted = 0 ORDER BY id DESC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, offsetId, blogId, limit)

	if err != nil {
		log.Errorf("queryx in SelectListByBlogId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogCommentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogCommentsDO{}
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

func (dao *BlogCommentsDAO) SelectListByCommentId(ctx context.Context, commentId, offsetId, limit int32) (rValue []*dataobject.BlogCommentsDO, err error) {
	var (
		query = "select id, user_id, type, text, blog_id, comment_id, reply_id, `date`, `deleted` from blog_comments where id > ? and type = 1 and comment_id = ? and deleted = 0 ORDER BY id ASC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, offsetId, commentId, limit)

	if err != nil {
		log.Errorf("queryx in SelectListByCommentId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogCommentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogCommentsDO{}
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

func (dao *BlogCommentsDAO) LikeTx(tx *sqlx.Tx, id int32, liked bool) (rowsAffected int64, err error) {
	var (
		query   = "update blog_comments set likes=likes+1 where id = ? and deleted = 0"
		rResult sql.Result
	)
	if !liked {
		query = "update blog_comments set likes=likes-1 where id = ? and deleted = 0"
	}
	rResult, err = tx.Exec(query, id)

	if err != nil {
		log.Errorf("exec in LikeTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in LikeTx(_), error: %v", err)
	}

	return
}
