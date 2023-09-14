package dataobject

type SecretChatMessagesDO struct {
	Id           int64  `db:"id"`
	SenderUserId int32  `db:"sender_user_id"`
	ChatId       int32  `db:"chat_id"`
	PeerId       int32  `db:"peer_id"`
	RandomId     int64  `db:"random_id"`
	MessageType  int8   `db:"message_type"`
	MessageData  string `db:"message_data"`
	Date2        int32  `db:"date2"`
}
