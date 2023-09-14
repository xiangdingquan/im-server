package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

// ServiceUser
type (
	TUsersInfo struct {
		Uids []uint32 `json:"uIds"`
	}

	TUsersSearchByPhone struct {
		Phone string `json:"phone"`
	}

	TGetUserInfoExt struct {
		UserId int32 `json:"userId"`
	}

	TSetUserInfoExt struct {
		Gender      int32  `json:"gender"`
		Birth       string `json:"birth"`
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		Province    string `json:"province"`
		City        string `json:"city"`
		CityCode    string `json:"cityCode"`
	}

	ServiceUsers interface {
		Info(context.Context, *grpc_util.RpcMetadata, *TUsersInfo) *helper.ResultJSON
		QueryPrivacySettings(context.Context, *grpc_util.RpcMetadata, *TUsersInfo) *helper.ResultJSON
		SearchByPhone(context.Context, *grpc_util.RpcMetadata, *TUsersSearchByPhone) *helper.ResultJSON
		SetUserInfoExt(context.Context, *grpc_util.RpcMetadata, *TSetUserInfoExt) *helper.ResultJSON
		GetUserInfoExt(context.Context, *grpc_util.RpcMetadata, *TGetUserInfoExt) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterUsers(s ServiceUsers) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//获取用户姓名
		"users.info": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TUsersInfo{}, func(data interface{}) *helper.ResultJSON {
				return s.Info(ctx, md, data.(*TUsersInfo))
			})
		},

		//获取隐私设置
		"users.queryPrivacySettings": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TUsersInfo{}, func(data interface{}) *helper.ResultJSON {
				return s.QueryPrivacySettings(ctx, md, data.(*TUsersInfo))
			})
		},

		//根据手机号获取用户id
		"users.searchByPhone": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TUsersSearchByPhone{}, func(data interface{}) *helper.ResultJSON {
				return s.SearchByPhone(ctx, md, data.(*TUsersSearchByPhone))
			})
		},
		//用户信息
		"users.setUserInfoExt": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TSetUserInfoExt{}, func(data interface{}) *helper.ResultJSON {
				return s.SetUserInfoExt(ctx, md, data.(*TSetUserInfoExt))
			})
		},
		"users.getUserInfoExt": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TGetUserInfoExt{}, func(data interface{}) *helper.ResultJSON {
				return s.GetUserInfoExt(ctx, md, data.(*TGetUserInfoExt))
			})
		},
	}
}
