package core

import (
	"context"

	"open.chat/app/json/services/handler/redpacket/dao"
	model2 "open.chat/model"
	"open.chat/pkg/database/sqlx"
)

// RedPacketCore .
type RedPacketCore struct {
	*dao.Dao
}

// New .
func New(d *dao.Dao) *RedPacketCore {
	if d == nil {
		d = dao.New()
	}
	return &RedPacketCore{d}
}

type GiveRedPacket struct {
	Rid   uint32  `json:"redpacketId"` //红包id
	Uid   int32   `json:"userId"`      //用户
	Price float64 `json:"price"`       //金额
}

func (m *RedPacketCore) GetRedPacketTimeoutList(ctx context.Context, date int32) (list []GiveRedPacket) {
	list = make([]GiveRedPacket, 0)
	redPackets, _ := m.RedPacketDAO.SelectTimeoutList(ctx, date)
	for _, r := range redPackets {
		list = append(list, GiveRedPacket{
			Rid:   r.ID,
			Uid:   int32(r.OwnerUID),
			Price: r.RemainPrice,
		})
	}
	return list
}

// GetRedPacket 领取红包.
func (m *RedPacketCore) GiveBackRedPacket(ctx context.Context, redpacketID, userID uint32, Price float64) error {
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//设置红包
		_, err := m.RedPacketDAO.SetCompletedTx(tx, redpacketID)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = true
		//增加返还余额和日志
		_, err = m.Wallet.IncBalanceTx(tx, userID, model2.WalletRecordType_GivebackRedPacket, Price, redpacketID, "退回红包")
		if err != nil {
			result.Err = err
			return
		}
	})

	return tR.Err
}
