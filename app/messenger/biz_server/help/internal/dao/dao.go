package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/messenger/biz_server/help/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AppConfigsDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:            db,
		AppConfigsDAO: mysql_dao.NewAppConfigsDAO(db),
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
