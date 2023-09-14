package dataobject

type WalletDO struct {
	ID       int32   `db:"id"`
	UID      int32   `db:"uid"`
	Address  string  `db:"address"`
	Password string  `db:"password"`
	Balance  float64 `db:"balance"`
	Data     int32   `db:"date"`
	Deleted  bool    `db:"deleted"`
}
