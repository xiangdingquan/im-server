package dao

import (
	"context"
	"time"

	"open.chat/app/json/db/dbo"
	"open.chat/model"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// TUserRecord .
type TUserRecord struct {
	UserID uint32  `json:"userId"` //用户id
	Price  float64 `json:"price"`  //获取到的金额
	GotAt  uint32  `json:"gotAt"`  //领取的时间
}

type TRedPacketAllRecord struct {
	RedPacketID uint32        `db:"red_packet_id" json:"redPacketId"`
	ChatID      uint32        `db:"chat_id" json:"chatId"`          //IM会话标识
	OwnerID     uint32        `db:"owner_uid" json:"from"`          //创建者
	Type        uint8         `db:"type" json:"type"`               //1.单聊红包 2.拼手气红包 3.普通红包
	Title       string        `db:"title" json:"title"`             //标题
	Price       float64       `db:"price" json:"price"`             //单个红包金额
	TotalPrice  float64       `db:"total_price" json:"total_price"` //红包总金额
	Count       uint32        `db:"total_count" json:"count"`       //红包数量
	CreateTime  time.Time     `db:"created_at" json:"-"`            //创建红包的时间
	IsExpire    bool          `db:"-" json:"isExpire"`              //是否过期
	CreateAt    uint32        `db:"-" json:"createAt"`              //创建时间
	Users       []TUserRecord `db:"-" json:"users"`                 //用户记录
}

type TRedPacketStatistics struct {
	TotalPrice    float64 `json:"total_price"`
	Count         uint32  `json:"count"`
	TopPriceCount uint32  `json:"top_price_count"`
}

// CreateRedPacket 创建红包.
func (d *Dao) CreateRedPacket(ctx context.Context, chatID, ownerUID uint32, honbaoType uint8, title string, price, total_price float64, total_count uint32) (uint32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		hbDO := &dbo.RedPacketDO{
			ChatID:      chatID,
			OwnerUID:    ownerUID,
			Type:        honbaoType,
			Titel:       title,
			Price:       price,
			TotalPrice:  total_price,
			RemainPrice: total_price,
			TotalCount:  total_count,
			RemainCount: total_count,
			CreateDate:  uint32(time.Now().Unix()),
		}

		redpacketID, _, err := d.RedPacketDAO.InsertTx(tx, hbDO)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = (uint32)(redpacketID)

		//减少余额
		_, err = d.Wallet.DecBalanceTx(tx, ownerUID, model.WalletRecordType_CreateRedPacket, total_price, uint32(redpacketID), "创建红包")
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

// GetRedPacket 领取红包.
func (d *Dao) GetRedPacket(ctx context.Context, redpacketID, userID uint32, Price float64) (uint32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//领取红包
		_, err := d.RedPacketDAO.DecreaseOneTx(tx, redpacketID, Price)
		if err != nil {
			result.Err = err
			return
		}

		//增加红包记录
		hbrDO := &dbo.TRedPacketRecordDo{
			RedPacketID: redpacketID,
			UserID:      userID,
			Price:       Price,
		}
		rID, _, err := d.RedPacketRecordsDAO.InsertTx(tx, hbrDO)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = (uint32)(rID)

		//增加余额和日志
		_, err = d.Wallet.IncBalanceTx(tx, userID, model.WalletRecordType_GetRedPacket, Price, redpacketID, "领取红包")
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

// QueryRedPacketsRecord .
func (d *Dao) QueryRedPacketsRecord(ctx context.Context, redPacketIDs []uint32) (rList []TRedPacketAllRecord, err error) {
	if len(redPacketIDs) == 0 {
		return []TRedPacketAllRecord{}, nil
	}
	var (
		query = "SELECT id AS red_packet_id, chat_id, owner_uid, type, title, price, total_price, total_count, created_at FROM red_packets WHERE id IN (?)"
		rows  *sqlx.Rows
	)

	query, args, err := sqlx.In(query, redPacketIDs)
	if err != nil {
		log.Error("sqlx.In in QueryRedPacketsRecord(_), error: %v", err)
		return
	}

	rows, err = d.Query(ctx, query, args...)
	if err != nil {
		log.Errorf("queryx in QueryRedPacketsRecord(_), error: %v", err)
		return
	}

	defer rows.Close()

	now := time.Now().Unix()
	var redPackets map[uint32]*TRedPacketAllRecord = make(map[uint32]*TRedPacketAllRecord)
	for rows.Next() {
		v := TRedPacketAllRecord{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in QueryRedPacketsRecord(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		v.IsExpire = now >= int64(v.CreateAt+24*3600)
		v.Users = make([]TUserRecord, 0)
		redPackets[v.RedPacketID] = &v
	}

	rdos, err := d.RedPacketRecordsDAO.SelectsByRids(ctx, redPacketIDs)
	for _, do := range rdos {
		rp, ok := redPackets[do.RedPacketID]
		if ok {
			rp.Users = append(rp.Users, TUserRecord{
				UserID: do.UserID,
				Price:  do.Price,
				GotAt:  do.CreateAt,
			})
		}
	}

	for _, id := range redPacketIDs {
		rp, ok := redPackets[id]
		if ok {
			rList = append(rList, *rp)
		}
	}

	return
}

// QueryMeRecords .
// type_ 1.创建红包 2.领取红包
func (d *Dao) QueryMeRecords(ctx context.Context, uID uint32, type_, count, page uint32) (rList []TRedPacketAllRecord, err error) {
	if type_ == 0 {
		type_ = 1
	}
	if count == 0 {
		count = 20
	}
	if page == 0 {
		page = 1
	}
	var startPs = count * (page - 1)
	query := "SELECT id AS red_packet_id FROM red_packets WHERE owner_uid = ? AND deleted = 0 ORDER BY id DESC LIMIT ?,?"
	if type_ == 2 {
		query = "SELECT red_packet_id FROM red_packet_records WHERE user_id = ? AND deleted = 0 ORDER BY id DESC LIMIT ?,?"
	}
	redPacketIDs := make([]uint32, 0)
	err = d.DB.Select(ctx, &redPacketIDs, query, uID, startPs, count)
	if err != nil {
		log.Errorf("queryx in QueryMeRecords(_), error: %v", err)
		return
	}

	return d.QueryRedPacketsRecord(ctx, redPacketIDs)
}

func (d *Dao) QueryStatistics(ctx context.Context, uID uint32, type_, year uint32) (stat *TRedPacketStatistics, err error) {
	var do *dbo.TRedPacketStatistics
	if type_ == 1 {
		do, err = d.RedPacketDAO.Statistics(ctx, uID, year)
	} else {
		do, err = d.RedPacketRecordsDAO.Statistics(ctx, uID, year)
	}

	if err != nil {
		return
	}

	stat = &TRedPacketStatistics{
		TotalPrice:    do.TotalPrice,
		Count:         do.Count,
		TopPriceCount: do.TopPriceCount,
	}

	return
}
