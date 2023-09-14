package admin_log_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/job/admin_log/adminlogpb"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type AdminLogClient struct {
	client *databus.Databus
}

func New() *AdminLogClient {
	var (
		dbus struct {
			Adminlog *databus.Config
		}
	)
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))
	log.Debugf("adminlog: %v", dbus.Adminlog)
	return &AdminLogClient{
		client: databus.New(dbus.Adminlog),
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *AdminLogClient) PutChannelAdminLogEventAction(ctx context.Context, i *adminlogpb.ChannelAdminLogEventData) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("PutChannelAdminLogEventAction.send(updates:%v).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}
