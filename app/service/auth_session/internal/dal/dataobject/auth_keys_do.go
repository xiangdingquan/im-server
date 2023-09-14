package dataobject

type AuthKeysDO struct {
	Id        int32  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	AuthKey   string `db:"auth_key"`
	Body      string `db:"body"`
	Deleted   int8   `db:"deleted"`
}
