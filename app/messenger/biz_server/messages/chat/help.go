package chat

import (
	"open.chat/app/messenger/biz_server/messages/chat/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New()
	}
	return s
}
