package service

import (
	"context"
	"math"

	"open.chat/model"
	"open.chat/pkg/math2"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) getUnreadMentions(ctx context.Context, selfUserId int32, peer *model.PeerUtil, offsetId, addOffset, limit, maxId, minId int32) (boxList model.MessageBoxList) {
	loadType := loadTypeBackward
	if addOffset >= 0 {
		loadType = loadTypeBackward
	} else if addOffset+limit > 0 {
		loadType = loadTypeFirstAroundDate
	} else {
		loadType = loadTypeForward
	}

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	switch loadType {
	case loadTypeBackward:
		if offsetId == 0 {
			offsetId = math.MaxInt32
		}
		boxList = s.MessageFacade.GetOffsetIdBackwardUnreadMentions(ctx, selfUserId, peer, offsetId, minId, maxId, addOffset+limit)
	case loadTypeFirstAroundDate:
		boxList1 := s.MessageFacade.GetOffsetIdForwardUnreadMentions(ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset)
		for i, j := 0, len(boxList1)-1; i < j; i, j = i+1, j-1 {
			boxList1[i], boxList1[j] = boxList1[j], boxList1[i]
		}
		boxList = append(boxList, boxList1...)
		boxList2 := s.MessageFacade.GetOffsetIdBackwardUnreadMentions(ctx, selfUserId, peer, offsetId, minId, maxId, limit+addOffset)
		boxList = append(boxList, boxList2...)
	case loadTypeForward:
		boxList = s.MessageFacade.GetOffsetIdForwardUnreadMentions(ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset)
		for i, j := 0, len(boxList)-1; i < j; i, j = i+1, j-1 {
			boxList[i], boxList[j] = boxList[j], boxList[i]
		}
	}
	return
}

func (s *Service) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getUnreadMentions - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err              error
		peer             = model.FromInputPeer2(md.UserId, request.GetPeer())
		chat             *model.MutableChat
		channel          *model.MutableChannel
		isChannel        bool
		messagesMessages *mtproto.Messages_Messages
		minId            = request.MinId
		limit            = request.Limit
	)

	if limit > 50 {
		limit = 50
	}

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_CHAT:
		// 400	CHAT_ID_INVALID	The provided chat id is invalid
		if chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId); err != nil {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("messages.getHistory - error: %v", err)
			return nil, err
		}
		_ = chat
	case model.PEER_CHANNEL:
		// 400	CHANNEL_INVALID	The provided channel is invalid
		// 400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup
		if channel, err = s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId); err != nil {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("messages.getHistory - error: %v", err)
			return nil, err
		}

		me := channel.GetImmutableChannelParticipant(md.UserId)
		if me != nil {
			minId = math2.Int32Max(me.AvailableMinId, request.MinId)
		}

		isChannel = true
		_ = channel
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	boxList := s.getUnreadMentions(ctx, md.UserId, peer, request.OffsetId, request.AddOffset, limit, request.MaxId, minId)

	if !isChannel {
		if int32(len(boxList)) == limit {
			messagesMessages = mtproto.MakeTLMessagesMessagesSlice(&mtproto.Messages_Messages{
				Count: int32(len(boxList)),
			}).To_Messages_Messages()
		} else {
			messagesMessages = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
				Count: int32(len(boxList)),
			}).To_Messages_Messages()
		}
	} else {
		messagesMessages = mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
			Messages: nil,
			Chats:    nil,
			Users:    nil,
			Inexact:  false,
			Count:    0,
			NextRate: nil,
			Pts:      channel.Channel.Pts,
		}).To_Messages_Messages()
	}
	messagesMessages.Messages, messagesMessages.Users, messagesMessages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	messagesMessages.Count = int32(len(messagesMessages.Messages))

	log.Debugf("messages.getUnreadMentions - reply: %s", messagesMessages.DebugString())
	return messagesMessages, nil
}
