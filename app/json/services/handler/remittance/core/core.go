package core

import (
	"context"

	"open.chat/app/json/services/handler/remittance/dao"
	"open.chat/model"
	"open.chat/pkg/database/sqlx"
)

const (
	RemittanceStatusInitial          = 0
	RemittanceStatusReceived         = 1
	RemittanceStatusRefundedByUser   = 2
	RemittanceStatusRefundedBySystem = 3

	RemittanceTypeSingle = 1
	RemittanceTypeGroup  = 2

	MidOfRemittance     = 5
	SidOfRemit          = 100
	SidOfReceive        = 200
	SidOfRefundByUser   = 300
	SidOfRemind         = 400
	SidOfRefundBySystem = 500

	SearchTypePayer = 1
	SearchTypePayee = 2
)

type RemittanceCore struct {
	*dao.Dao
}

func New(d *dao.Dao) *RemittanceCore {
	if d == nil {
		d = dao.New()
	}
	return &RemittanceCore{d}
}

type TransferInformation struct {
	Rid   uint32  `json:"remittanceID"` //转账标识
	Uid   int32   `json:"userId"`       //对方标识
	Price float64 `json:"price"`        //金额
}

// 由系统根据过期时间获取订单
func (m *RemittanceCore) GetRemittanceTimeoutList(ctx context.Context, date int32) (list []TransferInformation) {
	list = make([]TransferInformation, 0)
	remittances, _ := m.RemittanceDao.SelectTimeoutList(ctx, date)
	for _, r := range remittances {
		list = append(list, TransferInformation{
			Rid:   r.ID,
			Uid:   int32(r.PayerUID),
			Price: r.Amount,
		})
	}
	return list
}

// 由系统调用超时订单退回
func (d *RemittanceCore) RefundByTimeout(ctx context.Context, remittanceID uint32, payerUID uint32, amount float64) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//退款
		err := d.UpdateRemittanceStatusTx(tx, remittanceID, 2)
		if err != nil {
			result.Err = err
			return
		}

		//退回余额和日志
		_, err = d.Wallet.IncBalanceTx(tx, payerUID, model.WalletRecordType_RefundRemittanceBySystem, amount, remittanceID, "转账-系统退款")
		if err != nil {
			result.Err = err
			return
		}
	})

	return tR.Err

}
