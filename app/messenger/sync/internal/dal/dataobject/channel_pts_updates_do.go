package dataobject

type ChannelPtsUpdatesDO struct {
	Id           int64  `db:"id"`
	ChannelId    int32  `db:"channel_id"`
	Pts          int32  `db:"pts"`
	PtsCount     int32  `db:"pts_count"`
	UpdateType   int8   `db:"update_type"`
	NewMessageId int32  `db:"new_message_id"`
	UpdateData   string `db:"update_data"`
	Date2        int32  `db:"date2"`
}
