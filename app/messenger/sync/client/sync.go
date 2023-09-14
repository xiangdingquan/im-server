package sync_client

import (
	"context"

	sync_facade "open.chat/app/messenger/sync/facade"
	_ "open.chat/app/messenger/sync/facade/databus"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

// ///////////////////////////////////////////////////////////////////////////////////////////
// type Instance func() Initializer
var facade sync_facade.SyncFacade

func New() sync_facade.SyncFacade {
	if facade == nil {
		facade, _ = sync_facade.NewSyncFacade("esync")
	}
	return facade
}

func SyncUpdatesMe(ctx context.Context, userId int32, authKeyId, sessionId int64, serverId string, updates *mtproto.Updates) error {
	return facade.SyncUpdatesMe(ctx, userId, authKeyId, sessionId, serverId, updates)
}

func SyncUpdatesNotMe(ctx context.Context, userId int32, authKeyId int64, updates *mtproto.Updates) error {
	return facade.SyncUpdatesNotMe(ctx, userId, authKeyId, updates)
}

func PushUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) error {
	return facade.PushUpdates(ctx, userId, updates)
}

func PushBotUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) error {
	return facade.PushBotUpdates(ctx, userId, updates)
}

func PushRpcResult(ctx context.Context, authKeyId int64, serverId string, sessionId, clientReqMsgId int64, result []byte) error {
	return facade.PushRpcResult(ctx, authKeyId, serverId, sessionId, clientReqMsgId, result)
}

func PushUsersUpdates(ctx context.Context, updates ...*model.UserUpdates) error {
	for _, uUpdates := range updates {
		PushUpdates(ctx, uUpdates.UserId, uUpdates.Updates)
	}
	return nil
}

func broadcastChatUpdates(ctx context.Context, chat *model.MutableChat, updates *mtproto.Updates, excludeId ...int32) error {
	if chat == nil {
		return nil
	}

	chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
		if ok, _ := util.Contains(userId, excludeId); !ok {
			facade.PushUpdates(ctx, userId, updates)
		}
		return nil
	})
	return nil
}

func BroadcastChatUpdates(ctx context.Context, chatId int32, updates *mtproto.Updates, excludeId ...int32) error {
	return facade.BroadcastUpdates(ctx, model.BroadcastTypeChat, chatId, excludeId, updates)
}

func BroadcastChannelUpdates(ctx context.Context, channelId int32, updates *mtproto.Updates, excludeId ...int32) error {
	return facade.BroadcastUpdates(ctx, model.BroadcastTypeChannel, channelId, excludeId, updates)
}

func BroadcastChannelAdminsUpdates(ctx context.Context, channelId int32, updates *mtproto.Updates, excludeId ...int32) error {
	return facade.BroadcastUpdates(ctx, model.BroadcastTypeChannelAdmins, channelId, excludeId, updates)
}
