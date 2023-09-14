package mysql_util

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"open.chat/pkg/database/sqlx"
)

var _self *sqlx.DB

func GetSingletonSqlxDB() *sqlx.DB {
	if _self == nil {
		var (
			dc struct {
				Mysql *sqlx.Config
			}
		)

		checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
		_self = sqlx.NewMySQL(dc.Mysql)
	}

	return _self
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
