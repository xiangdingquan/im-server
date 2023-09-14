package dbo

type (
	// WalletDO .
	WalletDO struct {
		ID       uint32  `db:"id"`
		UID      uint32  `db:"uid"`      //用户id
		Address  string  `db:"address"`  //钱包地址
		Password string  `db:"password"` //钱包地址
		Balance  float64 `db:"balance"`  //钱包余额
		Data     int32   `db:"date"`
		Deleted  bool    `db:"deleted"`
	}

	// WalletRecordDO .
	WalletRecordDO struct {
		ID      uint32  `db:"id"`
		UID     uint32  `db:"uid"`     //用户id
		Type    int8    `db:"type"`    //类型 1.充值/提现 2.红包 3.收款/转账
		Amount  float64 `db:"amount"`  //
		Related uint32  `db:"related"` //关联的id 1.充值/提现id 2.红包id 3.收款/转账id
		Remarks string  `db:"remarks"` //备注
		Date    int32   `db:"date"`
		Deleted bool    `db:"deleted"`
	}
)
