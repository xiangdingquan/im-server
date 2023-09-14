package core

import (
	"context"

	"open.chat/app/service/biz_service/channel/dao"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type ChannelCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *ChannelCore {
	return &ChannelCore{
		Dao: dao,
	}
}

// GetChannelListBySelfAndIDList
func (m *ChannelCore) GetChannelListByIdList(ctx context.Context, selfUserId int32, idList []int32) (chats []*mtproto.Chat) {
	if len(idList) == 0 {
		return []*mtproto.Chat{}
	}

	chats = make([]*mtproto.Chat, 0, len(idList))

	for _, id := range idList {
		channel, err := m.GetMutableChannel(ctx, id, selfUserId)
		if err != nil {
			log.Errorf("getChatListBySelfIDList - not find chat_id: %d", id)
		} else {
			chats = append(chats, channel.ToUnsafeChat(selfUserId))
		}
	}
	return
}

func (m *ChannelCore) GetChannelById(ctx context.Context, selfUserId, channelId int32) (chat *mtproto.Chat) {
	channel, err := m.GetMutableChannel(ctx, channelId, selfUserId)
	if err != nil {
		log.Errorf("GetChannelBySelfID - not find chat_id: %d", channelId)
	} else {
		chat = channel.ToUnsafeChat(selfUserId)
	}

	return
}

func (m *ChannelCore) GetChannelParticipantIdList(ctx context.Context, channelId int32) []int32 {
	doList2, _ := m.ChannelParticipantsDAO.SelectByChannelId(ctx, channelId)
	idList := make([]int32, 0, len(doList2))
	for i := 0; i < len(doList2); i++ {
		if doList2[i].State == 0 {
			idList = append(idList, doList2[i].UserId)
		}
	}
	return idList
}

func (m *ChannelCore) GetChannelAdminParticipantIdList(ctx context.Context, channelId int32) []int32 {
	doList2, _ := m.ChannelParticipantsDAO.SelectAdminList(ctx, channelId)
	idList := make([]int32, 0, len(doList2))
	for i := 0; i < len(doList2); i++ {
		if doList2[i].State == 0 {
			idList = append(idList, doList2[i].UserId)
		}
	}
	return idList
}

func (m *ChannelCore) GetTopMessageListByIdList(ctx context.Context, idList []int32) (topMessages map[int32]int32) {
	topMessages = make(map[int32]int32)

	if len(idList) > 0 {
		doList, _ := m.ChannelsDAO.SelectByIdList(ctx, idList)
		for i := 0; i < len(doList); i++ {
			topMessages[doList[i].Id] = doList[i].TopMessage
		}
	}

	return
}

func (m *ChannelCore) GetAvailableMinId(ctx context.Context, userId, channelId int32) int32 {
	id, _ := m.ChannelParticipantsDAO.SelectAvailableMinId(ctx, userId, channelId)
	return id
}

func (m *ChannelCore) SearchChannelByTitle(ctx context.Context, q string) (idList []int32) {
	idList, _ = m.ChannelsDAO.SelectByTitle(ctx, q)
	return
}

func (m *ChannelCore) GetUsersChannelIdList(ctx context.Context, id []int32) map[int32][]int32 {
	var idList2 = make(map[int32][]int32, len(id))

	chatDOList, err := m.ChannelParticipantsDAO.SelectUsersChannelIdList(ctx, id)
	if err != nil {
		log.Errorf("getMyChatIdList - error: %v", err)
	}
	for i := 0; i < len(chatDOList); i++ {
		idList2[chatDOList[i].UserId] = append(idList2[chatDOList[i].UserId], chatDOList[i].ChannelId)
	}

	return idList2
}

func (m *ChannelCore) SearchPublicChannel(ctx context.Context, offset int32, limit int32) (idList []int32) {
	idList, _ = m.ChannelsDAO.SelectPublicChannels(ctx, offset, limit)
	return
}
