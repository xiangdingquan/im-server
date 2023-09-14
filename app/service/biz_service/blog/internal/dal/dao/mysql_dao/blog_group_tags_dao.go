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

type BlogGroupTagsDAO struct {
	db *sqlx.DB
}

func NewBlogGroupTagsDAO(db *sqlx.DB) *BlogGroupTagsDAO {
	return &BlogGroupTagsDAO{db}
}

func (dao *BlogGroupTagsDAO) Insert(ctx context.Context, do *dataobject.BlogGroupTagsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_group_tags(user_id, title, member_uids, `date`) values (:user_id, :title, :member_uids, :date)"
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

func (dao *BlogGroupTagsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogGroupTagsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_group_tags(user_id, title, member_uids, `date`) values (:user_id, :title, :member_uids, :date)"
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

func (dao *BlogGroupTagsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogGroupTagsDO, err error) {
	var (
		query = "select id, user_id, title, member_uids, `date`, `deleted` from blog_group_tags where id = ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, id)
	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogGroupTagsDO{}
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

func (dao *BlogGroupTagsDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_group_tags set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogGroupTagsDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_group_tags set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogGroupTagsDAO) SelectByUserAndTags(ctx context.Context, fromUserId int32, ids []int32) (rValue []*dataobject.BlogGroupTagsDO, err error) {
	var (
		query = "select id, user_id, title, member_uids, `date`, `deleted` from blog_group_tags where id in (?) and user_id = ? and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogGroupTagsDO, 0)
	if len(ids) == 0 {
		return rValue, nil
	}

	query, a, err = sqlx.In(query, ids, fromUserId)
	if err != nil {
		log.Errorf("sqlx.In in SelectByUserAndTags(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectByUserAndTags(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogGroupTagsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserAndTags(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogGroupTagsDAO) SelectByUser(ctx context.Context, fromUserId int32) (rValue []*dataobject.BlogGroupTagsDO, err error) {
	var (
		query = "select id, user_id, title, member_uids, `date`, `deleted` from blog_group_tags where user_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, fromUserId)
	if err != nil {
		log.Errorf("queryx in SelectByUser(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogGroupTagsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogGroupTagsDO{}
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

func (dao *BlogGroupTagsDAO) UpdateMembers(ctx context.Context, fromUserId int32, tagId int32, memberList string) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE `blog_group_tags` SET member_uids = ? WHERE id = ? AND user_id = ? AND deleted = 0"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, memberList, tagId, fromUserId)

	if err != nil {
		log.Errorf("exec in UpdateMembers(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMembers(_), error: %v", err)
	}

	return
}

func (dao *BlogGroupTagsDAO) UpdateMembersTx(tx *sqlx.Tx, fromUserId int32, tagId int32, memberList string) (rowsAffected int64, err error) {
	var (
		query   = "UPDATE `blog_group_tags` SET member_uids = ? WHERE id = ? AND user_id = ? AND deleted = 0"
		rResult sql.Result
	)

	rResult, err = tx.Exec(query, memberList, tagId, fromUserId)

	if err != nil {
		log.Errorf("exec in UpdateMembersTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMembersTx(_), error: %v", err)
	}

	return
}

func (dao *BlogGroupTagsDAO) DeleteGroupTags(tx *sqlx.Tx, fromUserId int32, ids []int32) (rowsAffected int64, err error) {
	var (
		query   = "update blog_group_tags set deleted = 1 where id in (?) and user_id = ? and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)
	query, a, err = sqlx.In(query, ids, fromUserId)
	if err != nil {
		log.Errorf("sqlx.In in DeleteGroupTags(_), error: %v", err)
		return
	}

	rResult, err = tx.Exec(query, a...)
	if err != nil {
		log.Errorf("exec in DeleteGroupTags(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteGroupTags(_), error: %v", err)
	}

	return
}
