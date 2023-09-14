package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/private/internal/dal/dao/mysql_dao"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ConversationsDAO
	*mysql_dao.DialogFiltersDAO
	*sqlx.CommonDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:               db,
		ConversationsDAO: mysql_dao.NewConversationsDAO(db),
		DialogFiltersDAO: mysql_dao.NewDialogFiltersDAO(db),
		CommonDAO:        sqlx.NewCommonDAO(db),
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

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
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
