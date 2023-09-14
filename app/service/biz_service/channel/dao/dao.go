package dao

import (
	"context"
)

// Dao dao.
type Dao struct {
	*Mysql
	*Redis
}

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Mysql.Close()
	d.Redis.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.Redis.Ping(ctx); err != nil {
		return err
	}
	return d.Mysql.Ping(ctx)
}
