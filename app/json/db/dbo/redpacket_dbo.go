package dbo

import "time"

type (
	// RedPacketDO .
	RedPacketDO struct {
		ID          uint32    `db:"id"`
		ChatID      uint32    `db:"chat_id"`
		OwnerUID    uint32    `db:"owner_uid"`
		Type        uint8     `db:"type"`         //1.单聊红包 2.拼手气红包 3.普通红包
		Titel       string    `db:"title"`        //标题
		Price       float64   `db:"price"`        //单个金额
		TotalPrice  float64   `db:"total_price"`  //总金额
		TotalCount  uint32    `db:"total_count"`  //红包总数量
		RemainPrice float64   `db:"remain_price"` //剩余红包金额
		RemainCount uint32    `db:"remain_count"` //剩余红包数量
		CreateDate  uint32    `db:"create_date"`  //创建时间
		Completed   bool      `db:"completed"`
		Deleted     bool      `db:"deleted"`
		CreateTime  time.Time `db:"created_at"`
		CreateAt    uint32    `db:"-"`
	}

	// TRedPacketRecordDo .
	TRedPacketRecordDo struct {
		ID          uint32    `db:"id"`
		RedPacketID uint32    `db:"red_packet_id"`
		UserID      uint32    `db:"user_id"`
		Price       float64   `db:"price"`      //获取到的金额
		Deleted     bool      `db:"deleted"`    //
		CreateTime  time.Time `db:"created_at"` //
		CreateAt    uint32    `db:"-"`
	}

	TRedPacketStatistics struct {
		TotalPrice    float64 `db:"total_price"`
		Count         uint32  `db:"cnt"`
		TopPriceCount uint32  `db:"top_price_count"`
	}
)
