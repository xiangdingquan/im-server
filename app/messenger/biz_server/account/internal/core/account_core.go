package core

import (
	"context"

	"open.chat/app/messenger/biz_server/account/internal/dao"
	"open.chat/pkg/log"
)

type AccountCore struct {
	*dao.Dao
}

func New(o *dao.Dao) *AccountCore {
	if o == nil {
		o = dao.New()
	}
	return &AccountCore{o}
}

func (c *AccountCore) IsBannedIp(ctx context.Context, ip string) bool {
	if len(ip) == 0 {
		return false
	}
	ipdo, err := c.BannedIpDAO.Select(ctx, ip)
	if ipdo == nil || err != nil {
		return false
	}
	log.Warnf("IsBannedIp: %s", ip)
	return true
}
