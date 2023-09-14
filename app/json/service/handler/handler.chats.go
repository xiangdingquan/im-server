package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	TChatID struct {
		ChatID int32 `json:"chatId"` //群聊id 普通群-id 超级群+id
	}

	TChatIDModifyBannedRights struct {
		ChatID             int32 `json:"chatId"`             //群聊id 普通群-id 超级群+id
		BanWhisper         bool  `json:"banWhisper"`         //禁止私聊
		BanSendWebLink     bool  `json:"banSendWebLink"`     //禁止发送网页链接
		BanSendQRcode      bool  `json:"banSendQRcode"`      //禁止发送二维码
		BanSendKeyword     bool  `json:"banSendKeyword"`     //禁止发送关键字
		BanSendDmMention   bool  `json:"banSendDmMention"`   //禁止发送dm@
		KickWhoSendKeyword bool  `json:"kickWhoSendKeyword"` //发送敏感词移出群聊
		ShowKickMessage    bool  `json:"showKickMessage"`    //敏感词移出群聊提示
	}

	TChatIDSetFilterKeywords struct {
		ChatID   int32    `json:"chatId"`   //群聊id 普通群-id 超级群+id
		Keywords []string `json:"keywords"` //关键字列表
	}

	TChatIdSetNickname struct {
		ChatID   int32  `json:"chatId"` //群聊id 普通群-id 超级群+id
		Nickname string `json:"nickname"`
	}

	// ServiceChats
	ServiceChats interface {
		GetBannedRightEx(context.Context, *grpc_util.RpcMetadata, *TChatID) *helper.ResultJSON
		ModifyBannedRightEx(context.Context, *grpc_util.RpcMetadata, *TChatIDModifyBannedRights) *helper.ResultJSON
		GetFilterKeywords(context.Context, *grpc_util.RpcMetadata, *TChatID) *helper.ResultJSON
		SetFilterKeywords(context.Context, *grpc_util.RpcMetadata, *TChatIDSetFilterKeywords) *helper.ResultJSON
		DisableInviteLink(context.Context, *grpc_util.RpcMetadata, *TChatID) *helper.ResultJSON
		KickWhoSendKeyword(context.Context, *grpc_util.RpcMetadata, *TChatID) *helper.ResultJSON
		SetNickname(context.Context, *grpc_util.RpcMetadata, *TChatIdSetNickname) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterChats(s ServiceChats) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		"chats.getBannedRightex": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatID{}, func(data interface{}) *helper.ResultJSON {
				return s.GetBannedRightEx(ctx, md, data.(*TChatID))
			})
		},

		"chats.modifyBannedRightex": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatIDModifyBannedRights{}, func(data interface{}) *helper.ResultJSON {
				return s.ModifyBannedRightEx(ctx, md, data.(*TChatIDModifyBannedRights))
			})
		},

		"chats.getFilterKeywords": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatID{}, func(data interface{}) *helper.ResultJSON {
				return s.GetFilterKeywords(ctx, md, data.(*TChatID))
			})
		},

		"chats.setFilterKeywords": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatIDSetFilterKeywords{}, func(data interface{}) *helper.ResultJSON {
				return s.SetFilterKeywords(ctx, md, data.(*TChatIDSetFilterKeywords))
			})
		},

		"chats.disableInviteLink": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatID{}, func(data interface{}) *helper.ResultJSON {
				return s.DisableInviteLink(ctx, md, data.(*TChatID))
			})
		},

		"chats.kickWhoSendKeyword": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatID{}, func(data interface{}) *helper.ResultJSON {
				return s.KickWhoSendKeyword(ctx, md, data.(*TChatID))
			})
		},

		"chats.setNickname": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TChatIdSetNickname{}, func(data interface{}) *helper.ResultJSON {
				return s.SetNickname(ctx, md, data.(*TChatIdSetNickname))
			})
		},
	}
}
