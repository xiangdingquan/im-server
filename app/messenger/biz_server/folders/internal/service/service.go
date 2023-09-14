package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	sync_client "open.chat/app/messenger/sync/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	media_client "open.chat/app/service/media/client"
)

type Service struct {
	private_facade.PrivateFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	user_client.UserFacade
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

	s.PrivateFacade, err = private_facade.NewPrivateFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	sync_client.New()
	media_client.New()
	return s
}

func (s *Service) Close() {

}
