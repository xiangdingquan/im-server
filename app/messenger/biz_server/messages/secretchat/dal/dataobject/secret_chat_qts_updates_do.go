package dataobject

type SecretChatQtsUpdatesDO struct {
	Id            int64 `db:"id"`
	UserId        int32 `db:"user_id"`
	AuthKeyId     int64 `db:"auth_key_id"`
	ChatId        int32 `db:"chat_id"`
	Qts           int32 `db:"qts"`
	ChatMessageId int64 `db:"chat_message_id"`
	Date2         int32 `db:"date2"`
}
