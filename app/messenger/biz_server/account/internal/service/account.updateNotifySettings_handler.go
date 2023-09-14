package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateNotifySettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err         error
		peerUser    *mtproto.User
		peerChat    *mtproto.Chat
		peerChannel *mtproto.Chat
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.unregisterDevice - error: %v", err)
		return nil, err
	}

	settings, err := model.MakePeerNotifySettings(request.Settings)
	if err != nil {
		log.Errorf("account.updateNotifySettings - error: %v", err)
		return nil, err
	}

	peer := model.FromInputNotifyPeer(md.UserId, request.GetPeer())
	switch peer.PeerType {
	case model.PEER_USERS:
	case model.PEER_CHATS:
	case model.PEER_BROADCASTS:
	case model.PEER_USER:
		peerUser, err = s.UserFacade.GetUserById(ctx, md.UserId, peer.PeerId)
		if err != nil {
			log.Errorf("account.updateNotifySettings - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
	case model.PEER_CHAT:
		peerChat = s.ChatFacade.GetChatBySelfId(ctx, md.UserId, peer.PeerId)
		if peerChat.PredicateName == mtproto.Predicate_chatEmpty {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("account.updateNotifySettings - error: %v", err)
			return nil, err
		}
	case model.PEER_CHANNEL:
		channel, err2 := s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId)
		if err2 != nil {
			log.Errorf("account.updateNotifySettings - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
		me := channel.GetImmutableChannelParticipant(md.UserId)
		if me == nil || me.IsKicked() {
			err = mtproto.ErrChannelPrivate
			log.Errorf("account.updateNotifySettings - error: %v", err)
			return nil, err
		}
		peerChannel = channel.ToUnsafeChat(md.UserId)
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("account.updateNotifySettings - error: %v", err)
		return nil, err
	}

	err = s.UserFacade.SetNotifySettings(ctx, md.UserId, peer, settings)
	if err != nil {
		log.Errorf("setNotifySettings error - %v", err)
		return nil, err
	}

	go func() {
		//// sync
		updateNotifySettings := mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
			Peer_NOTIFYPEER: peer.ToNotifyPeer(),
			NotifySettings:  settings,
		})
		notifySettingUpdates := model.NewUpdatesLogic(md.UserId)
		notifySettingUpdates.AddUpdate(updateNotifySettings.To_Update())

		switch peer.PeerType {
		case model.PEER_USER:
			// user, _ := s.UserFacade.GetUserById(ctx2, md.UserId, peer.PeerId)
			notifySettingUpdates.AddUser(peerUser)
		case model.PEER_CHAT:
			notifySettingUpdates.AddChat(peerChat)
		case model.PEER_CHANNEL:
			notifySettingUpdates.AddChat(peerChannel)
		case model.PEER_USERS:
		case model.PEER_CHATS:
		case model.PEER_BROADCASTS:
		default:
			return
		}

		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, notifySettingUpdates.ToUpdates())
	}()

	log.Debugf("account.updateNotifySettings#84be5b93 - reply: {true}")
	return mtproto.ToBool(true), nil
}
