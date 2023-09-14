package dataobject

type BlogUserPrivaciesDO struct {
	Id      int32  `db:"id"`
	UserId  int32  `db:"user_id"`
	KeyType int8   `db:"key_type"`
	Rules   string `db:"rules"`
}
