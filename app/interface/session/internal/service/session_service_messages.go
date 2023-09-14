package service

import (
	"context"
	"math/rand"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (c *session) checkContainer(msgId int64, seqno int32, container *mtproto.TLMsgContainer) int32 {
	if c.inQueue.Lookup(msgId) != nil {
		log.Errorf("checkContainer - msgId collision: {msg_id: %d, seqno: %d}", msgId, seqno)
		return kMsgIdCollision
	}

	if len(container.Messages) == 0 {
		return 0
	}

	for _, v := range container.Messages {
		if v.Seqno > seqno {
			log.Errorf("checkContainer - v.seqno(%s) > seqno({msg_id: %d, seqno: %d})", v.DebugString(), msgId, seqno)
			return kInvalidContainer
		}
		if v.MsgId >= msgId {
			log.Errorf("checkContainer - v.MsgId(%s) > msgId({msg_id: %d, seqno: %d})", v.DebugString(), msgId, seqno)
			return kInvalidContainer
		}

		if _, ok := v.Object.(*mtproto.TLMsgContainer); ok {
			log.Errorf("checkContainer - is container: %v", v)
			return kInvalidContainer
		}
	}

	return 0
}

func (c *session) onNewSessionCreated(gatewayId string, msgId int64) {
	log.Debugf("onNewSessionCreated - request data: %d", msgId)
	newSessionCreated := mtproto.MakeTLNewSessionCreated(&mtproto.NewSession{
		FirstMsgId: msgId,
		UniqueId:   rand.Int63(),
		ServerSalt: c.cb.getCacheSalt().GetSalt(),
	})

	log.Debugf("onNewSessionCreated - reply: {%v}", newSessionCreated)

	c.sendDirectToGateway(gatewayId, true, newSessionCreated, func(sentRaw *mtproto.TLMessageRawData) {
		id2 := c.authSessions.getNextNotifyId()
		sentMsg := c.outQueue.AddNotifyMsg(id2, true, sentRaw)
		sentMsg.sent = 0
	})
}

func (c *session) onDestroyAuthKey(msgId *inboxMsg, destroyAuthKey *mtproto.TLDestroyAuthKey) {
	log.Debugf("onDestroyAuthKey - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), destroyAuthKey.DebugString())
	res := mtproto.MakeTLDestroyAuthKeyOk(nil).To_DestroyAuthKeyRes()
	c.sendRpcResultToQueue(msgId.msgId, res)
	msgId.state = RECEIVED | ACKNOWLEDGED
}

func (c *session) onPing(msgId *inboxMsg, ping *mtproto.TLPing) {
	log.Debugf("onPing - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), ping.DebugString())
	pong := &mtproto.TLPong{Data2: &mtproto.Pong{
		MsgId:  msgId.msgId,
		PingId: ping.PingId,
	}}

	c.sendRawToQueue(msgId.msgId, false, pong)
	msgId.state = RECEIVED | NEED_NO_ACK

	c.closeDate = time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout
}

func (c *session) onPingDelayDisconnect(msgId *inboxMsg, pingDelayDisconnect *mtproto.TLPingDelayDisconnect) {
	log.Debugf("onPingDelayDisconnect - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), pingDelayDisconnect.DebugString())

	pong := &mtproto.TLPong{Data2: &mtproto.Pong{
		MsgId:  msgId.msgId,
		PingId: pingDelayDisconnect.PingId,
	}}

	c.sendRawToQueue(msgId.msgId, false, pong)
	msgId.state = RECEIVED | NEED_NO_ACK

	willCloseDate := time.Now().Unix() + int64(pingDelayDisconnect.DisconnectDelay) + kPingAddTimeout
	if willCloseDate > c.closeDate {
		c.closeDate = willCloseDate
	}
}

func (c *session) onHttpWait(msgId int64, seqNo int32, request *mtproto.TLHttpWait) {
	log.Debugf("onHttpWait - request data: {sess: %s, msg_id: %d, seq_no: %d, req: {%s}}",
		c, msgId, seqNo, request.DebugString())

	_ = request.GetMaxDelay()
	_ = request.GetWaitAfter()
	t := request.GetMaxWait() / 1000
	if t == 0 {
		t = 1
	}
}

func (c *session) onDestroySession(msgId *inboxMsg, request *mtproto.TLDestroySession) {
	log.Debugf("onDestroySession - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
	if request.SessionId == c.sessionId {
		log.Error("the result of this being applied to the current session is undefined.")
		return
	}

	if c.cb.destroySession(request.GetSessionId()) {
		destroySessionOk := mtproto.MakeTLDestroySessionOk(&mtproto.DestroySessionRes{
			SessionId: request.SessionId,
		}).To_DestroySessionRes()
		c.sendRawToQueue(msgId.msgId, false, destroySessionOk)
	} else {
		destroySessionNone := mtproto.MakeTLDestroySessionNone(&mtproto.DestroySessionRes{
			SessionId: request.SessionId,
		}).To_DestroySessionRes()
		c.sendRawToQueue(msgId.msgId, false, destroySessionNone)
	}

	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onGetFutureSalts(msgId *inboxMsg, request *mtproto.TLGetFutureSalts) {
	log.Debugf("onGetFutureSalts - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())

	salts, err := c.GetFutureSalts(context.Background(), c.cb.getAuthKeyId(), request.Num)
	if err != nil {
		log.Errorf("getFutureSalts error: %v", err)
		return
	}

	futureSalts := mtproto.MakeTLFutureSalts(&mtproto.FutureSalts{
		ReqMsgId: msgId.msgId,
		Now:      int32(time.Now().Unix()),
		Salts:    salts,
	}).To_FutureSalts()

	log.Debugf("onGetFutureSalts - reply data: %s", futureSalts.DebugString())
	c.sendRawToQueue(msgId.msgId, false, futureSalts)
	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onRpcDropAnswer(msgId *inboxMsg, request *mtproto.TLRpcDropAnswer) {
	log.Debugf("onRpcDropAnswer - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())

	rpcAnswer := &mtproto.RpcDropAnswer{}
	var found = false
	if !found {
		rpcAnswer.Constructor = mtproto.CRC32_rpc_answer_unknown
	}

	rpcAnswer = &mtproto.RpcDropAnswer{
		PredicateName: mtproto.Predicate_rpc_answer_unknown,
	}

	c.sendRpcResultToQueue(msgId.msgId, rpcAnswer)
	msgId.state = RECEIVED | ACKNOWLEDGED
}

func (c *session) onContestSaveDeveloperInfo(msgId *inboxMsg, request *mtproto.TLContestSaveDeveloperInfo) {
	log.Warnf("onContestSaveDeveloperInfo - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), request.DebugString())
}
