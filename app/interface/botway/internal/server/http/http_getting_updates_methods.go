package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func getUpdates(c *bm.Context) {
	req := new(botapi.GetUpdates2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetUpdates(c, token, req)
	})
}

func setWebhook(c *bm.Context) {
	req := new(botapi.SetWebhook2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetWebhook(c, token, req)
	})
}

func deleteWebhook(c *bm.Context) {
	req := new(botapi.DeleteWebhook2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.DeleteWebhook(c, token, req)
	})
}

func getWebhookInfo(c *bm.Context) {
	req := new(botapi.GetWebhookInfo2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetWebhookInfo(c, token, req)
	})
}
