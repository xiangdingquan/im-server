package service

import (
	secretchat_core "open.chat/app/messenger/biz_server/messages/secretchat/core"
	secretchat_dao "open.chat/app/messenger/biz_server/messages/secretchat/dao"
	channel_core "open.chat/app/service/biz_service/channel/core"
	channel_dao "open.chat/app/service/biz_service/channel/dao"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	updates_facade "open.chat/app/service/biz_service/updates/facade"
	user_client "open.chat/app/service/biz_service/user/client"
)

type Service struct {
	user_client.UserFacade
	updates_facade.UpdatesFacade
	chat_facade.ChatFacade

	*channel_core.ChannelCore
	*secretchat_core.SecretChatCore
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		err error
		s   = new(Service)
	)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	s.UpdatesFacade, err = updates_facade.NewUpdatesFacade("local")
	checkErr(err)

	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)

	s.ChannelCore = channel_core.New(channel_dao.New())
	s.SecretChatCore = secretchat_core.New(secretchat_dao.New())

	return s
}

func (s *Service) Close() {

}
