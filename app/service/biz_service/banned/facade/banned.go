package banned_facade

import (
	"context"
	"fmt"
)

type BannedFacade interface {
	CheckPhoneNumberBanned(ctx context.Context, phoneNumber string) bool
	GetBannedByPhoneList(ctx context.Context, phoneList []string) map[string]bool

	Ban(ctx context.Context, phoneNumber string, expires int32, reason string) bool
	UnBan(ctx context.Context, phoneNumber string) bool
}

type Instance func() BannedFacade

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

func NewBannedFacade(name string) (inst BannedFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
