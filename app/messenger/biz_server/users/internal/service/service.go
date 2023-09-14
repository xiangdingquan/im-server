package service

import (
	banned_facade "open.chat/app/service/biz_service/banned/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	user_client "open.chat/app/service/biz_service/user/client"
)

type Service struct {
	user_client.UserFacade
	private_facade.PrivateFacade
	chat_facade.ChatFacade
	banned_facade.BannedFacade
}

func New() *Service {
	s := new(Service)
	s.UserFacade, _ = user_client.NewUserFacade("local")
	s.PrivateFacade, _ = private_facade.NewPrivateFacade("local")
	s.ChatFacade, _ = chat_facade.NewChatFacade("local")
	s.BannedFacade, _ = banned_facade.NewBannedFacade("local")
	return s
}

func (s *Service) Close() {
}
