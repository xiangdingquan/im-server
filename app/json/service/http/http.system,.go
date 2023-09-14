package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
)

// ServiceUser
type (
	TSystemSendCode struct {
		Type        consts.SmsCodeType `json:"type"` //短信类型 2注册账号
		PhoneNumber string             `json:"phone_number"`
	}

	ServiceSystem interface {
		SendSmsCode(context.Context, *TSystemSendCode) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterSystem(s ServiceSystem, rg *bm.RouterGroup) {
	rg2 := rg.Group("/system")
	//curl -i -H "Content-Type: application/json" -X POST -d '{"chatId":1073741888}' http://172.192.168.102:40101/json/system/sendCode
	rg2.POST("/sendCode", func(c *bm.Context) {
		helper.DoHttpJson(c, &TSystemSendCode{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.SendSmsCode(ctx, data.(*TSystemSendCode))
		})
	})
}
