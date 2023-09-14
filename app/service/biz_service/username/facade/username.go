package username_facade

import (
	"context"
	"fmt"

	"open.chat/app/service/biz_service/username/internal/dal/dataobject"
)

type UsernameFacade interface {
	GetAccountUsername(ctx context.Context, userId int32) (string, error)
	CheckAccountUsername(ctx context.Context, userId int32, username string) (int, error)
	GetChannelUsername(ctx context.Context, channelId int32) (string, error)
	CheckChannelUsername(ctx context.Context, channelId int32, username string) (int, error)
	UpdateUsernameByPeer(ctx context.Context, peerType, peerId int32, username string) (bool, error)

	CheckUsername(ctx context.Context, username string) (int, error)
	UpdateUsername(ctx context.Context, peerType, peerId int32, username string) (bool, error)
	DeleteUsername(ctx context.Context, username string) (bool, error)
	ResolveUsername(ctx context.Context, username string) (int32, int32, error)

	GetListByUsernameList(ctx context.Context, names []string) (map[string]*dataobject.UsernameDO, error)

	DeleteUsernameByPeer(ctx context.Context, peerType, peerId int32) error
}

type Instance func() UsernameFacade

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

func NewUsernameFacade(name string) (inst UsernameFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
