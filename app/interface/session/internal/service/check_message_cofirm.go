package service

import (
	"open.chat/mtproto"
)

func checkMessageConfirm(msg mtproto.TLObject) bool {
	switch msg.(type) {
	case *mtproto.TLMsgContainer,
		*mtproto.TLMsgsAck,
		*mtproto.TLHttpWait,
		*mtproto.TLBadMsgNotification,
		*mtproto.TLMsgsAllInfo,
		*mtproto.TLMsgsStateInfo,
		*mtproto.TLMsgDetailedInfo,
		*mtproto.TLMsgNewDetailedInfo:

		return false
	default:
		return true
	}
}
