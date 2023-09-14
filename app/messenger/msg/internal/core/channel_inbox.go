package core

import (
	"context"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func (m *MsgCore) makeChannelMessageBox(ctx context.Context, selfUserId int32, do *dataobject.ChannelMessagesDO) (box *model.MessageBox) {
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
	return
}

func (m *MsgCore) GetLastChannelMessage(ctx context.Context, selfUserId, channelId int32) (*model.MessageBox, error) {
	cDO, _ := m.ChannelsDAO.Select(ctx, channelId)
	if cDO != nil {
		myDO, err := m.ChannelMessagesDAO.SelectByMessageId(ctx, channelId, cDO.TopMessage)
		if err != nil {
			return nil, err
		} else if myDO == nil {
			return nil, mtproto.ErrMessageIdInvalid
		}
		return m.makeChannelMessageBox(ctx, selfUserId, myDO), nil
	} else {
		return nil, nil
	}
}
