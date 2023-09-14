package dataobject

type AppConfigsDO struct {
	Id      int32  `db:"id"`
	Key2    string `db:"key2"`
	Type2   string `db:"type2"`
	Value2  string `db:"value2"`
	Deleted int8   `db:"deleted"`
}
