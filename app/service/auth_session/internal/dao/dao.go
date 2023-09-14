package dao

import (
	"context"
	"flag"
	"github.com/oschwald/geoip2-golang"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/auth_session/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

var (
	mmdb string
)

func init() {
	flag.StringVar(&mmdb, "mmdb", "./GeoLite2-City.mmdb", "mmdb")
}

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthKeysDAO
	*mysql_dao.AuthOpLogsDAO
	*mysql_dao.AuthUsersDAO
	*mysql_dao.AuthsDAO
	*mysql_dao.DevicesDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:            db,
		AuthKeysDAO:   mysql_dao.NewAuthKeysDAO(db),
		AuthOpLogsDAO: mysql_dao.NewAuthOpLogsDAO(db),
		AuthUsersDAO:  mysql_dao.NewAuthUsersDAO(db),
		AuthsDAO:      mysql_dao.NewAuthsDAO(db),
		DevicesDAO:    mysql_dao.NewDevicesDAO(db),
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
	MMDB *geoip2.Reader
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	var (
		err error
	)

	dao.MMDB, err = geoip2.Open(mmdb)
	checkErr(err)

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.MMDB.Close()
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
