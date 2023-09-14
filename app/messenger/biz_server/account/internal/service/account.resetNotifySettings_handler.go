package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.resetNotifySettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var err error

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.resetNotifySettings - error: %v", err)
		return nil, err
	}

	err = s.UserFacade.ResetNotifySettings(ctx, md.UserId)
	if err != nil {
		log.Errorf("getNotifySettings error - %v", err)
		// We ignore error
		return mtproto.ToBool(false), nil
	}

	pushNotifySettings := func(peerType int32) {
		peer := &model.PeerUtil{
			PeerType: peerType,
			PeerId:   0,
		}
		syncUpdates := model.NewUpdatesLogicByUpdate(md.UserId, mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
			Peer_NOTIFYPEER: peer.ToNotifyPeer(),
			NotifySettings:  model.MakeDefaultPeerNotifySettings(peerType),
		}).To_Update()).ToUpdates()
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, syncUpdates)
	}

	go pushNotifySettings(model.PEER_USERS)
	go pushNotifySettings(model.PEER_CHATS)
	go pushNotifySettings(model.PEER_BROADCASTS)

	log.Debugf("account.resetNotifySettings - reply: {true}")
	return mtproto.ToBool(true), nil
}
