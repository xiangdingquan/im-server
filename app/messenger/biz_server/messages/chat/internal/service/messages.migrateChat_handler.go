package service

import (
	"context"
	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesMigrateChat(ctx context.Context, request *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.migrateChat - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	chat, err := s.ChatFacade.GetMutableChat(ctx, request.ChatId, md.UserId)
	if err != nil {
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	me := chat.GetImmutableChatParticipant(md.UserId)
	if me == nil {
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	if !me.IsChatMemberAdmin() && !me.IsChatMemberCreator() {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	key := crypto.CreateAuthKey()
	_, err = s.RPCSessionClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey: &authsessionpb.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        model.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})
	if err != nil {
		log.Errorf("create channel secret key error")
		return nil, err
	}

	channel, err := s.ChannelFacade.MigrateFromChat(ctx, md.UserId, 0, chat)
	if err != nil {
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	// sendChannelMessage
	migrateFromChatUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(channel.GetChannelId()),
		&msgpb.OutboxMessage{
			NoWebpage:  true,
			Background: false,
			RandomId:   rand.Int63(),
			Message: channel.MakeMessageService(md.UserId,
				false,
				0,
				model.MakeMessageActionChannelMigrateFrom(channel.Channel.Title, chat.Chat.Id)),
			ScheduleDate: nil,
		},
	)
	if err != nil {
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	err = s.ChatFacade.MigratedToChannel(ctx, chat, channel.GetId(), channel.GetAccessHash())
	if err != nil {
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	migratedToChannelUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChatPeerUtil(chat.Chat.Id),
		&msgpb.OutboxMessage{
			NoWebpage:    false,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatMigrateTo(channel.GetId())),
			ScheduleDate: nil,
		})
	if err != nil {
		log.Errorf("messages.migrateChat - error: %v", err)
		return nil, err
	}

	updateNotifySettings := mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
		Peer_NOTIFYPEER: mtproto.MakeTLNotifyPeer(&mtproto.NotifyPeer{
			Peer: mtproto.MakeTLPeerChannel(&mtproto.Peer{
				ChannelId: channel.GetId(),
			}).To_Peer(),
		}).To_NotifyPeer(),
		NotifySettings: model.MakeDefaultPeerNotifySettings(model.PEER_CHANNEL),
	}).To_Update()

	updateReadHistoryOutbox := mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
		Peer_PEER: model.MakePeerChat(chat.Chat.Id),
		MaxId:     me.Dialog.TopMessage,
		Pts_INT32: int32(idgen.NextPtsId(context.Background(), md.UserId)),
		PtsCount:  1,
	}).To_Update()

	// 1. updateNotifySettings
	go func() {
		updateHelper := model.MakeUpdatesHelper(
			mtproto.MakeTLUpdateChannel(&mtproto.Update{
				ChannelId: channel.GetId(),
			}).To_Update(),
			updateNotifySettings)

		syncNotMe := updateHelper.ToPushUpdates(ctx, md.UserId, nil, nil, nil)
		sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, syncNotMe)
	}()

	// merge updates
	//
	rUpdates := model.MakeUpdatesByUpdates(updateNotifySettings)

	for _, upd := range migrateFromChatUpdates.Updates {
		if upd.PredicateName != mtproto.Predicate_updateMessageID {
			rUpdates.Updates = append(rUpdates.Updates, upd)
		}
	}
	for _, upd := range migratedToChannelUpdates.Updates {
		if upd.PredicateName != mtproto.Predicate_updateMessageID {
			rUpdates.Updates = append(rUpdates.Updates, upd)
		}
	}
	rUpdates.Updates = append(rUpdates.Updates, updateReadHistoryOutbox)

	for _, u2 := range migrateFromChatUpdates.Users {
		found := false
		for _, u := range rUpdates.Users {
			if u.Id == u2.Id {
				found = true
				break
			}
		}
		if !found {
			rUpdates.Users = append(rUpdates.Users, u2)
		}
	}

	for _, u2 := range migratedToChannelUpdates.Users {
		found := false
		for _, u := range rUpdates.Users {
			if u.Id == u2.Id {
				found = true
				break
			}
		}
		if !found {
			rUpdates.Users = append(rUpdates.Users, u2)
		}
	}

	for _, u2 := range migrateFromChatUpdates.Users {
		found := false
		for _, u := range rUpdates.Users {
			if u.Id == u2.Id {
				found = true
				break
			}
		}
		if !found {
			rUpdates.Users = append(rUpdates.Users, u2)
		}
	}

	for _, c2 := range migratedToChannelUpdates.Chats {
		found := false
		for _, c := range rUpdates.Chats {
			if c.Id == c2.Id {
				found = true
				break
			}
		}
		if !found {
			rUpdates.Chats = append(rUpdates.Chats, c2)
		}
	}

	for _, c2 := range migrateFromChatUpdates.Chats {
		found := false
		for _, c := range rUpdates.Chats {
			if c.Id == c2.Id {
				found = true
				break
			}
		}
		if !found {
			rUpdates.Chats = append(rUpdates.Chats, c2)
		}
	}

	// readHistory
	go func() {
		sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, model.MakeUpdatesByUpdates(updateReadHistoryOutbox))
	}()

	log.Debugf("messages.migrateChat#15a3b8e3 - reply: {%s}", rUpdates.DebugString())
	return rUpdates, nil
}
