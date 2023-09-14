package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/messenger/biz_server/messages/secretchat/core"
	"open.chat/app/messenger/biz_server/messages/secretchat/dao"
	sync_client "open.chat/app/messenger/sync/client"
	report_facade "open.chat/app/service/biz_service/report/facade"
	user_client "open.chat/app/service/biz_service/user/client"
)

type Service struct {
	*core.SecretChatCore
	user_client.UserFacade
	report_facade.ReportFacade
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			c *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.ReportFacade, err = report_facade.NewReportFacade("local")
	checkErr(err)

	s.SecretChatCore = core.New(dao.New())

	sync_client.New()
	return s
}

func (s *Service) Close() {

}
