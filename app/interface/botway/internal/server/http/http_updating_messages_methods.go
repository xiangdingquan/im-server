package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func editMessageText(c *bm.Context) {
	req := new(botapi.EditMessageText2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.EditMessageText(c, token, req)
	})
}

func editMessageCaption(c *bm.Context) {
	req := new(botapi.EditMessageCaption2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.EditMessageCaption(c, token, req)
	})
}

func editMessageMedia(c *bm.Context) {
	req := new(botapi.EditMessageMedia2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.EditMessageMedia(c, token, req)
	})
}

func editMessageReplyMarkup(c *bm.Context) {
	req := new(botapi.EditMessageReplyMarkup2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.EditMessageReplyMarkup(c, token, req)
	})
}

func stopPoll(c *bm.Context) {
	req := new(botapi.StopPoll2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.StopPoll(c, token, req)
	})
}

func deleteMessage(c *bm.Context) {
	req := new(botapi.DeleteMessage2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.DeleteMessage(c, token, req)
	})
}
