package core

import (
	"context"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (m *MessageCore) makeChannelMessageBox(ctx context.Context, selfUserId int32, do *dataobject.ChannelMessagesDO) (box *model.MessageBox) {
	box = &model.MessageBox{
		SelfUserId:        selfUserId,
		SendUserId:        do.SenderUserId,
		MessageId:         do.ChannelMessageId,
		DialogId:          model.MakeDialogId(selfUserId, model.PEER_CHANNEL, do.ChannelId),
		DialogMessageId:   0,
		MessageDataId:     do.MessageDataId,
		RandomId:          do.RandomId,
		Pts:               do.Pts,
		PtsCount:          0,
		MessageFilterType: 0,
		MessageBoxType:    model.MESSAGE_BOX_TYPE_CHANNEL,
		MessageType:       do.MessageType,
		Message:           nil,
		TtlSeconds:        do.TtlSeconds,
	}

	box.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
	box.Message.MediaUnread = util.Int8ToBool(do.MediaUnread)
	box.Message.Mentioned = m.CommonDAO.CheckExists(ctx, "channel_unread_mentions", map[string]interface{}{
		"user_id":              selfUserId,
		"channel_id":           do.ChannelId,
		"mentioned_message_id": do.ChannelMessageId,
		"deleted":              0,
	})
	pollId, _ := model.GetPollIdByMessage(box.Message.Media)
	if pollId != 0 {
		poll, _ := m.PollFacade.GetMediaPoll(ctx, selfUserId, pollId)
		if poll != nil {
			box.Message.Media = poll.ToMessageMedia()
		}
	}
	return
}

func (m *MessageCore) GetChannelMessage(ctx context.Context, selfUserId, channelId, id int32) (*model.MessageBox, error) {
	myDO, err := m.ChannelMessagesDAO.SelectByMessageId(ctx, channelId, id)
	if err != nil {
		return nil, err
	} else if myDO == nil {
		return nil, mtproto.ErrMessageIdInvalid
	}
	return m.makeChannelMessageBox(ctx, selfUserId, myDO), nil
}

func (m *MessageCore) GetChannelMessageList(ctx context.Context, selfUserId, channelId int32, idList []int32) (boxList []*model.MessageBox) {
	boxList = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	doList, err := m.ChannelMessagesDAO.SelectByMessageIdList(ctx, selfUserId, channelId, idList)
	if err != nil {
		log.Errorf("getChannelMessageList - error: %v", err)
	} else {
		for i := 0; i < len(doList); i++ {
			boxList = append(boxList, m.makeChannelMessageBox(ctx, selfUserId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetChannelMessageListByDataIdList(ctx context.Context, userId int32, idList []int64) (boxList model.MessageBoxList) {
	boxList = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	doList, err := m.ChannelMessagesDAO.SelectByMessageDataIdList(ctx, idList)
	if err != nil {
		log.Errorf("getChannelMessageListByDataIdList - error: %v", err)
	} else {
		for i := 0; i < len(doList); i++ {
			boxList = append(boxList, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}

	return
}
