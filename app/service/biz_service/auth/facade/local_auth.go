package facade

import (
	"context"
	"open.chat/app/service/biz_service/auth/internal/core"
	"open.chat/app/service/biz_service/auth/internal/dao"
)

type localAuthFacade struct {
	*core.AuthCore
}

func localAuthFacadeInstance() AuthFacade {
	return &localAuthFacade{
		AuthCore: core.New(dao.New()),
	}
}

func (c localAuthFacade) GetPlatform(ctx context.Context, authKeyId int64) (int32, error) {
	return c.AuthCore.GetPlatform(ctx, authKeyId)
}

func init() {
	Register("local", localAuthFacadeInstance)
}
