package dataobject

type BlogBannedUsersDo struct {
	UserId  int32 `db:"user_id"`
	BanFrom int32 `db:"ban_from"`
	BanTo   int32 `db:"ban_to"`
}
