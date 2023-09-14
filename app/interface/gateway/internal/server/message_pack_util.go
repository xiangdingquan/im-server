package server

import (
	"open.chat/mtproto"
)

func ParseFromIncomingMessage(b []byte) (msgId int64, obj mtproto.TLObject, err error) {
	dBuf := mtproto.NewDecodeBuf(b)

	msgId = dBuf.Long()
	oLen := dBuf.Int()
	_ = oLen
	obj = dBuf.Object()
	err = dBuf.GetError()

	return
}

func SerializeToBuffer(msgId int64, obj mtproto.TLObject) []byte {
	oBuf := obj.Encode(0)
	x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.Long(0)
	x.Long(msgId)
	x.Int(int32(len(oBuf)))
	x.Bytes(oBuf)

	return x.GetBuf()
}
