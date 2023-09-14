package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.SecretChatsDAO
	*mysql_dao.SecretChatMessagesDAO
	*mysql_dao.SecretChatQtsUpdatesDAO
	*mysql_dao.SecretChatsCloseRequestsDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := sqlx.NewMySQL(c)
	return &Mysql{
		DB:                          db,
		SecretChatsDAO:              mysql_dao.NewSecretChatsDAO(db),
		SecretChatMessagesDAO:       mysql_dao.NewSecretChatMessagesDAO(db),
		SecretChatQtsUpdatesDAO:     mysql_dao.NewSecretChatQtsUpdatesDAO(db),
		SecretChatsCloseRequestsDAO: mysql_dao.NewSecretChatsCloseRequestsDAO(db),
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
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() (dao *Dao) {
	var (
		dc struct {
			Mysql *sqlx.Config
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	dao = &Dao{
		Mysql: newMysqlDao(dc.Mysql),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.Mysql.Ping(ctx)
}
