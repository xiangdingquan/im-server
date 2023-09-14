package stickers

import (
	"open.chat/app/messenger/biz_server/stickers/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New()
	}
	return s
}
