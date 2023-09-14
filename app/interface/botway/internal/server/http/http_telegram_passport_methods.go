package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func setPassportDataErrors(c *bm.Context) {
	req := new(botapi.SetPassportDataErrors2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetPassportDataErrors(c, token, req)
	})
}
