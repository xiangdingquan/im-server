package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func sendInvoice(c *bm.Context) {
	req := new(botapi.SendInvoice2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendInvoice(c, token, req)
	})
}

func answerShippingQuery(c *bm.Context) {
	req := new(botapi.AnswerShippingQuery2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.AnswerShippingQuery(c, token, req)
	})
}

func answerPreCheckoutQuery(c *bm.Context) {
	req := new(botapi.AnswerPreCheckoutQuery2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.AnswerPreCheckoutQuery(c, token, req)
	})
}
