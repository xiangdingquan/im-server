package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/blog/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.BlogsDAO
	*mysql_dao.BlogMomentsDAO
	*mysql_dao.BlogLikesDAO
	*mysql_dao.BlogGroupTagsDAO
	*mysql_dao.BlogFollowsDAO
	*mysql_dao.BlogCommentsDAO
	*mysql_dao.BlogMomentDeletesDAO
	*mysql_dao.BlogPtsUpdatesDAO
	*mysql_dao.BlogUserPrivaciesDAO
	*mysql_dao.BlogTopicsDAO
	*mysql_dao.BlogTopicMappingsDAO
	*mysql_dao.BlogBannedUsersDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                   db,
		BlogsDAO:             mysql_dao.NewBlogsDAO(db),
		BlogMomentsDAO:       mysql_dao.NewBlogMomentsDAO(db),
		BlogLikesDAO:         mysql_dao.NewBlogLikesDAO(db),
		BlogGroupTagsDAO:     mysql_dao.NewBlogGroupTagsDAO(db),
		BlogFollowsDAO:       mysql_dao.NewBlogFollowsDAO(db),
		BlogCommentsDAO:      mysql_dao.NewBlogCommentsDAO(db),
		BlogMomentDeletesDAO: mysql_dao.NewBlogMomentDeletesDAO(db),
		BlogPtsUpdatesDAO:    mysql_dao.NewBlogPtsUpdatesDAO(db),
		BlogUserPrivaciesDAO: mysql_dao.NewBlogUserPrivaciesDAO(db),
		BlogTopicsDAO:        mysql_dao.NewBlogTopicsDAO(db),
		BlogTopicMappingsDAO: mysql_dao.NewBlogTopicMappingsDAO(db),
		BlogBannedUsersDAO:   mysql_dao.NewBlogBannedUsersDAO(db),
		CommonDAO:            sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
