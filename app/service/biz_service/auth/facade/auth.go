package facade

import (
	"context"
	"fmt"
)

type AuthFacade interface {
	GetPlatform(ctx context.Context, authKeyId int64) (int32, error)
}

type Instance func() AuthFacade

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

func NewAuthFacade(name string) (inst AuthFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
