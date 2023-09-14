package mysql_dao

import (
	"context"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogBannedUsersDAO struct {
	db *sqlx.DB
}

func NewBlogBannedUsersDAO(db *sqlx.DB) *BlogBannedUsersDAO {
	return &BlogBannedUsersDAO{db}
}

func (dao *BlogBannedUsersDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BlogBannedUsersDo) (err error) {
	var (
		query = "INSERT INTO blog_banned_users (user_id,ban_from,ban_to) VALUES (:user_id,FROM_UNIXTIME(:ban_from),FROM_UNIXTIME(:ban_to)) ON DUPLICATE KEY UPDATE ban_from=VALUES(ban_from), ban_to=VALUES(ban_to)"
	)
	_, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(_), error: %v", err)
	}
	return
}

func (dao *BlogBannedUsersDAO) Select(ctx context.Context, offset, limit int32) (rList []*dataobject.BlogBannedUsersDo, err error) {
	var (
		query = "SELECT user_id,UNIX_TIMESTAMP(ban_from) ban_from,UNIX_TIMESTAMP(ban_to) ban_to FROM blog_banned_users ORDER BY ban_to DESC LIMIT ?,?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, offset, limit)
	if err != nil {
		log.Errorf("queryx in BlogBannedUsersDAO.Select(%d, %d), error: %v", offset, limit, err)
		return
	}

	defer rows.Close()

	l := make([]*dataobject.BlogBannedUsersDo, 0)
	if rows.Next() {
		do := &dataobject.BlogBannedUsersDo{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in BlogBannedUsersDAO.Select(%d, %d), error: %v", offset, limit, err)
			return
		}
		l = append(l, do)
	}
	rList = l

	return
}

func (dao *BlogBannedUsersDAO) SelectByUsers(ctx context.Context, uidList []int32) (rList []*dataobject.BlogBannedUsersDo, err error) {
	var (
		query = "SELECT user_id,UNIX_TIMESTAMP(ban_from) ban_from,UNIX_TIMESTAMP(ban_to) ban_to FROM blog_banned_users WHERE user_id in (?)"
		a     []interface{}
	)

	query, a, err = sqlx.In(query, uidList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByUsers(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		log.Errorf("select in SelectByUsers(_), error: %v", err)
	}

	return
}
