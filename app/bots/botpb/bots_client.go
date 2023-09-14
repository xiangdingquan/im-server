package botpb

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type BotsClient struct {
	client *databus.Databus
}

func New(name string) (c *BotsClient) {
	var (
		dbus struct {
			Bots       *databus.Config
			Gif        *databus.Config
			Pic        *databus.Config
			Foursquare *databus.Config
			Bing       *databus.Config
		}
	)
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))

	switch name {
	case "bots":
		c = &BotsClient{
			client: databus.New(dbus.Bots),
		}
	case "gif":
		c = &BotsClient{
			client: databus.New(dbus.Gif),
		}
	case "pic":
		c = &BotsClient{
			client: databus.New(dbus.Pic),
		}
	case "foursquare":
		c = &BotsClient{
			client: databus.New(dbus.Foursquare),
		}
	case "bing":
		c = &BotsClient{
			client: databus.New(dbus.Bing),
		}
	default:
		panic("invalid name: " + name)
	}

	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *BotsClient) PushBotUpdates(ctx context.Context, i *BotUpdates) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("PushBotUpdates.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}
