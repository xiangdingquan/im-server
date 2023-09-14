package dataobject

type MessageBoxesDO struct {
	Id                int32 `db:"id"`
	UserId            int32 `db:"user_id"`
	UserMessageBoxId  int32 `db:"user_message_box_id"`
	DialogId          int64 `db:"dialog_id"`
	DialogMessageId   int32 `db:"dialog_message_id"`
	MessageDataId     int64 `db:"message_data_id"`
	Pts               int32 `db:"pts"`
	PtsCount          int32 `db:"pts_count"`
	MessageBoxType    int8  `db:"message_box_type"`
	ReplyToMsgId      int32 `db:"reply_to_msg_id"`
	Mentioned         int8  `db:"mentioned"`
	MediaUnread       int8  `db:"media_unread"`
	MessageFilterType int8  `db:"message_filter_type"`
	Date2             int32 `db:"date2"`
	Deleted           int8  `db:"deleted"`
}
