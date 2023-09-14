package remittance

import (
	"context"
	"open.chat/app/json/db/dbo"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/handler"
	"open.chat/app/json/services/handler/remittance/core"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

type (
	Config struct {
		MaxMoney float64
	}

	cls struct {
		*core.RemittanceCore
		user_client.UserFacade
	}
)

var G_RemittanceCfg *Config = nil

func New(s *svc.Service) {
	service := &cls{
		RemittanceCore: core.New(nil),
	}
	var err error
	service.UserFacade, err = user_client.NewUserFacade("local")
	helper.CheckErr(err)
	s.AppendServices(handler.RegisterRemittance(service))
}

func (s *cls) Remit(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceRemit) *helper.ResultJSON {
	// 现在只支持单聊转账
	if r.Type != core.RemittanceTypeSingle {
		return &helper.ResultJSON{Code: -1, Msg: "type error"}
	}

	if r.Type == core.RemittanceTypeSingle {
		// 单聊
		//检查对方是否为黑名单
		if s.IsBlockedByUser(ctx, int32(r.ChatId), md.UserId) {
			return &helper.ResultJSON{Code: 501, Msg: "user is blacklist,please check!"}
		}
	}

	if r.Payee == uint32(md.UserId) {
		log.Errorf("remit remittance, invalid payee")
		return &helper.ResultJSON{Code: -5, Msg: "invalid payee"}
	}

	if r.Amount > G_RemittanceCfg.MaxMoney {
		return &helper.ResultJSON{Code: 402, Msg: "remittance amount exceed limit"}
	}
	if r.Amount < 1e-2 {
		return &helper.ResultJSON{Code: 401, Msg: "too little amount"}
	}

	if r.Password == "" {
		return &helper.ResultJSON{Code: -2, Msg: "please submit your wallet password"}
	}

	wdo, err := s.Wallet.SelectByUid(ctx, uint32(md.UserId))
	if wdo == nil || err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "get wallet info fail"}
	}

	if r.Password != wdo.Password {
		return &helper.ResultJSON{Code: 404, Msg: "your wallet password is wrong"}
	}

	if wdo.Balance < r.Amount {
		return &helper.ResultJSON{Code: 400, Msg: "insufficient balance"}
	}

	payerUID := uint32(md.GetUserId())
	rID, err := s.Dao.CreateRemittance(ctx, r.ChatId, payerUID, r.Payee, r.Type, core.RemittanceStatusInitial, r.Description, r.Amount)
	if rID == 0 || err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "create remittance fail"}
	}

	err = s.sendMsg(ctx, md.UserId, md.GetAuthId(), core.MidOfRemittance, core.SidOfRemit, r.ChatId, r.Type, rID, payerUID, r.Payee, r.Amount)
	if err != nil {
		log.Errorf("remittance.Create[%s]", err.Error())
		return &helper.ResultJSON{Code: -5, Msg: "send message fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

//var _mutex sync.Mutex

func (s *cls) Receive(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceReceive) *helper.ResultJSON {
	if r.RemittanceId == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "invalid remittance id"}
	}

	//_mutex.Lock()
	//defer _mutex.Unlock()

	rdo, err := s.RemittanceDao.SelectByID(ctx, r.RemittanceId)
	if err != nil {
		log.Errorf("select remittance failed, remittanceID:%d, error: %v", r.RemittanceId, err)
		return &helper.ResultJSON{Code: -2, Msg: "select remittance failed"}
	}

	if errJson := s.checkRemittance(rdo, uint32(md.UserId)); errJson != nil {
		return errJson
	}

	err = s.Dao.ReceiveRemittance(ctx, r.RemittanceId, rdo.PayeeUID, rdo.Amount)
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "get remittance failed"}
	}

	var chatID uint32
	if rdo.Type == core.RemittanceTypeSingle {
		chatID = rdo.PayerUID
	} else {
		chatID = rdo.ChatID
	}
	err = s.sendMsg(ctx, md.UserId, md.GetAuthId(), core.MidOfRemittance, core.SidOfReceive, chatID, rdo.Type, rdo.ID, rdo.PayerUID, rdo.PayeeUID, rdo.Amount)

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) Refund(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceRefund) *helper.ResultJSON {
	if r.RemittanceId == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "invalid remittance id"}
	}

	//_mutex.Lock()
	//defer _mutex.Unlock()

	rdo, err := s.RemittanceDao.SelectByID(ctx, r.RemittanceId)
	if err != nil {
		log.Errorf("select remittance failed, remittanceID:%d, error: %v", r.RemittanceId, err)
		return &helper.ResultJSON{Code: -2, Msg: "select remittance failed"}
	}

	if errJson := s.checkRemittance(rdo, uint32(md.UserId)); errJson != nil {
		return errJson
	}

	err = s.RefundRemittance(ctx, r.RemittanceId, rdo.PayerUID, rdo.Amount)
	if err != nil {
		log.Errorf("refund remittance failed, remittanceId:%d, error: %v", r.RemittanceId, err)
		return &helper.ResultJSON{Code: -3, Msg: "refund failed"}
	}

	if errJson := s.checkRemittance(rdo, uint32(md.UserId)); errJson != nil {
		return errJson
	}
	var chatID uint32
	if rdo.Type == core.RemittanceTypeSingle {
		chatID = rdo.PayerUID
	} else {
		chatID = rdo.ChatID
	}
	err = s.sendMsg(ctx, md.UserId, md.GetAuthId(), core.MidOfRemittance, core.SidOfRefundByUser, chatID, rdo.Type, rdo.ID, rdo.PayerUID, rdo.PayeeUID, rdo.Amount)

	return &helper.ResultJSON{Code: 200, Msg: "success"}

}

func (s *cls) GetRecords(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceGetRecords) *helper.ResultJSON {
	var (
		doList []*dbo.RemittanceDO
		err    error
		out    *helper.ResultJSON
	)

	switch r.Type {
	case core.SearchTypePayer:
		if doList, err = s.RemittanceDao.SelectByPayer(ctx, uint32(md.UserId), r.FromId, r.Limit); err != nil {
			log.Errorf("select remittances by payer failed, error: %v", err)
		} else {
			out = &helper.ResultJSON{
				Code: 200,
				Msg:  "success",
				Data: s.dbToJsonList(doList),
			}
		}
	case core.SearchTypePayee:
		if doList, err = s.RemittanceDao.SelectByPayee(ctx, uint32(md.UserId), r.FromId, r.Limit); err != nil {
			log.Errorf("select remittances by payee failed, error: %v", err)
		} else {
			out = &helper.ResultJSON{
				Code: 200,
				Msg:  "success",
				Data: s.dbToJsonList(doList),
			}
		}
	default:
		log.Errorf("get remittance with invalid type:%d", r.Type)
		out = &helper.ResultJSON{Code: -1, Msg: "invalid type"}
	}
	return out
}

func (s *cls) GetRecord(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceGetRecord) *helper.ResultJSON {
	rdo, err := s.RemittanceDao.SelectByID(ctx, r.RemittanceId)
	if err != nil {
		log.Errorf("get remittance record with id(%d) failed, error: %v", r.RemittanceId, err)
		return &helper.ResultJSON{Code: -1, Msg: "get record failed"}
	}

	return &helper.ResultJSON{
		Code: 200,
		Msg:  "success",
		Data: s.dbToJson(rdo),
	}
}

func (s *cls) Remind(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TRemittanceRemind) *helper.ResultJSON {
	rdo, err := s.RemittanceDao.SelectByID(ctx, r.RemittanceId)
	if err != nil {
		log.Errorf("get remittance record with id(%d) failed, error: %v", r.RemittanceId, err)
		return &helper.ResultJSON{Code: -1, Msg: "get record failed"}
	}

	if rdo.PayerUID != uint32(md.UserId) {
		log.Errorf("remind remittance, uid not match in , payer uid is %d, but request uid is %d", rdo.PayeeUID, md.UserId)
		return &helper.ResultJSON{Code: -2, Msg: "invalid uid"}
	}

	if rdo.Status != core.RemittanceStatusInitial {
		log.Errorf("remind remittance, invalid status (%d)", rdo.Status)
		return &helper.ResultJSON{Code: -3, Msg: "invalid status"}
	}

	if rdo.Type != core.RemittanceTypeSingle {
		log.Errorf("remind remittance, unsupported type (%d)", rdo.Type)
		return &helper.ResultJSON{Code: -4, Msg: "invalid type"}
	}

	sender := helper.MakeSender(uint32(md.UserId), md.GetAuthId(), core.MidOfRemittance, core.SidOfRemind)
	msg := struct {
		RemittanceId uint32 `json:"remittanceId"`
	}{
		RemittanceId: r.RemittanceId,
	}

	err = sender.SendToUser(ctx, rdo.ChatID, &msg)

	if err != nil {
		log.Errorf("send message filed, chatID:%d, remittanceID:%d", rdo.ChatID, r.RemittanceId)
		return &helper.ResultJSON{Code: -5, Msg: "send message failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) sendMsg(ctx context.Context, uid int32, aid int64, mid, sid uint32, chatID uint32, rType uint8, rID uint32, payerUID, payeeUID uint32, amount float64) error {
	sender := helper.MakeSender(uint32(uid), aid, mid, sid)
	msg := struct {
		RemittanceId uint32  `json:"remittanceId"`
		Payer        uint32  `json:"payer"`
		Payee        uint32  `json:"payee"`
		Amount       float64 `json:"amount"`
	}{
		RemittanceId: rID,
		Payer:        payerUID,
		Payee:        payeeUID,
		Amount:       amount,
	}

	var err error
	if rType == 1 {
		err = sender.SendToUser(ctx, chatID, &msg)
	} else {
		err = sender.SendToChannel(ctx, chatID, &msg)
	}

	if err != nil {
		log.Errorf("send message filed, chatID:%d, remittanceID:%d", chatID, rID)
	}

	return err
}

func (s *cls) checkRemittance(rdo *dbo.RemittanceDO, payeeUID uint32) *helper.ResultJSON {
	if rdo == nil {
		return &helper.ResultJSON{Code: -100, Msg: "remittance not found"}
	}

	if rdo.PayeeUID != payeeUID {
		return &helper.ResultJSON{Code: 400, Msg: "you are not payee"}
	}

	if rdo.Status != 0 {
		return &helper.ResultJSON{Code: -101, Msg: "invalid status"}
	}
	return nil
}

func (s *cls) dbToJson(rdo *dbo.RemittanceDO) *handler.TRemittanceRecord {
	out := &handler.TRemittanceRecord{
		ID:          rdo.ID,
		ChatId:      rdo.ChatID,
		PayerUID:    rdo.PayerUID,
		PayeeUID:    rdo.PayeeUID,
		Amount:      rdo.Amount,
		Status:      rdo.Status,
		Type:        rdo.Type,
		Description: rdo.Description,
	}

	out.RemittedAt = uint32(rdo.CreateTime.Unix())
	switch rdo.Status {
	case core.RemittanceStatusReceived:
		out.ReceivedAt = uint32(rdo.UpdateTime.Unix())
	case core.RemittanceStatusRefundedByUser:
		fallthrough
	case core.RemittanceStatusRefundedBySystem:
		out.RefundedAt = uint32(rdo.UpdateTime.Unix())
	}

	return out
}

func (s *cls) dbToJsonList(l []*dbo.RemittanceDO) (out []*handler.TRemittanceRecord) {
	out = make([]*handler.TRemittanceRecord, len(l))
	for i, v := range l {
		out[i] = s.dbToJson(v)
	}
	return
}
