package account_facade

import (
	"context"
	"strconv"

	"open.chat/app/service/biz_service/account/internal/core"
	"open.chat/app/service/biz_service/account/internal/dao"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type localAccountFacade struct {
	*core.AccountCore
}

func localAccountFacadeInstance() AccountFacade {
	return &localAccountFacade{
		AccountCore: core.New(dao.New()),
	}
}

func (c *localAccountFacade) RecoverPassword(ctx context.Context, userId int32, code string) error {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return err
	}

	return passwordLogic.RecoverPassword(ctx, code)
}

func (c *localAccountFacade) RequestPasswordRecovery(ctx context.Context, userId int32) (*mtproto.Auth_PasswordRecovery, error) {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return nil, err
	}

	return passwordLogic.RequestPasswordRecovery()
}

func (c *localAccountFacade) CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error) {
	return c.AccountCore.CheckSessionPasswordNeeded(ctx, userId)
}

func (c *localAccountFacade) CheckPassword(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP) (bool, error) {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return false, err
	}

	return passwordLogic.CheckPassword(password), nil
}

func (c *localAccountFacade) GetPassword(ctx context.Context, userId int32) (*mtproto.Account_Password, error) {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return nil, err
	}

	return passwordLogic.GetPassword(), nil
}

func (c *localAccountFacade) GetPasswordSetting(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP) (*mtproto.Account_PasswordSettings, error) {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return nil, err
	}

	return passwordLogic.GetPasswordSetting(password)
}

func (c *localAccountFacade) UpdatePasswordSetting(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP, newSettings *mtproto.Account_PasswordInputSettings) error {
	passwordLogic, err := c.AccountCore.MakePasswordData(ctx, userId)
	if err != nil {
		log.Errorf("makePasswordData error: %v", err)
		return err
	}

	return passwordLogic.UpdatePasswordSetting(ctx, password, newSettings)
}

func (c *localAccountFacade) SetSettingValue(ctx context.Context, userId int32, key, value string) error {
	return c.AccountCore.SetSettingValue(ctx, userId, key, value)
}

func (c *localAccountFacade) GetSettingValueString(ctx context.Context, userId int32, key string, defaultValue string) (v string) {
	val, ok := c.AccountCore.GetSettingValue(ctx, userId, key)
	if ok {
		v = val
	} else {
		v = defaultValue
	}
	return
}

func (c *localAccountFacade) GetSettingValueInt32(ctx context.Context, userId int32, key string, defaultValue int32) (v int32) {
	val, ok := c.AccountCore.GetSettingValue(ctx, userId, key)
	if !ok {
		return defaultValue
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		log.Errorf("GetSettingValueInt32, error: %v", err)
		return defaultValue
	}

	return int32(i)
}

func (c *localAccountFacade) GetSettingValueBool(ctx context.Context, userId int32, key string, defaultValue bool) (v bool) {
	val, ok := c.AccountCore.GetSettingValue(ctx, userId, key)
	if !ok {
		return defaultValue
	}

	return val == "true"
}

func init() {
	Register("local", localAccountFacadeInstance)
}
