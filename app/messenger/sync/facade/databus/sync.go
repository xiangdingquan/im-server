package databus

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"open.chat/app/infra/databus/pkg/queue/databus"
	sync_facade "open.chat/app/messenger/sync/facade"
	"open.chat/app/messenger/sync/syncpb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type syncClient struct {
	syncBus *databus.Databus
}

func New() sync_facade.SyncFacade {
	var (
		dbus struct {
			Sync *databus.Config
		}
	)
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))
	return &syncClient{
		syncBus: databus.New(dbus.Sync),
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *syncClient) SyncUpdatesMe(ctx context.Context, userId int32, authKeyId, sessionId int64, serverId string, updates *mtproto.Updates) (err error) {
	m := &syncpb.TLSyncSyncUpdates{
		UserId:    userId,
		AuthKeyId: authKeyId,
		ServerId:  &types.StringValue{Value: serverId},
		SessionId: &types.Int64Value{Value: sessionId},
		Updates:   updates,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("SyncUpdatesMe.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func (c *syncClient) SyncUpdatesNotMe(ctx context.Context, userId int32, authKeyId int64, updates *mtproto.Updates) (err error) {
	m := &syncpb.TLSyncSyncUpdates{
		UserId:    userId,
		AuthKeyId: authKeyId,
		Updates:   updates,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("SyncUpdatesNotMe.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func (c *syncClient) PushUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) (err error) {
	m := &syncpb.TLSyncPushUpdates{
		UserId:  userId,
		Updates: updates,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("PushUpdates.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func (c *syncClient) PushBotUpdates(ctx context.Context, userId int32, updates *mtproto.Updates) (err error) {
	m := &syncpb.TLSyncPushUpdates{
		UserId:  userId,
		IsBot:   true,
		Updates: updates,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("PushUpdates.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func (c *syncClient) PushRpcResult(ctx context.Context, authKeyId int64, serverId string, sessionId, clientReqMsgId int64, result []byte) (err error) {
	m := &syncpb.TLSyncPushRpcResult{
		ServerId:  serverId,
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		ReqMsgId:  clientReqMsgId,
		Result:    result,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("PushRpcResult.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func (c *syncClient) BroadcastUpdates(ctx context.Context, broadcastType int32, chatId int32, excludeIdList []int32, updates *mtproto.Updates) (err error) {
	m := &syncpb.TLSyncBroadcastUpdates{
		BroadcastType: broadcastType,
		ChatId:        chatId,
		ExcludeIds:    excludeIdList,
		Updates:       updates,
	}

	if err = c.syncBus.Send(ctx, proto.MessageName(m), m); err != nil {
		log.Errorf("BroadcastUpdates.send(updates:%v).error(%v)", logger.JsonDebugData(m), err)
	}

	return
}

func init() {
	sync_facade.Register("esync", New)
}
