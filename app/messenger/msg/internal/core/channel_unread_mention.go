package core

import (
	"context"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
)

func (m *MsgCore) GetChannelUnreadMentionCount(ctx context.Context, userId, channelId int32) int32 {
	params := map[string]interface{}{
		"user_id":    userId,
		"channel_id": channelId,
		"deleted":    0,
	}
	return int32(m.CommonDAO.CalcSize(ctx, "channel_unread_mentions", params))
}

func (m *MsgCore) InsertChannelUnreadMentions(ctx context.Context, channelId int32, idList []int32, msgId int32) {
	if len(idList) == 0 {
		return
	}

	unreadMentions := make([]*dataobject.ChannelUnreadMentionsDO, len(idList))
	for i := 0; i < len(idList); i++ {
		unreadMentions[i] = &dataobject.ChannelUnreadMentionsDO{
			UserId:             idList[i],
			ChannelId:          channelId,
			MentionedMessageId: msgId,
			Deleted:            0,
		}
	}

	m.ChannelUnreadMentionsDAO.InsertBulk(ctx, unreadMentions)
}

func (m *MsgCore) UpdateChannelUnreadReadMention(ctx context.Context, userId int32, channelId, mentionedMessageId int32) {
	m.ChannelUnreadMentionsDAO.Delete(ctx, userId, channelId, mentionedMessageId)
}

func (m *MsgCore) UpdateChannelMediaUnread(ctx context.Context, channelId int32, id int32) error {
	_, err := m.ChannelMessagesDAO.Update(ctx,
		map[string]interface{}{
			"media_unread": 0,
		},
		channelId,
		id)
	return err
}
