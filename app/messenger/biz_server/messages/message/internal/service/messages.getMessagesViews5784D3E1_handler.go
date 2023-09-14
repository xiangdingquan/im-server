package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (s *Service) MessagesGetMessagesViews5784D3E1(ctx context.Context, request *mtproto.TLMessagesGetMessagesViews5784D3E1) (*mtproto.Messages_MessageViews, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getMessagesViews5784D3E1 - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	var (
		boxMsgList    model.MessageBoxList
		channelIds    = make(map[int32][]int32)
		reqIndexViews = make(map[int32]int64)
		increment     = mtproto.FromBool(request.GetIncrement())
		peer          = model.FromInputPeer2(md.UserId, request.Peer)
	)
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		switch peer.PeerType {
		case model.PEER_USER:
			mutableUsers := s.UserFacade.GetMutableUsers(ctx, md.UserId, peer.PeerId)
			if len(mutableUsers) != 2 {
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.getMessagesViews - error: %v, not found peer(%v) == 0", err, peer)
				return nil, err
			}
		case model.PEER_CHAT:
			mutableChat, err := s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId)
			if err != nil {
				log.Errorf("messages.getMessagesViews - error: not found chat by id(%v) type", peer)
				return nil, mtproto.ErrChatIdInvalid
			} else if !mutableChat.CheckParticipantExist(md.UserId) {
				log.Errorf("messages.getMessagesViews - error: not in chat(%v) type", request.Id)
				return nil, mtproto.ErrChannelPrivate
			}
		}
		boxMsgList = s.MessageFacade.GetUserMessageList(ctx, md.UserId, request.Id)
		if len(request.Id) != len(boxMsgList) {
			log.Errorf("messages.getMessagesViews - error: boxMsgList empty by id(%v) type", request.Id)
			return nil, mtproto.ErrMsgIdInvalid
		}
		for _, box := range boxMsgList {
			fwdFrom := box.Message.GetFwdFrom()
			if fwdFrom == nil {
				log.Errorf("messages.getMessagesViews - error: boxMsgList empty by id(%v) type", request.Id)
				return nil, mtproto.ErrMsgIdInvalid
			} else {
				if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
					channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
				} else {
					idList = append(idList, fwdFrom.GetChannelPost().GetValue())
					channelIds[fwdFrom.GetChannelId().GetValue()] = idList
				}
				reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
			}
		}
	case model.PEER_CHANNEL:
		channel, err := s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId)
		if err != nil {
			log.Errorf("messages.getMessagesViews - error: not found chat by id(%v) type", peer)
			return nil, mtproto.ErrChannelInvalid
		}
		boxMsgList = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, peer.PeerId, request.Id)
		boxMsgList = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, peer.PeerId, request.Id)
		if channel.Channel.IsMegagroup() {
			for _, box := range boxMsgList {
				fwdFrom := box.Message.GetFwdFrom()
				if fwdFrom == nil {
					log.Errorf("messages.getMessagesViews - error: boxMsgList empty by id(%v) type", request.Id)
					return nil, mtproto.ErrMsgIdInvalid
				} else {
					if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
						channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
					} else {
						idList = append(idList, box.Message.GetFwdFrom().GetChannelPost().GetValue())
						channelIds[fwdFrom.GetChannelId().GetValue()] = idList
					}
					reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
				}
			}
		} else if channel.Channel.IsBroadcast() {
			for _, box := range boxMsgList {
				fwdFrom := box.Message.GetFwdFrom()
				if fwdFrom == nil {
					if idList, ok := channelIds[box.MessageId]; !ok {
						channelIds[channel.Channel.Id] = []int32{box.MessageId}
					} else {
						idList = append(idList, box.MessageId)
						channelIds[channel.Channel.Id] = idList
					}
				} else {
					if idList, ok := channelIds[fwdFrom.GetChannelId().GetValue()]; !ok {
						channelIds[fwdFrom.GetChannelId().GetValue()] = []int32{fwdFrom.GetChannelPost().GetValue()}
					} else {
						idList = append(idList, box.Message.GetFwdFrom().GetChannelPost().GetValue())
						channelIds[fwdFrom.GetChannelId().GetValue()] = idList
					}
					reqIndexViews[box.MessageId] = int64(fwdFrom.GetChannelId().GetValue())<<32 | int64(fwdFrom.GetChannelPost().GetValue())
				}
			}
		} else {
			log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
			return nil, mtproto.ErrMsgIdInvalid
		}
	default:
		log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
		return nil, mtproto.ErrInputRequestInvalid
	}

	if len(request.Id) == 0 {
		err := mtproto.ErrInputRequestInvalid
		log.Errorf("messages.getMessagesViews - error: %v, invalid request, len(id) == 0", err)
		return nil, err
	}

	var allViews map[int32][]*mtproto.MessageViews = make(map[int32][]*mtproto.MessageViews)
	for k, v := range channelIds {
		allViews[k] = s.ChannelFacade.GetChannelMessagesViews2(ctx, k, v, increment)
		if increment {
			s.ChannelFacade.IncrementChannelMessagesViews(ctx, k, v)
		}
	}

	var views []*mtproto.MessageViews
	for _, id := range request.Id {
		if indexV, ok := reqIndexViews[id]; !ok {
			log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
			return nil, mtproto.ErrMsgIdInvalid
		} else {
			rcId := int32(indexV >> 32)
			rmId := int32(indexV & 0xffffffff)
			var v *mtproto.MessageViews
			for i, id2 := range channelIds[rcId] {
				if id2 == rmId {
					v = allViews[rcId][i]
					break
				}
			}
			if v == nil {
				log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
				return nil, mtproto.ErrMsgIdInvalid
			}
			views = append(views, v)
		}
	}

	mtproto.MakeTLMessagesMessageViews(&mtproto.Messages_MessageViews{
		Views: views,
		Chats: nil,
		Users: nil,
	}).To_Messages_MessageViews()

	// Sorry: not impl MessagesGetMessagesViews5784D3E1 logic
	log.Warn("messages.getMessagesViews5784D3E1 - error: method MessagesGetMessagesViews5784D3E1 not impl")

	return nil, mtproto.ErrMethodNotImpl
}
