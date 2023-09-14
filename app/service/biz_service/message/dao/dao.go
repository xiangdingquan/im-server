package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/message/dal/dao/mysql_dao"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.MessageDatasDAO
	*mysql_dao.MessageBoxesDAO
	*mysql_dao.ConversationsDAO
	*mysql_dao.MessagesDAO
	*mysql_dao.ChannelMessagesDAO
	*mysql_dao.ScheduledMessagesDAO
	*mysql_dao.ChannelUnreadMentionsDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                       db,
		MessageDatasDAO:          mysql_dao.NewMessageDatasDAO(db),
		MessageBoxesDAO:          mysql_dao.NewMessageBoxesDAO(db),
		ConversationsDAO:         mysql_dao.NewConversationsDAO(db),
		MessagesDAO:              mysql_dao.NewMessagesDAO(db),
		ChannelMessagesDAO:       mysql_dao.NewChannelMessagesDAO(db),
		ScheduledMessagesDAO:     mysql_dao.NewScheduledMessagesDAO(db),
		ChannelUnreadMentionsDAO: mysql_dao.NewChannelUnreadMentionsDAO(db),
		CommonDAO:                sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

// Dao dao.
type Dao struct {
	*Mysql
	*Redis
}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	idgen.NewUUID()
	idgen.NewSeqIDGen()

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Redis.Close()
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.Redis.Ping(ctx); err != nil {
		return
	}
	return d.Mysql.Ping(ctx)
}
