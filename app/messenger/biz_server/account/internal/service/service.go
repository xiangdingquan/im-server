package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/messenger/biz_server/account/internal/core"
	push_facade "open.chat/app/messenger/push/facade"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	_ "open.chat/app/service/biz_service/account/facade"
	account_facade "open.chat/app/service/biz_service/account/facade"
	banned_facade "open.chat/app/service/biz_service/banned/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	report_facade "open.chat/app/service/biz_service/report/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	*core.AccountCore
	authsessionpb.RPCSessionClient
	user_client.UserFacade
	chat_facade.ChatFacade
	username_facade.UsernameFacade
	channel_facade.ChannelFacade
	account_facade.AccountFacade
	report_facade.ReportFacade
	push_facade.PushFacade
	banned_facade.BannedFacade
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	s.AccountCore = core.New(nil)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	s.RPCSessionClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.AccountFacade, err = account_facade.NewAccountFacade("local")
	checkErr(err)
	s.ReportFacade, err = report_facade.NewReportFacade("local")
	checkErr(err)
	s.PushFacade, err = push_facade.NewPushFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.BannedFacade, err = banned_facade.NewBannedFacade("local")
	checkErr(err)
	sync_client.New()
	return s
}

func (s *Service) Close() {
}
