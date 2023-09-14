package message_facade

import (
	"context"
	"fmt"

	"open.chat/model"
)

type MessageFacade interface {
	GetPeerUserMessageId(ctx context.Context, userId, messageId, peerUserId int32) int32
	GetPeerUserMessage(ctx context.Context, userId, messageId, peerUserId int32) (*model.MessageBox, error)
	GetPeerChatMessageList(ctx context.Context, userId, messageId, peerChatId int32) map[int32]*model.MessageBox

	GetUserMessage(ctx context.Context, userId int32, id int32) (*model.MessageBox, error)
	GetUserMessageList(ctx context.Context, userId int32, idList []int32) model.MessageBoxList
	GetUserMessageListByDataIdList(ctx context.Context, userId int32, id []int64) model.MessageBoxList

	// channel
	GetChannelMessage(ctx context.Context, userId, channelId, id int32) (*model.MessageBox, error)
	GetChannelMessageList(ctx context.Context, userId, channelId int32, idList []int32) model.MessageBoxList
	GetChannelMessageListByDataIdList(ctx context.Context, userId int32, idList []int64) model.MessageBoxList

	// history
	GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList)
	GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList)

	GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList)
	GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList)

	GetScheduledMessageListByIdList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (messages model.MessageBoxList)
	DeleteScheduledMessageList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (err error)
	DeleteSendedScheduledMessageList(ctx context.Context, idList []int64) (err error)
	GetScheduledMessageHistory(ctx context.Context, userId int32, peer *model.PeerUtil) (messages model.MessageBoxList)
	GetScheduledTimeoutMessageList(ctx context.Context, date int32) (messages model.MessageBoxList)

	UpdateMediaUnread(ctx context.Context, userId int32, id int32)

	SearchByMediaType(ctx context.Context, userId int32, peer *model.PeerUtil, mediaType model.MediaType, minId, offset, limit int32) (messages model.MessageBoxList)
	Search(ctx context.Context, userId int32, peer *model.PeerUtil, q string, minId, offset, limit int32) (messages model.MessageBoxList)
	SearchGlobal(ctx context.Context, userId int32, q string, offset, limit int32) (messages model.MessageBoxList)

	GetHistoryMessagesCount(ctx context.Context, userId int32, peer *model.PeerUtil) int32

	// unread_mentions
	GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList)
	GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList)

	ReadEphemeralMsgByBetween(ctx context.Context, userId int32, peer *model.PeerUtil, minId, maxId int32) bool
	GetEphemeralExpireList(ctx context.Context) []*model.CountDownMessage
	DelEphemeralList(ctx context.Context, messages []*model.CountDownMessage) bool
}

type Instance func() MessageFacade

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

func NewMessageFacade(name string) (inst MessageFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
