package dao

import (
	"context"
	"encoding/base64"

	"open.chat/app/messenger/push/internal/dal/dataobject"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (d *Dao) RegisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, noMuted, appSandbox bool, secret []byte, otherUids []int32) error {
	do := &dataobject.DevicesDO{
		AuthKeyId:  authKeyId,
		UserId:     userId,
		TokenType:  int8(tokenType),
		Token:      token,
		NoMuted:    util.BoolToInt8(noMuted),
		AppSandbox: util.BoolToInt8(appSandbox),
		Secret:     base64.RawStdEncoding.EncodeToString(secret),
		OtherUids:  util.JoinInt32List(otherUids, ":"),
	}
	_, _, err := d.DevicesDAO.InsertOrUpdate(ctx, do)

	return err
}

func (d *Dao) UnregisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, otherUids []int32) error {
	_, err := d.DevicesDAO.UpdateState(ctx, 1, authKeyId, userId)
	return err
}

func (d *Dao) UpdateDeviceLockedPeriod(ctx context.Context, userId int32, authKeyId int64, period int32) error {
	_, err := d.DevicesDAO.UpdateLockedPeriod(ctx, period, authKeyId, userId)
	return err
}

func (d *Dao) GetToken(ctx context.Context, userId int32, authKeyId int64, tokenType int32) (string, error) {
	var token string

	if do, err := d.DevicesDAO.Select(ctx, authKeyId, userId, int8(tokenType)); err != nil {
		log.Errorf("db error - %v", err)
		return "", err
	} else if do == nil {
		log.Errorf("not find token - {userId: %d, keyId: %d}", userId, authKeyId)
	} else {
		token = do.Token
	}

	return token, nil
}
