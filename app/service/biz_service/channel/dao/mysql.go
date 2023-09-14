package dao

import (
	"context"
	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/channel/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ChannelsDAO
	*mysql_dao.ChannelParticipantsDAO
	*mysql_dao.ChannelMessagesDAO
	*mysql_dao.ChannelMediaUnreadDAO
	*mysql_dao.ChannelPtsUpdatesDAO
	*mysql_dao.ChannelAdminLogsDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                     db,
		ChannelsDAO:            mysql_dao.NewChannelsDAO(db),
		ChannelParticipantsDAO: mysql_dao.NewChannelParticipantsDAO(db),
		ChannelMessagesDAO:     mysql_dao.NewChannelMessagesDAO(db),
		ChannelMediaUnreadDAO:  mysql_dao.NewChannelMediaUnreadDAO(db),
		ChannelPtsUpdatesDAO:   mysql_dao.NewChannelPtsUpdatesDAO(db),
		ChannelAdminLogsDAO:    mysql_dao.NewChannelAdminLogsDAO(db),
		CommonDAO:              sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
