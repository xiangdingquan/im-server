package service

import (
	json_service "open.chat/app/json"
	jsvc "open.chat/app/json/service"
	user_client "open.chat/app/service/biz_service/user/client"
)

// Service .
type Service struct {
	user_client.UserFacade
	*jsvc.Service
}

// New .
func New() *Service {
	s := new(Service)
	s.UserFacade, _ = user_client.NewUserFacade("local")
	s.Service = json_service.New()
	return s
}

// Close .
func (s *Service) Close() {
}
