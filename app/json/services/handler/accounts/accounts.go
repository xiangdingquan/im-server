package accounts

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/handler"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	account_facade "open.chat/app/service/biz_service/account/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	account_facade.AccountFacade
	authsessionpb.RPCSessionClient
	user_client.UserFacade
}

func New(s *svc.Service) {
	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		c   = new(cls)
		err error
	)

	c.AccountFacade, err = account_facade.NewAccountFacade("local")
	checkErr(err)
	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	c.RPCSessionClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)
	c.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.AppendServices(handler.RegisterAccounts(c))
}

func (s *cls) CheckPassword(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TCheckPassword) *helper.ResultJSON {
	if len(r.Password) != 32 {
		log.Errorf("accounts.checkPassword, invalid len")
		return &helper.ResultJSON{Code: 401, Msg: "invalid button type"}
	}

	password, err := s.GetPasswordById(ctx, md.UserId)
	if err != nil {
		log.Errorf("accounts.checkPassword, err: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "internal error"}
	}

	if password == "" {
		log.Errorf("accounts.checkPassword, password not set")
		return &helper.ResultJSON{Code: 403, Msg: "password not set"}
	}

	if r.Password != password {
		log.Errorf("accounts.checkPassword, invalid password")
		return &helper.ResultJSON{Code: 404, Msg: "invalid password"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}
