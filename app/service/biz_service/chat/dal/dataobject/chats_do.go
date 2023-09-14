package dataobject

type ChatsDO struct {
	Id                   int32  `db:"id"`
	CreatorUserId        int32  `db:"creator_user_id"`
	AccessHash           int64  `db:"access_hash"`
	RandomId             int64  `db:"random_id"`
	ParticipantCount     int32  `db:"participant_count"`
	Title                string `db:"title"`
	About                string `db:"about"`
	Notice               string `db:"notice"`
	Link                 string `db:"link"`
	Photo                string `db:"photo"`
	ChatPhoto            string `db:"chat_photo"`
	PhotoId              int64  `db:"photo_id"`
	AdminsEnabled        int8   `db:"admins_enabled"`
	DefaultBannedRights  int32  `db:"default_banned_rights"`
	MigratedToId         int32  `db:"migrated_to_id"`
	MigratedToAccessHash int64  `db:"migrated_to_access_hash"`
	Deactivated          int8   `db:"deactivated"`
	Version              int32  `db:"version"`
	Date                 int32  `db:"date"`
}
