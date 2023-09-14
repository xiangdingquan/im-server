package wallet

import (
	"open.chat/app/messenger/biz_server/wallet/internal/service"
)

var s *service.Service

func New() *service.Service {
	if s == nil {
		s = service.New()
	}
	return s
}
