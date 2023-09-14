package dataobject

type WallPapersDO struct {
	Id        int32  `db:"id"`
	Type      int8   `db:"type"`
	Title     string `db:"title"`
	Color     int32  `db:"color"`
	BgColor   int32  `db:"bg_color"`
	PhotoId   int64  `db:"photo_id"`
	DeletedAt int64  `db:"deleted_at"`
}
