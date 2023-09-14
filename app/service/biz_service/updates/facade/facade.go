package updates_facade

import (
	"context"
	"fmt"

	"open.chat/mtproto"
)

type UpdatesFacade interface {
	GetState(ctx context.Context, authKeyId int64, userId int32) (*mtproto.Updates_State, error)
	GetDifference(ctx context.Context, authKeyId int64, userId, pts, limit int32) (*mtproto.Updates_Difference, error)
}

type Instance func() UpdatesFacade

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

func NewUpdatesFacade(name string) (inst UpdatesFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
