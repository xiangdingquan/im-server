package http

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/json/helper"
)

type (
	TAccountSignUp struct {
		Inviter     uint32 `json:"userId,omitempty"`
		PhoneNumber string `json:"phoneNumber"`
		PhoneCode   string `json:"phoneCode,omitempty"`
		UserName    string `json:"userName"`
		Password    string `json:"password"`
		NickName    string `json:"nickName"`
	}

	TAccountUid struct {
		Uid uint32 `json:"userId"` //用户id
	}

	TAccountBan struct {
		Uid     uint32 `json:"userId"`            //用户id
		Expires uint32 `json:"expires,omitempty"` //封号时长
		Reason  string `json:"reason,omitempty"`  //封号原因
	}

	TAccountUpdateUsername struct {
		Uid      uint32 `json:"userId"`
		Username string `json:"username,omitempty"` // 昵称
	}

	TAccountUpdateProfile struct {
		Uid       uint32 `json:"userId"`
		FirstName string `json:"first_name,omitempty"` //名
		LastName  string `json:"last_name,omitempty"`  //姓
		About     string `json:"about,omitempty"`      //简介
	}

	TAccountUpdatePhoto struct {
		Uid   uint32 `json:"userId"`
		Photo string `json:"photo"`
	}

	TAccountGetPhoto struct {
		Uid      uint32 `json:"userId"` //用户id
		BigPhoto bool   `json:"isBig"`  //大图
	}

	TCreateVirtual struct {
		NickName string `json:"nickName"`
	}

	// ServiceAccount
	ServiceAccount interface {
		SignUp(context.Context, *TAccountSignUp) *helper.ResultJSON
		UpdateUsername(context.Context, *TAccountUpdateUsername) *helper.ResultJSON
		UpdateProfile(context.Context, *TAccountUpdateProfile) *helper.ResultJSON
		UpdatePhoto(context.Context, *TAccountUpdatePhoto) *helper.ResultJSON
		ToggleBan(context.Context, *TAccountBan) *helper.ResultJSON
		GetPhoto(context.Context, *TAccountGetPhoto) *[]byte
		CreateVirtual(context.Context, *TCreateVirtual) *helper.ResultJSON
	}
)

// RegisterUser .
func RegisterAccount(s ServiceAccount, rg *bm.RouterGroup) {
	rg2 := rg.Group("/account")

	//curl -i -H "Content-Type: application/json" -X POST -d '{"uid":136817694,"authId":"1231232123","data":{"amount":1.2}}' http://172.192.168.102:40101/json/account/toggleBan
	rg2.POST("/signUp", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAccountSignUp{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.SignUp(ctx, data.(*TAccountSignUp))
		})
	})

	rg2.POST("/updateUsername", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAccountUpdateUsername{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.UpdateUsername(ctx, data.(*TAccountUpdateUsername))
		})
	})

	rg2.POST("/updateProfile", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAccountUpdateProfile{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.UpdateProfile(ctx, data.(*TAccountUpdateProfile))
		})
	})

	rg2.POST("/updatePhoto", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAccountUpdatePhoto{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.UpdatePhoto(ctx, data.(*TAccountUpdatePhoto))
		})
	})

	rg2.POST("/getPhoto", func(c *bm.Context) {
		helper.DoHttpDownload(c, &TAccountGetPhoto{}, func(ctx context.Context, data interface{}) *[]byte {
			return s.GetPhoto(ctx, data.(*TAccountGetPhoto))
		})
	})

	rg2.POST("/toggleBan", func(c *bm.Context) {
		helper.DoHttpJson(c, &TAccountBan{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.ToggleBan(ctx, data.(*TAccountBan))
		})
	})

	rg2.POST("/createVirtual", func(c *bm.Context) {
		helper.DoHttpJson(c, &TCreateVirtual{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.CreateVirtual(ctx, data.(*TCreateVirtual))
		})
	})
}
