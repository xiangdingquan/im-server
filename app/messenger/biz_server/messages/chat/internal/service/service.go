package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	user_client.UserFacade
	username_facade.UsernameFacade
	msg_facade.MsgFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	authsessionpb.RPCSessionClient
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

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.RPCSessionClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	sync_client.New()
	return s
}

func (s *Service) Close() {
}
