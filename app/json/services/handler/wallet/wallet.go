package wallet

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/json/db/dbo"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	"open.chat/model"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/services/handler/wallet/core"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/pkg/grpc_util"

	"open.chat/app/json/service/handler"
)

type cls struct {
	*core.WalletCore
	user_client.UserFacade
	AuthSessionRpcClient authsessionpb.RPCSessionClient
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New .
func New(s *svc.Service) {
	var err error
	service := &cls{
		WalletCore: core.New(nil),
	}
	service.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	var ac struct {
		Wardenclient *warden.ClientConfig
	}
	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	service.AuthSessionRpcClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)
	s.AppendServices(handler.RegisterWallet(service))
}

func (s *cls) SetPassword(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TWalletPassword) *helper.ResultJSON {
	if r.SmsCode == "" {
		return &helper.ResultJSON{Code: -1, Msg: "please post sms code"}
	}
	if r.Password == "" {
		return &helper.ResultJSON{Code: -2, Msg: "please set password"}
	}
	if len(r.Password) != 32 {
		return &helper.ResultJSON{Code: -3, Msg: "the password format is wrong"}
	}

	me, err := s.UserFacade.GetUserSelf(ctx, md.UserId)
	if err != nil || me == nil {
		return &helper.ResultJSON{Code: -1, Msg: "get self info fail"}
	}

	phoneNumber := me.GetPhone().GetValue()
	if helper.VerifyCode(ctx, consts.SmsCodeType_SetWalletPassword, phoneNumber, r.SmsCode) != nil {
		return &helper.ResultJSON{Code: 400, Msg: "verify Code is wrong"}
	}

	wdo, err := s.WalletDAO.SelectByUid(ctx, uint32(md.GetUserId()))
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "get wallet info fail"}
	}

	if wdo == nil {
		walletID, err := s.CreateWallet(ctx, uint32(md.GetUserId()), r.Password)
		if walletID == 0 || err != nil {
			return &helper.ResultJSON{Code: -5, Msg: "create wallet fail"}
		}
		wdo, err = s.WalletDAO.SelectByID(ctx, walletID)
		if wdo == nil || err != nil {
			return &helper.ResultJSON{Code: -6, Msg: "get wallet info fail"}
		}
	} else {
		_, err := s.WalletDAO.UpdatePassword(ctx, wdo.ID, r.Password)
		if err != nil {
			return &helper.ResultJSON{Code: -7, Msg: "update password fail"}
		}
		wdo.Password = r.Password
	}

	var data = struct {
		Address string `json:"address"`
	}{
		Address: wdo.Address,
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) Info(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	wdo, err := s.WalletDAO.SelectByUid(ctx, uint32(md.GetUserId()))
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet info fail"}
	}

	if wdo == nil {
		return &helper.ResultJSON{Code: 400, Msg: "you did not create a wallet"}
	}

	var data = struct {
		Address     string  `json:"address"`
		Balance     float64 `json:"balance"`
		HasPassword bool    `json:"hasPaymentPassword"`
	}{
		Address:     wdo.Address,
		Balance:     wdo.Balance,
		HasPassword: len(wdo.Password) != 0,
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

type channels struct {
	ChannelID   uint32 `json:"channelID"`
	ChannelName string `json:"channelName"`
	ChannelUrl  string `json:"channelUrl"`
}

func (s *cls) GetThirdChannels(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	res := &struct {
		Url   string `json:"payUrl"`
		CSUid uint32 `json:"csUserId"`
	}{}
	code, msg, err := helper.WebInterface("recharge", &struct {
		Uid    uint32      `json:"uid"`
		AuthId int64       `json:"authId"`
		Data   interface{} `json:"data"`
	}{
		Uid:    uint32(md.GetUserId()),
		AuthId: md.AuthId,
		Data:   &map[string]interface{}{},
	}, res)

	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "web server request error"}
	}

	var data = struct {
		CSUid    uint32     `json:"csUserId"`
		Channels []channels `json:"channels"`
	}{
		CSUid:    res.CSUid, //777000
		Channels: make([]channels, 0),
	}

	if code == 200 && res.Url != "" {
		data.Channels = append(data.Channels, channels{
			ChannelID:  uint32(len(data.Channels) + 1),
			ChannelUrl: res.Url,
		})
	}

	return &helper.ResultJSON{Code: 200, Msg: msg, Data: data}
}

func (s *cls) TopUp(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TWalletTopUp) *helper.ResultJSON {
	wdo, err := s.WalletDAO.SelectByUid(ctx, uint32(md.GetUserId()))
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet info fail"}
	}

	if wdo == nil {
		return &helper.ResultJSON{Code: 400, Msg: "you did not create a wallet"}
	}

	res := &struct {
		CsUserId uint32 `json:"csUserId"`
		Url      string `json:"payUrl"`
	}{}
	code, msg, err := helper.WebInterface("recharge", &struct {
		Uid    uint32      `json:"uid"`
		AuthId int64       `json:"authId"`
		Data   interface{} `json:"data"`
	}{
		Uid:    wdo.UID,
		AuthId: md.AuthId,
		Data:   &map[string]interface{}{},
	}, res)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "web server request error"}
	}

	if code != 200 {
		return &helper.ResultJSON{Code: -3, Msg: msg}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: struct {
		PayUrl string `json:"payUrl"`
	}{PayUrl: res.Url}}
}

func (s *cls) Withdraw(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TWalletWithdraw) *helper.ResultJSON {
	wdo, err := s.WalletDAO.SelectByUid(ctx, uint32(md.GetUserId()))
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet info fail"}
	}

	if wdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "you did not create a wallet"}
	}

	if wdo.Balance < r.Amount {
		return &helper.ResultJSON{Code: 400, Msg: "insufficient balance"}
	}

	res := &struct {
		CsUserId uint32 `json:"csUserId"`
		Url      string `json:"checkoutUrl"`
	}{}
	code, msg, err := helper.WebInterface("withdraw", &struct {
		Uid    uint32      `json:"uid"`
		AuthId int64       `json:"authId"`
		Data   interface{} `json:"data"`
	}{
		Uid:    wdo.UID,
		AuthId: md.AuthId,
		Data: map[string]interface{}{
			"amount": r.Amount,
		},
	}, res)
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "web server request error"}
	}

	if code != 200 {
		return &helper.ResultJSON{Code: -4, Msg: msg}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: struct {
		CheckoutUrl string `json:"checkoutUrl"`
	}{CheckoutUrl: res.Url}}
}

func (s *cls) Records(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TWalletRecords) *helper.ResultJSON {
	wrdos, err := s.WalletRecordDAO.SelectsByUid(ctx, uint32(md.GetUserId()), r.Count, r.Page)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet records info fail"}
	}

	var data = make([]struct {
		ID       uint32   `json:"id"`
		Type     int8     `json:"type"`     //类型 1.充值 2提现 3.收款 4.转账 5.创建红包 6.领取红包 7.红包退回
		Amount   float64  `json:"amount"`   //变动金额
		Related  uint32   `json:"related"`  //关联的id 1.充值/提现id 2.红包id 3.收款/转账id
		Remarks  string   `json:"remarks"`  //备注
		CreateAt uint32   `json:"createAt"` //创建时间
		ExtData  *extData `json:"extData"`
	}, len(wrdos))

	for i, wr := range wrdos {
		data[i].ID = wr.ID
		data[i].Type = wr.Type
		data[i].Amount = wr.Amount
		data[i].Related = wr.Related
		//data[i].Remarks = wr.Remarks
		data[i].Remarks = s.getRemark(ctx, wr.Type)
		data[i].CreateAt = uint32(wr.Date)

		ext, err := s.getExtData(ctx, wr.Related, wr.Type)
		if err == nil && ext != nil {
			data[i].ExtData = ext
		}

	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

type extData struct {
	RemittanceInfo *dbo.RemittanceDO `json:"remittanceInfo"`
}

func (s *cls) getExtData(ctx context.Context, relatedId uint32, recordType int8) (*extData, error) {
	switch recordType {
	case model.WalletRecordType_RemitRemittance:
		fallthrough
	case model.WalletRecordType_ReceiveRemittance:
		fallthrough
	case model.WalletRecordType_RefundRemittanceByUser:
		fallthrough
	case model.WalletRecordType_RefundRemittanceBySystem:
		remittanceDo, err := s.RemittanceDao.SelectByID(ctx, relatedId)
		if err != nil {
			return nil, err
		}
		return &extData{RemittanceInfo: remittanceDo}, nil

	default:
		return nil, nil
	}
}

func (s *cls) getRemark(ctx context.Context, recordType int8) string {
	return model.Localize(ctx, s.AuthSessionRpcClient, model.WalletRecordTypeToRemark(recordType))
}
