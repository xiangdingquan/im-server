package dataobject

type BlogMomentsDO struct {
	Id             int32   `db:"id"`
	UserId         int32   `db:"user_id"`
	BlogId         int32   `db:"blog_id"`
	Text           string  `db:"text"`
	Entities       string  `db:"entities"`
	Video          string  `db:"video"`
	Photos         string  `db:"photos"`
	MentionUserIds string  `db:"mention_uids"`
	ShareType      int8    `db:"share_type"`
	MemberUserIds  string  `db:"member_uids"`
	HasGeo         bool    `db:"has_geo"`
	Lat            float64 `db:"lat"`
	Long           float64 `db:"long"`
	Address        string  `db:"address"`
	Likes          int32   `db:"likes"`
	Commits        int32   `db:"commits"`
	Date           int32   `db:"date"`
	Topics         string  `db:"topics"`
	Sort           int32   `db:"sort"`
	Deleted        bool    `db:"deleted"`
}
