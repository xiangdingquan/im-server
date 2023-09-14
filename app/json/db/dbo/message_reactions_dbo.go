package dbo

type MessageReactionDo struct {
	Type       int8  `db:"type"`
	ChatID     int64 `db:"chat_id"`
	MessageId  int32 `db:"message_id"`
	UserId     int32 `db:"user_id"`
	ReactionId int8  `db:"reaction_id"`
}
