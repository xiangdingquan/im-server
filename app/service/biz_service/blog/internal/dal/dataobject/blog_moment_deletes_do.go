package dataobject

type BlogMomentDeletesDO struct {
	Id      int32 `db:"id"`
	UserId  int32 `db:"user_id"`
	BlogId  int32 `db:"blog_id"`
	Date    int32 `db:"date"`
	Deleted bool  `db:"deleted"`
}
