package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
	"log"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
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
)

var rc struct {
	server *warden.ServerConfig
}

func init() {

	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
}

func main() {

	conn, err := grpc.Dial(rc.server.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mtproto.NewRPCBotsClient(conn)

	//r := TRedPacketRecordPage{
	//	Type:  1,
	//	Count: 20,
	//	Page:  1,
	//}
	//method := "redpacket.record"
	r := TRedPacketCreate{
		ChatId:     1073741842,
		Count:      2,
		Password:   "c33367701511b4f6020ec61ded352059",
		Price:      25.0,
		Title:      "自动创建",
		TotalPrice: 50.0,
		Type:       2,
	}
	method := "redpacket.create"
	//r := TRedPacketID{
	//	RedPacketID: 2,
	//}
	//method := "redpacket.get"
	strJSON, err := json.Marshal(r)
	if err != nil {
		strJSON = []byte(`{"code":"-1","msg":"error","data":{}}`)
	}

	tlData := mtproto.MakeTLDataJSON(&mtproto.DataJSON{
		Data:        string(strJSON),
		Constructor: 2104790276,
	})

	ctx, err := grpc_util.RpcMetadataToOutgoing(context.TODO(), &grpc_util.RpcMetadata{
		ServerId:    "session001",
		ClientAddr:  "39.144.138.189",
		AuthId:      -2591078482711196039,
		SessionId:   2131127024233412874,
		ReceiveTime: time.Now().Unix(),
		UserId:      136817692,
		ClientMsgId: 7179096495420328340,
		Layer:       1202,
	})
	if err != nil {
		log.Println("======")
		log.Fatal(err)
	}
	reply, err := client.BotsSendCustomRequest(ctx, &mtproto.TLBotsSendCustomRequest{
		Constructor:  -1440257555,
		CustomMethod: method,
		Params:       tlData.To_DataJSON(),
	})
	if err != nil {
		log.Println("1111111111")
		log.Fatal(err)
	}
	fmt.Println(reply)
}
