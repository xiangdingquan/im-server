package service

import (
	"math"
	"reflect"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (c *session) onMsgsAck(msgId int64, seqno int32, request *mtproto.TLMsgsAck) {
	log.Debugf("onMsgsAck - request data: {sess: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c, msgId, seqno, request.DebugString())

	c.outQueue.OnMsgsAck(request.GetMsgIds(), func(inMsgId int64) {
		if inMsgId > math.MaxInt32 {
			c.inQueue.ChangeAckReceived(inMsgId)
		} else {
		}
	})
}

func (c *session) checkBadServerSalt(gatewayId string, salt int64, msg *mtproto.TLMessage2) bool {
	valid := false

	if salt == c.cb.getCacheSalt().GetSalt() {
		valid = true
	} else {
		if c.cb.getCacheSalt() != nil {
			if salt == c.cb.getCacheSalt().GetSalt() {
				date := int32(time.Now().Unix())
				if c.cb.getCacheSalt().GetValidUntil()+300 >= date {
					valid = true
				}
			}
		}
	}

	if !valid {
		badServerSalt := mtproto.MakeTLBadServerSalt(&mtproto.BadMsgNotification{
			BadMsgId:      msg.MsgId,
			ErrorCode:     kServerSaltIncorrect,
			BadMsgSeqno:   msg.Seqno,
			NewServerSalt: c.cb.getCacheSalt().GetSalt(),
		}).To_BadMsgNotification()
		log.Warnf("invalid salt: %d, send badServerSalt: {%v}, cacheSalt: %v", salt, badServerSalt, c.cb.getCacheSalt())

		c.sendDirectToGateway(gatewayId, false, badServerSalt, func(sentRaw *mtproto.TLMessageRawData) {
		})
		return false
	}

	return valid
}

func (c *session) checkBadMsgNotification(gatewayId string, excludeMsgIdToo bool, msg *mtproto.TLMessage2) bool {
	var errorCode int32 = 0
	serverTime := time.Now().Unix()
	for {
		clientTime := int64(msg.MsgId / 4294967296.0)
		if !excludeMsgIdToo {
			if clientTime+300 < serverTime {
				errorCode = kMsgIdTooLow
				log.Errorf("bad server time from msg_id: %d, my time: %d", clientTime, serverTime)

				break
			}
			if clientTime > serverTime+60 {
				errorCode = kMsgIdTooHigh
				log.Errorf("bad server time from msg_id: %d, my time: %d", clientTime, serverTime)
				break
			}
		}
		if msg.MsgId%4 != 0 {
			errorCode = kMsgIdMod4
			break
		}

		if msg.MsgId < c.inQueue.GetMinMsgId() {
			errorCode = kMsgIdTooOld
			break
		}
		if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
			errorCode = c.checkContainer(msg.MsgId, msg.Seqno, msgContainer)
			if errorCode != 0 {
				break
			}
		}

		if checkMessageConfirm(msg.Object) {
			if msg.Seqno%2 == 0 {
			}
		} else {
			if msg.Seqno%2 != 0 {
			}
		}

		break
	}

	if errorCode != 0 {
		badMsgNotification := mtproto.MakeTLBadMsgNotification(&mtproto.BadMsgNotification{
			BadMsgId:    msg.MsgId,
			BadMsgSeqno: msg.Seqno,
			ErrorCode:   errorCode,
		}).To_BadMsgNotification()
		log.Warnf("errorCode - ", errorCode, ", msg: ", reflect.TypeOf(msg.Object))
		c.sendDirectToGateway(gatewayId, false, badMsgNotification, func(sentRaw *mtproto.TLMessageRawData) {

		})
		return false
	}
	return true
}

func (c *session) onMsgsStateReq(msgId *inboxMsg, request *mtproto.TLMsgsStateReq) {
	log.Debugf("onMsgsStateReq - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
	msgIds := request.GetMsgIds()
	info := make([]byte, len(msgIds))
	for i := 0; i < len(msgIds); i++ {
		if msgIds[i] < c.inQueue.GetMinMsgId() {
			info[i] = UNKNOWN
			continue
		}
		if msgIds[i] > c.inQueue.GetMaxMsgId() {
			info[i] = NOT_RECEIVED_SURE
			continue
		}

		iMsgId := c.inQueue.Lookup(msgIds[i])
		if iMsgId == nil {
			info[i] = NOT_RECEIVED
			continue
		}

		info[i] = iMsgId.state
	}

	msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
		ReqMsgId: msgId.msgId,
		Info:     string(info),
	}).To_MsgsStateInfo()

	c.sendRawToQueue(msgId.msgId, false, msgsStateInfo)
	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onMsgsStateInfo(msgId *inboxMsg, request *mtproto.TLMsgsStateInfo) {
	log.Debugf("onMsgsStateInfo - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
	reqMsgId := request.GetReqMsgId()
	oMsg := c.outQueue.Lookup(reqMsgId)
	if oMsg == nil {
		log.Errorf("not found reqMsgId - %d", reqMsgId)
		return
	}

	var (
		msgIds []int64
		info   = []byte(request.GetInfo())
	)

	dBuf := mtproto.NewDecodeBuf(oMsg.msg.Body)
	r := dBuf.Object()
	switch tl := r.(type) {
	case *mtproto.TLMsgsStateReq:
		msgIds = tl.GetMsgIds()
	case *mtproto.TLMsgResendReq:
		msgIds = tl.GetMsgIds()
	default:
		log.Errorf("not found reqMsgId - %d", reqMsgId)
		return
	}

	if len(msgIds) != len(info) {
		log.Errorf("invalid msgIds, len(msgIds) != len(info)")
		return
	}

	ackIds := make([]int64, 0, len(msgIds))
	ackIds = append(ackIds, reqMsgId)
	resendIds := make([]int64, 0, len(msgIds))
	for i := 0; i < len(msgIds); i++ {
		if info[i] == UNKNOWN || info[i] == NOT_RECEIVED_SURE {
		} else if info[i] == NOT_RECEIVED {
			resendIds = append(resendIds, msgIds[i])
		} else {
			ackIds = append(ackIds, msgIds[i])
		}
	}
	c.outQueue.OnMsgsAck(ackIds, func(inMsgId int64) {
		if inMsgId <= math.MaxInt32 {
		}
	})

	if len(resendIds) > 0 {
	}

	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onMsgsAllInfo(msgId *inboxMsg, request *mtproto.TLMsgsAllInfo) {
	log.Debugf("onMsgsAllInfo - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
	var (
		msgIds = request.GetMsgIds()
		info   = []byte(request.GetInfo())
	)

	ackIds := make([]int64, 0, len(msgIds))
	resendIds := make([]int64, 0, len(msgIds))

	for i := 0; i < len(msgIds); i++ {
		if info[i] == UNKNOWN || info[i] == NOT_RECEIVED_SURE {
		} else if info[i] == NOT_RECEIVED {
			resendIds = append(resendIds, msgIds[i])
		} else {
			ackIds = append(ackIds, msgIds[i])
		}
	}
	c.outQueue.OnMsgsAck(ackIds, func(inMsgId int64) {
		if inMsgId <= math.MaxInt32 {
		}
	})

	if len(resendIds) > 0 {
	}

	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onMsgResendReq(msgId *inboxMsg, request *mtproto.TLMsgResendReq) {
	log.Warnf("onMsgResendReq - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
	var (
		msgIds         = request.GetMsgIds()
		info           = make([]byte, len(msgIds))
		bMsgsStateInfo = false
	)

	for i := 0; i < len(msgIds); i++ {
		if msgIds[i] < c.inQueue.GetMinMsgId() {
			bMsgsStateInfo = true
			break
		}
		if msgIds[i] > c.inQueue.GetMaxMsgId() {
			bMsgsStateInfo = true
			break
		}

		iMsgId := c.inQueue.Lookup(msgIds[i])
		if iMsgId == nil {
			bMsgsStateInfo = true
			break
		}
	}

	if bMsgsStateInfo {
		for i := 0; i < len(msgIds); i++ {
			if msgIds[i] < c.inQueue.GetMinMsgId() {
				info[i] = UNKNOWN
				continue
			}
			if msgIds[i] > c.inQueue.GetMaxMsgId() {
				info[i] = NOT_RECEIVED_SURE
				continue
			}
			iMsgId := c.inQueue.Lookup(msgIds[i])
			if iMsgId == nil {
				info[i] = NOT_RECEIVED
				continue
			}

			info[i] = iMsgId.state
		}

		msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
			ReqMsgId: msgId.msgId,
			Info:     string(info),
		}).To_MsgsStateInfo()

		c.sendRawToQueue(msgId.msgId, false, msgsStateInfo)
		msgId.state = RECEIVED | NEED_NO_ACK
	} else {
		for i := 0; i < len(msgIds); i++ {
		}
	}
}

func (c *session) onMsgDetailInfo(msgId *inboxMsg, request *mtproto.TLMsgDetailedInfo) {
	log.Warnf("onMsgDetailInfo - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
}

func (c *session) onMsgNewDetailInfo(msgId *inboxMsg, request *mtproto.TLMsgNewDetailedInfo) {
	log.Warnf("onMsgNewDetailInfo - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
}

func (c *session) notifyMsgsStateInfo(gatewayId string, inMsg *inboxMsg) {
	msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
		ReqMsgId: inMsg.msgId,
		Info:     string([]byte{inMsg.state}),
	})
	c.sendDirectToGateway(gatewayId, false, msgsStateInfo, func(sentRaw *mtproto.TLMessageRawData) {

	})
}

func (c *session) notifyMsgsStateReq() {
}

func (c *session) notifyMsgsAllInfo() {
}

func (c *session) notifyMsgDetailedInfo(inMsg *inboxMsg) {
}

func (c *session) notifyNewMsgDetailedInfo() {
}

func (c *session) notifyMsgResendAnsSeq() {
}
