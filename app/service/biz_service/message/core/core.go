package core

import (
	"open.chat/app/service/biz_service/message/dao"
	poll_facade "open.chat/app/service/biz_service/poll/facade"
)

const (
	MESSAGE_TYPE_UNKNOWN          = 0
	MESSAGE_TYPE_MESSAGE_EMPTY    = 1
	MESSAGE_TYPE_MESSAGE          = 2
	MESSAGE_TYPE_MESSAGE_44F9B43D = 2
	MESSAGE_TYPE_MESSAGE_SERVICE  = 3
	MESSAGE_TYPE_MESSAGE_452C0E65 = 4
)
const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
	MESSAGE_BOX_TYPE_CHANNEL  = 2
)

type MessageCore struct {
	*dao.Dao
	poll_facade.PollFacade
}

func New(dao *dao.Dao) *MessageCore {
	poll, _ := poll_facade.NewPollFacade("local")
	return &MessageCore{
		Dao:        dao,
		PollFacade: poll,
	}
}
