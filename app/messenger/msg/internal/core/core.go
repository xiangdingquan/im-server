package core

import (
	"open.chat/app/messenger/msg/internal/dal/dataobject"
	"open.chat/app/messenger/msg/internal/dao"
	"open.chat/model"
	"open.chat/pkg/hack"
)

type MsgCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *MsgCore {
	return &MsgCore{
		Dao: dao,
	}
}

func makeMessageBox(do *dataobject.MessagesDO) (box *model.MessageBox) {
	box = &model.MessageBox{
		SelfUserId:        do.UserId,
		SendUserId:        do.SenderUserId,
		MessageId:         do.UserMessageBoxId,
		DialogId:          do.DialogId,
		DialogMessageId:   do.DialogMessageId,
		RandomId:          do.RandomId,
		Pts:               do.Pts,
		PtsCount:          do.PtsCount,
		MessageFilterType: int8(do.MessageDataType),
		MessageBoxType:    do.MessageBoxType,
		MessageType:       do.MessageType,
		Message:           nil,
		TtlSeconds:        do.TtlSeconds,
	}
	box.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
	return
}
