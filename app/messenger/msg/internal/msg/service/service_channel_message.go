package service

import (
	"context"

	idgen "open.chat/app/service/idgen/client"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"math/rand"

	chats_core "open.chat/app/json/services/handler/chats/core"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const channel_Bot int32 = 136817688

func (s *Service) checkDmUser(ctx context.Context, sendUserId int32, channelId int32, dmUsers []int32) []int32 {
	var visibles model.IDList
	if len(dmUsers) == 0 {
		return visibles
	}
	chatsCore := chats_core.New(nil)
	bannedRights := chatsCore.GetChannelBannedRights(ctx, uint32(channelId))
	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, append(dmUsers, sendUserId)...)
	if err == nil {
		participant := channel.GetImmutableChannelParticipant(sendUserId)
		if (!bannedRights.BanWhisper && !bannedRights.BanSendDmMention) || sendUserId == channel_Bot || participant.IsCreatorOrAdmin() {
			for _, uid := range dmUsers {
				participant := channel.GetImmutableChannelParticipant(uid)
				if participant != nil && uid != sendUserId {
					visibles.AddIfNot(uid)
				}
			}
		}
	}
	return visibles
}

func (s *Service) SendChannelMessage(ctx context.Context, r *msgpb.ChannelMessage) (*mtproto.Updates, error) {
	log.Debugf("sendChannelMessage - %v", r)
	hasDuplicateMessage, err := s.MsgCore.HasDuplicateMessage(ctx, r.From.Id, r.RandomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := s.MsgCore.GetDuplicateMessage(ctx, r.From.Id, r.RandomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, r.PeerChannelId, r.From.Id)
	if err != nil {
		log.Errorf("not found channel by id: %d", r.PeerChannelId)
		return nil, err
	}

	iMsg := r.Message
	if channel.Channel.Signatures {
		fName, lName, err := s.UserFacade.GetFirstAndLastName(ctx, r.From.Id)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}

		var postAuthor string
		if fName != "" && lName != "" {
			postAuthor = fName + " " + lName
		} else {
			if fName != "" {
				postAuthor = fName
			} else if lName != "" {
				postAuthor = lName
			}
		}
		log.Debugf("post_author: {%s, %s, %s}", fName, lName, postAuthor)
		iMsg.PostAuthor = &types.StringValue{Value: postAuthor}
	}
	if channel.Channel.IsBroadcast() {
		iMsg.FromId_FLAGPEER = nil
		iMsg.Post = true
		iMsg.Views = &types.Int32Value{Value: 1}
		iMsg.Forwards = &types.Int32Value{Value: 1}
		if !channel.Channel.Signatures {
			iMsg.FromId_FLAGPEER = nil
		}
	}

	dmUsers := s.checkDmUser(ctx, r.From.Id, r.PeerChannelId, model.PickDmUsers(iMsg))
	outBox, err := s.MsgCore.SendChannelMessage(ctx, r.From.Id, r.PeerChannelId, r.RandomId, iMsg, dmUsers)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())

	syncNotMe := updatesHelper.ToSyncNotMeUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.SyncUpdatesNotMe(ctx, r.From.Id, r.From.AuthKeyId, syncNotMe)
	if err != nil {
		return nil, err
	}

	outBoxMessage := proto.Clone(outBox.Message).(*mtproto.Message)
	outBoxMessage.Out = false
	updateNewChannelMessage := mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBoxMessage),
	}).To_Update()

	go func() {
		var (
			ctx2                = context.Background()
			unreadMentionIdList []int32
			idList              []int32 = outBox.DmUsers
		)
		if len(outBox.DmUsers) == 0 {
			idList = s.ChannelFacade.GetChannelParticipantIdList(ctx2, r.PeerChannelId)
		}
		for _, id := range idList {
			if id != r.From.Id {
				var (
					pushUpdatesHelper *model.UpdatesHelper
				)

				outBoxMessage.MediaUnread = true
				outBoxMessage.Mentioned = model.CheckHasMention(outBoxMessage.Entities, id)

				if id == outBox.ReplyOwnerId {
					outBoxMessage.Mentioned = true
					outBoxMessage.MediaUnread = true
				}

				if outBoxMessage.Mentioned {
					unreadMentionIdList = append(unreadMentionIdList, id)
				}

				if outBoxMessage.GetAction().GetPredicateName() == mtproto.Predicate_messageActionChannelMigrateFrom {
					maxId2 := s.ChatFacade.GetTopMessage(ctx2, id, outBoxMessage.GetAction().GetChatId())
					pts3 := int32(idgen.NextPtsId(ctx2, id))
					updateReadHistoryInbox := mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
						FolderId:         nil,
						Peer_PEER:        model.MakePeerChat(outBoxMessage.GetAction().GetChatId()),
						MaxId:            maxId2,
						StillUnreadCount: 0,
						Pts_INT32:        pts3,
						PtsCount:         1,
					}).To_Update()

					pushUpdatesHelper = model.MakeUpdatesHelper(updateReadHistoryInbox)
				} else {
					pushUpdatesHelper = model.MakeUpdatesHelper(updateNewChannelMessage)
				}

				pushUpdates := pushUpdatesHelper.ToPushUpdates(ctx2, id, s.UserFacade, s.ChatFacade, s.ChannelFacade)

				if outBoxMessage.GetFwdFrom() != nil && outBoxMessage.GetFwdFrom().GetChannelId().GetValue() != 0 {
					pUser, _ := s.UserFacade.GetUserById(ctx2, id, channel_Bot)
					if pUser != nil {
						pushUpdates.Users = append(pushUpdates.Users, pUser)
					}
				}
				sync_client.PushUpdates(ctx2, id, pushUpdates)
			}
		}

		if len(unreadMentionIdList) > 0 {
			s.MsgCore.InsertChannelUnreadMentions(ctx2, r.PeerChannelId, unreadMentionIdList, outBox.MessageId)
		}
	}()

	go func() {
		if channel.Channel.Broadcast && channel.Channel.LinkedChatId > 0 {
			log.Debugf("linkedChatId - %v", channel.Channel)

			ctx2 := context.Background()

			outBoxMessage := proto.Clone(outBox.Message).(*mtproto.Message)
			outBoxMessage.ToId = model.MakePeerChannel(channel.Channel.LinkedChatId)
			outBoxMessage.Post = false
			if outBoxMessage.FwdFrom == nil {
				fwdFrom := mtproto.MakeTLMessageFwdHeader(&mtproto.MessageFwdHeader{
					Date: outBoxMessage.GetDate(),
				}).To_MessageFwdHeader()

				fwdFrom.ChannelId = &types.Int32Value{Value: channel.Channel.Id}
				fwdFrom.ChannelPost = &types.Int32Value{Value: outBoxMessage.Id}
				fwdFrom.PostAuthor = outBoxMessage.PostAuthor
				outBoxMessage.PostAuthor = nil
				outBoxMessage.FwdFrom = fwdFrom
			}
			outBoxMessage.FwdFrom.SavedFromPeer = model.MakePeerChannel(channel.Channel.Id)
			outBoxMessage.FwdFrom.SavedFromMsgId = &types.Int32Value{Value: outBoxMessage.Id}
			outBox, err := s.MsgCore.SendChannelMessage(ctx2, channel_Bot, channel.Channel.LinkedChatId, rand.Int63(), outBoxMessage, outBox.DmUsers)
			if err != nil {
				log.Error(err.Error())
				return
			}

			outBoxMessage.Out = false
			var idList []int32 = outBox.DmUsers
			if len(outBox.DmUsers) == 0 {
				idList = s.ChannelFacade.GetChannelParticipantIdList(ctx2, channel.Channel.LinkedChatId)
			}
			for _, id := range idList {
				pushUpdatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
					Pts_INT32:       outBox.Pts,
					PtsCount:        outBox.PtsCount,
					RandomId:        outBox.RandomId,
					Message_MESSAGE: model.MessageUpdate(outBoxMessage),
				}).To_Update())

				pushUpdates := pushUpdatesHelper.ToPushUpdates(ctx2, id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
				pushUpdates.Users = []*mtproto.User{}
				pUser, _ := s.UserFacade.GetUserById(ctx2, id, channel_Bot)
				if pUser != nil {
					pushUpdates.Users = append(pushUpdates.Users, pUser)
				}

				sync_client.PushUpdates(ctx2, id, pushUpdates)
			}
		}
	}()

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	s.MsgCore.PutDuplicateMessage(ctx, r.From.Id, r.RandomId, replyUpdates)

	return replyUpdates, nil
}

func (s *Service) SendChannelMultiMessage(ctx context.Context, r *msgpb.ChannelMultiMessage) (*mtproto.Updates, error) {
	log.Debugf("sendChannelMultiMessage - %v", r)

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, r.PeerChannelId, r.From.Id)
	if err != nil {
		log.Errorf("not found channel by id: %d", r.PeerChannelId)
		return nil, err
	}

	iMsgs := r.Message
	if channel.Channel.IsSignatures() {
		fName, lName, err := s.UserFacade.GetFirstAndLastName(ctx, r.From.Id)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}

		var postAuthor string
		if fName != "" && lName != "" {
			postAuthor = fName + " " + lName
		} else {
			if fName != "" {
				postAuthor = fName
			} else if lName != "" {
				postAuthor = lName
			}
		}
		log.Debugf("post_author: {%s, %s, %s}", fName, lName, postAuthor)
		for _, msg := range iMsgs {
			msg.PostAuthor = &types.StringValue{Value: postAuthor}
		}
	}
	if channel.Channel.IsBroadcast() {
		for _, msg := range iMsgs {
			msg.FromId_FLAGPEER = nil
			msg.Post = true
			msg.Views = &types.Int32Value{Value: 1}
			msg.Forwards = &types.Int32Value{Value: 1}
			if channel.Channel.IsSignatures() {
				msg.FromId_FLAGPEER = nil
			}
		}
	}

	updatesHelper := model.MakeUpdatesHelper()
	pushUpdates := make([]*mtproto.Update, 0, len(iMsgs))
	for i := 0; i < len(iMsgs); i++ {
		dmUsers := s.checkDmUser(ctx, r.From.Id, r.PeerChannelId, model.PickDmUsers(iMsgs[i]))
		outBoxMsg, err := s.MsgCore.SendChannelMessage(ctx, r.From.Id, r.PeerChannelId, r.RandomId[i], iMsgs[i], dmUsers)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		updatesHelper.PushBackUpdate(mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
			Pts_INT32:       outBoxMsg.Pts,
			PtsCount:        outBoxMsg.PtsCount,
			RandomId:        outBoxMsg.RandomId,
			Message_MESSAGE: model.MessageUpdate(outBoxMsg.Message),
		}).To_Update())

		inboxMsg := proto.Clone(outBoxMsg.Message).(*mtproto.Message)
		inboxMsg.Out = false
		pushUpdates = append(pushUpdates, mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
			Pts_INT32:       outBoxMsg.Pts,
			PtsCount:        outBoxMsg.PtsCount,
			RandomId:        outBoxMsg.RandomId,
			Message_MESSAGE: model.MessageUpdate(inboxMsg),
			Users:           outBoxMsg.DmUsers,
		}).To_Update())
	}

	err = sync_client.SyncUpdatesNotMe(ctx,
		r.From.Id,
		r.From.AuthKeyId,
		updatesHelper.ToSyncNotMeUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		allParticipant := s.ChannelFacade.GetChannelParticipantIdList(ctx2, r.PeerChannelId)
		updatesHelpers := map[int32]*model.UpdatesHelper{}
		for _, up := range pushUpdates {
			idList := up.GetUsers()
			if len(idList) == 0 {
				idList = allParticipant
			}
			for _, id := range idList {
				if id != r.From.Id {
					pushUpdatesHelper, ok := updatesHelpers[id]
					if !ok {
						updatesHelpers[id] = model.MakeUpdatesHelper(up)
					} else {
						pushUpdatesHelper.PushBackUpdate(up)
					}
				}
			}
		}
		for uid, upHelper := range updatesHelpers {
			pushUpdates := upHelper.ToPushUpdates(ctx2, uid, s.UserFacade, s.ChatFacade, s.ChannelFacade)
			pUser, _ := s.UserFacade.GetUserById(ctx2, uid, channel_Bot)
			if pUser != nil {
				pushUpdates.Users = append(pushUpdates.Users, pUser)
			}
			sync_client.PushUpdates(ctx2, uid, pushUpdates)
		}
	}()

	go func() {
		if channel.Channel.Broadcast && channel.Channel.LinkedChatId > 0 {
			log.Debugf("linkedChatId - %v", channel.Channel)

			ctx2 := context.Background()

			discussUpdates := make([]*mtproto.Update, 0, len(iMsgs))
			for i := 0; i < len(iMsgs); i++ {
				outBoxMessage := proto.Clone(iMsgs[i]).(*mtproto.Message)
				outBoxMessage.ToId = model.MakePeerChannel(channel.Channel.LinkedChatId)
				outBoxMessage.Post = false
				if outBoxMessage.FwdFrom == nil {
					fwdFrom := mtproto.MakeTLMessageFwdHeader(&mtproto.MessageFwdHeader{
						Date: outBoxMessage.GetDate(),
					}).To_MessageFwdHeader()

					fwdFrom.ChannelId = &types.Int32Value{Value: channel.Channel.Id}
					fwdFrom.ChannelPost = &types.Int32Value{Value: outBoxMessage.Id}
					fwdFrom.PostAuthor = outBoxMessage.PostAuthor
					outBoxMessage.PostAuthor = nil
					outBoxMessage.FwdFrom = fwdFrom
				}
				outBoxMessage.FwdFrom.SavedFromPeer = model.MakePeerChannel(channel.Channel.Id)
				outBoxMessage.FwdFrom.SavedFromMsgId = &types.Int32Value{Value: outBoxMessage.Id}
				dmUsers := s.checkDmUser(ctx, channel_Bot, channel.Channel.LinkedChatId, model.PickDmUsers(outBoxMessage))
				outBox, err := s.MsgCore.SendChannelMessage(ctx2, channel_Bot, channel.Channel.LinkedChatId, rand.Int63(), outBoxMessage, dmUsers)
				if err != nil {
					log.Error(err.Error())
					return
				}

				outBoxMessage.Out = false
				discussUpdates = append(discussUpdates, mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
					Pts_INT32:       outBox.Pts,
					PtsCount:        outBox.PtsCount,
					RandomId:        outBox.RandomId,
					Message_MESSAGE: model.MessageUpdate(outBoxMessage),
					Users:           outBox.DmUsers,
				}).To_Update())
			}

			allParticipant := s.ChannelFacade.GetChannelParticipantIdList(ctx2, channel.Channel.LinkedChatId)
			updatesHelpers := map[int32]*model.UpdatesHelper{}
			for _, up := range discussUpdates {
				idList := up.GetUsers()
				if len(idList) == 0 {
					idList = allParticipant
				}
				for _, id := range idList {
					if id != r.From.Id {
						pushUpdatesHelper, ok := updatesHelpers[id]
						if !ok {
							updatesHelpers[id] = model.MakeUpdatesHelper(up)
						} else {
							pushUpdatesHelper.PushBackUpdate(up)
						}
					}
				}
			}
			for uid, upHelper := range updatesHelpers {
				pushUpdates := upHelper.ToPushUpdates(ctx2, uid, s.UserFacade, s.ChatFacade, s.ChannelFacade)
				pushUpdates.Users = []*mtproto.User{}
				pUser, _ := s.UserFacade.GetUserById(ctx2, uid, channel_Bot)
				if pUser != nil {
					pushUpdates.Users = append(pushUpdates.Users, pUser)
				}
				sync_client.PushUpdates(ctx2, uid, pushUpdates)
			}
		}
	}()

	return updatesHelper.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade), nil
}

func (s *Service) EditChannelMessage(ctx context.Context, r *msgpb.ChannelMessage) (*mtproto.Updates, error) {
	return nil, nil
}

func (s *Service) sendChannelOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	users := s.UserFacade.GetMutableUsers(ctx, r.From.Id)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}

	reply, err = s.SendChannelMessage(ctx, &msgpb.ChannelMessage{
		From:          r.From,
		PeerChannelId: r.PeerId,
		RandomId:      r.Message.RandomId,
		Message:       r.Message.Message,
	})
	return
}

func (s *Service) sendChannelOutgoingMultiMessage(ctx context.Context, r *msgpb.OutgoingMultiMessage) (reply *mtproto.Updates, err error) {
	users := s.UserFacade.GetMutableUsers(ctx, r.From.Id)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}

	multiRequest := &msgpb.ChannelMultiMessage{
		From:          r.From,
		PeerChannelId: r.PeerId,
		RandomId:      make([]int64, 0, len(r.MultiMessage)),
		Message:       make([]*mtproto.Message, 0, len(r.MultiMessage)),
	}
	for _, m := range r.MultiMessage {
		multiRequest.RandomId = append(multiRequest.RandomId, m.RandomId)
		multiRequest.Message = append(multiRequest.Message, m.Message)
	}
	reply, err = s.SendChannelMultiMessage(ctx, multiRequest)
	return
}

func (s *Service) editChannelOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	outBox, err2 := s.MsgCore.EditChannelOutboxMessage(ctx, r.From.Id, r.PeerId, r.Message.Message)
	if err2 != nil {
		err = err2
		return
	}

	updateEditChannelMessage := mtproto.MakeTLUpdateEditChannelMessage(&mtproto.Update{
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
		Pts_INT32:       outBox.Pts,
		PtsCount:        1,
	}).To_Update()
	meUpdates := model.MakeUpdatesHelper(updateEditChannelMessage)

	go func() {
		ctx2 := context.Background()
		sync_client.SyncUpdatesNotMe(ctx2, r.From.Id, r.From.AuthKeyId, meUpdates.ToSyncNotMeUpdates(
			context.Background(),
			r.From.Id,
			s.UserFacade,
			s.ChatFacade,
			s.ChannelFacade))
		idList := s.ChannelFacade.GetChannelParticipantIdList(ctx2, r.PeerId)
		for _, id := range idList {
			if id != r.From.Id {
				pushUpdates := model.MakeUpdatesHelper(updateEditChannelMessage)
				sync_client.PushUpdates(ctx2, id, pushUpdates.ToPushUpdates(ctx2, id, s.UserFacade, s.ChatFacade, s.ChannelFacade))
			}
		}
	}()

	reply = meUpdates.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	return
}

func (s *Service) deleteChannelUserHistory(ctx context.Context, r *msgpb.DeleteHistoryRequest) (reply *mtproto.Messages_AffectedHistory, err error) {
	var (
		pts, ptsCount int32
	)

	deleteIds := s.GetChannelMessageIdListBySenderUserId(ctx, r.ChannelId, r.PeerId)
	if len(deleteIds) == 0 {
		pts = int32(idgen.CurrentChannelPtsId(ctx, r.ChannelId))
		ptsCount = 0
		return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
			Pts:      pts,
			PtsCount: ptsCount,
			Offset:   0,
		}).To_Messages_AffectedHistory(), nil
	}

	if pts, err = s.MsgCore.DeleteChannelMessages(ctx, r.ChannelId, deleteIds); err != nil {
		return nil, err
	}
	ptsCount = int32(len(deleteIds))

	updateDeleteChannelMessages := mtproto.MakeTLUpdateDeleteChannelMessages(&mtproto.Update{
		ChannelId: r.ChannelId,
		Messages:  deleteIds,
		Pts_INT32: pts,
		PtsCount:  ptsCount,
	}).To_Update()

	go func() {
		ctx2 := context.Background()
		sync_client.SyncUpdatesNotMe(ctx2, r.From.Id, r.From.AuthKeyId, model.MakeUpdatesByUpdates(updateDeleteChannelMessages))
		sync_client.BroadcastChannelUpdates(ctx2, r.ChannelId, model.MakeUpdatesByUpdates(updateDeleteChannelMessages), r.From.Id)
	}()

	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

func (s *Service) deleteChannelUserMessages(ctx context.Context, r *msgpb.DeleteMessagesRequest) (reply *mtproto.Messages_AffectedMessages, err error) {
	var (
		pts      int32
		ptsCount int32 = int32(len(r.Id))
	)

	if r.Revoke {
		if pts, err = s.MsgCore.DeleteChannelMessages(ctx, r.PeerId, r.Id); err != nil {
			return nil, err
		}
	} else {
		if pts, err = s.MsgCore.DeleteMessagesJustSelf(ctx, r.From.Id, r.PeerId, r.Id); err != nil {
			return nil, err
		}
	}

	updateDeleteChannelMessages := mtproto.MakeTLUpdateDeleteChannelMessages(&mtproto.Update{
		ChannelId: r.PeerId,
		Messages:  r.Id,
		Pts_INT32: pts,
		PtsCount:  ptsCount,
	}).To_Update()

	go func(revoke bool) {
		ctx2 := context.Background()
		sync_client.SyncUpdatesNotMe(ctx2, r.From.Id, r.From.AuthKeyId, model.MakeUpdatesByUpdates(updateDeleteChannelMessages))
		if revoke {
			sync_client.BroadcastChannelUpdates(ctx2, r.PeerId, model.MakeUpdatesByUpdates(updateDeleteChannelMessages), r.From.Id)
		}
	}(r.Revoke)

	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages(), nil
}
