package core

import (
	"context"
	"math"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (m *ChannelCore) GetChannelParticipants(ctx context.Context, channel *model.ImmutableChannel, filter *mtproto.ChannelParticipantsFilter, offset, limit int32) (int32, []*model.ImmutableChannelParticipant) {
	var (
		count        int32
		participants []*model.ImmutableChannelParticipant
	)

	switch filter.GetPredicateName() {
	case mtproto.Predicate_channelParticipantsRecent:
		participants = m.getChannelParticipantOffsetLimitRecentList(ctx, channel, offset, limit)
		count = m.GetParticipantCount(ctx, channel.Id) //int32(len(participants))
	case mtproto.Predicate_channelParticipantsAdmins:
		participants = m.getChannelParticipantOffsetLimitAdminList(ctx, channel, offset, limit)
		count = int32(len(participants))
	case mtproto.Predicate_channelParticipantsKicked:
		participants = m.getChannelParticipantOffsetLimitKickedList(ctx, channel, filter.Q_STRING, offset, limit)
		count = int32(len(participants))
	case mtproto.Predicate_channelParticipantsBots:
		participants = m.getChannelParticipantOffsetLimitBotList(ctx, channel, offset, limit)
		count = int32(len(participants))
	case mtproto.Predicate_channelParticipantsBanned:
		participants = m.getChannelParticipantOffsetLimitBannedList(ctx, channel, filter.Q_STRING, offset, limit)
		count = int32(len(participants))
	case mtproto.Predicate_channelParticipantsSearch:
		participants = m.getChannelParticipantOffsetLimitListBySearch(ctx, channel, filter.Q_STRING, offset, limit)
		count = int32(len(participants))
	case mtproto.Predicate_channelParticipantsContacts:
		participants = m.getChannelParticipantOffsetLimitContactList(ctx, channel, filter.Q_STRING, offset, limit)
		count = int32(len(participants))
	default:
		log.Errorf("getChannelParticipants - invalid filter: %s", filter.DebugString())
		return 0, nil
	}

	if participants == nil {
		participants = []*model.ImmutableChannelParticipant{}
	}
	return count, participants
}

func (m *ChannelCore) getChannelParticipantOffsetLimitRecentList(ctx context.Context, channel *model.ImmutableChannel, offset, limit int32) []*model.ImmutableChannelParticipant {
	pDOList, err := m.ChannelParticipantsDAO.SelectRecentListOffsetLimit(ctx, channel.Id, offset, limit)
	if err != nil {
		log.Errorf("selectRecentListOffsetLimit error: %v", err)
		return []*model.ImmutableChannelParticipant{}
	}
	pList := make([]*model.ImmutableChannelParticipant, 0, len(pDOList))
	for i := 0; i < len(pDOList); i++ {
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) GetChannelParticipantRecentList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	pDOList, _ := m.ChannelParticipantsDAO.SelectRecentList(ctx, channel.Id)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) getChannelParticipantOffsetLimitAdminList(ctx context.Context, channel *model.ImmutableChannel, offset, limit int32) []*model.ImmutableChannelParticipant {
	pDOList, _ := m.ChannelParticipantsDAO.SelectAdminList(ctx, channel.Id)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) GetChannelParticipantAdminList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	pDOList, _ := m.ChannelParticipantsDAO.SelectAdminList(ctx, channel.Id)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) getChannelParticipantOffsetLimitKickedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit int32) []*model.ImmutableChannelParticipant {
	pDOList, _ := m.ChannelParticipantsDAO.SelectKickedList(ctx, channel.Id)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) GetChannelParticipantKickedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	pDOList, _ := m.ChannelParticipantsDAO.SelectKickedList(ctx, channel.Id)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) getChannelParticipantOffsetLimitBotList(ctx context.Context, channel *model.ImmutableChannel, offset, limit int32) []*model.ImmutableChannelParticipant {
	return []*model.ImmutableChannelParticipant{}
}

func (m *ChannelCore) GetChannelParticipantBotList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	return []*model.ImmutableChannelParticipant{}
}

func (m *ChannelCore) getChannelParticipantOffsetLimitBannedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit int32) []*model.ImmutableChannelParticipant {
	pDOList, _ := m.ChannelParticipantsDAO.SelectBannedList(ctx, channel.Id, math.MaxInt32)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) GetChannelParticipantBannedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	pDOList, _ := m.ChannelParticipantsDAO.SelectBannedList(ctx, channel.Id, math.MaxInt32)
	pList := make([]*model.ImmutableChannelParticipant, 0, limit)
	for i := int(offset); i < len(pDOList); i++ {
		if i >= int(limit+offset) {
			break
		}
		pList = append(pList, makeImmutableChannelParticipant(channel, &pDOList[i]))
	}
	return pList
}

func (m *ChannelCore) getChannelParticipantOffsetLimitListBySearch(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit int32) []*model.ImmutableChannelParticipant {
	return []*model.ImmutableChannelParticipant{}
}

func (m *ChannelCore) GetChannelParticipantListBySearch(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	return []*model.ImmutableChannelParticipant{}
}

func (m *ChannelCore) getChannelParticipantOffsetLimitContactList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit int32) []*model.ImmutableChannelParticipant {
	return m.getChannelParticipantOffsetLimitRecentList(ctx, channel, offset, limit)
}

func (m *ChannelCore) GetChannelParticipantContactList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	if limit > 50 {
		limit = 50
	}
	return []*model.ImmutableChannelParticipant{}
}
