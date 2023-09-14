package dataobject

type ChannelParticipantsDO struct {
	Id                        int64  `db:"id"`
	ChannelId                 int32  `db:"channel_id"`
	UserId                    int32  `db:"user_id"`
	IsCreator                 int8   `db:"is_creator"`
	IsPinned                  int8   `db:"is_pinned"`
	OrderPinned               int64  `db:"order_pinned"`
	ReadInboxMaxId            int32  `db:"read_inbox_max_id"`
	ReadOutboxMaxId           int32  `db:"read_outbox_max_id"`
	UnreadCount               int32  `db:"unread_count"`
	UnreadMentionsCount       int32  `db:"unread_mentions_count"`
	UnreadMark                int8   `db:"unread_mark"`
	DraftType                 int8   `db:"draft_type"`
	DraftMessageData          string `db:"draft_message_data"`
	FolderId                  int32  `db:"folder_id"`
	FolderPinned              int32  `db:"folder_pinned"`
	FolderOrderPinned         int64  `db:"folder_order_pinned"`
	InviterUserId             int32  `db:"inviter_user_id"`
	PromotedBy                int32  `db:"promoted_by"`
	AdminRights               int32  `db:"admin_rights"`
	HiddenPrehistory          int8   `db:"hidden_prehistory"`
	HiddenPrehistoryMessageId int32  `db:"hidden_prehistory_message_id"`
	KickedBy                  int32  `db:"kicked_by"`
	BannedRights              int32  `db:"banned_rights"`
	BannedUntilDate           int32  `db:"banned_until_date"`
	MigratedFromMaxId         int32  `db:"migrated_from_max_id"`
	AvailableMinId            int32  `db:"available_min_id"`
	AvailableMinPts           int32  `db:"available_min_pts"`
	Rank                      string `db:"rank"`
	State                     int8   `db:"state"`
	Date2                     int32  `db:"date2"`
}
