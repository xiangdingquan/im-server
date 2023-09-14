package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

// ServiceDiscover .
type ServiceDiscover interface {
	List(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
}

// RegisterDiscover .
func RegisterDiscover(s ServiceDiscover) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//发现列表菜单
		"discover.list": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(interface{}) *helper.ResultJSON {
				return s.List(ctx, md)
			})
		},
	}
}
