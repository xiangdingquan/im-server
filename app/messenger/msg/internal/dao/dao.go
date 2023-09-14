package dao

import (
	"context"

	"open.chat/app/messenger/msg/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.MessagesDAO
	*mysql_dao.ConversationsDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.ChannelMessagesDAO
	*mysql_dao.ChannelMessagesDeleteDAO
	*mysql_dao.ChannelMessageVisiblesDAO
	*mysql_dao.ChannelsDAO
	*mysql_dao.ChannelParticipantsDAO
	*mysql_dao.ChannelPtsUpdatesDAO
	*mysql_dao.ScheduledMessagesDAO
	*mysql_dao.ChannelUnreadMentionsDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                        db,
		MessagesDAO:               mysql_dao.NewMessagesDAO(db),
		ConversationsDAO:          mysql_dao.NewConversationsDAO(db),
		ChatParticipantsDAO:       mysql_dao.NewChatParticipantsDAO(db),
		ChannelMessagesDAO:        mysql_dao.NewChannelMessagesDAO(db),
		ChannelMessagesDeleteDAO:  mysql_dao.NewChannelMessagesDeleteDAO(db),
		ChannelMessageVisiblesDAO: mysql_dao.NewChannelMessageVisiblesDAO(db),
		ChannelsDAO:               mysql_dao.NewChannelsDAO(db),
		ChannelParticipantsDAO:    mysql_dao.NewChannelParticipantsDAO(db),
		ChannelPtsUpdatesDAO:      mysql_dao.NewChannelPtsUpdatesDAO(db),
		ScheduledMessagesDAO:      mysql_dao.NewScheduledMessagesDAO(db),
		ChannelUnreadMentionsDAO:  mysql_dao.NewChannelUnreadMentionsDAO(db),
		CommonDAO:                 sqlx.NewCommonDAO(db),
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
