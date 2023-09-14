package banned_facade

import (
	"context"

	"open.chat/app/service/biz_service/banned/internal/core"
	"open.chat/app/service/biz_service/banned/internal/dao"
)

type localBannedFacade struct {
	*core.BannedCore
}

func New() BannedFacade {
	return &localBannedFacade{
		BannedCore: core.New(dao.New()),
	}
}

func (c *localBannedFacade) CheckPhoneNumberBanned(ctx context.Context, phoneNumber string) bool {
	return c.BannedCore.CheckPhoneNumberBanned(ctx, phoneNumber)
}

func (c *localBannedFacade) GetBannedByPhoneList(ctx context.Context, phoneList []string) map[string]bool {
	return c.BannedCore.GetBannedByPhoneList(ctx, phoneList)
}

func (c *localBannedFacade) Ban(ctx context.Context, phoneNumber string, expires int32, reason string) bool {
	return c.BannedCore.Ban(ctx, phoneNumber, expires, reason)
}

func (c *localBannedFacade) UnBan(ctx context.Context, phoneNumber string) bool {
	return c.BannedCore.UnBan(ctx, phoneNumber)
}

func init() {
	Register("local", New)
}
