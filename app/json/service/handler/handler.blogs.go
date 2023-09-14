package handler

import (
	"context"
	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TBlogsUserPrivacyRule struct {
		Rule   int8    `json:"rule"`
		Users  []int32 `json:"users"`
		Days   int32   `json:"days"`
		Counts int32   `json:"counts"`
	}

	TBlogsUserPrivacy struct {
		Key   int8                    `json:"key"`
		Rules []TBlogsUserPrivacyRule `json:"rules"`
	}

	TBlogsUserPrivacyList struct {
		PrivacyList []TBlogsUserPrivacy `json:"privacyList"`
	}

	TBlogsSetPrivacy struct {
		Privacy TBlogsUserPrivacy `json:"privacy"`
	}

	TBlogsModifyPrivacyUsers struct {
		IsAdding bool    `json:"isAdding"`
		Key      int8    `json:"key"`
		Users    []int32 `json:"users"`
	}

	TBlogsUIDList struct {
		UIDList []int32 `json:"uidList"`
	}

	ServiceBlogs interface {
		SetPrivacy(context.Context, *grpc_util.RpcMetadata, *TBlogsSetPrivacy) *helper.ResultJSON
		ModifyPrivacyUsers(context.Context, *grpc_util.RpcMetadata, *TBlogsModifyPrivacyUsers) *helper.ResultJSON
		GetAllPrivacy(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
	}
)

func RegisterBlogs(s ServiceBlogs) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//设置权限
		"blogs.setPrivacy": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TBlogsSetPrivacy{}, func(data interface{}) *helper.ResultJSON {
				return s.SetPrivacy(ctx, md, data.(*TBlogsSetPrivacy))
			})
		},
		//更改权限相关用户
		"blogs.modifyPrivacyUsers": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TBlogsModifyPrivacyUsers{}, func(data interface{}) *helper.ResultJSON {
				return s.ModifyPrivacyUsers(ctx, md, data.(*TBlogsModifyPrivacyUsers))
			})
		},
		//查询权限
		"blogs.getAllPrivacy": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.GetAllPrivacy(ctx, md)
			})
		},
	}
}
