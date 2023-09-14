package dataobject

type ScheduledMessagesDO struct {
	Id               int64  `db:"id"`
	UserId           int32  `db:"user_id"`
	UserMessageBoxId int32  `db:"user_message_box_id"`
	PeerType         int8   `db:"peer_type"`
	PeerId           int32  `db:"peer_id"`
	DialogId         int64  `db:"dialog_id"`
	RandomId         int64  `db:"random_id"`
	MessageType      int8   `db:"message_type"`
	MessageDataType  int8   `db:"message_data_type"`
	MessageData      string `db:"message_data"`
	MessageBoxType   int8   `db:"message_box_type"`
	ScheduledDate    int32  `db:"scheduled_date"`
	Date2            int32  `db:"date2"`
	State            int8   `db:"state"`
}
