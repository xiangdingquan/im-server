package system

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	"open.chat/model"
	"strings"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/http"
	"open.chat/pkg/phonenumber"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	AuthSessionRpcClient authsessionpb.RPCSessionClient
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{}

	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		err error
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	service.AuthSessionRpcClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	http.RegisterSystem(service, rg)
}

func (s *cls) SendSmsCode(ctx context.Context, r *http.TSystemSendCode) *helper.ResultJSON {
	phoneNumber := strings.Trim(r.PhoneNumber, " ")
	if phoneNumber == "" {
		return &helper.ResultJSON{Code: -1, Msg: "please post phone number"}
	}

	if r.Type != consts.SmsCodeType_RegistAccount {
		return &helper.ResultJSON{Code: -2, Msg: "error in type"}
	}

	pNumber, err := phonenumber.MakePhoneNumberHelper(r.PhoneNumber, "")
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "phone number invalid"}
	}

	phone := pNumber.GetNormalizeDigits()
	code, err := helper.SendVerifyCode(ctx, r.Type, phone, model.GetLangType(ctx, s.AuthSessionRpcClient))
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "send verify code fail"}
	}

	if code == "" {
		return &helper.ResultJSON{Code: -5, Msg: "send verify code fail"}
	}

	return &helper.ResultJSON{Code: 0, Msg: "success"}
}
