package dao

import (
	"context"

	"open.chat/app/infra/databus/internal/conf"
	"open.chat/pkg/database/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: sqlx.NewMySQL(c.MySQL),
	}
	return
}

func (d *Dao) Ping(c context.Context) error {
	return d.db.Ping(c)
}

func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
