package dataobject

type UserPresencesDO struct {
	Id         int32 `db:"id"`
	UserId     int32 `db:"user_id"`
	LastSeenAt int64 `db:"last_seen_at"`
}
