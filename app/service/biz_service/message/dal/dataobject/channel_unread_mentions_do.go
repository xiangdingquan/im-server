package dataobject

type ChannelUnreadMentionsDO struct {
	Id                 int64 `db:"id"`
	UserId             int32 `db:"user_id"`
	ChannelId          int32 `db:"channel_id"`
	MentionedMessageId int32 `db:"mentioned_message_id"`
	Deleted            int8  `db:"deleted"`
}
