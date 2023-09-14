package dataobject

type ChatParticipantsDO struct {
	Id                  int32  `db:"id"`
	ChatId              int32  `db:"chat_id"`
	UserId              int32  `db:"user_id"`
	ParticipantType     int8   `db:"participant_type"`
	IsPinned            int8   `db:"is_pinned"`
	OrderPinned         int64  `db:"order_pinned"`
	TopMessage          int32  `db:"top_message"`
	PinnedMsgId         int32  `db:"pinned_msg_id"`
	ReadInboxMaxId      int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId     int32  `db:"read_outbox_max_id"`
	UnreadCount         int32  `db:"unread_count"`
	UnreadMentionsCount int32  `db:"unread_mentions_count"`
	UnreadMark          int8   `db:"unread_mark"`
	DraftType           int8   `db:"draft_type"`
	DraftMessageData    string `db:"draft_message_data"`
	FolderId            int32  `db:"folder_id"`
	FolderPinned        int32  `db:"folder_pinned"`
	FolderOrderPinned   int64  `db:"folder_order_pinned"`
	InviterUserId       int32  `db:"inviter_user_id"`
	InvitedAt           int32  `db:"invited_at"`
	KickedAt            int32  `db:"kicked_at"`
	LeftAt              int32  `db:"left_at"`
	HasScheduled        int32  `db:"has_scheduled"`
	State               int8   `db:"state"`
	Date2               int32  `db:"date2"`
}
