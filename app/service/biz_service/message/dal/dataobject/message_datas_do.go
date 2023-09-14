package dataobject

type MessageDatasDO struct {
	Id              int64  `db:"id"`
	MessageDataId   int64  `db:"message_data_id"`
	DialogId        int64  `db:"dialog_id"`
	DialogMessageId int32  `db:"dialog_message_id"`
	SenderUserId    int32  `db:"sender_user_id"`
	PeerType        int8   `db:"peer_type"`
	PeerId          int32  `db:"peer_id"`
	RandomId        int64  `db:"random_id"`
	MessageType     int8   `db:"message_type"`
	MessageData     string `db:"message_data"`
	MediaUnread     int8   `db:"media_unread"`
	HasMediaUnread  int8   `db:"has_media_unread"`
	Date            int32  `db:"date"`
	EditMessage     string `db:"edit_message"`
	EditDate        int32  `db:"edit_date"`
	Deleted         int8   `db:"deleted"`
}
