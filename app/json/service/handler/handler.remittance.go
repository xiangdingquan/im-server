package handler

import (
	"context"
	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TRemittanceRemit struct {
		ChatId      uint32  `json:"chatId"`      //IM会话标识
		Type        uint8   `json:"type"`        //1.单聊转账 2.群聊转账
		Description string  `json:"description"` //转账说明
		Amount      float64 `json:"amount"`      //转账金额
		Password    string  `json:"password"`    //钱包密码(MD5)
		Payee       uint32  `json:"payee"`       //收款人
	}

	TRemittanceReceive struct {
		RemittanceId uint32 `json:"remittanceId"` //转帐id
	}

	TRemittanceRefund struct {
		RemittanceId uint32 `json:"remittanceId"`
	}

	TRemittanceGetRecord struct {
		RemittanceId uint32 `json:"remittanceId"`
	}

	TRemittanceGetRecords struct {
		Type   uint32 `json:"type"` //1.支出 2.收入
		FromId uint32 `json:"fromId"`
		Limit  uint32 `json:"limit"`
	}

	TRemittanceRemind struct {
		RemittanceId uint32 `json:"remittanceId"`
	}

	TRemittanceRecord struct {
		ID          uint32  `json:"id"`
		ChatId      uint32  `json:"chatId"`
		PayerUID    uint32  `json:"payerUID"`
		PayeeUID    uint32  `json:"payeeUID"`
		Amount      float64 `json:"amount"`
		Status      uint8   `json:"status"`
		Type        uint8   `json:"type"`
		Description string  `json:"description"`
		RemittedAt  uint32  `json:"remittedAt"`
		ReceivedAt  uint32  `json:"receivedAt"`
		RefundedAt  uint32  `json:"refundedAt"`
	}

	ServiceRemittance interface {
		Remit(context.Context, *grpc_util.RpcMetadata, *TRemittanceRemit) *helper.ResultJSON
		Receive(context.Context, *grpc_util.RpcMetadata, *TRemittanceReceive) *helper.ResultJSON
		Refund(context.Context, *grpc_util.RpcMetadata, *TRemittanceRefund) *helper.ResultJSON
		GetRecord(context.Context, *grpc_util.RpcMetadata, *TRemittanceGetRecord) *helper.ResultJSON
		GetRecords(context.Context, *grpc_util.RpcMetadata, *TRemittanceGetRecords) *helper.ResultJSON
		Remind(context.Context, *grpc_util.RpcMetadata, *TRemittanceRemind) *helper.ResultJSON
	}
)

func RegisterRemittance(s ServiceRemittance) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//创建转帐
		"remittance.remit": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceRemit{}, func(data interface{}) *helper.ResultJSON {
				return s.Remit(ctx, md, data.(*TRemittanceRemit))
			})
		},
		//转账-收款
		"remittance.receive": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceReceive{}, func(data interface{}) *helper.ResultJSON {
				return s.Receive(ctx, md, data.(*TRemittanceReceive))
			})
		},
		//转账-退款
		"remittance.refund": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceRefund{}, func(data interface{}) *helper.ResultJSON {
				return s.Refund(ctx, md, data.(*TRemittanceRefund))
			})
		},
		//查询记录
		"remittance.getRecords": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceGetRecords{}, func(data interface{}) *helper.ResultJSON {
				return s.GetRecords(ctx, md, data.(*TRemittanceGetRecords))
			})
		},
		//查询记录列表
		"remittance.getRecord": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceGetRecord{}, func(data interface{}) *helper.ResultJSON {
				return s.GetRecord(ctx, md, data.(*TRemittanceGetRecord))
			})
		},
		//发送提醒消息
		"remittance.remind": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRemittanceRemind{}, func(data interface{}) *helper.ResultJSON {
				return s.Remind(ctx, md, data.(*TRemittanceRemind))
			})
		},
	}
}
