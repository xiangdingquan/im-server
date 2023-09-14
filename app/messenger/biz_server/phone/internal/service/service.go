package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/messenger/biz_server/phone/internal/core"
	"open.chat/app/messenger/biz_server/phone/internal/dao"
	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	user_client "open.chat/app/service/biz_service/user/client"
)

var fingerprint uint64 = 12240908862933197005

type Service struct {
	*core.PhoneCallCore
	user_client.UserFacade
	msg_facade.MsgFacade
	RelayIp string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			RelayIp string
			c       *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.RelayIp = ac.RelayIp
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)

	s.PhoneCallCore = core.New(dao.New())

	sync_client.New()
	return s
}

func (s *Service) Close() {

}
