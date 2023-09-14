package service

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/service/auth_session/internal/core"
	"open.chat/app/service/auth_session/internal/dao"
)

// Service service.
type Service struct {
	ac *paladin.Map
	*dao.Dao
	*core.AuthSessionCore
}

// New new a service and return.
func New() (s *Service) {
	var (
		ac  = new(paladin.TOML)
		err error
	)

	if err = paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	dao := dao.New()
	s = &Service{
		ac:              ac,
		Dao:             dao,
		AuthSessionCore: core.New(dao),
	}

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.Dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.Dao.Close()
}
