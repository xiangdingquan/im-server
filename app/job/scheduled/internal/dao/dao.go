package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/job/scheduled/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ScheduledMessagesDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := sqlx.NewMySQL(c)
	return &Mysql{
		DB:                   db,
		ScheduledMessagesDAO: mysql_dao.NewScheduledMessagesDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

type Dao struct {
	*Mysql
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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

func (d *Dao) GetUserLangCode(ctx context.Context, user_id int32) (rValue string) {
	var (
		query = "select auth_key_id from auth_users where user_id = ? and deleted = 0 order by updated_at desc limit 1"
	)
	rValue = "en"
	var auth_key_id int64 = 0
	err := d.Get(ctx, &auth_key_id, query, user_id)

	if err != nil {
		log.Errorf("get in SelectAuthKey(_), error: %v", err)
		return
	}

	if err == nil && auth_key_id > 0 {
		var query = "select system_lang_code from auths where auth_key_id = ? limit 1"
		err = d.Get(ctx, &rValue, query, auth_key_id)
		if err != nil {
			log.Errorf("get in SelectAuthKey(_), error: %v", err)
			return
		}
	}
	return
}
