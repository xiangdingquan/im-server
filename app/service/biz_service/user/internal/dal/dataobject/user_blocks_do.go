package dataobject

type UserBlocksDO struct {
	Id      int32 `db:"id"`
	UserId  int32 `db:"user_id"`
	BlockId int32 `db:"block_id"`
	Date    int32 `db:"date"`
	Deleted int8  `db:"deleted"`
}
