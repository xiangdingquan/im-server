package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	sync_client "open.chat/app/messenger/sync/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	user_client.UserFacade
	username_facade.UsernameFacade
	channel_facade.ChannelFacade
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
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)

	sync_client.New()

	return s
}

func (s *Service) Close() {

}
