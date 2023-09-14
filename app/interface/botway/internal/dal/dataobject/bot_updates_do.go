package dataobject

type BotUpdatesDO struct {
	Id         int64  `db:"id"`
	BotId      int32  `db:"bot_id"`
	UpdateId   int32  `db:"update_id"`
	UpdateType int8   `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date2      int64  `db:"date2"`
}
