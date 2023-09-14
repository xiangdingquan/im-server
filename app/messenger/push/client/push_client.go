package push_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/messenger/push/pushpb"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type PushClient struct {
	client *databus.Databus
}

func New() *PushClient {
	var (
		dbus struct {
			Push *databus.Config
		}
	)
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))

	return &PushClient{
		client: databus.New(dbus.Push),
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *PushClient) PushUpdatesIfNot(ctx context.Context, i *pushpb.PushUpdatesIfNot) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("PushUpdatesIfNot.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *PushClient) PushUpdates(ctx context.Context, i *pushpb.PushUpdates) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("PushUpdates.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}
