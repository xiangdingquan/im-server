package dataobject

type UserSettingsDO struct {
	Id      int32  `db:"id"`
	UserId  int32  `db:"user_id"`
	Key2    string `db:"key2"`
	Value   string `db:"value"`
	Deleted int8   `db:"deleted"`
}
