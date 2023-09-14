package msg_facade

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/pkg/env2"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util/client"
	"open.chat/pkg/log"
)

var (
	_self msgpb.RPCMsgClient
)

type eMsgFacade struct {
	client msgpb.RPCMsgClient
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func eMessageFacadeInstance() MsgFacade {
	if _self == nil {
		var (
			ac struct {
				Wardenclient *warden.ClientConfig
			}
		)

		checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
		conn, err := client.NewRPCClient(env2.MessengerOutboxId, ac.Wardenclient)
		checkErr(err)
		_self = msgpb.NewRPCMsgClient(conn.GetClientConn())
	}

	return &eMsgFacade{
		client: _self,
	}
}

func init() {
	Register("emsg", eMessageFacadeInstance)
}

// facade
func (c *eMsgFacade) SendUserMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error) {
	userMessage := &msgpb.UserMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerUserId: toUserId,
		RandomId:   randomId,
		Message:    message,
	}

	return c.client.SendUserMessage(ctx, userMessage)
}

func (c *eMsgFacade) SendChatMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error) {
	chatMessage := &msgpb.ChatMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChatId: chatId,
		RandomId:   randomId,
		Message:    message,
	}

	return c.client.SendChatMessage(ctx, chatMessage)
}

func (c *eMsgFacade) SendChannelMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error) {
	channelMessage := &msgpb.ChannelMessage{
		From:          &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChannelId: channelId,
		RandomId:      randomId,
		Message:       message,
	}

	return c.client.SendChannelMessage(ctx, channelMessage)
}

func (c *eMsgFacade) SendUserMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomId []int64, message []*mtproto.Message) (*mtproto.Updates, error) {
	userMessage := &msgpb.UserMultiMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerUserId: toUserId,
		RandomId:   randomId,
		Message:    message,
	}

	return c.client.SendUserMultiMessage(ctx, userMessage)
}

func (c *eMsgFacade) SendChatMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, randomId []int64, message []*mtproto.Message) (*mtproto.Updates, error) {
	chatMessage := &msgpb.ChatMultiMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChatId: chatId,
		RandomId:   randomId,
		Message:    message,
	}

	return c.client.SendChatMultiMessage(ctx, chatMessage)
}

func (c *eMsgFacade) SendChannelMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, randomId []int64, message []*mtproto.Message) (*mtproto.Updates, error) {
	channelMessage := &msgpb.ChannelMultiMessage{
		From:          &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChannelId: channelId,
		RandomId:      randomId,
		Message:       message,
	}

	return c.client.SendChannelMultiMessage(ctx, channelMessage)
}

func (c *eMsgFacade) PushUserMessage(ctx context.Context, pushType, fromUserId int32, toUserId int32, randomId int64, message *mtproto.Message) error {
	userMessage := &msgpb.UserMessage{
		From:       &msgpb.Sender{Id: fromUserId, Type: pushType},
		PeerUserId: toUserId,
		RandomId:   randomId,
		Message:    message,
	}

	_, err := c.client.PushUserMessage(ctx, userMessage)
	if err != nil {
		log.Errorf("pushUserMessage error - %v", err)
	}
	return err
}

func (c *eMsgFacade) EditUserMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, message *mtproto.Message) (*mtproto.Updates, error) {
	userMessage := &msgpb.UserMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerUserId: toUserId,
		Message:    message,
	}

	return c.client.EditUserMessage(ctx, userMessage)
}

func (c *eMsgFacade) EditChatMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, message *mtproto.Message) (*mtproto.Updates, error) {
	chatMessage := &msgpb.ChatMessage{
		From:       &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChatId: chatId,
		Message:    message,
	}

	return c.client.SendChatMessage(ctx, chatMessage)
}

func (c *eMsgFacade) EditChannelMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, message *mtproto.Message) (*mtproto.Updates, error) {
	channelMessage := &msgpb.ChannelMessage{
		From:          &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerChannelId: channelId,
		Message:       message,
	}

	return c.client.SendChannelMessage(ctx, channelMessage)
}

func (c *eMsgFacade) SendMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, message *msgpb.OutboxMessage) (*mtproto.Updates, error) {
	outgoingMessage := &msgpb.OutgoingMessage{
		From:     &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
		Message:  message,
	}
	return c.client.SendMessage(ctx, outgoingMessage)
}

func (c *eMsgFacade) SendMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, multiMessage []*msgpb.OutboxMessage) (*mtproto.Updates, error) {
	outgoingMultiMedia := &msgpb.OutgoingMultiMessage{
		From:         &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType:     peer.PeerType,
		PeerId:       peer.PeerId,
		MultiMessage: multiMessage,
	}
	return c.client.SendMultiMessage(ctx, outgoingMultiMedia)
}

func (c *eMsgFacade) PushMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, message *msgpb.OutboxMessage) (*mtproto.Bool, error) {
	outgoingMessage := &msgpb.OutgoingMessage{
		From:     &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType: model.PEER_USER,
		PeerId:   toUserId,
		Message:  message,
	}
	return c.client.PushMessage(ctx, outgoingMessage)
}

func (c *eMsgFacade) EditMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, message *msgpb.OutboxMessage) (*mtproto.Updates, error) {
	outgoingMessage := &msgpb.OutgoingMessage{
		From:     &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
		Message:  message,
	}
	return c.client.EditMessage(ctx, outgoingMessage)
}

func (c *eMsgFacade) DeleteMessages(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, revoke bool, id []int32) (*mtproto.Messages_AffectedMessages, error) {
	r := &msgpb.DeleteMessagesRequest{
		From:     &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
		Revoke:   revoke,
		Id:       id,
	}
	return c.client.DeleteMessages(ctx, r)
}

func (c *eMsgFacade) DeleteHistory(ctx context.Context, fromUserId int32, fromAuthKeyId int64, justClear, revoke bool, peer *model.PeerUtil, maxId int32) (*mtproto.Messages_AffectedHistory, error) {
	r := &msgpb.DeleteHistoryRequest{
		From:      &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		ChannelId: 0,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		JustClear: justClear,
		Revoke:    revoke,
		MaxId:     maxId,
	}
	return c.client.DeleteHistory(ctx, r)
}

func (c *eMsgFacade) DeleteChannelUserHistory(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, peer *model.PeerUtil) (*mtproto.Messages_AffectedHistory, error) {
	r := &msgpb.DeleteHistoryRequest{
		From:      &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		ChannelId: channelId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		JustClear: false,
		Revoke:    false,
		MaxId:     0,
	}

	affectedHistory, err := c.client.DeleteHistory(ctx, r)
	if err != nil {
		return nil, err
	}

	return affectedHistory, nil
}

func (c *eMsgFacade) ReadMessageContents(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, id []*msgpb.ContentMessage) (*mtproto.Messages_AffectedMessages, error) {
	r := &msgpb.ReadMessageContentsRequest{
		From:     &msgpb.Sender{Id: fromUserId, AuthKeyId: fromAuthKeyId},
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
		Id:       id,
	}

	affectedMessages, err := c.client.ReadMessageContents(ctx, r)
	if err != nil {
		return nil, err
	}

	return affectedMessages, nil
}
