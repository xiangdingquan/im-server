package dataobject

type ChannelsDO struct {
	Id                  int32  `db:"id"`
	CreatorUserId       int32  `db:"creator_user_id"`
	AccessHash          int64  `db:"access_hash"`
	RandomId            int64  `db:"random_id"`
	TopMessage          int32  `db:"top_message"`
	PinnedMsgId         int32  `db:"pinned_msg_id"`
	Date2               int32  `db:"date2"`
	Pts                 int32  `db:"pts"`
	ParticipantsCount   int32  `db:"participants_count"`
	AdminsCount         int32  `db:"admins_count"`
	KickedCount         int32  `db:"kicked_count"`
	BannedCount         int32  `db:"banned_count"`
	Title               string `db:"title"`
	About               string `db:"about"`
	PhotoId             int64  `db:"photo_id"`
	Public              int8   `db:"public"`
	Username            string `db:"username"`
	Link                string `db:"link"`
	Broadcast           int8   `db:"broadcast"`
	Verified            int8   `db:"verified"`
	Megagroup           int8   `db:"megagroup"`
	Democracy           int8   `db:"democracy"`
	Signatures          int8   `db:"signatures"`
	AdminsEnabled       int8   `db:"admins_enabled"`
	DefaultBannedRights int32  `db:"default_banned_rights"`
	MigratedFromChatId  int32  `db:"migrated_from_chat_id"`
	PreHistoryHidden    int8   `db:"pre_history_hidden"`
	Deactivated         int8   `db:"deactivated"`
	Version             int32  `db:"version"`
	Date                int32  `db:"date"`
	Deleted             int8   `db:"deleted"`
}
