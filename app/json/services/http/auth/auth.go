package auth

import (
	"context"
	"math/rand"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/messenger/biz_server/auth"
	"open.chat/mtproto"

	"open.chat/app/json/service/http"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

type cls struct {
	mtproto.RPCAuthServer
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{
		RPCAuthServer: auth.New(),
	}
	http.RegisterAuth(service, rg)
}

func (s *cls) SendCode(ctx context.Context, r *http.TAuthSendCode) *helper.ResultJSON {
	authKeyId := rand.Int63()
	ctx, err := helper.DefaultMetadata(ctx, 0, authKeyId)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}

	t := &mtproto.TLAuthSendCode{
		Constructor: mtproto.CRC32_auth_sendCode_86aef0ec,
		PhoneNumber: r.Phone,
	}

	reply, err := s.RPCAuthServer.AuthSendCode(ctx, t)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: err.Error()}
	}

	data := &struct {
		AuthKeyId     int64  `json:"authKeyId"`
		PhoneCodeHash string `json:"phone_code_hash"`
	}{
		AuthKeyId:     authKeyId,
		PhoneCodeHash: reply.PhoneCodeHash,
	}

	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}

func (s *cls) SignUp(ctx context.Context, r *http.TAuthSignUp) *helper.ResultJSON {
	ctx, err := helper.DefaultMetadata(ctx, 0, r.AuthKeyId)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}

	t := &mtproto.TLAuthSignUp{
		PhoneNumber:   r.PhoneNumber,
		PhoneCodeHash: r.PhoneCodeHash,
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		PhoneCode:     r.PhoneCode,
	}

	reply, err := s.RPCAuthServer.AuthSignUp(ctx, t)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: err.Error()}
	}

	data := &struct {
		UId       uint32 `json:"userId"`
		Username  string `json:"username"`
		Phone     string `json:"phone"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{
		UId:       uint32(reply.User.Id),
		Username:  reply.User.Username.GetValue(),
		Phone:     reply.User.Phone.GetValue(),
		FirstName: reply.User.FirstName.GetValue(),
		LastName:  reply.User.LastName.GetValue(),
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}

func (s *cls) LogoutUsers(ctx context.Context, r *http.TAuthLogoutUsers) *helper.ResultJSON {
	if len(r.UserIds) == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "please post user list"}
	}
	uids := make([]uint32, 0)
	for _, uid := range r.UserIds {
		ctx, err := helper.DefaultMetadata(context.TODO(), uid, 0)
		if err != nil {
			continue
		}
		s.RPCAuthServer.AuthResetAuthorizations(ctx, &mtproto.TLAuthResetAuthorizations{})
		s.RPCAuthServer.AuthLogOut(ctx, &mtproto.TLAuthLogOut{})
		uids = append(uids, uid)
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: uids}
}
