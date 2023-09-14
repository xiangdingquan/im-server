package service

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/messenger/biz_server/messages/sticker/internal/core"
	"open.chat/app/messenger/biz_server/messages/sticker/internal/dao"
)

type Service struct {
	ac *paladin.Map
	*core.StickerCore
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:          ac,
		StickerCore: core.New(dao.New()),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.Dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	// s.Dao.Close()
}
