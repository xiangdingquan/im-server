package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/dfs/facade"
	_ "open.chat/app/service/dfs/facade/cachefs"
	_ "open.chat/app/service/dfs/facade/dfs"
	"open.chat/app/service/media/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.DocumentsDAO
	*mysql_dao.EncryptedFilesDAO
	*mysql_dao.PhotoDatasDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                db,
		DocumentsDAO:      mysql_dao.NewDocumentsDAO(db),
		EncryptedFilesDAO: mysql_dao.NewEncryptedFilesDAO(db),
		PhotoDatasDAO:     mysql_dao.NewPhotoDatasDAO(db),
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
	dfs_facade.DfsFacade
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
	dao.DfsFacade, _ = dfs_facade.NewDfsFacade("dfs")

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
