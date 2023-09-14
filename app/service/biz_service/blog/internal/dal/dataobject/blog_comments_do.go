package dataobject

type BlogCommentsDO struct {
	Id        int32  `db:"id"`
	UserId    int32  `db:"user_id"`
	Type      int8   `db:"type"`
	Text      string `db:"text"`
	BlogId    int32  `db:"blog_id"`
	CommentId int32  `db:"comment_id"`
	ReplyId   int32  `db:"reply_id"`
	Date      int32  `db:"date"`
	Deleted   bool   `db:"deleted"`
}
