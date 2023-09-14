package upload

import (
	"open.chat/app/messenger/biz_server/upload/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New("/opt/nbfs")
	}
	return s
}
