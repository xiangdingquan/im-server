package fcm

import (
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func cryptoPushData(jsonString string, pushAuthKey []byte) ([]byte, error) {
	b := hack.Bytes(jsonString)
	x := mtproto.NewEncodeBuf(len(b) + 4)
	x.Int(int32(len(b)))
	x.Bytes(b)

	return []byte{}, nil
}
