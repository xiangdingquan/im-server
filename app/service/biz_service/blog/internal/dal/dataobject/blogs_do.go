package dataobject

type BlogsDO struct {
	Id           int32 `db:"id"`
	UserId       int32 `db:"user_id"`
	UpdateMaxId  int32 `db:"read_update_id"`
	BlogMaxId    int32 `db:"read_blog_id"`
	CommentMaxId int32 `db:"read_comment_id"`
	Moments      int32 `db:"moments"`
	Follows      int32 `db:"follows"`
	Fans         int32 `db:"fans"`
	Comments     int32 `db:"comments"`
	Likes        int32 `db:"likes"`
	Data         int32 `db:"date"`
	Deleted      bool  `db:"deleted"`
}
