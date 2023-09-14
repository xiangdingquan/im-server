package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/chat/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ChatsDAO
	*mysql_dao.ChatParticipantsDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                  db,
		ChatsDAO:            mysql_dao.NewChatsDAO(db),
		ChatParticipantsDAO: mysql_dao.NewChatParticipantsDAO(db),
		CommonDAO:           sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
