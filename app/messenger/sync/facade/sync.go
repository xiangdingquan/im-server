package sync_facade

import (
	"context"
	"fmt"

	"open.chat/mtproto"
)

type SyncFacade interface {
	SyncUpdatesMe(ctx context.Context, userId int32, authKeyId, sessionId int64, serverId string, updates *mtproto.Updates) error
	SyncUpdatesNotMe(ctx context.Context, userId int32, authKeyId int64, updates *mtproto.Updates) error
	PushUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) error
	PushBotUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) error
	PushRpcResult(ctx context.Context, authKeyId int64, serverId string, sessionId, clientReqMsgId int64, result []byte) error
	BroadcastUpdates(ctx context.Context, broadcastType int32, chatId int32, excludeIdList []int32, updates *mtproto.Updates) error
}

type Instance func() SyncFacade

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

func NewSyncFacade(name string) (inst SyncFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
