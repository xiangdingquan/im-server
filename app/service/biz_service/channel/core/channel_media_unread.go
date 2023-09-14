package core

import (
	"context"

	"open.chat/pkg/util"
)

func (m *ChannelCore) GetMediaUnread(ctx context.Context, userId, channelId, channelMessageId int32) bool {
	mediaUnreadDO, _ := m.ChannelMediaUnreadDAO.SelectMediaUnread(ctx, userId, channelId, channelMessageId)
	if mediaUnreadDO == nil {
		return util.Int8ToBool(mediaUnreadDO.MediaUnread)
	} else {
		return false
	}
}
