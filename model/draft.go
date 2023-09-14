package model

import (
	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
)

func MakeDraftMessageEmpty(date int32) (draft *mtproto.DraftMessage) {
	draft = mtproto.MakeTLDraftMessageEmpty(&mtproto.DraftMessage{}).To_DraftMessage()
	if date != 0 {
		draft.Date_FLAGINT32 = &types.Int32Value{Value: date}
	}
	return
}
