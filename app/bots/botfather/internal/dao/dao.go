package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/bots/botfather/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/app/service/auth_session/client"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.BotsDAO
	*mysql_dao.UsersDAO
	*mysql_dao.BotCommandsDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:             db,
		BotsDAO:        mysql_dao.NewBotsDAO(db),
		UsersDAO:       mysql_dao.NewUsersDAO(db),
		BotCommandsDAO: mysql_dao.NewBotCommandsDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

type Dao struct {
	*Mysql
	*Redis
	authsessionpb.RPCSessionClient
}

func New(c *warden.ClientConfig) (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	var err error
	dao.RPCSessionClient, err = authsession_client.New(c)
	if err != nil {
		panic(err)
	}
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
