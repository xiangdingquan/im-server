package handler

import (
	"context"
	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TBatchSend struct {
		Users   []int32 `json:"users"`
		Message string  `json:"message"`
	}

	TGetBatchSend struct {
		FromId int32 `json:"fromId"`
		Limit  int32 `json:"limit"`
	}

	TSendReaction struct {
		MessageId  int32 `json:"messageId"`
		ChatId     int64 `json:"chatId"`
		Type       int8  `json:"type"`
		ReactionId int8  `json:"reactionId"`
	}

	TGetMessagesReactions struct {
		MessageIds []int32 `json:"messageIds"`
		ChatId     int64   `json:"chatId"`
		Type       int8    `json:"type"`
	}

	ServiceMessages interface {
		BatchSend(context.Context, *grpc_util.RpcMetadata, *TBatchSend) *helper.ResultJSON
		GetBatchSend(context.Context, *grpc_util.RpcMetadata, *TGetBatchSend) *helper.ResultJSON
		ClearBatchSend(context.Context, *grpc_util.RpcMetadata) *helper.ResultJSON
		SendReaction(ctx context.Context, metadata *grpc_util.RpcMetadata, reaction *TSendReaction) *helper.ResultJSON
		GetMessagesReactions(ctx context.Context, metadata *grpc_util.RpcMetadata, reactions *TGetMessagesReactions) *helper.ResultJSON
	}
)

func RegisterMessages(s ServiceMessages) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//群发助手
		"messages.batchSend": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TBatchSend{}, func(data interface{}) *helper.ResultJSON {
				return s.BatchSend(ctx, md, data.(*TBatchSend))
			})
		},
		"messages.getBatchSend": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TGetBatchSend{}, func(data interface{}) *helper.ResultJSON {
				return s.GetBatchSend(ctx, md, data.(*TGetBatchSend))
			})
		},
		"messages.clearBatchSend": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(nil, func(data interface{}) *helper.ResultJSON {
				return s.ClearBatchSend(ctx, md)
			})
		},
		//消息表情评论
		"messages.sendReaction": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TSendReaction{}, func(data interface{}) *helper.ResultJSON {
				return s.SendReaction(ctx, md, data.(*TSendReaction))
			})
		},
		"messages.getMessagesReactions": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TGetMessagesReactions{}, func(data interface{}) *helper.ResultJSON {
				return s.GetMessagesReactions(ctx, md, data.(*TGetMessagesReactions))
			})
		},
	}
}
