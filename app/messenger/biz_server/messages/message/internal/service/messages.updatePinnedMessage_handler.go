package service

import (
	"context"
	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUpdatePinnedMessage(ctx context.Context, request *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.updatePinnedMessage - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer          = model.FromInputPeer2(md.UserId, request.Peer)
		resultUpdates *mtproto.Updates
	)

	switch peer.PeerType {
	case model.PEER_SELF:
		if request.Id != 0 {
			_, err := s.MessageFacade.GetUserMessage(ctx, md.UserId, request.Id)
			if err != nil {
				log.Errorf("messages.updatePinnedMessage - error: %v", err)
				return nil, err
			}
		}

		s.PrivateFacade.UpdateUserPinnedMessage(ctx, md.UserId, md.UserId, request.Id)

		resultUpdates = model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPinnedMessage(&mtproto.Update{
			UserId:   md.UserId,
			Id_INT32: request.Id,
		}).To_Update())
		sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, resultUpdates)
	case model.PEER_CHAT:
		var (
			chat                    *model.MutableChat
			err                     error
			updateChatPinnedMessage *mtproto.Update
		)

		if request.Id == 0 {
			chat, err = s.ChatFacade.UpdateUnChatPinnedMessage(ctx, md.UserId, peer.PeerId)
			if err != nil {
				log.Errorf("messages.updatePinnedMessage - error: %v", err)
				return nil, err
			}

			updateChatPinnedMessage = mtproto.MakeTLUpdateChatPinnedMessage(&mtproto.Update{
				ChatId:   peer.PeerId,
				Id_INT32: 0,
				Version:  chat.Chat.Version,
			}).To_Update()
			resultUpdates = model.MakeUpdatesByUpdates(updateChatPinnedMessage)

			go func() {
				chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
					if userId != md.UserId {
						sync_client.PushUpdates(
							context.Background(),
							userId,
							model.MakeUpdatesByUpdates(updateChatPinnedMessage))
					}
					return nil
				})
			}()

		} else {
			pinnedList := map[int32]int32{md.UserId: request.Id}
			pinnedMsgList := s.MessageFacade.GetPeerChatMessageList(ctx, md.UserId, request.Id, peer.PeerId)
			for id, box := range pinnedMsgList {
				pinnedList[id] = box.Message.Id
			}
			chat, err = s.ChatFacade.UpdateChatPinnedMessage(ctx, md.UserId, peer.PeerId, pinnedList)
			if err != nil {
				log.Errorf("messages.updatePinnedMessage - error: %v", err)
				return nil, err
			}

			updateChatPinnedMessage = mtproto.MakeTLUpdateChatPinnedMessage(&mtproto.Update{
				ChatId:   peer.PeerId,
				Id_INT32: request.Id,
				Version:  chat.Chat.Version,
			}).To_Update()

			resultUpdates, _ = s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, &msgpb.OutboxMessage{
				NoWebpage:  false,
				Background: false,
				RandomId:   rand.Int63(),
				Message:    chat.MakePinnedMessageService(md.UserId, request.Id),
			})

			resultUpdates.Updates = append([]*mtproto.Update{updateChatPinnedMessage}, resultUpdates.Updates...)

			go func() {
				for k, v := range pinnedList {
					if k != md.UserId {
						sync_client.PushUpdates(
							context.Background(),
							k,
							model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateChatPinnedMessage(&mtproto.Update{
								ChatId:   peer.PeerId,
								Id_INT32: v,
								Version:  chat.Chat.Version,
							}).To_Update()))
					}
				}
			}()
		}
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, model.MakeUpdatesByUpdates(updateChatPinnedMessage))

	case model.PEER_CHANNEL:
		channel, err := s.ChannelFacade.UpdatePinnedMessage(ctx, peer.PeerId, md.UserId, request.Id)
		if err != nil {
			log.Errorf("messages.updatePinnedMessage - error: %v", err)
			return nil, err
		}

		updatePinnedMessage := mtproto.MakeTLUpdateChannelPinnedMessage(&mtproto.Update{
			ChannelId: peer.PeerId,
			Id_INT32:  request.Id,
		}).To_Update()
		updatePinnedMessage1 := mtproto.MakeTLUpdatePinnedChannelMessages(&mtproto.Update{
			Pinned:    !request.Unpin,
			ChannelId: peer.PeerId,
			Messages:  []int32{request.Id},
		}).To_Update()

		if request.Id == 0 {
			resultUpdates = model.MakeUpdatesByUpdates(updatePinnedMessage, updatePinnedMessage1)
		} else {
			resultUpdates, _ = s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, &msgpb.OutboxMessage{
				NoWebpage:    false,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      channel.MakeMessageService(md.UserId, false, request.Id, model.MakeMessageActionPinMessage()),
				ScheduleDate: nil,
			})
			resultUpdates.Updates = append([]*mtproto.Update{updatePinnedMessage, updatePinnedMessage1}, resultUpdates.Updates...)
		}

		go func() {
			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, model.MakeUpdatesByUpdates(updatePinnedMessage, updatePinnedMessage1))

			pushUpdates := model.MakeUpdatesByUpdates(updatePinnedMessage, updatePinnedMessage1)
			pIdList := s.ChannelFacade.GetChannelParticipantIdList(context.Background(), peer.PeerId)
			log.Infof("messages.updatePinnedMessage - pIdList: %#v", pIdList)
			for _, id := range pIdList {
				if id != md.UserId {
					sync_client.PushUpdates(context.Background(), id, pushUpdates)
				}
			}
		}()

	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.updatePinnedMessage - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.updatePinnedMessage#d2aaf7ec - reply: %s", resultUpdates.DebugString())
	return resultUpdates, nil
}
