package account_facade

import (
	"context"
	"fmt"

	"open.chat/mtproto"
)

type AccountFacade interface {
	RecoverPassword(ctx context.Context, userId int32, code string) error
	RequestPasswordRecovery(ctx context.Context, userId int32) (*mtproto.Auth_PasswordRecovery, error)
	CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error)
	CheckPassword(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP) (bool, error)
	GetPassword(ctx context.Context, userId int32) (*mtproto.Account_Password, error)
	GetPasswordSetting(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP) (*mtproto.Account_PasswordSettings, error)
	UpdatePasswordSetting(ctx context.Context, userId int32, password *mtproto.InputCheckPasswordSRP, newSettings *mtproto.Account_PasswordInputSettings) error

	SetSettingValue(ctx context.Context, userId int32, key, value string) error
	GetSettingValueString(ctx context.Context, userId int32, key string, defaultValue string) (v string)
	GetSettingValueInt32(ctx context.Context, userId int32, key string, defaultValue int32) (v int32)
	GetSettingValueBool(ctx context.Context, userId int32, key string, defaultValue bool) (v bool)
}

type Instance func() AccountFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewAccountFacade(name string) (inst AccountFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
