package dao

import (
	"context"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.LangPackStringsDAO
	*mysql_dao.LangPackLanguagesDAO
	*mysql_dao.LanguagesDAO
	*mysql_dao.AppLanguagesDAO
	*mysql_dao.StringsDAO
	*sqlx.CommonDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                   db,
		LangPackStringsDAO:   mysql_dao.NewLangPackStringsDAO(db),
		LangPackLanguagesDAO: mysql_dao.NewLangPackLanguagesDAO(db),
		LanguagesDAO:         mysql_dao.NewLanguagesDAO(db),
		AppLanguagesDAO:      mysql_dao.NewAppLanguagesDAO(db),
		StringsDAO:           mysql_dao.NewStringsDAO(db),
		CommonDAO:            sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
