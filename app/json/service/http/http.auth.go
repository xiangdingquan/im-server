package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/helper"
)

// ServiceUser
type (
	TAuthSendCode struct {
		Phone string `json:"phone_number"`
	}

	TAuthSignUp struct {
		AuthKeyId     int64  `json:"authKeyId"`            //授权id
		PhoneNumber   string `json:"phone_number"`         //手机号
		PhoneCodeHash string `json:"phone_code_hash"`      //发送的验证码标识
		FirstName     string `json:"first_name,omitempty"` //名字
		LastName      string `json:"last_name"`            //姓氏
		PhoneCode     string `json:"phone_code"`           //收到的验证码
	}

	TAuthLogoutUsers struct {
		UserIds []uint32 `json:"userIds"`
	}

	TAuthLogOut struct {
	}

	ServiceAuth interface {
		SendCode(context.Context, *TAuthSendCode) *helper.ResultJSON
		SignUp(context.Context, *TAuthSignUp) *helper.ResultJSON
		LogoutUsers(context.Context, *TAuthLogoutUsers) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterAuth(s ServiceAuth, rg *bm.RouterGroup) {
	rg2 := rg.Group("/auth")
	//rg2.POST("/sendcode", func(c *bm.Context) {
	//	helper.DoHttpJson(c, &TAuthSendCode{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
	//		return s.SendCode(ctx, data.(*TAuthSendCode))
	//	})
	//})

	//rg2.POST("/signUp", func(c *bm.Context) {
	//	helper.DoHttpJson(c, &TAuthSignUp{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
	//		return s.SignUp(ctx, data.(*TAuthSignUp))
	//	})
	//})

	rg2.POST("/logoutUsers", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAuthLogoutUsers{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.LogoutUsers(ctx, data.(*TAuthLogoutUsers))
		})
	})
}
