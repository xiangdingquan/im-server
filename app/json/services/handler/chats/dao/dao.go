package dao

import (
	"context"

	mysqldao "open.chat/app/json/db/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
)

// Mysql .
type Mysql struct {
	*sqlx.DB
	*mysqldao.ChatsDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:       db,
		ChatsDAO: mysqldao.NewChatsDAO(db),
	}
}

// Close .
func (d *Mysql) Close() error {
	return d.DB.Close()
}

// Ping .
func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

// Dao .
type Dao struct {
	*Mysql
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
