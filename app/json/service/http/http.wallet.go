package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/helper"
)

// ServiceUser
type (
	TWalletUid struct {
		Uid uint32 `json:"userId"`
	}

	TWalletModifyBalance struct {
		Uid     uint32  `json:"userId"`
		Amount  float64 `json:"amount"`  //操作金额
		Remarks string  `json:"remarks"` //修改说明
	}

	TWalletPaymentResult struct {
		Uid     uint32  `json:"userId"`
		Type    uint8   `json:"type"`    //1.充值下单返回 2.充值结果 3提现下单返回 4.提现结果
		Success bool    `json:"success"` //0失败 1成功
		Amount  float64 `json:"amount"`  //操作金额
		Related uint32  `json:"related"` //相关联的id
		Remarks string  `json:"remarks"` //结果备注
	}

	TWalletCheckWithdraw struct {
		Uid     uint32  `json:"userId"`  //用户id
		Amount  float64 `json:"amount"`  //提现金额
		Related uint32  `json:"related"` //相关联的id
	}

	ServiceWallet interface {
		QueryBalance(context.Context, *TWalletUid) *helper.ResultJSON
		ModifyBalance(context.Context, *TWalletModifyBalance) *helper.ResultJSON
		PaymentResult(context.Context, *TWalletPaymentResult) *helper.ResultJSON
		CheckWithdraw(context.Context, *TWalletCheckWithdraw) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterWallet(s ServiceWallet, rg *bm.RouterGroup) {
	rg2 := rg.Group("/wallet")
	rg2.POST("/queryBalance", func(c *bm.Context) {
		helper.DoHttpJson(c, &TWalletUid{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.QueryBalance(ctx, data.(*TWalletUid))
		})
	})

	rg2.POST("/modifyBalance", func(c *bm.Context) {
		helper.DoHttpJson(c, &TWalletModifyBalance{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.ModifyBalance(ctx, data.(*TWalletModifyBalance))
		})
	})

	rg2.POST("/paymentResult", func(c *bm.Context) {
		helper.DoHttpJson(c, &TWalletPaymentResult{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.PaymentResult(ctx, data.(*TWalletPaymentResult))
		})
	})

	rg2.POST("/checkWithdraw", func(c *bm.Context) {
		helper.DoHttpJson(c, &TWalletCheckWithdraw{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.CheckWithdraw(ctx, data.(*TWalletCheckWithdraw))
		})
	})
}
