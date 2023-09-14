package dataobject

type WalletRecordDO struct {
	ID      int32   `db:"id"`
	UID     int32   `db:"uid"`
	Type    int8    `db:"type"`
	Amount  float64 `db:"amount"`
	Related int32   `db:"related"`
	Remarks string  `db:"remarks"`
	Deleted bool    `db:"deleted"`
	Date    int32   `db:"date"`
}
