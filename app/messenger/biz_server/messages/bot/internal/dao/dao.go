package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/bots/botpb"
	"open.chat/app/pkg/env2"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/grpc_util/client"
	"open.chat/pkg/log"
)

// Dao dao.
type Dao struct {
	*Redis
	BotsClient botpb.RPCBotsClient
	GifClient  botpb.RPCBotsClient
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New(c *warden.ClientConfig) (dao *Dao, err error) {
	dao = &Dao{
		Redis: newRedisDao(),
	}

	if conn, err2 := client.NewClient(env2.BotsBotFatherId, c); err2 != nil {
		log.Errorf("fail to dial: %v", err)
		return nil, err2
	} else {
		dao.BotsClient = botpb.NewRPCBotsClient(conn)
	}

	if conn, err2 := client.NewClient(env2.BotsGifId, c); err2 != nil {
		log.Errorf("fail to dial: %v", err)
		return nil, err2
	} else {
		dao.GifClient = botpb.NewRPCBotsClient(conn)
	}

	idgen.NewUUID()
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
