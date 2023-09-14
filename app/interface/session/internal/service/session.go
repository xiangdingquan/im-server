package service

import (
	"context"
	"fmt"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	kDefaultPingTimeout  = 60
	kPingAddTimeout      = 15
	kCacheSessionTimeout = 3 * 60
	waitMsgAcksTimeout   = 30
)

const (
	kStateNew = iota
	kStateOnline
	kStateOffline
	kStateClose
)

const (
	kSessionStateNew = iota
	kSessionStateCreated
)

const (
	kServerSaltIncorrect = int32(48)
)

const (
	kMsgIdTooLow    = int32(16)
	kMsgIdTooHigh   = int32(17)
	kMsgIdMod4      = int32(18)
	kMsgIdCollision = int32(19)

	kMsgIdTooOld = int32(20)

	kSeqNoTooLow  = int32(32)
	kSeqNoTooHigh = int32(33)
	kSeqNoNotEven = int32(34)
	kSeqNoNotOdd  = int32(35)

	kInvalidContainer = int32(64)
)

var emptyMsgContainer = mtproto.NewTLMsgRawDataContainer()
var androidPushTooLong = mtproto.MakeTLUpdatesTooLong(nil)

type messageData struct {
	confirmFlag  bool
	compressFlag bool
	obj          mtproto.TLObject
}

func (m *messageData) String() string {
	return fmt.Sprintf("{confirmFlag: %v, compressFlag: %v, obj: {%s}}", m.confirmFlag, m.compressFlag, m.obj)
}

type serverIdCtx struct {
	gatewayId       string
	lastReceiveTime int64
}

func (c serverIdCtx) Equal(id string) bool {
	return c.gatewayId == id
}

type sessionCallback interface {
	getCacheSalt() *mtproto.TLFutureSalt

	getAuthKeyId() int64
	getTempAuthKeyId() int64

	getUserId() int32
	setUserId(userId int32)

	getLayer() int32
	setLayer(layer int32)

	setClient(c string)
	getClient() string

	setIpAddress(c string)
	getIpAddress() string

	destroySession(sessionId int64) bool

	sendToRpcQueue(rpcMessage *rpcApiMessage)

	onBindPushSessionId(sessionId int64)
	setOnline()
	trySetOffline()
}

type session struct {
	sessionId       int64
	sessionState    int
	gatewayIdList   []serverIdCtx
	nextSeqNo       uint32
	firstMsgId      int64
	connState       int
	closeDate       int64
	lastReceiveTime int64
	isAndroidPush   bool
	isGeneric       bool
	inQueue         *sessionInboundQueue
	outQueue        *sessionOutgoingQueue
	pendingQueue    *sessionRpcResultWaitingQueue
	pushQueue       *sessionPushQueue
	isHttp          bool
	httpQueue       *httpRequestQueue
	cb              sessionCallback
	*authSessions
}

func newSession(sessionId int64, sesses *authSessions) *session {
	sess := &session{
		sessionId:       sessionId,
		gatewayIdList:   make([]serverIdCtx, 0, 1),
		sessionState:    kSessionStateNew,
		closeDate:       time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout,
		connState:       kStateNew,
		lastReceiveTime: time.Now().UnixNano(),
		inQueue:         newSessionInboundQueue(),
		outQueue:        newSessionOutgoingQueue(),
		pendingQueue:    newSessionRpcResultWaitingQueue(),
		pushQueue:       newSessionPushQueue(),
		isHttp:          false,
		httpQueue:       newHttpRequestQueue(),
		cb:              sesses,
		authSessions:    sesses,
	}

	return sess
}

func (c *session) String() string {
	return fmt.Sprintf("{user_id: %d, auth_key_id: %d, session_id: %d, state: %d, conn_state: %d, conn_id_list: %#v}",
		c.authSessions.AuthUserId,
		c.authSessions.authKeyId,
		c.sessionId,
		c.sessionState,
		c.connState,
		c.gatewayIdList)
}

func (c *session) addGatewayId(gateId string) {
	for _, id := range c.gatewayIdList {
		if id.Equal(gateId) {
			return
		}
	}
	c.gatewayIdList = append(c.gatewayIdList, serverIdCtx{gatewayId: gateId, lastReceiveTime: time.Now().Unix()})
}

func (c *session) getGatewayId() string {
	if len(c.gatewayIdList) > 0 {
		return c.gatewayIdList[len(c.gatewayIdList)-1].gatewayId
	} else {
		return ""
	}
}

func (c *session) checkGatewayIdExist(gateId string) bool {
	for _, id := range c.gatewayIdList {
		if id.Equal(gateId) {
			return true
		}
	}
	return false
}

func (c *session) changeConnState(state int) {
	c.connState = state
	if c.isAndroidPush || c.isGeneric {
		if state == kStateOnline {
			c.cb.setOnline()
		} else if state == kStateOffline {
			c.cb.trySetOffline()
		}
	}
}

func (c *session) onSessionConnNew(id string) {
	if c.connState != kStateOnline {
		c.changeConnState(kStateOnline)
		c.addGatewayId(id)
	}
}

func (c *session) onSessionMessageData(gatewayId, clientIp string, salt int64, msg *mtproto.TLMessage2) {
	if !c.checkBadServerSalt(gatewayId, salt, msg) {
		return
	}

	willCloseDate := time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout
	if willCloseDate > c.closeDate {
		c.closeDate = willCloseDate
	}

	if !c.checkBadMsgNotification(gatewayId, false, msg) {
		return
	}

	var msgs []*mtproto.TLMessage2
	if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
		for _, m2 := range msgContainer.Messages {
			msgs = append(msgs, &mtproto.TLMessage2{
				MsgId:  m2.MsgId,
				Seqno:  m2.Seqno,
				Bytes:  m2.Bytes,
				Object: m2.Object,
			})
		}
		c.inQueue.AddMsgId(msg.MsgId)
	} else {
		msgs = append(msgs, msg)
	}

	for i := 0; i < len(msgs); i++ {
		if packed, ok := msgs[i].Object.(*mtproto.TLGzipPacked); ok {
			msgs[i] = &mtproto.TLMessage2{
				MsgId:  msgs[i].MsgId,
				Seqno:  msgs[i].Seqno,
				Bytes:  int32(len(packed.PackedData)),
				Object: packed.Obj,
			}
		}
	}

	minMsgId := msg.MsgId
	for _, m2 := range msgs {
		if minMsgId < m2.MsgId {
			minMsgId = m2.MsgId
		}
	}

	if c.sessionState == kSessionStateNew || minMsgId < c.firstMsgId {
		log.Debugf("onNewSessionCreated - %#v,c: %s", msgs, c)
		c.onNewSessionCreated(gatewayId, minMsgId)
		if c.firstMsgId != 0 {
			c.firstMsgId = minMsgId
		}
		c.sessionState = kSessionStateCreated
	}

	defer func() {
		c.sendQueueToGateway(gatewayId)
		c.inQueue.Shrink()
	}()

	for _, m2 := range msgs {
		if !c.checkBadMsgNotification(gatewayId, true, m2) {
			continue
		}

		if m2.Object == nil {
			log.Errorf("obj is nil: %v", m2)
			continue
		}

		switch m2.Object.(type) {
		case *mtproto.TLMsgsAck:
			c.onMsgsAck(m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLMsgsAck))

		case *mtproto.TLHttpWait:
			c.onHttpWait(m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLHttpWait))

		default:
			inMsg := c.inQueue.AddMsgId(m2.MsgId)
			if inMsg.state == NONE {
				c.processMsg(clientIp, inMsg, m2.Object)
			} else {
				continue
			}
		}
	}
}

func (c *session) processMsg(clientIp string, inMsg *inboxMsg, r mtproto.TLObject) {
	switch r := r.(type) {
	case *mtproto.TLDestroyAuthKey:
		c.onDestroyAuthKey(inMsg, r)
	case *mtproto.TLRpcDropAnswer:
		c.onRpcDropAnswer(inMsg, r)
	case *mtproto.TLGetFutureSalts:
		c.onGetFutureSalts(inMsg, r)
	case *mtproto.TLPing:
		c.onPing(inMsg, r)
	case *mtproto.TLPingDelayDisconnect:
		c.onPingDelayDisconnect(inMsg, r)
	case *mtproto.TLDestroySession:
		c.onDestroySession(inMsg, r)
	case *mtproto.TLMsgsStateReq:
		c.onMsgsStateReq(inMsg, r)
	case *mtproto.TLMsgsStateInfo:
		c.onMsgsStateInfo(inMsg, r)
	case *mtproto.TLMsgsAllInfo:
		c.onMsgsAllInfo(inMsg, r)
	case *mtproto.TLMsgResendReq:
		c.onMsgResendReq(inMsg, r)
	case *mtproto.TLMsgDetailedInfo:
		c.onMsgDetailInfo(inMsg, r)
	case *mtproto.TLMsgNewDetailedInfo:
		c.onMsgNewDetailInfo(inMsg, r)
	case *mtproto.TLContestSaveDeveloperInfo:
		c.onContestSaveDeveloperInfo(inMsg, r)
	case *mtproto.TLInvokeWithLayer:
		c.onInvokeWithLayer(clientIp, inMsg, r)
	case *mtproto.TLInvokeAfterMsg:
		c.onInvokeAfterMsg(clientIp, inMsg, r)
	case *mtproto.TLInvokeAfterMsgs:
		c.onInvokeAfterMsgs(clientIp, inMsg, r)
	case *mtproto.TLInvokeWithoutUpdates:
		c.onInvokeWithoutUpdates(clientIp, inMsg, r)
	case *mtproto.TLInvokeWithMessagesRange:
		c.onInvokeWithMessagesRange(clientIp, inMsg, r)
	case *mtproto.TLInvokeWithTakeout:
		c.onInvokeWithTakeout(clientIp, inMsg, r)
	case *mtproto.TLGzipPacked:
		c.onRpcRequest(clientIp, inMsg, r.Obj)
	default:
		c.onRpcRequest(clientIp, inMsg, r)
	}
}

func (c *session) onSessionConnClose(id string) {
	var (
		idx = -1
	)

	for i, cId := range c.gatewayIdList {
		if cId.Equal(id) {
			idx = i
			break
		}
	}

	if idx != -1 {
		c.gatewayIdList = append(c.gatewayIdList[:idx], c.gatewayIdList[idx+1:]...)
	}

	if len(c.gatewayIdList) == 0 {
		c.changeConnState(kStateOffline)
	}
}

func (c *session) sessionOnline() bool {
	return c.connState == kStateOnline
}

func (c *session) sessionClosed() bool {
	return c.connState == kStateClose
}

func (c *session) onTimer() bool {
	date := time.Now().Unix()
	gatewayId := c.getGatewayId()

	timeoutIdList := c.pendingQueue.OnTimer()
	for _, id := range timeoutIdList {
		c.sendRpcResult(&mtproto.TLRpcResult{
			ReqMsgId: id,
			Result: &mtproto.TLRpcError{Data2: &mtproto.RpcError{
				ErrorCode:    -503,
				ErrorMessage: "Timeout",
			}},
		})
	}

	httpTimeOutList := c.httpQueue.PopTimeoutList()
	if len(httpTimeOutList) > 0 {
		log.Debugf("timeoutList: %d", len(httpTimeOutList))
	}
	for _, ch := range httpTimeOutList {
		c.sendHttpDirectToGateway(ch, false, emptyMsgContainer, func(sentRaw *mtproto.TLMessageRawData) {
		})
	}

	if c.connState == kStateOnline {
		if date >= c.closeDate {
			c.changeConnState(kStateOffline)
		} else {
			c.sendQueueToGateway(gatewayId)
		}
	} else if c.connState == kStateOffline || c.connState == kStateNew {
		if date >= c.closeDate+kCacheSessionTimeout {
			c.changeConnState(kStateClose)
		}
	}
	return true
}

func (c *session) encodeMessage(messageId int64, confirm bool, tl mtproto.TLObject) ([]byte, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	seqNo := c.generateMessageSeqNo(confirm)

	if messageId == 0 {
		messageId = nextMessageId(false)
	}

	return SerializeToBuffer(salt, c.sessionId, &mtproto.TLMessage2{
		MsgId:  messageId,
		Seqno:  seqNo,
		Object: tl,
	}, c.cb.getLayer()), nil
}

func (c *session) generateMessageSeqNo(increment bool) int32 {
	value := c.nextSeqNo
	if increment {
		c.nextSeqNo++
		return int32(value*2 + 1)
	} else {
		return int32(value * 2)
	}
}

func (c *session) sendRpcResultToQueue(reqMsgId int64, result mtproto.TLObject) {
	rpcResult := &mtproto.TLRpcResult{
		ReqMsgId: reqMsgId,
		Result:   result,
	}
	b := rpcResult.Encode(c.cb.getLayer())
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(true),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(b)),
		Body:  b,
	}
	c.outQueue.AddRpcResultMsg(reqMsgId, rawMsg)
}

func (c *session) sendPushRpcResultToQueue(reqMsgId int64, result []byte) {
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(true),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(result)),
		Body:  result,
	}
	c.outQueue.AddRpcResultMsg(reqMsgId, rawMsg)
}

func (c *session) sendPushToQueue(pushMsgId int64, pushMsg mtproto.TLObject) {
	rawBytes := pushMsg.Encode(c.cb.getLayer())
	if len(rawBytes) > 256 {
		gzipPacked := &mtproto.TLGzipPacked{
			PackedData: rawBytes,
		}
		rawBytes = gzipPacked.Encode(c.cb.getLayer())
	}

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(rawBytes)),
		Body:  rawBytes,
	}
	c.outQueue.AddPushUpdates(pushMsgId, rawMsg)
}

func (c *session) sendRawToQueue(msgId int64, confirm bool, rawMsg mtproto.TLObject) {
	b := rawMsg.Encode(c.cb.getLayer())
	rawMsg2 := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}
	c.outQueue.AddNotifyMsg(msgId, confirm, rawMsg2)
}

func (c *session) sendHttpDirectToGateway(ch chan interface{}, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	b := obj.Encode(c.cb.getLayer())

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}

	rB, err := c.SendHttpDataToGateway(
		context.Background(),
		ch,
		c.cb.getTempAuthKeyId(),
		salt,
		c.sessionId,
		rawMsg)

	if err != nil {
		log.Errorf("sendHttpDirectToGateway - %v", err)
	}

	cb(rawMsg)
	return rB, err
}

func (c *session) sendDirectToGateway(gatewayId string, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	b := obj.Encode(c.cb.getLayer())

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}

	var (
		rB  bool
		err error
	)

	if !c.isHttp {
		rB, err = c.SendDataToGateway(
			context.Background(),
			gatewayId,
			c.cb.getTempAuthKeyId(),
			salt,
			c.sessionId,
			rawMsg)
	} else {
		if ch := c.httpQueue.Pop(); ch != nil {
			rB, err = c.SendHttpDataToGateway(
				context.Background(),
				ch,
				c.cb.getTempAuthKeyId(),
				salt,
				c.sessionId,
				rawMsg)
		}
	}

	if err != nil {
		log.Errorf("sendToClient - %v", err)
	}

	cb(rawMsg)
	return rB, err
}

func (c *session) sendRawDirectToGateway(gatewayId string, raw *mtproto.TLMessageRawData) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()

	var (
		rB  bool
		err error
	)
	if !c.isHttp {
		rB, err = c.SendDataToGateway(
			context.Background(),
			gatewayId,
			c.cb.getTempAuthKeyId(),
			salt,
			c.sessionId,
			raw)
	} else {
		if ch := c.httpQueue.Pop(); ch != nil {
			rB, err = c.SendHttpDataToGateway(
				context.Background(),
				ch,
				c.cb.getTempAuthKeyId(),
				salt,
				c.sessionId,
				raw)
		}
	}

	if err != nil {
		log.Errorf("sendRawDirectToGateway - %v", err)
	}
	return rB, err
}

func (c *session) sendQueueToGateway(gatewayId string) {
	if gatewayId == "" {
		return
	}

	if c.outQueue.oMsgs.Len() == 0 {
		return
	}

	var (
		pendings = make([]*outboxMsg, 0)
		b        = false
		err      error
	)
	for e := c.outQueue.oMsgs.Front(); e != nil; e = e.Next() {
		if e.Value.(*outboxMsg).sent == 0 || time.Now().Unix() >= e.Value.(*outboxMsg).sent+waitMsgAcksTimeout {
			pendings = append(pendings, e.Value.(*outboxMsg))
		}
	}

	if len(pendings) == 1 {
		log.Debugf("sendRawDirectToGateway - pendings[0]")
		b, err = c.sendRawDirectToGateway(gatewayId, pendings[0].msg)

	} else if len(pendings) > 1 {
		msgContainer := &mtproto.TLMsgRawDataContainer{
			Messages: make([]*mtproto.TLMessageRawData, 0, len(pendings)),
		}
		for _, m := range pendings {
			msgContainer.Messages = append(msgContainer.Messages, m.msg)
		}

		log.Debugf("sendRawDirectToGateway - TLMsgRawDataContainer")
		b, err = c.sendDirectToGateway(gatewayId, false, msgContainer, func(sentRaw *mtproto.TLMessageRawData) {
		})
	}

	if err == nil && b {
		for _, m := range pendings {
			log.Debugf("need_no_ack: %d", m.msgId)
			if m.state == NEED_NO_ACK {
				c.outQueue.Remove(m.msgId)
			} else {
				m.sent = time.Now().Unix()
			}
		}
	}
}
