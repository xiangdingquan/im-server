package push_facade

import (
	"context"

	"open.chat/app/messenger/push/internal/dao"
)

type localPushFacade struct {
	*dao.Dao
}

func localPushFacadeInstance() PushFacade {
	return &localPushFacade{
		Dao: dao.New(),
	}
}

func (c *localPushFacade) RegisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, noMuted, appSandbox bool, secret []byte, otherUids []int32) error {
	return c.Dao.RegisterDevice(ctx, userId, authKeyId, tokenType, token, noMuted, appSandbox, secret, otherUids)
}

func (c *localPushFacade) UnregisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, otherUids []int32) error {
	return c.Dao.UnregisterDevice(ctx, userId, authKeyId, tokenType, token, otherUids)
}

func (c *localPushFacade) UpdateDeviceLockedPeriod(ctx context.Context, userId int32, authKeyId int64, period int32) error {
	return c.Dao.UpdateDeviceLockedPeriod(ctx, userId, authKeyId, period)
}

func (c *localPushFacade) GetToken(ctx context.Context, userId int32, authKeyId int64, tokenType int32) (string, error) {
	return c.Dao.GetToken(ctx, userId, authKeyId, tokenType)
}

func init() {
	Register("local", localPushFacadeInstance)
}
