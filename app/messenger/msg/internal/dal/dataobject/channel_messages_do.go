package dataobject

type ChannelMessagesDO struct {
	Id               int64  `db:"id"`
	ChannelId        int32  `db:"channel_id"`
	ChannelMessageId int32  `db:"channel_message_id"`
	SenderUserId     int32  `db:"sender_user_id"`
	RandomId         int64  `db:"random_id"`
	Pts              int32  `db:"pts"`
	MessageDataId    int64  `db:"message_data_id"`
	MessageType      int8   `db:"message_type"`
	MessageData      string `db:"message_data"`
	Message          string `db:"message"`
	MediaType        int8   `db:"media_type"`
	MediaUnread      int8   `db:"media_unread"`
	HasMediaUnread   int8   `db:"has_media_unread"`
	EditMessage      string `db:"edit_message"`
	EditDate         int32  `db:"edit_date"`
	Views            int32  `db:"views"`
	FromScheduled    int32  `db:"from_scheduled"`
	TtlSeconds       int32  `db:"ttl_seconds"`
	HasRemove        int8   `db:"has_remove"`
	HasDM            int8   `db:"has_dm"`
	Date             int32  `db:"date"`
	Deleted          int8   `db:"deleted"`
}
