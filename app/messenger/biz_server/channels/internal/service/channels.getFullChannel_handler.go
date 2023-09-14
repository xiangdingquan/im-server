package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/app/pkg/env2"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetFullChannel(ctx context.Context, request *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getFullChannel#8736a09 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.getFullChannel - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.getFullChannel - error: %v", err)
		return nil, err
	}

	hasUsername := channel.Channel.Username != ""
	me := channel.GetImmutableChannelParticipant(md.UserId)

	if !hasUsername {
		if me == nil || !me.IsStateOk() {
			err := mtproto.ErrChannelPrivate
			log.Errorf("channels.getFullChannel - error: %v", err)
			return nil, err
		}
	} else {
		if me != nil && me.IsKicked() {
			err := mtproto.ErrChannelPrivate
			log.Errorf("channels.getFullChannel - error: %v", err)
			return nil, err
		}
	}

	fullChat := mtproto.MakeTLChannelFull(&mtproto.ChatFull{
		CanViewParticipants:  false,
		CanSetUsername:       false,
		CanSetStickers:       false,
		HiddenPrehistory:     channel.Channel.HiddenPrehistory,
		CanViewStats:         false,
		CanSetLocation:       false,
		HasScheduled:         false,
		Id:                   channel.GetChannelId(),
		About:                channel.Channel.About,
		Notice:               channel.Channel.Notice,
		ParticipantsCount:    &types.Int32Value{Value: channel.Channel.ParticipantsCount},
		AdminsCount:          nil,
		KickedCount:          nil,
		BannedCount:          nil,
		OnlineCount:          &types.Int32Value{Value: 0},
		ReadInboxMaxId:       0,
		ReadOutboxMaxId:      0,
		UnreadCount:          0,
		ChatPhoto:            channel.Channel.ChatPhoto,
		NotifySettings:       mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings(),
		ExportedInvite:       model.ExportedChatInviteEmpty,
		BotInfo:              nil,
		MigratedFromChatId:   nil,
		MigratedFromMaxId:    nil,
		PinnedMsgId:          nil,
		Stickerset:           nil,
		AvailableMinId:       nil,
		FolderId:             nil,
		LinkedChatId:         nil,
		Location:             channel.Channel.Location,
		SlowmodeSeconds:      nil,
		SlowmodeNextSendDate: nil,
		Pts:                  channel.Channel.Pts,
	}).To_ChatFull()

	if channel.Channel.IsBroadcast() && me != nil && me.IsCreatorOrAdmin() {
		fullChat.CanViewParticipants = true
	} else if channel.Channel.IsMegagroup() {
		fullChat.CanViewParticipants = true
	}
	if channel.Channel.IsBroadcast() {
		fullChat.ReadOutboxMaxId = channel.Channel.TopMessage
		fullChat.ReadInboxMaxId = 0
	}

	if me != nil && me.CanAdminInviteUsers() {
		fullChat.CanSetUsername = true
	}

	count, adminsCount, kickedCount, bannedCount := s.ChannelFacade.GetParticipantCounts(ctx, channel.GetChannelId())
	fullChat.ParticipantsCount = &types.Int32Value{Value: count}
	if me != nil && me.CanAdminInviteUsers() {
		fullChat.AdminsCount = &types.Int32Value{Value: adminsCount}
		fullChat.KickedCount = &types.Int32Value{Value: kickedCount}
		fullChat.BannedCount = &types.Int32Value{Value: bannedCount}
	}

	if me != nil && me.IsStateOk() {
		fullChat.ReadInboxMaxId = me.ReadInboxMaxId
		fullChat.UnreadCount = me.UnreadCount
		fullChat.ReadOutboxMaxId = channel.Channel.ReadOutboxMaxId
	}

	if channel.Channel.Link != "" {
		fullChat.ExportedInvite = mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
			Link: env2.T_ME + "/joinchat?link=" + channel.Channel.Link,
		}).To_ExportedChatInvite()
	}

	fullChat.BotInfo = []*mtproto.BotInfo{}

	if me != nil && me.MigratedFromMaxId != 0 {
		fullChat.MigratedFromMaxId = &types.Int32Value{Value: me.MigratedFromMaxId}
		fullChat.MigratedFromChatId = &types.Int32Value{Value: channel.Channel.MigratedFromChatId}
	}

	if channel.Channel.PinnedMsgId != 0 {
		fullChat.PinnedMsgId = &types.Int32Value{Value: channel.Channel.PinnedMsgId}
	}

	if me != nil && me.IsStateOk() && me.FolderId != 0 {
		fullChat.AvailableMinId = &types.Int32Value{Value: me.AvailableMinId}
	}

	if me != nil && me.IsStateOk() && me.FolderId != 0 {
		fullChat.FolderId = &types.Int32Value{Value: me.FolderId}
	}

	if channel.Channel.LinkedChatId > 0 {
		fullChat.LinkedChatId = &types.Int32Value{Value: channel.Channel.LinkedChatId}
	}

	if channel.Channel.SlowmodeSeconds > 0 {
		fullChat.SlowmodeSeconds = &types.Int32Value{Value: channel.Channel.SlowmodeSeconds}
	}

	if channel.Channel.SlowmodeNextSendDate > 0 {
		fullChat.SlowmodeNextSendDate = &types.Int32Value{Value: channel.Channel.SlowmodeNextSendDate}
	}

	if err != nil {
		log.Errorf("channels.getFullChannel - error: %v", err)
		return nil, err
	}

	messagesChatFull := mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: fullChat,
		Chats:    []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
		Users:    []*mtproto.User{},
	}).To_Messages_ChatFull()

	if channel.Channel.LinkedChatId > 0 {
		if linkChat, _ := s.ChannelFacade.GetMutableChannel(ctx, channel.Channel.LinkedChatId, md.UserId); linkChat != nil {
			messagesChatFull.Chats = append(messagesChatFull.Chats, linkChat.ToUnsafeChat(md.UserId))
		}
	}

	log.Debugf("channels.getFullChannel#8736a09 - reply: %s", messagesChatFull.DebugString())
	return messagesChatFull, nil
}
