package json_service

import (
	svc "open.chat/app/json/service"
	"open.chat/app/json/services"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var (
	srvs *svc.Service
)

// New .
func New() *svc.Service {
	if srvs == nil {
		srvs = svc.New()
		services.RegistHandler(srvs)
	}
	return srvs
}

func RegistRouter(rg *bm.RouterGroup) {
	if srvs == nil {
		New()
	}
	services.RegistRouter(srvs, rg)
}
