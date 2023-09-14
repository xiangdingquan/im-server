package core

import (
	"context"
	"time"

	"open.chat/model"
)

var countsTable = [8][8][]int32{
	// UNKNOWN
	{nil, {1, 0, 0, 0}, nil, {1, 1, 0, 0}, nil, nil, nil, nil}, //

	// STATE_PARTICIPANT
	{nil, nil, {0, 1, 0, 0}, nil, {0, 0, 0, 1}, {-1, 0, 1, 0}, nil, {-1, 0, 0, 0}},

	// STATE_ADMIN
	{nil, {0, -1, 0, 0}, {0, 0, 0, 0}, nil, {0, -1, 0, 1}, {-1, -1, 1, 0}, nil, {-1, -1, 0, 0}},

	// STATE_CREATOR
	{nil, nil, nil, nil, nil, nil, nil, {-1, -1, 0, 0}},

	// STATE_RESTRICTED
	{nil, {0, 0, 0, -1}, {0, 1, 0, -1}, nil, {0, 0, 0, 0}, {-1, 0, 1, -1}, {-1, 0, 0, -1}, {-1, 0, 0, -1}},

	// STATE_BANNED
	{nil, nil, nil, nil, nil, nil, {0, 0, -1, 0}, nil},

	// STATE_REMOVED
	{nil, {1, 0, 0, 0}, nil, nil, nil, {1, 0, 0, 1}, nil, nil},

	// STATE_LEFT
	{nil, {1, 0, 0, 0}, nil, {1, 1, 0, 0}, nil, nil, nil, nil},
}

func (m *ChannelCore) GetParticipantCounts(ctx context.Context, id int32) (int32, int32, int32, int32) {
	var (
		count       int
		adminsCount int32
		kickedCount int32
		bannedCount int32
	)

	doList, _ := m.ChannelParticipantsDAO.SelectRecentList(ctx, id)
	for _, p := range doList {
		count++
		isAdmin := p.IsCreator == 1 || p.AdminRights != 0
		isKicked := p.State == model.ChatMemberStateKicked
		isRestricted := (p.BannedRights > 0 && (p.BannedRights&model.BAN_VIEW_MESSAGES == 0)) && int32(time.Now().Unix()) <= p.BannedUntilDate
		if isAdmin {
			adminsCount += 1
		}
		if isKicked {
			kickedCount += 1
		}
		if isRestricted {
			bannedCount += 1
		}
	}
	return int32(count), adminsCount, kickedCount, bannedCount
}

func (m *ChannelCore) GetParticipantCount(ctx context.Context, id int32) int32 {
	return int32(m.CommonDAO.CalcSize(ctx, "channel_participants", map[string]interface{}{
		"channel_id": id,
		"state":      0,
	}))
}
