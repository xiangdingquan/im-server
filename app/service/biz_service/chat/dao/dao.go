package dao

import (
	"context"

	idgen "open.chat/app/service/idgen/client"
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

	idgen.NewUUID()
	idgen.NewSeqIDGen()

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Redis.Close()
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.Redis.Ping(ctx); err != nil {
		return
	}
	return d.Mysql.Ping(ctx)
}
