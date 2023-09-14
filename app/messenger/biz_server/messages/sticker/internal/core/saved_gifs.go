package core

import (
	"context"

	"open.chat/app/messenger/biz_server/messages/sticker/internal/dal/dataobject"
)

func (m *StickerCore) SaveGif(ctx context.Context, userId int32, gifId int64) {
	m.SavedGifsDAO.InsertIgnore(ctx, &dataobject.SavedGifsDO{
		UserId: userId,
		GifId:  gifId,
	})
}

func (m *StickerCore) GetSavedGifs(ctx context.Context, userId int32) (idList []int64) {
	idList, _ = m.SavedGifsDAO.SelectAll(ctx, userId)
	return
}

func (m *StickerCore) DeleteSavedGif(ctx context.Context, userId int32, id int64) {
	m.SavedGifsDAO.Delete(ctx, userId, id)
}
