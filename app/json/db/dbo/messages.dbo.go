package dbo

type BatchSendMessageDo struct {
	ID      uint32 `db:"id"`
	UID     uint32 `db:"uid"`
	Message string `db:"message"`
	ToUsers string `db:"to_users"`
}

type MessagesDo struct {
	UserMessageBoxId int32 `db:"user_message_box_id"`
	DialogId         int64 `db:"dialog_id"`
	DialogMessageId  int32 `db:"dialog_message_id"`
}
