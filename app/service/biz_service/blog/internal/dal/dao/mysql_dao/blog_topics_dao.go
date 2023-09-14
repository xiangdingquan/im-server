package mysql_dao

import (
	"context"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogTopicsDAO struct {
	db *sqlx.DB
}

func NewBlogTopicsDAO(db *sqlx.DB) *BlogTopicsDAO {
	return &BlogTopicsDAO{db}
}

func (dao *BlogTopicsDAO) InsertOrUpdate(ctx context.Context, doList []*dataobject.BlogTopicsDO) (err error) {
	var (
		query = "insert into blog_topics (name) values (:name) on duplicate key update ranking=ranking+1"
	)

	_, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", doList, err)
	}
	return
}

func (dao *BlogTopicsDAO) SelectByName(ctx context.Context, nameList []string) (rList []*dataobject.BlogTopicsDO, err error) {
	var (
		query = "select id,name,ranking from blog_topics where name in (?)"
		a     []interface{}
	)

	query, a, err = sqlx.In(query, nameList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByName(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		log.Errorf("select in SelectByName(_), error: %v", err)
	}

	return
}

func (dao *BlogTopicsDAO) Select(ctx context.Context, fromTopicId, limit int32) (rList []*dataobject.BlogTopicsDO, err error) {
	var (
		query = "select id,name,ranking from blog_topics where id<? order by id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, fromTopicId, limit)

	if err != nil {
		log.Errorf("queryx in BlogTopicsDAO.Select(%d, %d), error: %v", fromTopicId, limit, err)
		return
	}

	defer rows.Close()

	var values []*dataobject.BlogTopicsDO
	for rows.Next() {
		v := &dataobject.BlogTopicsDO{}
		err = rows.StructScan(v)
		if err != nil {
			log.Errorf("structScan in BlogTopicsDAO.Select(%d, %d), error: %v", fromTopicId, limit, err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *BlogTopicsDAO) Count(ctx context.Context) (count int32, err error) {
	var (
		query = "select count(*) from blog_topics"
	)

	err = dao.db.Get(ctx, &count, query)
	if err != nil {
		log.Errorf("get in BlogTopicsDAO.Count(_), error: %v", err)
		return 0, err
	}

	return count, nil
}

func (dao *BlogTopicsDAO) SelectByIds(ctx context.Context, ids []int32) (rList []*dataobject.BlogTopicsDO, err error) {
	var (
		query = "select id,name,ranking from blog_topics where id in (?)"
		a     []interface{}
	)

	query, a, err = sqlx.In(query, ids)
	if err != nil {
		log.Errorf("sqlx.In in SelectByIds(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		log.Errorf("select in SelectByIds(_), error: %v", err)
	}

	return
}

func (dao *BlogTopicsDAO) SelectOrdered(ctx context.Context, limit int32) (rList []*dataobject.BlogTopicsDO, err error) {
	var (
		query = "select id,name,ranking from blog_topics order by ranking desc limit ?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, limit)

	if err != nil {
		log.Errorf("queryx in BlogTopicsDAO.SelectOrdered(%d), error: %v", limit, err)
		return
	}

	defer rows.Close()

	var values []*dataobject.BlogTopicsDO
	for rows.Next() {
		v := &dataobject.BlogTopicsDO{}
		err = rows.StructScan(v)
		if err != nil {
			log.Errorf("structScan in BlogTopicsDAO.SelectOrdered(%d), error: %v", limit, err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return

}
