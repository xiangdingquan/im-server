package redpacket

import (
	"context"
	"math"
	"math/rand"
	"sync"

	"open.chat/app/json/db/dbo"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/services/handler/redpacket/core"
	"open.chat/app/json/services/handler/redpacket/dao"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/app/json/service/handler"
	user_client "open.chat/app/service/biz_service/user/client"
)

type (
	RedpacketConfig struct {
		MaxCount uint32
		MaxMoney float64
	}

	cls struct {
		*core.RedPacketCore
		user_client.UserFacade
	}
)

var G_RedpacketCfg *RedpacketConfig = nil

// New .
func New(s *svc.Service) {
	c := &cls{
		RedPacketCore: core.New(nil),
	}
	var err error
	c.UserFacade, err = user_client.NewUserFacade("local")
	helper.CheckErr(err)
	s.AppendServices(handler.RegisterRedPacket(c))
}

func (s *cls) Create(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRedPacketCreate) *helper.ResultJSON {
	if r.Type == 1 {
		r.Count = 1
		//检查对方是否为黑名单
		if s.IsBlockedByUser(ctx, int32(r.ChatId), md.UserId) {
			return &helper.ResultJSON{Code: 501, Msg: "user is blacklist,please check!"}
		}
	} else {
		if r.Type != 2 && r.Type != 3 {
			return &helper.ResultJSON{Code: -1, Msg: "type error"}
		}
		if r.Count < 1 {
			return &helper.ResultJSON{Code: -2, Msg: "redpacket number error"}
		}
		if r.Count > G_RedpacketCfg.MaxCount {
			return &helper.ResultJSON{Code: 401, Msg: "redpacket number error"}
		}
		if r.Type == 2 {
			r.Price = 0
		}
	}

	if r.Type != 2 {
		r.TotalPrice = r.Price * float64(r.Count)
	}

	if r.TotalPrice < 1e-2*float64(r.Count) {
		return &helper.ResultJSON{Code: 402, Msg: "too little amount"}
	}

	if r.Price > float64(G_RedpacketCfg.MaxMoney) || r.TotalPrice/float64(r.Count) > G_RedpacketCfg.MaxMoney {
		return &helper.ResultJSON{Code: 403, Msg: "a redpacket amount error"}
	}

	if r.Password == "" {
		return &helper.ResultJSON{Code: -8, Msg: "please submit your wallet password"}
	}

	wdo, err := s.Wallet.SelectByUid(ctx, uint32(md.UserId))
	if wdo == nil || err != nil {
		return &helper.ResultJSON{Code: -5, Msg: "get wallet info fail"}
	}

	if r.Password != wdo.Password {
		return &helper.ResultJSON{Code: 404, Msg: "your wallet password is wrong"}
	}

	if wdo.Balance < r.TotalPrice {
		return &helper.ResultJSON{Code: 400, Msg: "insufficient balance"}
	}

	owner := uint32(md.GetUserId())
	redpacketID, err := s.Dao.CreateRedPacket(ctx, r.ChatId, owner, r.Type, r.Title, r.Price, r.TotalPrice, r.Count)
	if redpacketID == 0 || err != nil {
		return &helper.ResultJSON{Code: -7, Msg: "create redpacket fail"}
	}

	sender := helper.MakeSender(owner, md.GetAuthId(), 2, 100)
	msg := struct {
		RedPacketID uint32 `json:"redPacketId"` //红包id
		From        uint32 `json:"from"`        //所有者
		Title       string `json:"title"`       //标题
	}{
		RedPacketID: redpacketID,
		From:        owner,
		Title:       r.Title,
	}
	if r.Type == 1 {
		err = sender.SendToUser(ctx, r.ChatId, &msg)
	} else {
		err = sender.SendToChannel(ctx, r.ChatId, &msg)
	}
	if err != nil {
		log.Errorf("redpacket.Create[%s]", err.Error())
		return &helper.ResultJSON{Code: -8, Msg: "send message fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) getRedPacket(ctx context.Context, hdo *dbo.RedPacketDO, meID uint32) bool {
	if hdo == nil || meID == 0 {
		return false
	}
	price := hdo.Price
	if hdo.Type == 2 {
		var min float64 = 0.01
		var max float64 = hdo.RemainPrice / float64(hdo.RemainCount) * 2
		price = hdo.RemainPrice
		//随机金额
		if hdo.RemainCount > 1 {
			price = rand.Float64() * max
			if price < min {
				price = min
			}
			price = math.Floor(price*100) / 100
		}
	}
	//保存记录
	rID, err := s.Dao.GetRedPacket(ctx, hdo.ID, meID, price)
	if rID == 0 || err != nil {
		return false
	}
	hdo.RemainPrice -= price
	hdo.RemainCount--
	return true
}

var _mutex sync.Mutex

func (s *cls) Get(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRedPacketID) *helper.ResultJSON {
	if r.RedPacketID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "redpacket id error"}
	}
	rdo, err := s.RedPacketRecordsDAO.Select(ctx, r.RedPacketID, uint32(md.UserId))
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "record query fail"}
	}
	if rdo != nil {
		return &helper.ResultJSON{Code: -3, Msg: "repeated get"}
	}
	_mutex.Lock()
	defer _mutex.Unlock()
	hdo, err := s.RedPacketDAO.SelectByID(ctx, r.RedPacketID)
	if hdo == nil || err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "not find redpacket record"}
	}
	if hdo.Type == 1 {
		if hdo.OwnerUID == uint32(md.UserId) {
			return &helper.ResultJSON{Code: 400, Msg: "can't get my own redpacket"}
		}
		if hdo.ChatID != uint32(md.UserId) {
			return &helper.ResultJSON{Code: 401, Msg: "this is not your redpacket"}
		}
	}
	result := &helper.ResultJSON{Code: 200, Msg: "success"}
	if hdo.Completed || hdo.RemainCount < 1 {
		result = &helper.ResultJSON{Code: 402, Msg: "redpacket is finished"}
	} else {
		//获取红包
		remainPrice := hdo.RemainPrice
		if !s.getRedPacket(ctx, hdo, uint32(md.UserId)) {
			return &helper.ResultJSON{Code: -5, Msg: "get redpacket error"}
		}
		sender := helper.MakeSender(uint32(md.UserId), md.GetAuthId(), 2, 200)
		msg := struct {
			RedPacketID uint32  `json:"redPacketId"` //红包id
			From        uint32  `json:"from"`        //红包所有人
			Type        uint8   `json:"type"`        //1.单聊红包 2.拼手气红包 3.普通红包
			Price       float32 `json:"price"`       //获得金额
			IsLast      bool    `json:"isLast"`      //是否最后一个红包
		}{
			RedPacketID: hdo.ID,
			From:        hdo.OwnerUID,
			Type:        hdo.Type,
			Price:       float32(remainPrice - hdo.RemainPrice),
			IsLast:      hdo.RemainCount == 0,
		}
		if hdo.Type == 1 {
			err = sender.SendToUser(ctx, hdo.OwnerUID, &msg)
		} else {
			err = sender.SendToChannel(ctx, hdo.ChatID, &msg)
		}
		if err != nil {
			log.Errorf("redpacket.Create[%s]", err.Error())
			return &helper.ResultJSON{Code: -7, Msg: "send message fail"}
		}
	}
	//获取完成
	rdos, err := s.RedPacketRecordsDAO.SelectsByRid(ctx, r.RedPacketID)
	if err != nil {
		return &helper.ResultJSON{Code: -6, Msg: "record query fail"}
	}
	users := make([]dao.TUserRecord, len(rdos))
	for i, do := range rdos {
		users[i] = dao.TUserRecord{
			UserID: do.UserID,
			Price:  do.Price,
			GotAt:  do.CreateAt,
		}
	}
	result.Data = users
	return result
}

func (s *cls) Detail(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRedPacketID) *helper.ResultJSON {
	if r.RedPacketID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "redpacket id error"}
	}
	hdos, err := s.QueryRedPacketsRecord(ctx, []uint32{r.RedPacketID})
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "find redpacket record error"}
	}
	if len(hdos) == 0 {
		return &helper.ResultJSON{Code: -3, Msg: "not find redpacket record"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: hdos[0]}
}

func (s *cls) Record(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRedPacketRecordPage) *helper.ResultJSON {
	hdos, err := s.QueryMeRecords(ctx, uint32(md.UserId), r.Type, r.Count, r.Page)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "find redpacket record error"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: hdos}
}

func (s *cls) Statistics(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRedPacketStatisticsReq) *helper.ResultJSON {
	if r.Type != 1 && r.Type != 2 {
		return &helper.ResultJSON{Code: -1, Msg: "Invalid type"}
	}

	stat, err := s.QueryStatistics(ctx, uint32(md.UserId), r.Type, r.Year)
	if err != nil {
		log.Errorf("[redpacket] query statistics failed, error: %v", err)
		return &helper.ResultJSON{Code: -2, Msg: "query statistics failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: stat}
}
