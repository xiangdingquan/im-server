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

type BlogsDAO struct {
	db *sqlx.DB
}

func NewBlogsDAO(db *sqlx.DB) *BlogsDAO {
	return &BlogsDAO{db}
}

func (dao *BlogsDAO) Insert(ctx context.Context, do *dataobject.BlogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blogs(user_id, read_update_id, read_blog_id, read_comment_id, moments, follows, fans, comments, likes, `date`) values (:user_id, :read_update_id, :read_blog_id, :read_comment_id, :moments, :follows, :fans, :comments, :likes, :date)"
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

func (dao *BlogsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blogs(user_id, read_update_id, read_blog_id, read_comment_id, moments, follows, fans, comments, likes, `date`) values (:user_id, :read_update_id, :read_blog_id, :read_comment_id, :moments, :follows, :fans, :comments, :likes, :date)"
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

func (dao *BlogsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogsDO, err error) {
	var (
		query = "select id, user_id, read_update_id, read_blog_id, read_comment_id, moments, follows, fans, comments, likes, `date`, `deleted` from blogs where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogsDO{}
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

func (dao *BlogsDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blogs set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogsDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blogs set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogsDAO) SelectByUser(ctx context.Context, userId int32) (rValue *dataobject.BlogsDO, err error) {
	var (
		query = "select id, user_id, read_update_id, read_blog_id, read_comment_id, moments, follows, fans, comments, likes, `date`, `deleted` from blogs where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUser(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BlogsDAO) IncMoments(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set moments = moments + ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in IncMoments(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncMoments(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) DecMoments(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set moments = moments - ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in DecMoments(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecMoments(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) IncFollows(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set follows = follows + ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in IncFollows(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncFollows(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) DecFollows(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set follows = follows - ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in DecFollows(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecFollows(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) IncFans(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set fans = fans + ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in IncFans(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncFans(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) DecFans(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set fans = fans - ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in DecFans(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecFans(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) IncLikes(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set likes = likes + ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in IncLikes(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncLikes(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) DecLikes(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set likes = likes - ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in DecLikes(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecLikes(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) IncComments(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set comments = comments + ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in IncComments(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in IncComments(_), error: %v", err)
	}

	return
}

func (dao *BlogsDAO) DecComments(tx *sqlx.Tx, userId, n int32) (rowsAffected int64, err error) {
	var (
		query   = "update blogs set comments = comments - ? where user_id = ?"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, n, userId)

	if err != nil {
		log.Errorf("exec in DecComments(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecComments(_), error: %v", err)
	}

	return
}
