package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/messenger/biz_server/langpack/internal/core"
	"open.chat/app/messenger/biz_server/langpack/internal/dao"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Service struct {
	*core.LangPackCore
}

func New() *Service {
	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		s = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	s.LangPackCore = core.New(dao.New())

	return s
}

func (s *Service) Close() error {
	return nil
}
