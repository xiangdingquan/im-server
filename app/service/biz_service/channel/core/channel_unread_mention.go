package core

import (
	"context"
)

func (m *ChannelCore) GetChannelUnreadMentionsCount(ctx context.Context, userId, channelId, maxId int32) int32 {
	return m.ChannelParticipantsDAO.SelectUnreadMentionsCount(ctx, channelId, userId, maxId)
}
