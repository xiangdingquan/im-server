package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"open.chat/app/interface/botway/internal/dao"
	"open.chat/app/service/dfs/facade"
	_ "open.chat/app/service/dfs/facade/dfs"
)

type Service struct {
	c        *Config
	sessions *botSessionManager
	dao      *dao.Dao
	dfs_facade.DfsFacade
}

func New() (s *Service) {
	var (
		ac  = &Config{}
		err error
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s = &Service{
		c:        ac,
		sessions: NewBotSessionManager(),
		dao:      dao.New(ac.WardenClient),
	}
	s.DfsFacade, err = dfs_facade.NewDfsFacade("dfs")
	if err != nil {
		panic(err)
	}

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
