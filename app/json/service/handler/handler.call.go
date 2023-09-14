package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	// TAvCallToken .
	TAvCallToken struct {
		ChannelName string `json:"channelName"` //通道名
		UserID      uint32 `json:"uid"`         //用户id
	}

	// TAvCallCreate .
	TAvCallCreate struct {
		ChannelName string   `json:"channelName"` //通话标识
		To          []uint32 `json:"to"`          //接收者 userids
		ChatID      uint32   `json:"chatId"`      //归属会话 可为0
		IsMeetingAV bool     `json:"isMeetingAV"` //当前会话是否为多人
		IsVideo     bool     `json:"isVideo"`     //是否开启视频
	}

	// TAvOnInvite .
	TAvOnInvite struct {
		CallID      uint32   `json:"callId"`      //记录标识
		ChannelName string   `json:"channelName"` //通话标识
		From        uint32   `json:"from"`        //发起人
		To          []uint32 `json:"to"`          //邀请的成员
		ChatID      uint32   `json:"chatId"`      //归属会话 可为0
		IsMeetingAV bool     `json:"isMeetingAV"` //当前会话是否为多人
		IsVideo     bool     `json:"isVideo"`     //是否开启视频
		CreateAt    uint32   `json:"createAt"`    //创建成功的时间
	}

	// TAvChannelName .
	TAvChannelName struct {
		ChannelName string `json:"channelName"` //通话标识
	}

	// TAvCallID .
	TAvCallID struct {
		CallID uint32 `json:"callId"` //通话标识
	}

	// TAvCallTimeOut .
	TAvCallTimeOut struct {
		Timeout uint32 `json:"timeOut"` //超时间 秒
	}

	// TAvRecordPage .
	TAvRecordPage struct {
		Type  int8   `json:"type"`  //0:单聊呼出和接听 1单聊呼出 2单聊接听 3会议
		Count uint32 `json:"count"` //每页数量
		Page  uint32 `json:"page"`  //获取第N页
	}

	// ServiceCall 视频语音通话
	ServiceCall interface {
		CreateToken(context.Context, *grpc_util.RpcMetadata, *TAvCallToken) *helper.ResultJSON    //创建音视频授权码
		Create(context.Context, *grpc_util.RpcMetadata, *TAvCallCreate) *helper.ResultJSON        //创建音视频通话
		GetInfo(context.Context, *grpc_util.RpcMetadata, *TAvCallID) *helper.ResultJSON           //通话信息
		Cancel(context.Context, *grpc_util.RpcMetadata, *TAvCallID) *helper.ResultJSON            //取消通话
		Start(context.Context, *grpc_util.RpcMetadata, *TAvCallID) *helper.ResultJSON             //开始通话
		Stop(context.Context, *grpc_util.RpcMetadata, *TAvCallID) *helper.ResultJSON              //结束通话
		QueryOffline(context.Context, *grpc_util.RpcMetadata, *TAvCallTimeOut) *helper.ResultJSON //查询离线通话
		QueryRecord(context.Context, *grpc_util.RpcMetadata, *TAvRecordPage) *helper.ResultJSON   //查询通话记录

		AckInvite(context.Context, *grpc_util.RpcMetadata, *TAvCallID) *helper.ResultJSON //确认邀请
	}
)

// RegisterCall 注册通话服务
func RegisterCall(s ServiceCall) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//创建token
		"call.token.create": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallToken{}, func(data interface{}) *helper.ResultJSON {
				return s.CreateToken(ctx, md, data.(*TAvCallToken))
			})
		},
		//创建
		"call.create": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallCreate{}, func(data interface{}) *helper.ResultJSON {
				return s.Create(ctx, md, data.(*TAvCallCreate))
			})
		},
		//详情
		"call.getInfo": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallID{}, func(data interface{}) *helper.ResultJSON {
				return s.GetInfo(ctx, md, data.(*TAvCallID))
			})
		},
		//取消
		"call.cancel": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallID{}, func(data interface{}) *helper.ResultJSON {
				return s.Cancel(ctx, md, data.(*TAvCallID))
			})
		},
		//开始
		"call.start": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallID{}, func(data interface{}) *helper.ResultJSON {
				return s.Start(ctx, md, data.(*TAvCallID))
			})
		},
		//结束
		"call.stop": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallID{}, func(data interface{}) *helper.ResultJSON {
				return s.Stop(ctx, md, data.(*TAvCallID))
			})
		},
		//查询离线
		"call.queryOffline": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallTimeOut{}, func(data interface{}) *helper.ResultJSON {
				return s.QueryOffline(ctx, md, data.(*TAvCallTimeOut))
			})
		},
		//查询通话记录
		"call.queryRecord": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvRecordPage{}, func(data interface{}) *helper.ResultJSON {
				return s.QueryRecord(ctx, md, data.(*TAvRecordPage))
			})
		},
		//确认被邀请
		"call.ack.invite": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TAvCallID{}, func(data interface{}) *helper.ResultJSON {
				return s.AckInvite(ctx, md, data.(*TAvCallID))
			})
		},
	}
}
