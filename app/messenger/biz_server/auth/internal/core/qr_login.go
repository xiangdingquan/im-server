package core

import (
	"context"

	"open.chat/app/messenger/biz_server/auth/internal/model"
)

func (c *AuthCore) GetQRCode(ctx context.Context, authKeyId int64) (*model.QRCodeTransaction, error) {
	return c.Dao.GetCacheQRLoginCode(ctx, authKeyId)
}

func (c *AuthCore) DeleteQRCode(ctx context.Context, authKeyId int64) error {
	return c.Dao.DeleteCacheQRLoginCode(ctx, authKeyId)
}

func (c *AuthCore) UpdateQRCode(ctx context.Context, authKeyId int64, values map[string]interface{}) error {
	return c.Dao.UpdateCacheQRLoginCode(ctx, authKeyId, values)
}

func (c *AuthCore) PutQRCode(ctx context.Context, authKeyId int64, code *model.QRCodeTransaction, tm int) error {
	return c.Dao.PutCacheQRLoginCode(ctx, authKeyId, code, tm)
}
