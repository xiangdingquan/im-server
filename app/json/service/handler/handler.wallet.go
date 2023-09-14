package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	// TWalletPassword .
	TWalletPassword struct {
		SmsCode  string `json:"smsCode"`
		Password string `json:"password"`
	}

	// TWalletTopUp .
	TWalletTopUp struct {
		Amount float64 `json:"amount"`
	}

	// TWalletTopUp .
	TWalletWithdraw struct {
		Amount float64 `json:"amount"`
	}

	TWalletRecords struct {
		Count uint32 `json:"count"` //每页数量
		Page  uint32 `json:"page"`  //获取第N页
	}

	// ServiceWallet
	ServiceWallet interface {
		SetPassword(context.Context, *grpc_util.RpcMetadata, *TWalletPassword) *helper.ResultJSON
		Info(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
		Records(context.Context, *grpc_util.RpcMetadata, *TWalletRecords) *helper.ResultJSON
		GetThirdChannels(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
		TopUp(context.Context, *grpc_util.RpcMetadata, *TWalletTopUp) *helper.ResultJSON
		Withdraw(context.Context, *grpc_util.RpcMetadata, *TWalletWithdraw) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterWallet(s ServiceWallet) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//修改密码或创建钱包
		"wallet.setPassword": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TWalletPassword{}, func(data interface{}) *helper.ResultJSON {
				return s.SetPassword(ctx, md, data.(*TWalletPassword))
			})
		},

		//查询钱包信息
		"wallet.info": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(interface{}) *helper.ResultJSON {
				return s.Info(ctx, md)
			})
		},

		//查询钱包记录
		"wallet.records": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TWalletRecords{}, func(data interface{}) *helper.ResultJSON {
				return s.Records(ctx, md, data.(*TWalletRecords))
			})
		},

		//获取三方通道列表
		"wallet.getThirdChannels": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.GetThirdChannels(ctx, md)
			})
		},

		//获取充值地址
		"wallet.topUp": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TWalletTopUp{}, func(data interface{}) *helper.ResultJSON {
				return s.TopUp(ctx, md, data.(*TWalletTopUp))
			})
		},

		//获取提现地址
		"wallet.withdraw": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TWalletWithdraw{}, func(data interface{}) *helper.ResultJSON {
				return s.Withdraw(ctx, md, data.(*TWalletWithdraw))
			})
		},
	}
}
