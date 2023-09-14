package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/wallet/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.WalletDAO
	*mysql_dao.WalletRecordDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:              db,
		WalletDAO:       mysql_dao.NewWalletDAO(db),
		WalletRecordDAO: mysql_dao.NewWalletRecordDAO(db),
		CommonDAO:       sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
