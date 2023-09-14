package dataobject

type ImportedContactsDO struct {
	Id             int64 `db:"id"`
	UserId         int32 `db:"user_id"`
	ImportedUserId int32 `db:"imported_user_id"`
	Deleted        int8  `db:"deleted"`
}
