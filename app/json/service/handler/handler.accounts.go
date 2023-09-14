package handler

import (
	"context"
	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TToggleMultiOnline struct {
		IsOn bool `json:"isOn"`
	}

	TCheckPassword struct {
		Password string `json:"password"`
	}

	TNotificationSettings struct {
		ShowNotification bool  `json:"showNotification"`
		ShowPreview      bool  `json:"showPreview"`
		Sound            int32 `json:"sound"`
		InAppSound       bool  `json:"inAppSound"`
		InAppVibrate     bool  `json:"inAppVibration"`
		InAppPreview     bool  `json:"inAppPreview"`
		CountClosed      bool  `json:"countClosed"`
	}

	ServiceAccounts interface {
		ToggleMultiOnline(context.Context, *grpc_util.RpcMetadata, *TToggleMultiOnline) *helper.ResultJSON
		GetMultiOnline(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
		CheckPassword(context.Context, *grpc_util.RpcMetadata, *TCheckPassword) *helper.ResultJSON
		ModifyNotificationSettings(context.Context, *grpc_util.RpcMetadata, *TNotificationSettings) *helper.ResultJSON
		GetNotificationSettings(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
	}
)

func RegisterAccounts(s ServiceAccounts) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//多端登录设置
		"accounts.toggleMultiOnline": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TToggleMultiOnline{}, func(data interface{}) *helper.ResultJSON {
				return s.ToggleMultiOnline(ctx, md, data.(*TToggleMultiOnline))
			})
		},
		"accounts.getMultiOnline": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.GetMultiOnline(ctx, md)
			})
		},
		// 检查密码
		"accounts.checkPassword": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TCheckPassword{}, func(data interface{}) *helper.ResultJSON {
				return s.CheckPassword(ctx, md, data.(*TCheckPassword))
			})
		},
		//通知设置
		"accounts.modifyNotificationSettings": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TNotificationSettings{}, func(data interface{}) *helper.ResultJSON {
				return s.ModifyNotificationSettings(ctx, md, data.(*TNotificationSettings))
			})
		},
		"accounts.getNotificationSettings": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.GetNotificationSettings(ctx, md)
			})
		},
	}
}
