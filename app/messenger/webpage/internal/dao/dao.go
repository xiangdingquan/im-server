package dao

import (
	"context"

	idgen "open.chat/app/service/idgen/client"
)

type Dao struct {
	*Redis
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() (dao *Dao) {
	dao = new(Dao)
	dao.Redis = newRedisDao()

	idgen.NewUUID()
	idgen.NewSeqIDGen()
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Redis.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.Redis.Ping(ctx)
}
