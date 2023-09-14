package core

import (
	"context"

	"open.chat/app/messenger/biz_server/auth/internal/dao"
	"open.chat/pkg/log"
)

type AuthCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *AuthCore {
	return &AuthCore{dao}
}

func (c *AuthCore) CheckApiIdAndHash(apiId int32, apiHash string) error {
	return nil
}

func (c *AuthCore) IsBannedIp(ctx context.Context, ip string) bool {
	if len(ip) == 0 {
		return true
	}
	ipdo, err := c.BannedIpDAO.Select(ctx, ip)
	if ipdo == nil || err != nil {
		return false
	}
	log.Warnf("IsBannedIp: %s", ip)
	return true
}

func (c *AuthCore) IsBindIp(ctx context.Context, userId int32, ip string) bool {
	if len(ip) == 0 {
		return true
	}
	ubdo, err := c.UserBindIpsDAO.Select(ctx, userId)
	if ubdo == nil || err != nil || len(ubdo.IpList) == 0 {
		return true
	}
	for _, i := range ubdo.IpList {
		if ip == i {
			return true
		}
	}
	log.Warnf("IsBindIp: %d : %s", userId, ip)
	return false
}

func (c *AuthCore) GetUserListByIp(ctx context.Context, ip string) ([]int32, error) {
	return c.AuthsDAO.SelectUserListByIp(ctx, ip)
}

func (c *AuthCore) IsPasswordFlood(ctx context.Context, userId int32) (bool, error) {
	return c.Redis.IsPasswordFlood(ctx, userId)
}

func (c *AuthCore) GetPasswordFloodCount(ctx context.Context, userId int32) (int32, error) {
	return c.Redis.GetPasswordFloodCount(ctx, userId)
}

func (c *AuthCore) CleanPasswordFloodCount(ctx context.Context, userId int32) error {
	return c.Redis.CleanPasswordFloodCount(ctx, userId)
}

func (c *AuthCore) IncPFCount(ctx context.Context, uid int32, floodExpireInSecond, countExpireInSeocnd int32, limit int32) (isFlood bool, err error) {
	return c.Redis.IncPFCount(ctx, uid, floodExpireInSecond, countExpireInSeocnd, limit)
}
