package dialog

import (
	"open.chat/app/messenger/biz_server/messages/dialog/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New()
	}
	return s
}
