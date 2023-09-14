package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TChannelDeleteSelfMessage struct {
		ChannelID  uint32   `json:"channelId"`  //频道id
		MessageIds []uint32 `json:"messageIds"` //消息id
	}

	TChannelCountOnline struct {
		ChannelID int32 `json:"channelID"` //频道id
	}

	// ServiceChannel
	ServiceChannel interface {
		CleanMessages(context.Context, *grpc_util.RpcMetadata, *TChannelDeleteSelfMessage) *helper.ResultJSON
		CountOnline(context.Context, *grpc_util.RpcMetadata, *TChannelCountOnline) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterChannel(s ServiceChannel) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//查询清理消息
		"channel.cleanMessages": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChannelDeleteSelfMessage{}, func(data interface{}) *helper.ResultJSON {
				return s.CleanMessages(ctx, md, data.(*TChannelDeleteSelfMessage))
			})
		},
		//查询在线人数
		"channel.countOnline": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChannelCountOnline{}, func(data interface{}) *helper.ResultJSON {
				return s.CountOnline(ctx, md, data.(*TChannelCountOnline))
			})
		},
	}
}
