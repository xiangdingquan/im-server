package dataobject

type BlogLikesDO struct {
	Id        int32 `db:"id"`
	UserId    int32 `db:"user_id"`
	Type      int8  `db:"type"`
	BlogId    int32 `db:"blog_id"`
	CommentId int32 `db:"comment_id"`
	Date      int32 `db:"date"`
	Deleted   bool  `db:"deleted"`
}
