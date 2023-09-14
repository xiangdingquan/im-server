package model

import (
	"open.chat/mtproto"
)

// //////////////////////////////////////////////////////////////////////
type DialogFilterExt struct {
	Id           int32
	DialogFilter *mtproto.DialogFilter
	Order        int64
}

type DialogFilterExtList []*DialogFilterExt
