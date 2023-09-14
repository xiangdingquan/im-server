package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/messenger/biz_server/messages/bot/internal/dao"
	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	*dao.Dao
	user_client.UserFacade
	username_facade.UsernameFacade
	msg_facade.MsgFacade
	message_facade.MessageFacade
	private_facade.PrivateFacade
	chat_facade.ChatFacade
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
			Wardenclient *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.Dao, err = dao.New(ac.Wardenclient)
	checkErr(err)

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	s.PrivateFacade, err = private_facade.NewPrivateFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)

	sync_client.New()
	return s
}

func (s *Service) Close() {
	s.Dao.Close()
}
