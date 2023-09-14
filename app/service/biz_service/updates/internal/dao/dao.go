package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/updates/internal/dal/dao/mysql_dao"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthSeqUpdatesDAO
	*mysql_dao.ChannelPtsUpdatesDAO
	*mysql_dao.UserPtsUpdatesDAO
	*mysql_dao.UserQtsUpdatesDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                   db,
		AuthSeqUpdatesDAO:    mysql_dao.NewAuthSeqUpdatesDAO(db),
		ChannelPtsUpdatesDAO: mysql_dao.NewChannelPtsUpdatesDAO(db),
		UserPtsUpdatesDAO:    mysql_dao.NewUserPtsUpdatesDAO(db),
		UserQtsUpdatesDAO:    mysql_dao.NewUserQtsUpdatesDAO(db),
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

// New new a dao and return.
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

	idgen.NewUUID()
	idgen.NewSeqIDGen()

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
