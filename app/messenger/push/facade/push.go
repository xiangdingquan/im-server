package push_facade

import (
	"context"
	"fmt"
)

type PushFacade interface {
	RegisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, noMuted, appSandbox bool, secret []byte, otherUids []int32) error
	UnregisterDevice(ctx context.Context, userId int32, authKeyId int64, tokenType int, token string, otherUids []int32) error
	UpdateDeviceLockedPeriod(ctx context.Context, userId int32, authKeyId int64, period int32) error
	GetToken(ctx context.Context, userId int32, authKeyId int64, tokenType int32) (string, error)
}

type Instance func() PushFacade

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

func NewPushFacade(name string) (inst PushFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
