package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.checkChatInvite - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		chatInvite *mtproto.ChatInvite
		err        error
	)

	if request.Hash == "" {
		err := mtproto.ErrInviteHashEmpty
		log.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	} else {
	}

	chatInvite, err = s.checkChannelInvite(ctx, md, request.Hash)
	if err != nil {
		chatInvite, err = s.checkChatInvite(ctx, md, request.Hash)
	}

	if err != nil {
		return nil, err
	}

	// OK:
	log.Debugf("messages.checkChatInvite - reply: {%s}", chatInvite.DebugString())
	return chatInvite, nil
}

func (s *Service) checkChatInvite(ctx context.Context, md *grpc_util.RpcMetadata, hash string) (*mtproto.ChatInvite, error) {
	chat, err := s.ChatFacade.GetMutableChatByLink(ctx, hash, md.UserId)
	if err != nil {
		log.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}

	var chatInvite *mtproto.ChatInvite
	me := chat.GetImmutableChatParticipant(md.UserId)
	if me != nil && me.IsChatMemberStateNormal() {
		chatInvite = mtproto.MakeTLChatInviteAlready(&mtproto.ChatInvite{
			Chat: chat.ToUnsafeChat(md.UserId),
		}).To_ChatInvite()
	} else {
		chatInvite = mtproto.MakeTLChatInvite(&mtproto.ChatInvite{
			Channel:           false,
			Broadcast:         false,
			Public:            false,
			Megagroup:         false,
			Title:             chat.Chat.Title,
			Photo_PHOTO:       chat.Chat.ChatPhoto,
			ParticipantsCount: chat.Chat.ParticipantsCount,
			Participants: s.UserFacade.GetUserListByIdList(
				ctx,
				md.UserId,
				chat.ToChatParticipantIdList()),
		}).To_ChatInvite()
	}

	return chatInvite, nil
}

func (s *Service) checkChannelInvite(ctx context.Context, md *grpc_util.RpcMetadata, hash string) (*mtproto.ChatInvite, error) {
	channel, err := s.ChannelFacade.GetMutableChannelByLink(ctx, hash, md.UserId)
	if err != nil {
		log.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}
	//if channel.Channel.Username != "" {
	//	err = mtproto.ErrInviteHashExpired
	//	log.Errorf("messages.checkChatInvite - error: %v", err)
	//	return nil, err
	//}

	var chatInvite *mtproto.ChatInvite
	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me != nil {
		if me.IsKicked() {
			err = mtproto.ErrInviteHashExpired
			log.Errorf("messages.checkChatInvite - error: %v", err)
			return nil, err
		} else if me.IsStateOk() {
			chatInvite = mtproto.MakeTLChatInviteAlready(&mtproto.ChatInvite{
				Chat: channel.ToUnsafeChat(md.UserId),
			}).To_ChatInvite()
			return chatInvite, nil
		}
	}

	//channel.Channel.ParticipantsCount 不准确 暂时不知bug出在哪 临时处理
	participantsCount, _, _, _ := s.ChannelFacade.GetParticipantCounts(ctx, channel.GetId())

	chatInvite = mtproto.MakeTLChatInvite(&mtproto.ChatInvite{
		Channel:           true,
		Broadcast:         channel.Channel.IsBroadcast(),
		Public:            channel.Channel.Username != "",
		Megagroup:         channel.Channel.IsMegagroup(),
		Title:             channel.Channel.Title,
		Photo_PHOTO:       channel.Channel.ChatPhoto,
		ParticipantsCount: participantsCount, //channel.Channel.ParticipantsCount
	}).To_ChatInvite()

	var (
		userIdList model.IDList
	)

	if channel.Channel.IsMegagroup() {
		channel.FetchAndWalk(
			func() []*model.ImmutableChannelParticipant {
				return s.ChannelFacade.GetChannelParticipantRecentList(ctx, channel.Channel, 0, 50, 0)
			},
			func(participant *model.ImmutableChannelParticipant) {
				userIdList.AddIfNot(participant.UserId)
			})

		chatInvite.Participants = s.UserFacade.GetUserListByIdList(
			ctx,
			md.UserId,
			userIdList)
	}
	return chatInvite, nil
}
