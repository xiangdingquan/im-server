package dataobject

type ChannelAdminLogsDO struct {
	Id        int64  `db:"id"`
	UserId    int32  `db:"user_id"`
	ChannelId int32  `db:"channel_id"`
	Event     int32  `db:"event"`
	EventData string `db:"event_data"`
	Query     string `db:"query"`
	Date2     int32  `db:"date2"`
}
