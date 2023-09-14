package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	sync_client "open.chat/app/messenger/sync/client"

	"open.chat/app/messenger/webpage/internal/dao"
)

type Service struct {
	ac *paladin.Map
	*dao.Dao
}

func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	dao := dao.New()
	s = &Service{
		ac:  ac,
		Dao: dao,
	}

	sync_client.New()
	return s
}
