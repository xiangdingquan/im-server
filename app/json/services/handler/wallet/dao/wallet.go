package dao

import (
	"context"
	"math/rand"
	"time"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/random2"
)

// CreateWallet .
func (d *Dao) CreateWallet(ctx context.Context, uid uint32, password string) (uint32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		wDO := &dbo.WalletDO{
			UID:      uid,
			Password: password,
		}
		for i := 0; i < 5; i++ {
			wDO.Address = random2.RandomAlphanumeric(uint(rand.Intn(10) + 55))
			walletID, _, err := d.WalletDAO.InsertTx(tx, wDO)
			if err != nil {
				if sqlx.IsDuplicate(err) {
					continue
				} else {
					result.Err = err
					return
				}
			}
			result.Data = (uint32)(walletID)
		}
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return tR.Data.(uint32), nil
}

// IncBalance 增加余额.
// _type 1.充值/提现 2.红包 3.收款/转账
// related 关联的id 1.充值/提现id 2.红包id 3.收款/转账id
// remarks 备注
func (d *Dao) IncBalance(ctx context.Context, uid uint32, type_ int8, money float64, related uint32, remarks string) (uint32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//增加余额
		_, err := d.WalletDAO.IncreaseBalanceTx(tx, uid, money)
		if err != nil {
			result.Err = err
			return
		}
		//增加日志
		wrDo := &dbo.WalletRecordDO{
			UID:     uid,
			Type:    type_,
			Amount:  money,
			Related: related,
			Remarks: remarks,
			Date:    int32(time.Now().Unix()),
		}
		rid, _, err := d.WalletRecordDAO.InsertTx(tx, wrDo)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = (uint32)(rid)
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return tR.Data.(uint32), tR.Err
}

// IncreaseBalance 增加余额.
// _type 1.充值/提现 2.红包 3.收款/转账
// related 关联的id 1.充值/提现id 2.红包id 3.收款/转账id
// remarks 备注
func (d *Dao) IncBalanceTx(tx *sqlx.Tx, uid uint32, type_ int8, money float64, related uint32, remarks string) (uint32, error) {
	//增加余额
	_, err := d.WalletDAO.IncreaseBalanceTx(tx, uid, money)
	if err != nil {
		return 0, err
	}
	//增加日志
	wrDo := &dbo.WalletRecordDO{
		UID:     uid,
		Type:    type_,
		Amount:  money,
		Related: related,
		Remarks: remarks,
		Date:    int32(time.Now().Unix()),
	}
	rid, _, err := d.WalletRecordDAO.InsertTx(tx, wrDo)
	if err != nil {
		return 0, err
	}

	return (uint32)(rid), err
}

// IncreaseBalance 减少余额.
// _type 1.充值/提现 2.红包 3.收款/转账
// related 关联的id 1.充值/提现id 2.红包id 3.收款/转账id
// remarks 备注
func (d *Dao) DecBalanceTx(tx *sqlx.Tx, uid uint32, type_ int8, money float64, related uint32, remarks string) (uint32, error) {
	return d.IncBalanceTx(tx, uid, type_, -money, related, remarks)
}
