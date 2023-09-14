package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/helper"
)

// ServiceUser
type (
	TPusherMessage struct {
		Uids []uint32 `json:"userIds"`
		Msg  string   `json:"message"`
	}

	ServiceMessages interface {
		Messages(context.Context, *TPusherMessage) *helper.ResultJSON
		Notification(context.Context, *TPusherMessage) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterPusher(s ServiceMessages, rg *bm.RouterGroup) {
	rg2 := rg.Group("/pusher")
	//curl -i -H "Content-Type: application/json" -X POST -d '{"userIds":[136817712],"message":"this is test"}' http://172.192.168.102:40101/json/pusher/sendMessage
	rg2.POST("/sendMessage", func(c *bm.Context) {
		helper.DoHttpJson(c, &TPusherMessage{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.Messages(ctx, data.(*TPusherMessage))
		})
	})
	rg2.POST("/notification", func(c *bm.Context) {
		helper.DoHttpJson(c, &TPusherMessage{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.Notification(ctx, data.(*TPusherMessage))
		})
	})
}
