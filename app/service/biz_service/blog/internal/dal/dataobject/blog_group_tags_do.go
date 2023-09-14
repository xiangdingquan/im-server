package dataobject

type BlogGroupTagsDO struct {
	Id            int32  `db:"id"`
	UserId        int32  `db:"user_id"`
	Title         string `db:"title"`
	MemberUserIds string `db:"member_uids"`
	Date          int32  `db:"date"`
	Deleted       bool   `db:"deleted"`
}
