package dataobject

type PopularContactsDO struct {
	Id        int64  `db:"id"`
	Phone     string `db:"phone"`
	Importers int32  `db:"importers"`
	Deleted   int8   `db:"deleted"`
	UpdateAt  string `db:"update_at"`
}
