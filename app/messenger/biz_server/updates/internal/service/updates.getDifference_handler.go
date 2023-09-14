package service

import (
	"context"
	sync_client "open.chat/app/messenger/sync/client"
	"time"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("updates.getDifference#25939651 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var ptsTotalLimit = request.GetPtsTotalLimit().GetValue()

	if ptsTotalLimit <= 0 {
		ptsTotalLimit = 1000
	}

	difference, err := s.UpdatesFacade.GetDifference(ctx, md.AuthId, md.UserId, request.Pts, ptsTotalLimit)
	if err != nil {
		log.Errorf("sync.getDifference error - %v", err)
		return nil, err
	}

	lastDate, channelDifference, err := s.ChannelCore.GetDifference(ctx, md.UserId, request.Date)
	if err != nil {
		log.Errorf("sync.getDifference error - %v", err)
		return nil, err
	}

	lastQts, eIdList, newEncryptedMessages, err := s.SecretChatCore.GetDifference(ctx, md.AuthId, request.Qts)
	if err != nil {
		log.Errorf("sync.getDifference error - %v", err)
		return nil, err
	}

	if difference.PredicateName == mtproto.Predicate_updates_differenceEmpty &&
		len(channelDifference) == 0 &&
		len(newEncryptedMessages) == 0 {
		difference.Date = request.Date
	} else {
		messageList := difference.NewMessages
		for _, upd := range channelDifference {
			if upd.PredicateName == mtproto.Predicate_updateNewChannelMessage {
				msg := upd.GetMessage_MESSAGE()
				if msg.PredicateName == mtproto.Predicate_message {
					msg.Mentioned = model.CheckHasMention(msg.Entities, md.UserId)
				}
				messageList = append(messageList, msg)
			}
		}

		if len(messageList) > 0 {
			difference.OtherUpdates = append(difference.OtherUpdates, channelDifference...)
			difference.Date = lastDate

			userIdList, chatIdList, channelIdList := model.PickAllIdListByMessages(messageList)
			userIdList = append(userIdList, eIdList...)
			userList := s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList)
			chatList := s.ChatFacade.GetChatListByIdList(ctx, md.UserId, chatIdList)
			channelList := s.ChannelCore.GetChannelListByIdList(ctx, md.UserId, channelIdList)
			difference.Users = userList
			difference.Chats = append(chatList, channelList...)
		}

		if len(newEncryptedMessages) > 0 {
			difference.NewEncryptedMessages = newEncryptedMessages
			difference.State.Qts = lastQts
		}
	}

	go func() {
		ctx2 := context.Background()
		s.sendUpdateEncryption(ctx2, md.UserId)
	}()

	log.Debugf("updates.getDifference#25939651 - reply: %s", logger.JsonDebugData(difference))
	return difference, nil
}

func (s *Service) sendUpdateEncryption(ctx context.Context, selfId int32) {
	s.sendEncryptionRequest(ctx, selfId)
	s.sendEncryptionDiscard(ctx, selfId)
}

func (s *Service) sendEncryptionRequest(ctx context.Context, selfId int32) {
	l, err := s.SecretChatCore.GetRequested(ctx, selfId)
	if err != nil {
		return
	}

	pushUpdates := model.NewUpdatesLogic(selfId)

	uidList := make([]int32, 0)
	for _, data := range l {
		chat, _ := data.ToEncryptedChatRequested()
		updateRequestedEncryption := &mtproto.TLUpdateEncryption{Data2: &mtproto.Update{
			Date: int32(time.Now().Unix()),
			Chat: chat,
		}}
		pushUpdates.AddUpdate(updateRequestedEncryption.To_Update())
		uidList = append(uidList, data.AdminId)
	}

	uidList = append(uidList, selfId)
	users := s.UserFacade.GetUserListByIdList(ctx, selfId, uidList)
	pushUpdates.AddUsers(users)
	sync_client.PushUpdates(ctx, selfId, pushUpdates.ToUpdates())
}

func (s *Service) sendEncryptionDiscard(ctx context.Context, selfId int32) {
	chatIds, err := s.SecretChatCore.GetPendingClosed(ctx, selfId)
	if err != nil {
		log.Errorf("sendEncryptionDiscard failed, error:%v", err)
		return
	}

	for _, id := range chatIds {
		encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, id, 0)
		if err != nil {
			log.Errorf("sendEncryptionDiscard, secret chat id:%d, error: {%v}", id, err)
			continue
		}

		pushUpdates := model.NewUpdatesLogic(encryptChatData.AdminId)
		updateRequestedEncryption := &mtproto.TLUpdateEncryption{Data2: &mtproto.Update{
			Date: int32(time.Now().Unix()),
			Chat: &mtproto.EncryptedChat{
				Constructor: mtproto.CRC32_encryptedChatDiscarded,
				Id:          encryptChatData.Id,
			},
		}}
		pushUpdates.AddUpdate(updateRequestedEncryption.To_Update())

		users := s.UserFacade.GetUserListByIdList(ctx, selfId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx, selfId, pushUpdates.ToUpdates())
	}
}
