package service

import (
	"time"

	"open.chat/pkg/sync2"
)

var msgIdSeq = sync2.NewAtomicInt64(0)

func nextMessageId(isRpc bool) int64 {
	unixNano := time.Now().UnixNano()
	ts := unixNano / 1e9
	ms := (unixNano % 1e9) / 1e6
	sid := msgIdSeq.Add(1) & 0x1ffff
	msgIdSeq.CompareAndSwap(0x1ffff, 0)
	last := 1
	if !isRpc {
		last = 3
	}
	msgId := int64(ts<<32) | int64(ms)<<21 | int64(sid)<<3 | int64(last)
	return msgId
}
