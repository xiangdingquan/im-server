package model

import (
	"encoding/json"

	"open.chat/mtproto"
)

func TLObjectToJson(object mtproto.TLObject) (b []byte) {
	b, _ = json.Marshal(object)
	return
}
