package dataobject

type BlogFollowsDO struct {
	Id           int32 `db:"id"`
	UserId       int32 `db:"user_id"`
	TargetUserId int32 `db:"target_uid"`
	Date         int32 `db:"date"`
	Deleted      bool  `db:"deleted"`
}
