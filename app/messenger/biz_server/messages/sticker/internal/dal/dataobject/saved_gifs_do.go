package dataobject

type SavedGifsDO struct {
	Id     int64 `db:"id"`
	UserId int32 `db:"user_id"`
	GifId  int64 `db:"gif_id"`
}
