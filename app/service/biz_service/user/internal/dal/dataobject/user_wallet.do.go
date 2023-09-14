package dataobject

type UserWalletDO struct {
	ID       int32   `db:"id"`
	UID      int32   `db:"uid"`      //用户id
	Address  string  `db:"address"`  //钱包地址
	Password string  `db:"password"` //钱包地址
	Balance  float64 `db:"balance"`  //钱包余额
	Data     int32   `db:"date"`
	Deleted  int8    `db:"deleted"`
}
