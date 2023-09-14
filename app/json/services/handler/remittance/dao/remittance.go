package dao

import (
	"context"
	"time"

	"open.chat/app/json/db/dbo"
	"open.chat/model"
	"open.chat/pkg/database/sqlx"
)

// CreateRemittance 创建转账
func (d *Dao) CreateRemittance(ctx context.Context, chatID, payerUID, payeeUID uint32, remittanceType, status uint8, description string, amount float64) (uint32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		rDO := &dbo.RemittanceDO{
			ChatID:      chatID,
			PayerUID:    payerUID,
			PayeeUID:    payeeUID,
			Amount:      amount,
			Status:      status,
			Type:        remittanceType,
			Description: description,
			CreateDate:  uint32(time.Now().Unix()),
		}

		rID, _, err := d.RemittanceDao.InsertTx(tx, rDO)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = uint32(rID)

		_, err = d.Wallet.DecBalanceTx(tx, payerUID, model.WalletRecordType_RemitRemittance, amount, uint32(rID), "转账-付款")
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return tR.Data.(uint32), tR.Err
}

func (d *Dao) ReceiveRemittance(ctx context.Context, remittanceID uint32, payeeUID uint32, amount float64) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//收款
		err := d.UpdateRemittanceStatusTx(tx, remittanceID, 1)
		if err != nil {
			result.Err = err
			return
		}

		//增加余额和日志
		_, err = d.Wallet.IncBalanceTx(tx, payeeUID, model.WalletRecordType_ReceiveRemittance, amount, remittanceID, "转账-收款")
		if err != nil {
			result.Err = err
			return
		}
	})

	return tR.Err
}

func (d *Dao) RefundRemittance(ctx context.Context, remittanceID uint32, payerUID uint32, amount float64) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//退款
		err := d.UpdateRemittanceStatusTx(tx, remittanceID, 2)
		if err != nil {
			result.Err = err
			return
		}

		//增加余额和日志
		_, err = d.Wallet.IncBalanceTx(tx, payerUID, model.WalletRecordType_RefundRemittanceByUser, amount, remittanceID, "转账-用户退款")
		if err != nil {
			result.Err = err
			return
		}
	})

	return tR.Err

}
