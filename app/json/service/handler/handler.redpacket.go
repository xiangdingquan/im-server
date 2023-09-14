package handler

import (
	"context"

	"open.chat/app/json/helper"
	"open.chat/pkg/grpc_util"
)

type (
	// TRedPacketCreate .
	TRedPacketCreate struct {
		ChatId     uint32  `json:"chatId"`      //群聊id
		Type       uint8   `json:"type"`        //1.单聊红包 2.拼手气红包 3.普通红包
		Title      string  `json:"title"`       //标题
		Price      float64 `json:"price"`       //单个红包金额
		TotalPrice float64 `json:"total_price"` //红包总金额
		Count      uint32  `json:"count"`       //红包数量
		Password   string  `json:"password"`    //钱包密码
	}

	// TRedPacketID .
	TRedPacketID struct {
		RedPacketID uint32 `json:"redPacketId"` //红包id
	}

	// TRedPacketRecordPage .
	TRedPacketRecordPage struct {
		Type  uint32 `json:"type"`  //1创建红包 2领取红包 1|2 全部
		Count uint32 `json:"count"` //每页数量
		Page  uint32 `json:"page"`  //获取第N页
	}

	TRedPacketStatisticsReq struct {
		Type uint32 `json:"type"` //1创建红包 2领取红包 1|2 全部
		Year uint32 `json:"year"` //年份，传0为查找所有
	}

	// ServiceRedPacket .
	ServiceRedPacket interface {
		Create(context.Context, *grpc_util.RpcMetadata, *TRedPacketCreate) *helper.ResultJSON
		Get(context.Context, *grpc_util.RpcMetadata, *TRedPacketID) *helper.ResultJSON
		Detail(context.Context, *grpc_util.RpcMetadata, *TRedPacketID) *helper.ResultJSON
		Record(context.Context, *grpc_util.RpcMetadata, *TRedPacketRecordPage) *helper.ResultJSON
		Statistics(context.Context, *grpc_util.RpcMetadata, *TRedPacketStatisticsReq) *helper.ResultJSON
	}
)

// RegisterDiscover .
func RegisterRedPacket(s ServiceRedPacket) map[string]helper.MethodHandler {
	return map[string]helper.MethodHandler{
		//创建红包
		"redpacket.create": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRedPacketCreate{}, func(data interface{}) *helper.ResultJSON {
				return s.Create(ctx, md, data.(*TRedPacketCreate))
			})
		},
		//抢红包
		"redpacket.get": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRedPacketID{}, func(data interface{}) *helper.ResultJSON {
				return s.Get(ctx, md, data.(*TRedPacketID))
			})
		},
		//红包详情
		"redpacket.detail": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRedPacketID{}, func(data interface{}) *helper.ResultJSON {
				return s.Detail(ctx, md, data.(*TRedPacketID))
			})
		},
		//红包记录
		"redpacket.record": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRedPacketRecordPage{}, func(data interface{}) *helper.ResultJSON {
				return s.Record(ctx, md, data.(*TRedPacketRecordPage))
			})
		},
		//红包统计数据
		"redpacket.statistics": func(ctx context.Context, md *grpc_util.RpcMetadata, request *helper.DataJSON) (*helper.ResultJSON, error) {
			return request.JsonCall(&TRedPacketStatisticsReq{}, func(data interface{}) *helper.ResultJSON {
				return s.Statistics(ctx, md, data.(*TRedPacketStatisticsReq))
			})
		},
	}
}
