package handler

import (
	"context"

	"open.chat/pkg/grpc_util"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
)

type (
	// TSyncTime .
	TSyncTime struct {
		ClientTime uint64 `json:"client_time"` //客户端时间戳
	}

	//
	TSendSmsCode struct {
		Type consts.SmsCodeType `json:"type"` //短信类型 1修改钱包密码
	}

	TVerifySmsCode struct {
		Type consts.SmsCodeType `json:"type"`
		Code string             `json:"code"`
	}

	// ServiceSystem 实现设备转发的基础接口
	ServiceSystem interface {
		SyncTime(context.Context, *grpc_util.RpcMetadata, *TSyncTime) *helper.ResultJSON
		RelayData(context.Context, *grpc_util.RpcMetadata, *helper.TrelayData) *helper.ResultJSON
		SendSmsCode(context.Context, *grpc_util.RpcMetadata, *TSendSmsCode) *helper.ResultJSON
		VerifySmsCode(context.Context, *grpc_util.RpcMetadata, *TVerifySmsCode) *helper.ResultJSON
		GetOssConfig(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
	}
)

// RegisterSystem .
func RegisterSystem(s ServiceSystem) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//同步服务器时间
		"system.synctime": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TSyncTime{}, func(data interface{}) *helper.ResultJSON {
				return s.(ServiceSystem).SyncTime(ctx, md, data.(*TSyncTime))
			})
		},
		//转发数据
		"system.relayData": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&helper.TrelayData{}, func(data interface{}) *helper.ResultJSON {
				return s.(ServiceSystem).RelayData(ctx, md, data.(*helper.TrelayData))
			})
		},
		//发送验证码
		"system.sendSmsCode": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TSendSmsCode{}, func(data interface{}) *helper.ResultJSON {
				return s.(ServiceSystem).SendSmsCode(ctx, md, data.(*TSendSmsCode))
			})
		},
		//检查验证码
		"system.verifySmsCode": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TVerifySmsCode{}, func(data interface{}) *helper.ResultJSON {
				return s.(ServiceSystem).VerifySmsCode(ctx, md, data.(*TVerifySmsCode))
			})
		},
		"system.getOssConfig": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.(ServiceSystem).GetOssConfig(ctx, md)
			})
		},
	}
}
