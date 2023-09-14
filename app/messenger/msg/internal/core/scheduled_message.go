package core

import (
	"context"
	"fmt"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func (m *MsgCore) DeleteScheduledMessageList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (err error) {
	_, err = m.ScheduledMessagesDAO.UpdateStateByMessageId(ctx, model.ScheduledStateDelete, userId, int8(peer.PeerType), peer.PeerId, idList)
	return
}

func (m *MsgCore) SendScheduledMessage(ctx context.Context, userId int32, peer *model.PeerUtil, randomId int64, scheduledDate int32, message *mtproto.Message) (*model.MessageBox, error) {
	var (
		dialogId    = model.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		outBoxMsgId = int32(idgen.NextScheduledMessageBoxId(ctx, dialogId))
	)

	message.Out = true
	message.PeerId = peer.ToPeer()
	message.Id = outBoxMsgId
	message.Date = scheduledDate
	mType, mData := model.EncodeMessage(message)

	outBoxMsg := &model.MessageBox{
		SelfUserId:        userId,
		SendUserId:        userId,
		MessageId:         int32(outBoxMsgId),
		DialogId:          dialogId,
		DialogMessageId:   int32(outBoxMsgId),
		MessageDataId:     0,
		RandomId:          randomId,
		Pts:               0,
		PtsCount:          0,
		MessageFilterType: 0,
		MessageBoxType:    0,
		MessageType:       int8(mType),
		Message:           message,
	}

	switch peer.PeerType {
	case model.PEER_USER, model.PEER_CHAT:
		outBoxMsg.MessageBoxType = model.MESSAGE_BOX_TYPE_OUTGOING
	case model.PEER_CHANNEL:
		outBoxMsg.MessageBoxType = model.MESSAGE_BOX_TYPE_CHANNEL
	}

	scheduledMessageDO := &dataobject.ScheduledMessagesDO{
		UserId:           outBoxMsg.SendUserId,
		UserMessageBoxId: outBoxMsg.MessageId,
		PeerType:         int8(peer.PeerType),
		PeerId:           peer.PeerId,
		DialogId:         outBoxMsg.DialogId,
		RandomId:         outBoxMsg.RandomId,
		MessageType:      outBoxMsg.MessageType,
		MessageDataType:  int8(model.GetMediaType(message)),
		MessageData:      hack.String(mData),
		MessageBoxType:   outBoxMsg.MessageBoxType,
		ScheduledDate:    scheduledDate,
		Date2:            scheduledDate,
		State:            0,
	}
	log.Debugf("insert scheduledMessageDO: %v", scheduledMessageDO)

	lastInsertId, rowsAffected, err := m.ScheduledMessagesDAO.InsertOrReturnId(ctx, scheduledMessageDO)
	log.Debugf("insert result: %d, %d, %v", lastInsertId, rowsAffected, err)

	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		if lastInsertId > 0 {
			do, err := m.ScheduledMessagesDAO.SelectById(ctx, lastInsertId)
			if err != nil {
				return nil, err
			}
			if do != nil {
				outBoxMsg = &model.MessageBox{
					SelfUserId:        userId,
					SendUserId:        userId,
					MessageId:         do.UserMessageBoxId,
					DialogId:          dialogId,
					DialogMessageId:   do.UserMessageBoxId,
					MessageDataId:     0,
					RandomId:          randomId,
					Pts:               0,
					PtsCount:          0,
					MessageFilterType: 0,
					MessageBoxType:    do.MessageBoxType,
					MessageType:       int8(mType),
					Message:           nil,
				}
				outBoxMsg.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
			} else {
				log.Error("unknown error")
				return nil, fmt.Errorf("fatal unknown error")
			}
		} else {
			return nil, fmt.Errorf("insert error")
		}
	}

	return outBoxMsg, nil
}
