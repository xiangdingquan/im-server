package inbox_client

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type InboxClient struct {
	client *databus.Databus
}

func New(name string) (*InboxClient, error) {
	var (
		dbus struct {
			Inbox *databus.Config
			Bots  *databus.Config
		}

		c *databus.Databus
	)
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))
	if name == "inbox" && dbus.Inbox != nil {
		c = databus.New(dbus.Inbox)
	} else if name == "bots" && dbus.Bots != nil {
		c = databus.New(dbus.Bots)
	} else {
		return nil, fmt.Errorf("new inbox_client error")
	}

	return &InboxClient{
		client: c,
	}, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *InboxClient) SendUserMessageToInbox(ctx context.Context, i *msgpb.InboxUserMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("SendUserMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) SendChatMessageToInbox(ctx context.Context, i *msgpb.InboxChatMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("SendChatMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) SendUserMultiMessageToInbox(ctx context.Context, i *msgpb.InboxUserMultiMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("SendUserMultiMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) SendChatMultiMessageToInbox(ctx context.Context, i *msgpb.InboxChatMultiMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("SendChatMultiMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) EditUserMessageToInbox(ctx context.Context, i *msgpb.InboxUserEditMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("EditUserMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) EditChatMessageToInbox(ctx context.Context, i *msgpb.InboxChatEditMessage) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("EditChatMessageToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) DeleteMessagesToInbox(ctx context.Context, i *msgpb.InboxDeleteMessages) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("DeleteMessages.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) DeleteUserHistoryToInbox(ctx context.Context, i *msgpb.InboxUserDeleteHistory) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("DeleteHistory.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) DeleteChatHistoryToInbox(ctx context.Context, i *msgpb.InboxChatDeleteHistory) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("DeleteHistory.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) ReadUserMediaUnreadToInbox(ctx context.Context, i *msgpb.InboxUserReadMediaUnread) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("ReadUserMediaUnreadToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}

func (c *InboxClient) ReadChatMediaUnreadToInbox(ctx context.Context, i *msgpb.InboxChatReadMediaUnread) (err error) {
	if err := c.client.Send(ctx, proto.MessageName(i), i); err != nil {
		log.Errorf("ReadChatMediaUnreadToInbox.send(updates:%s).error(%v)", logger.JsonDebugData(i), err)
	}

	return
}
