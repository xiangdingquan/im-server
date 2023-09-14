package core

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

type ChannelMessage struct {
	OwnerId          int32
	SenderUserId     int32
	ChannelId        int32
	ChannelMessageId int32
	RandomId         int64
	Mentioned        bool
	MediaUnread      bool
	HasMediaUnread   bool
	Views            int32
	Pts              int32
	PtsCount         int32
	Message          *mtproto.Message
	EditMessage      string
	EditDate         int32
	Date2            int32
	*ChannelCore
}

func (m *ChannelMessage) ToMessage(ctx context.Context, toUserId int32) *mtproto.Message {
	message := proto.Clone(m.Message).(*mtproto.Message)
	message.Id = m.ChannelMessageId
	if m.Views != 0 {
		message.Views = &types.Int32Value{Value: m.Views}
		message.Forwards = &types.Int32Value{Value: m.Views}
	}

	if m.SenderUserId == toUserId {
		message.Out = true
		message.MediaUnread = false
	} else {
		message.Out = false
		if m.HasMediaUnread {
			message.MediaUnread = m.ChannelCore.GetMediaUnread(ctx, toUserId, m.ChannelId, m.ChannelMessageId)
		}
	}

	if m.EditDate != 0 {
		message.Message = m.EditMessage
		message.EditDate = &types.Int32Value{Value: m.EditDate}
	}

	return message
}

func (m *ChannelMessage) SaveEditMessage(ctx context.Context) error {
	_, err := m.ChannelMessagesDAO.Update(ctx, map[string]interface{}{
		"edit_date":    m.EditDate,
		"edit_message": m.EditMessage,
	}, m.ChannelId, m.ChannelMessageId)
	return err
}

func (m *ChannelCore) makeMessageByDO(boxDO *dataobject.ChannelMessagesDO) *ChannelMessage {
	boxMsg := &ChannelMessage{
		SenderUserId:     boxDO.SenderUserId,
		ChannelId:        boxDO.ChannelId,
		ChannelMessageId: boxDO.ChannelMessageId,
		RandomId:         boxDO.RandomId,
		Views:            boxDO.Views,
		EditDate:         boxDO.EditDate,
		EditMessage:      boxDO.EditMessage,
		ChannelCore:      m,
	}
	boxMsg.Message, _ = model.DecodeMessage(int(boxDO.MessageType), hack.Bytes(boxDO.MessageData))
	return boxMsg
}

func (m *ChannelCore) GetMessageListByIdList(ctx context.Context, toUserId, channelId int32, idList []int32) ([]*mtproto.Message, error) {
	if len(idList) == 0 {
		return []*mtproto.Message{}, nil
	}

	boxMsgDOList, err := m.ChannelMessagesDAO.SelectByMessageIdList(ctx, toUserId, channelId, idList)
	if err != nil {
		return nil, err
	}
	//log.Debugf("boxMsgDOList - %v", boxMsgDOList)

	messageList := make([]*mtproto.Message, 0, len(boxMsgDOList))

	for i := 0; i < len(boxMsgDOList); i++ {
		boxMsg := m.makeMessageByDO(&boxMsgDOList[i])
		messageList = append(messageList, boxMsg.ToMessage(ctx, toUserId))
	}

	return messageList, nil
}
