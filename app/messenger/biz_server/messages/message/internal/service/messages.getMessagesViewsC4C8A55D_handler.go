package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.getMessagesViews#c4c8a55d peer:InputPeer id:Vector<int> increment:Bool = Vector<int>;
func (s *Service) MessagesGetMessagesViewsC4C8A55D(ctx context.Context, request *mtproto.TLMessagesGetMessagesViewsC4C8A55D) (*mtproto.Vector_Int, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getMessagesViews - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		viewsList     []int32
		boxMsgList    model.MessageBoxList
		channelIds    = make(map[int32][]int32)
		reqIndexViews = make(map[int32]int64)
		views         = make(map[int32][]int32)
		increment     = mtproto.FromBool(request.GetIncrement())
		peer          = model.FromInputPeer2(md.UserId, request.Peer)
	)

	if len(request.Id) == 0 {
		err := mtproto.ErrInputRequestInvalid
		log.Errorf("messages.getMessagesViews - error: %v, invalid request, len(id) == 0", err)
		return nil, err
	}

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
				return nil, mtproto.ErrChatIdInvalid
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

		//me := channel.GetImmutableChannelParticipant(md.UserId)
		//if me == nil || !me.IsStateOk() {
		//	err = mtproto.ErrChannelPrivate
		//	log.Errorf("channels.updateUsername - error: %v", err)
		//	return nil, err
		//}

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

	for k, v := range channelIds {
		views[k] = s.ChannelFacade.GetChannelMessagesViews(ctx, k, v, increment)
		if increment {
			s.ChannelFacade.IncrementChannelMessagesViews(ctx, k, v)
		}
	}

	for _, id := range request.Id {
		if indexV, ok := reqIndexViews[id]; !ok {
			log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
			return nil, mtproto.ErrMsgIdInvalid
		} else {
			rcId := int32(indexV >> 32)
			rmId := int32(indexV & 0xffffffff)
			mId := int32(0)
			for i, id2 := range channelIds[rcId] {
				if id2 == rmId {
					mId = views[rcId][i]
					break
				}
			}
			if mId == 0 {
				log.Errorf("messages.getMessagesViews - error: invalid peer(%v) type", peer)
				return nil, mtproto.ErrMsgIdInvalid
			}
			viewsList = append(viewsList, mId)
		}
	}

	rViews := &mtproto.Vector_Int{
		Datas: viewsList,
	}
	log.Debugf("messages.getMessagesViews - reply: %s", rViews.DebugString())
	return rViews, nil
}
