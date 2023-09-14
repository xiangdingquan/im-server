package service

import (
	"context"

	"math"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/math2"
)

const (
	loadTypeBackward           = 0
	loadTypeForward            = 1
	loadTypeFirstUnread        = 2
	loadTypeFirstAroundMessage = 3
	loadTypeFirstAroundDate    = 4
	loadTypeLimit1             = 16
)

func calcLoadHistoryType(isChannel bool, offsetId, offsetDate, addOffset, limit, maxId, minId int32) int {
	if limit == 1 {
		return loadTypeLimit1
	}

	if isChannel && addOffset == -1 && maxId != 0 {
		return loadTypeBackward
	}

	if addOffset == 0 {
		return loadTypeBackward
	} else if addOffset == -1 {
		return loadTypeBackward
	} else if addOffset == -limit+5 {
		return loadTypeFirstAroundDate
	} else if addOffset == -limit/2 {
		return loadTypeFirstAroundMessage
	} else if addOffset == -limit-1 {
		return loadTypeForward
	} else if addOffset == -limit+6 {
		if maxId != 0 {
			return loadTypeFirstUnread
		}
	}
	return loadTypeForward
}

func (s *Service) getHistoryMessages(ctx context.Context, selfUserId int32, peer *model.PeerUtil, offsetId, offsetDate, addOffset, limit, maxId, minId, hash int32) (boxList model.MessageBoxList) {
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
		boxList = s.MessageFacade.GetOffsetIdBackwardHistoryMessages(ctx, selfUserId, peer, offsetId, minId, maxId, addOffset+limit, hash)
	case loadTypeFirstAroundDate:
		boxList1 := s.MessageFacade.GetOffsetIdForwardHistoryMessages(ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
		for i, j := 0, len(boxList1)-1; i < j; i, j = i+1, j-1 {
			boxList1[i], boxList1[j] = boxList1[j], boxList1[i]
		}
		boxList = append(boxList, boxList1...)
		// 降序
		boxList2 := s.MessageFacade.GetOffsetIdBackwardHistoryMessages(ctx, selfUserId, peer, offsetId, minId, maxId, limit+addOffset, hash)
		// log.Infof("%v", messages2)
		boxList = append(boxList, boxList2...)
	case loadTypeForward:
		boxList = s.MessageFacade.GetOffsetIdForwardHistoryMessages(ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
		for i, j := 0, len(boxList)-1; i < j; i, j = i+1, j-1 {
			boxList[i], boxList[j] = boxList[j], boxList[i]
		}
	}
	return
}

func (s *Service) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getHistory#dcbb8260 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

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
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		if peer.PeerType == model.PEER_CHAT {
			// 400	CHAT_ID_INVALID	The provided chat id is invalid
			if chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId); err != nil {
				err = mtproto.ErrPeerIdInvalid
				log.Errorf("messages.getHistory - error: %v", err)
				return nil, err
			}
			_ = chat
		}
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

	boxList := s.getHistoryMessages(ctx, md.UserId, peer, request.OffsetId, request.OffsetDate, request.AddOffset, limit, request.MaxId, minId, request.Hash)

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
			Count:    s.MessageFacade.GetHistoryMessagesCount(ctx, md.UserId, peer),
			NextRate: nil,
			Pts:      channel.Channel.Pts,
		}).To_Messages_Messages()
	}
	messagesMessages.Messages, messagesMessages.Users, messagesMessages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	log.Debugf("messages.getHistory - reply: %s", messagesMessages.DebugString())
	return messagesMessages, nil
}
