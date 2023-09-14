package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func answerInlineQuery(c *bm.Context) {
	req := new(botapi.AnswerInlineQuery2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.AnswerInlineQuery(c, token, req)
	})
}
