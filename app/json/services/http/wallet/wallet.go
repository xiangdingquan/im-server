package wallet

import (
	"context"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/services/handler/wallet/core"
	"open.chat/app/json/services/handler/wallet/dao"
	"open.chat/model"

	"open.chat/app/json/service/http"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

type cls struct {
	*core.WalletCore
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{
		WalletCore: core.New(dao.New()),
	}
	http.RegisterWallet(service, rg)
}

func (s *cls) QueryBalance(ctx context.Context, r *http.TWalletUid) *helper.ResultJSON {
	wdo, err := s.SelectByUid(ctx, r.Uid)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet info fail"}
	}

	if wdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "you did not create a wallet"}
	}

	var data = struct {
		Balance float64 `json:"balance"`
	}{
		Balance: wdo.Balance,
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}

func (s *cls) ModifyBalance(ctx context.Context, r *http.TWalletModifyBalance) *helper.ResultJSON {
	s.IncBalance(ctx, r.Uid, model.WalletRecordType_Modify, r.Amount, 0, r.Remarks)
	return s.QueryBalance(ctx, &http.TWalletUid{Uid: r.Uid})
}

func (s *cls) PaymentResult(ctx context.Context, r *http.TWalletPaymentResult) *helper.ResultJSON {
	if r.Type == 1 { //充值
		if r.Success {
			s.IncBalance(ctx, r.Uid, model.WalletRecordType_Recharge, r.Amount, r.Related, r.Remarks)
		}
	} else if r.Type == 2 { //提现
		if !r.Success {
			s.IncBalance(ctx, r.Uid, model.WalletRecordType_WithdrawFail, r.Amount, r.Related, r.Remarks)
		}
	} else {
		return &helper.ResultJSON{Code: -1, Msg: "type error"}
	}
	//通知处理
	return s.QueryBalance(ctx, &http.TWalletUid{Uid: r.Uid})
}

func (s *cls) CheckWithdraw(ctx context.Context, r *http.TWalletCheckWithdraw) *helper.ResultJSON {
	_, err := s.IncBalance(ctx, r.Uid, model.WalletRecordType_Withdrawal, -r.Amount, r.Related, "申请提现")
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "get wallet info fail"}
	}

	return &helper.ResultJSON{Code: 0, Msg: "success"}
}
