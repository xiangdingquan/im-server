package msg_facade

import (
	"context"
	"fmt"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
)

type MsgFacade interface {
	// send
	SendUserMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error)
	SendChatMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error)
	SendChannelMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, randomId int64, message *mtproto.Message) (*mtproto.Updates, error)
	SendUserMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomIds []int64, messages []*mtproto.Message) (*mtproto.Updates, error)
	SendChatMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, randomIds []int64, messages []*mtproto.Message) (*mtproto.Updates, error)
	SendChannelMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, randomIds []int64, messages []*mtproto.Message) (*mtproto.Updates, error)

	// push
	PushUserMessage(ctx context.Context, pushType, fromUserId int32, toUserId int32, randomId int64, message *mtproto.Message) error

	// edit
	EditUserMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, message *mtproto.Message) (*mtproto.Updates, error)
	EditChatMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, message *mtproto.Message) (*mtproto.Updates, error)
	EditChannelMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, message *mtproto.Message) (*mtproto.Updates, error)

	SendMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, message *msgpb.OutboxMessage) (*mtproto.Updates, error)
	SendMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, multiMessage []*msgpb.OutboxMessage) (*mtproto.Updates, error)

	PushMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, message *msgpb.OutboxMessage) (*mtproto.Bool, error)
	EditMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, message *msgpb.OutboxMessage) (*mtproto.Updates, error)
	DeleteMessages(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, revoke bool, id []int32) (*mtproto.Messages_AffectedMessages, error)
	DeleteHistory(ctx context.Context, fromUserId int32, fromAuthKeyId int64, justClear, revoke bool, peer *model.PeerUtil, maxId int32) (*mtproto.Messages_AffectedHistory, error)
	DeleteChannelUserHistory(ctx context.Context, fromUserId int32, fromAuthKeyId int64, channelId int32, peer *model.PeerUtil) (*mtproto.Messages_AffectedHistory, error)

	ReadMessageContents(ctx context.Context, fromUserId int32, fromAuthKeyId int64, peer *model.PeerUtil, id []*msgpb.ContentMessage) (*mtproto.Messages_AffectedMessages, error)
}

type Instance func() MsgFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewMsgFacade(name string) (inst MsgFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
