package dataobject

type ConversationsDO struct {
	Id                int32  `db:"id"`
	UserId            int32  `db:"user_id"`
	PeerId            int32  `db:"peer_id"`
	IsPinned          int8   `db:"is_pinned"`
	OrderPinned       int64  `db:"order_pinned"`
	TopMessage        int32  `db:"top_message"`
	PinnedMsgId       int32  `db:"pinned_msg_id"`
	ReadInboxMaxId    int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId   int32  `db:"read_outbox_max_id"`
	UnreadCount       int32  `db:"unread_count"`
	UnreadMark        int8   `db:"unread_mark"`
	DraftType         int8   `db:"draft_type"`
	DraftMessageData  string `db:"draft_message_data"`
	FolderId          int32  `db:"folder_id"`
	FolderPinned      int32  `db:"folder_pinned"`
	FolderOrderPinned int64  `db:"folder_order_pinned"`
	Date2             int32  `db:"date2"`
	Deleted           int8   `db:"deleted"`
}
