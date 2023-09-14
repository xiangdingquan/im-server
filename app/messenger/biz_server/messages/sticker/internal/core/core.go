package core

import (
	"open.chat/app/messenger/biz_server/messages/sticker/internal/dao"
)

type StickerCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *StickerCore {
	return &StickerCore{dao}
}
