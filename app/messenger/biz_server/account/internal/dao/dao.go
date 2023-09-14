package dao

import (
	"context"

	"open.chat/app/messenger/biz_server/account/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ThemesDAO
	*mysql_dao.ThemeFormatsDAO
	*mysql_dao.UserThemesDAO
	*mysql_dao.BannedIpDAO
	*mysql_dao.AuthsDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:              db,
		ThemesDAO:       mysql_dao.NewThemesDAO(db),
		ThemeFormatsDAO: mysql_dao.NewThemeFormatsDAO(db),
		UserThemesDAO:   mysql_dao.NewUserThemesDAO(db),
		BannedIpDAO:     mysql_dao.NewBannedIpDAO(db),
		AuthsDAO:        mysql_dao.NewAuthsDAO(db),
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
	dao = &Dao{
		Mysql: newMysqlDao(),
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
