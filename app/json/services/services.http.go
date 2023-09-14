package services

import (
	"context"
	"open.chat/app/json/services/http/blog"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"

	"open.chat/app/json/services/http/account"
	"open.chat/app/json/services/http/auth"
	"open.chat/app/json/services/http/chats"
	"open.chat/app/json/services/http/pusher"
	"open.chat/app/json/services/http/system"
	"open.chat/app/json/services/http/wallet"
)

func RegistRouter(s *svc.Service, rg *bm.RouterGroup) {
	rg.POST("/test", func(c *bm.Context) {
		helper.DoHttpJson(c, &map[string]interface{}{}, func(ctx context.Context, r interface{}) *helper.ResultJSON {
			data := r.(*map[string]interface{})
			return &helper.ResultJSON{
				Code: 200,
				Msg:  "success",
				Data: data,
			}
		})
	})
	account.New(s, rg)
	pusher.New(s, rg)
	wallet.New(s, rg)
	auth.New(s, rg)
	chats.New(s, rg)
	system.New(s, rg)
	blog.New(s, rg)
}
