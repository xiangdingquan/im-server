package dbo

import "time"

type (
	RemittanceDO struct {
		ID          uint32    `db:"id"`
		ChatID      uint32    `db:"chat_id"`
		PayerUID    uint32    `db:"payer_uid"` //付款人
		PayeeUID    uint32    `db:"payee_uid"` //收款人
		Amount      float64   `db:"amount"`    //转账金额
		Status      uint8     `db:"status"`    //状态 0.未收款 1.已收款 2.已退款
		Type        uint8     `db:"type"`      //1.单聊红包 2.群聊红包
		Description string    `db:"description"`
		RemittedAt  uint32    `db:"-"`
		ReceivedAt  uint32    `db:"-"`
		RefundedAt  uint32    `db:"-"`
		CreateDate  uint32    `db:"create_date"` //创建时间
		CreateTime  time.Time `db:"created_at"`
		UpdateTime  time.Time `db:"updated_at"`
	}
)
