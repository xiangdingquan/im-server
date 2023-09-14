package service

import (
	"container/list"
	"fmt"
	"math"
)

const (
	NONE                  byte = 0
	UNKNOWN               byte = 1
	NOT_RECEIVED          byte = 2
	NOT_RECEIVED_SURE     byte = 3
	RECEIVED              byte = 4
	ACKNOWLEDGED          byte = 8
	NEED_NO_ACK           byte = 16
	RPC_PROCESSING        byte = 32
	RESPONSE_GENERATED    byte = 64
	RESPONSE_ACKNOWLEDGED byte = 128
)

const (
	maxQueueSize = 400
)

type inboxMsg struct {
	msgId int64
	seqNo int32
	state byte
}

func (m *inboxMsg) ChangeState(s byte) {
	m.state = s
}

func (m *inboxMsg) DebugString() string {
	return fmt.Sprintf("{msg_id:%d, seqno: %d, state: %d}", m.msgId, m.seqNo, m.state)
}

func newInboxMsg(msgId int64) *inboxMsg {
	r := new(inboxMsg)
	r.msgId = msgId
	r.state = NONE
	return r
}

type sessionInboundQueue struct {
	firstMsgId int64
	minMsgId   int64
	maxMsgId   int64
	msgIds     *list.List
}

func newSessionInboundQueue() *sessionInboundQueue {
	q := new(sessionInboundQueue)
	q.msgIds = list.New()
	q.firstMsgId = 0
	q.minMsgId = 0
	q.maxMsgId = math.MaxInt64
	return q
}

func (q *sessionInboundQueue) AddMsgId(msgId int64) (r *inboxMsg) {
	if msgId < q.minMsgId {
		q.minMsgId = msgId
	}
	if msgId > q.maxMsgId {
		q.maxMsgId = msgId
	}

	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId > msgId {
			r = newInboxMsg(msgId)
			q.msgIds.InsertBefore(r, e)
			return
		} else if e.Value.(*inboxMsg).msgId == msgId {
			r = e.Value.(*inboxMsg)
			return
		}
	}
	r = newInboxMsg(msgId)
	q.msgIds.PushBack(r)

	return
}

func (q *sessionInboundQueue) GetMinMsgId() int64 {
	return q.minMsgId
}

func (q *sessionInboundQueue) GetMaxMsgId() int64 {
	return q.maxMsgId
}

func (q *sessionInboundQueue) ChangeAckReceived(msgId int64) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId == msgId {
			e.Value.(*inboxMsg).state = RECEIVED | RESPONSE_ACKNOWLEDGED
		}
	}
}

func (q *sessionInboundQueue) Lookup(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*inboxMsg).msgId {
			iMsg = e.Value.(*inboxMsg)
			return
		}
	}
	return
}

func (q *sessionInboundQueue) Shrink() {
	for q.msgIds.Len() > maxQueueSize {
		iMsg := q.msgIds.Remove(q.msgIds.Front())
		q.minMsgId = iMsg.(*inboxMsg).msgId
	}
}

func (q *sessionInboundQueue) FindLowerEntry(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Back(); e != nil; e = e.Prev() {
		if msgId >= e.Value.(*inboxMsg).msgId {
			return e.Value.(*inboxMsg)
		}
	}
	return nil
}

func (q *sessionInboundQueue) FindHigherEntry(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId >= msgId {
			return e.Value.(*inboxMsg)
		}
	}
	return nil
}

func (q *sessionInboundQueue) Length() int {
	return q.msgIds.Len()
}
