package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/helper"
)

// ServiceUser
type (
	TChatsDisband struct {
		ChatID int64 `json:"chatId"` //群组标识 <0普通群 >0超级群
	}

	TChatsDeleteDialog struct {
		UserId    uint32 `json:"userId"`    //用户id
		TargetUid uint32 `json:"targetUId"` //对方uid
	}

	TChatAndUser struct {
		ChatID  int64    `json:"chatId"`  //群组标识 <0普通群 >0超级群
		UserIds []uint32 `json:"userIds"` //用户列表
	}

	ServiceChats interface {
		Disband(context.Context, *TChatsDisband) *helper.ResultJSON
		DeleteDialog(context.Context, *TChatsDeleteDialog) *helper.ResultJSON
		Join(context.Context, *TChatAndUser) *helper.ResultJSON
		Leave(context.Context, *TChatAndUser) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterChats(s ServiceChats, rg *bm.RouterGroup) {
	rg2 := rg.Group("/chats")
	//curl -i -H "Content-Type: application/json" -X POST -d '{"chatId":1073741888}' http://172.192.168.102:40101/json/chats/disband
	rg2.POST("/disband", func(c *bm.Context) {
		helper.DoHttpJson(c, &TChatsDisband{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.Disband(ctx, data.(*TChatsDisband))
		})
	})
	rg2.POST("/deleteDialog", func(c *bm.Context) {
		helper.DoHttpJson(c, &TChatsDeleteDialog{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.DeleteDialog(ctx, data.(*TChatsDeleteDialog))
		})
	})
	rg2.POST("/join", func(c *bm.Context) {
		helper.DoHttpJson(c, &TChatAndUser{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.Join(ctx, data.(*TChatAndUser))
		})
	})
	rg2.POST("/leave", func(c *bm.Context) {
		helper.DoHttpJson(c, &TChatAndUser{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.Leave(ctx, data.(*TChatAndUser))
		})
	})
}
