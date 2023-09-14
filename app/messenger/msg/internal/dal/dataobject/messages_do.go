package dataobject

type MessagesDO struct {
	Id               int32  `db:"id"`
	UserId           int32  `db:"user_id"`
	UserMessageBoxId int32  `db:"user_message_box_id"`
	DialogId         int64  `db:"dialog_id"`
	DialogMessageId  int32  `db:"dialog_message_id"`
	SenderUserId     int32  `db:"sender_user_id"`
	PeerType         int8   `db:"peer_type"`
	PeerId           int32  `db:"peer_id"`
	RandomId         int64  `db:"random_id"`
	MessageType      int8   `db:"message_type"`
	MessageData      string `db:"message_data"`
	MessageDataId    int64  `db:"message_data_id"`
	MessageDataType  int32  `db:"message_data_type"`
	Message          string `db:"message"`
	Pts              int32  `db:"pts"`
	PtsCount         int32  `db:"pts_count"`
	MessageBoxType   int8   `db:"message_box_type"`
	ReplyToMsgId     int32  `db:"reply_to_msg_id"`
	Mentioned        int8   `db:"mentioned"`
	MediaUnread      int8   `db:"media_unread"`
	HasMediaUnread   int8   `db:"has_media_unread"`
	FromScheduled    int32  `db:"from_scheduled"`
	TtlSeconds       int32  `db:"ttl_seconds"`
	Date2            int32  `db:"date2"`
	Deleted          int8   `db:"deleted"`
}
