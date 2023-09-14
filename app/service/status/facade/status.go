package status_facade

import (
	"context"
	"fmt"
)

const (
	ONLINE_TIMEOUT = 60
)

type StatusFacade interface {
	AddOnline(ctx context.Context, userId int32, authKeyId int64, serverId string) error
	ExpireOnline(ctx context.Context, userId int32, authKeyId int64) (bool, error)
	DelOnline(ctx context.Context, userId int32, authKeyId int64) error

	GetOnlineListByKeyIdList(ctx context.Context, authKeyIds []int64) (res []string, err error)
	GetOnlineByKeyId(ctx context.Context, authKeyId int64) (res string, err error)
	GetOnlineListExcludeKeyId(ctx context.Context, userId int32, authKeyId int64) (res map[int64]string, err error)
	GetOnlineListByUser(ctx context.Context, userId int32) (res map[int64]string, err error)
	GetOnlineMapByUserList(ctx context.Context, userIdList []int32) (ress map[int64]string, onUserList []int32, err error)
}

type Instance func() StatusFacade

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("status_client: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("status_client: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewStatusClient(adapterName string) (adapter StatusFacade, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("status_client: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
