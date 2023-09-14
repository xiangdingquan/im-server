package dataobject

type ChannelMessagesDeleteDO struct {
	Id        int64 `db:"id"`
	UserId    int32 `db:"user_id"`
	ChannelId int32 `db:"channel_id"`
	MessageId int32 `db:"message_id"`
	Deleted   int8  `db:"deleted"`
}
