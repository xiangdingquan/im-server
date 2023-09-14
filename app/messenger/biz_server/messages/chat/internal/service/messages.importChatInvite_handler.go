package service

import (
	"context"
	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.importChatInvite - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		replyUpdates *mtproto.Updates
		err          error
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	if request.Hash == "" {
		err = mtproto.ErrInviteHashEmpty
		log.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	if replyUpdates, err = s.importChannelInvite(ctx, md, request.Hash); err != nil {
		if replyUpdates, err = s.importChatInvite(ctx, md, request.Hash); err != nil {
			return nil, err
		}
	}

	if replyUpdates == nil {
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INVITE_HASH_EXPIRED)
	}

	log.Debugf("messages.importChatInvite#6c50051c - reply: {%s}", replyUpdates.DebugString())
	return replyUpdates, nil
}

func (s *Service) importChatInvite(ctx context.Context, md *grpc_util.RpcMetadata, hash string) (*mtproto.Updates, error) {
	chat, err := s.ChatFacade.GetMutableChatByLink(ctx, hash)
	if err != nil {
		log.Warnf("messages.importChatInvite - GetMutableChatByLink error: %v", err)
		return nil, err
	}

	log.Debugf("chat: %v", chat.Chat)
	chat, err = s.ChatFacade.AddChatUser(ctx, chat.Chat.Id, chat.Chat.Creator, md.UserId)
	if err != nil {
		log.Warnf("messages.importChatInvite - addChatUser error: %v", err)
		return nil, err
	}

	replyUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChatPeerUtil(chat.Chat.Id),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatJoinByLink(chat.Chat.Creator)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("addChatUser error: %v", err)
		return nil, err
	}

	log.Debugf("messages.addChatUser#f9a0aa09 - reply: {%v}", replyUpdates.DebugString())
	return replyUpdates, nil
}

func (s *Service) importChannelInvite(ctx context.Context, md *grpc_util.RpcMetadata, hash string) (*mtproto.Updates, error) {
	channel, err := s.ChannelFacade.GetMutableChannelByLink(ctx, hash)
	if err != nil {
		log.Warnf("messages.importChatInvite - GetMutableChatByLink error: %v", err)
		return nil, err
	}

	channel, err = s.ChannelFacade.JoinChannel(ctx, channel.GetId(), md.UserId, false)
	if err != nil {
		log.Warnf("messages.importChatInvite - addChatUser error: %v", err)
		return nil, err
	}

	replyUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(channel.GetId()),
		&msgpb.OutboxMessage{
			NoWebpage:  true,
			Background: false,
			RandomId:   rand.Int63(),
			Message: channel.MakeMessageService(md.UserId,
				false,
				0,
				model.MakeMessageActionChatJoinByLink(channel.Channel.CreatorId)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("addChatUser error: %v", err)
		return nil, err
	}

	log.Debugf("messages.addChatUser#f9a0aa09 - reply: {%v}", replyUpdates.DebugString())
	return replyUpdates, nil
}
