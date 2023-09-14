package core

import (
	"context"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/model"
	"open.chat/pkg/hack"
)

func (m *MessageCore) makeScheduledMessageBox(ctx context.Context, do *dataobject.ScheduledMessagesDO) (box *model.MessageBox) {
	box = &model.MessageBox{
		SelfUserId:        do.UserId,
		SendUserId:        do.UserId,
		MessageId:         do.UserMessageBoxId,
		DialogId:          do.DialogId,
		DialogMessageId:   0,
		MessageDataId:     0,
		RandomId:          do.RandomId,
		Pts:               0,
		PtsCount:          0,
		MessageFilterType: 0,
		MessageBoxType:    do.MessageBoxType,
		MessageType:       do.MessageType,
		Message:           nil,
		Views:             0,
	}

	box.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
	return
}

func (m *MessageCore) GetScheduledMessageListByIdList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (messages model.MessageBoxList) {
	messages = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	doList, _ := m.ScheduledMessagesDAO.SelectByMessageIdList(ctx, userId, int8(peer.PeerType), peer.PeerId, idList)
	for i := 0; i < len(doList); i++ {
		messages = append(messages, m.makeScheduledMessageBox(ctx, &doList[i]))
	}

	return
}

func (m *MessageCore) DeleteScheduledMessageList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (err error) {
	_, err = m.ScheduledMessagesDAO.UpdateStateByMessageId(ctx, model.ScheduledStateDelete, userId, int8(peer.PeerType), peer.PeerId, idList)
	return
}

func (m *MessageCore) DeleteSendedScheduledMessageList(ctx context.Context, idList []int64) (err error) {
	_, err = m.ScheduledMessagesDAO.UpdateStateByIdList(ctx, model.ScheduledStateSended, idList)
	return
}

func (m *MessageCore) GetScheduledMessageHistory(ctx context.Context, userId int32, peer *model.PeerUtil) (messages model.MessageBoxList) {
	doList, _ := m.ScheduledMessagesDAO.SelectHistory(ctx, userId, int8(peer.PeerType), peer.PeerId)
	if len(doList) == 0 {
		messages = []*model.MessageBox{}
	} else {
		messages = make([]*model.MessageBox, 0, len(messages))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeScheduledMessageBox(ctx, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetScheduledTimeoutMessageList(ctx context.Context, date int32) (messages model.MessageBoxList) {
	doList, _ := m.ScheduledMessagesDAO.SelectScheduled(ctx, date)
	if len(doList) == 0 {
		messages = []*model.MessageBox{}
	} else {
		messages = make([]*model.MessageBox, 0, len(messages))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeScheduledMessageBox(ctx, &doList[i]))
		}
	}
	return
}
