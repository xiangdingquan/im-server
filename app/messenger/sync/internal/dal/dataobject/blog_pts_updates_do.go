package dataobject

type BlogPtsUpdatesDO struct {
	Id         int64  `db:"id"`
	UserId     int32  `db:"user_id"`
	Pts        int32  `db:"pts"`
	PtsCount   int32  `db:"pts_count"`
	UpdateType int8   `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date       int32  `db:"date"`
}
