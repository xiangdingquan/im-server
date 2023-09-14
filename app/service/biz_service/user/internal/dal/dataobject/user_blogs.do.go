package dataobject

type UserBlogsDO struct {
	Id        int32 `db:"id"`
	UserId    int32 `db:"user_id"`
	ReadMaxId int32 `db:"read_max_id"`
	Moments   int32 `db:"moments"`
	Follows   int32 `db:"follows"`
	Fans      int32 `db:"fans"`
	Likes     int32 `db:"likes"`
	Data      int32 `db:"date"`
	Deleted   bool  `db:"deleted"`
}
