package message_facade

import (
	"context"

	"open.chat/app/service/biz_service/message/core"
	"open.chat/app/service/biz_service/message/dao"
	"open.chat/model"
)

type localMessageFacade struct {
	*core.MessageCore
}

func localMessageFacadeInstance() MessageFacade {
	return &localMessageFacade{
		MessageCore: core.New(dao.New()),
	}
}

// message
func (c *localMessageFacade) GetPeerUserMessageId(ctx context.Context, userId, messageId, peerUserId int32) int32 {
	return c.MessageCore.GetPeerUserMessageId(ctx, userId, messageId, peerUserId)
}

func (c *localMessageFacade) GetPeerUserMessage(ctx context.Context, userId, messageId, peerUserId int32) (*model.MessageBox, error) {
	return c.MessageCore.GetPeerUserMessage(ctx, userId, messageId, peerUserId)
}

func (c *localMessageFacade) GetPeerChatMessageList(ctx context.Context, userId, messageId, peerChatId int32) map[int32]*model.MessageBox {
	return c.MessageCore.GetPeerChatMessageList(ctx, userId, messageId, peerChatId)
}

func (c *localMessageFacade) GetUserMessage(ctx context.Context, userId int32, id int32) (*model.MessageBox, error) {
	return c.MessageCore.GetUserMessage(ctx, userId, id)
}

func (c *localMessageFacade) GetUserMessageList(ctx context.Context, userId int32, idList []int32) model.MessageBoxList {
	return c.MessageCore.GetUserMessageList(ctx, userId, idList)
}

func (c *localMessageFacade) GetUserMessageListByDataIdList(ctx context.Context, userId int32, id []int64) model.MessageBoxList {
	return c.MessageCore.GetUserMessageListByDataIdList(ctx, userId, id)
}

// channel message
func (c *localMessageFacade) GetChannelMessage(ctx context.Context, userId, channelId, id int32) (*model.MessageBox, error) {
	return c.MessageCore.GetChannelMessage(ctx, userId, channelId, id)
}

func (c *localMessageFacade) GetChannelMessageList(ctx context.Context, userId, channelId int32, idList []int32) model.MessageBoxList {
	return c.MessageCore.GetChannelMessageList(ctx, userId, channelId, idList)
}

func (c *localMessageFacade) GetChannelMessageListByDataIdList(ctx context.Context, userId int32, idList []int64) model.MessageBoxList {
	return c.MessageCore.GetChannelMessageListByDataIdList(ctx, userId, idList)
}

// offset
func (c *localMessageFacade) GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetIdBackwardHistoryMessages(ctx, userId, peer, offsetId, minId, maxId, limit, hash)
}

func (c *localMessageFacade) GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetIdForwardHistoryMessages(ctx, userId, peer, offsetId, minId, maxId, limit, hash)
}

func (c *localMessageFacade) GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetDateBackwardHistoryMessages(ctx, userId, peer, offsetDate, minId, maxId, limit, hash)
}

func (c *localMessageFacade) GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetDateForwardHistoryMessages(ctx, userId, peer, offsetDate, minId, maxId, limit, hash)
}

// scheduled_message
func (c *localMessageFacade) GetScheduledMessageListByIdList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetScheduledMessageListByIdList(ctx, userId, peer, idList)
}

func (c *localMessageFacade) DeleteScheduledMessageList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (err error) {
	return c.MessageCore.DeleteScheduledMessageList(ctx, userId, peer, idList)
}

func (c *localMessageFacade) DeleteSendedScheduledMessageList(ctx context.Context, idList []int64) (err error) {
	return c.MessageCore.DeleteSendedScheduledMessageList(ctx, idList)
}

func (c *localMessageFacade) GetScheduledMessageHistory(ctx context.Context, userId int32, peer *model.PeerUtil) (messages model.MessageBoxList) {
	return c.MessageCore.GetScheduledMessageHistory(ctx, userId, peer)
}

func (c *localMessageFacade) GetScheduledTimeoutMessageList(ctx context.Context, date int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetScheduledTimeoutMessageList(ctx, date)
}

func (c *localMessageFacade) UpdateMediaUnread(ctx context.Context, userId int32, id int32) {
	c.MessageCore.UpdateMediaUnread(ctx, userId, id)
}

func (c *localMessageFacade) SearchByMediaType(ctx context.Context, userId int32, peer *model.PeerUtil, mediaType model.MediaType, minId, offset, limit int32) (messages model.MessageBoxList) {
	return c.MessageCore.SearchByMediaType(ctx, userId, peer, mediaType, minId, offset, limit)
}

func (c *localMessageFacade) Search(ctx context.Context, userId int32, peer *model.PeerUtil, q string, minId, offset, limit int32) (messages model.MessageBoxList) {
	return c.MessageCore.Search(ctx, userId, peer, q, minId, offset, limit)
}

func (c *localMessageFacade) SearchGlobal(ctx context.Context, userId int32, q string, offset, limit int32) (messages model.MessageBoxList) {
	return c.MessageCore.SearchGlobal(ctx, userId, q, offset, limit)
}

func (c *localMessageFacade) GetHistoryMessagesCount(ctx context.Context, userId int32, peer *model.PeerUtil) int32 {
	return c.MessageCore.GetHistoryMessagesCount(ctx, userId, peer)
}

func (c *localMessageFacade) GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetIdBackwardUnreadMentions(ctx, userId, peer, offsetId, minId, maxId, limit)
}

func (c *localMessageFacade) GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList) {
	return c.MessageCore.GetOffsetIdForwardUnreadMentions(ctx, userId, peer, offsetId, minId, maxId, limit)
}

func (c *localMessageFacade) ReadEphemeralMsgByBetween(ctx context.Context, userId int32, peer *model.PeerUtil, minId, maxId int32) bool {
	return c.MessageCore.ReadEphemeralMsgByBetween(ctx, userId, peer, minId, maxId)
}

func (c *localMessageFacade) GetEphemeralExpireList(ctx context.Context) []*model.CountDownMessage {
	return c.MessageCore.GetEphemeralExpireList(ctx)
}

func (c *localMessageFacade) DelEphemeralList(ctx context.Context, messages []*model.CountDownMessage) bool {
	return c.MessageCore.DelEphemeralList(ctx, messages)
}

func init() {
	Register("local", localMessageFacadeInstance)
}
