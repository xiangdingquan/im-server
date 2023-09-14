package core

import (
	"open.chat/app/service/media/internal/dao"
	"open.chat/mtproto"
)

type PhotoCallback interface {
	GetPhotoSizeList(photoId int64) (sizes []*mtproto.PhotoSize)
}

type MediaCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *MediaCore {
	return &MediaCore{Dao: dao}
}
