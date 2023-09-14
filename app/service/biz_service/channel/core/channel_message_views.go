package core

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

func (m *ChannelCore) GetChannelMessagesViews(ctx context.Context, channelId int32, idList []int32, increment bool) []int32 {
	viewsDOList, _ := m.ChannelMessagesDAO.SelectMessagesViews(ctx, channelId, idList)
	viewsList := make([]int32, 0, len(idList))

	for _, id := range idList {
		views := int32(1)
		for i := 0; i < len(viewsDOList); i++ {
			if viewsDOList[i].ChannelMessageId == id {
				if increment {
					views = viewsDOList[i].Views + 1
				} else {
					views = viewsDOList[i].Views
				}
				break
			}
		}
		viewsList = append(viewsList, views)
	}

	return viewsList
}

func (m *ChannelCore) IncrementChannelMessagesViews(ctx context.Context, channelId int32, idList []int32) {
	m.ChannelMessagesDAO.UpdateMessagesViews(ctx, channelId, idList)
}

func (m *ChannelCore) GetChannelMessagesViews2(ctx context.Context, channelId int32, idList []int32, increment bool) []*mtproto.MessageViews {
	viewsDOList, _ := m.ChannelMessagesDAO.SelectMessagesViews(ctx, channelId, idList)
	viewsList := make([]*mtproto.MessageViews, len(idList))
	for _, id := range idList {
		view := mtproto.MakeTLMessageViews(
			&mtproto.MessageViews{
				Views:    &types.Int32Value{Value: int32(1)},
				Forwards: &types.Int32Value{Value: int32(1)},
				Replies:  nil,
			},
		).To_MessageViews()
		for i := 0; i < len(viewsDOList); i++ {
			if viewsDOList[i].ChannelMessageId == id {
				if increment {
					view.GetViews().Value = viewsDOList[i].Views + 1
				} else {
					view.GetViews().Value = viewsDOList[i].Views
				}
				break
			}
		}
		viewsList = append(viewsList, view)
	}

	return viewsList
}
