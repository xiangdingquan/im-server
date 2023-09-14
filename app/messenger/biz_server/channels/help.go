package channels

import (
	"open.chat/app/messenger/biz_server/channels/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New()
	}
	return s
}
