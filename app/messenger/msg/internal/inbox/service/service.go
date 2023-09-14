package service

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	chats_core "open.chat/app/json/services/handler/chats/core"
	msg_core "open.chat/app/messenger/msg/internal/core"
	msg_dao "open.chat/app/messenger/msg/internal/dao"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	_ "open.chat/app/service/biz_service/channel/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	_ "open.chat/app/service/biz_service/chat/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

// Service is a service.
type Service struct {
	conf     *Config
	consumer *databus.Databus
	user_client.UserFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	*msg_core.MsgCore
	// *chat_core.ChatCore
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a service and return.
func New() *Service {
	var (
		err error
		ac  = &Config{}
		s   = new(Service)
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s.conf = ac
	s.consumer = databus.New(ac.Databus)

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)

	s.MsgCore = msg_core.New(msg_dao.New())

	sync_client.New()
	go s.consume()
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	// return s.Dao.Ping(ctx)
	return nil
}

// Close close the resources.
func (s *Service) Close() error {
	if err := s.consumer.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Service) consume() {
	msgs := s.consumer.Messages()
	for {
		msg, ok := <-msgs
		if !ok {
			log.Warn("[job] consumer has been closed")
			return
		}
		if msg.Topic != s.conf.Databus.Topic {
			log.Error("unknown message:%v", msg)
			continue
		}

		s.onInboxData(context.Background(), msg.Key, msg.Value)

		msg.Commit()
	}
}

func (s *Service) onInboxData(ctx context.Context, key string, value []byte) error {
	log.Debugf("recv {key: %s, value: %s", key, string(value))

	// var err error
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("handle panic: %s", debug.Stack())
		}
	}()

	switch key {
	case proto.MessageName((*msgpb.InboxUserMessage)(nil)):
		r := new(msgpb.InboxUserMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.SendUserMessageToInbox(ctx, r.From.Id, r.PeerUserId, r.DialogMessageId, r.MessageDataId, r.RandomId, r.Message)
	case proto.MessageName((*msgpb.InboxChatMessage)(nil)):
		r := new(msgpb.InboxChatMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.SendChatMessageToInbox(ctx, r.From.Id, r.PeerChatId, r.DialogMessageId, r.MessageDataId, r.RandomId, r.Message)
	case proto.MessageName((*msgpb.InboxUserMultiMessage)(nil)):
		r := new(msgpb.InboxUserMultiMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.SendUserMultiMessageToInbox(ctx, r.From.Id, r.PeerUserId, r.DialogMessageId, r.MessageDataId, r.RandomId, r.Message)
	case proto.MessageName((*msgpb.InboxChatMultiMessage)(nil)):
		r := new(msgpb.InboxChatMultiMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.SendChatMultiMessageToInbox(ctx, r.From.Id, r.PeerChatId, r.DialogMessageId, r.MessageDataId, r.RandomId, r.Message)
	case proto.MessageName((*msgpb.InboxUserEditMessage)(nil)):
		r := new(msgpb.InboxUserEditMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.EditUserInboxMessage(ctx, r.From.Id, r.PeerUserId, r.Message)
	case proto.MessageName((*msgpb.InboxChatEditMessage)(nil)):
		r := new(msgpb.InboxChatEditMessage)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.EditChatInboxMessage(ctx, r.From.Id, r.PeerChatId, r.Message)
	case proto.MessageName((*msgpb.InboxDeleteMessages)(nil)):
		r := new(msgpb.InboxDeleteMessages)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.DeleteInboxMessages(ctx, r.From.Id, r.Id)
	case proto.MessageName((*msgpb.InboxUserDeleteHistory)(nil)):
		r := new(msgpb.InboxUserDeleteHistory)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.DeleteUserInboxHistory(ctx, r.From.Id, r.PeerUserId, r.JustClear, r.MaxId)
	case proto.MessageName((*msgpb.InboxChatDeleteHistory)(nil)):
		r := new(msgpb.InboxChatDeleteHistory)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.DeleteChatInboxHistory(ctx, r.From.Id, r.PeerChatId, r.MaxId)
	case proto.MessageName((*msgpb.InboxUserReadMediaUnread)(nil)):
		r := new(msgpb.InboxUserReadMediaUnread)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.ReadUserMediaUnread(ctx, r.From.Id, r.Id)
	case proto.MessageName((*msgpb.InboxChatReadMediaUnread)(nil)):
		r := new(msgpb.InboxChatReadMediaUnread)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.ReadChatMediaUnread(ctx, r.From.Id, r.PeerChatId, r.Id)
	default:
		err := fmt.Errorf("invalid key: %s", key)
		log.Error(err.Error())
		return err
	}
}

func (s *Service) checkChatDmUser(ctx context.Context, sendUserId int32, chatId int32, dmUsers []int32) []int32 {
	var visibles model.IDList
	chat, err := s.ChatFacade.GetMutableChat(ctx, chatId, append(dmUsers, sendUserId)...)
	if err == nil {
		participant := chat.GetImmutableChatParticipant(sendUserId)
		chatsCore := chats_core.New(nil)
		bannedRights := chatsCore.GetChatBannedRights(ctx, uint32(chatId))
		if (!bannedRights.BanWhisper && !bannedRights.BanSendDmMention) || participant.IsChatCreatorOrAdmin() {
			for _, uid := range dmUsers {
				participant = chat.GetImmutableChatParticipant(uid)
				if participant != nil && uid != sendUserId {
					visibles.AddIfNot(uid)
				}
			}
		}
	}
	return visibles
}

func (s *Service) SendUserMessageToInbox(ctx context.Context, fromId, toId, dialogMessageId int32, messageDataId int64, clientRandomId int64, message *mtproto.Message) error {
	if fromId == toId {
		log.Warn("sendToSelfUser")
		return nil
	}

	inBox, err := s.MsgCore.SendUserMessageToInbox(ctx, fromId, toId, dialogMessageId, messageDataId, clientRandomId, message)
	if err != nil {
		return err
	}

	if inBox.DialogMessageId == 1 && (fromId != 42777 && fromId != 424000) {
		isContact, _ := s.UserFacade.GetContactAndMutual(ctx, toId, fromId)
		if !isContact {
			s.UserFacade.AddPeerSettings(ctx, toId, model.MakeUserPeerUtil(fromId), &mtproto.PeerSettings{
				ReportSpam:   true,
				AddContact:   true,
				BlockContact: true,
			})
		}
	}

	pushUpdates := s.makeUpdateNewMessageListUpdates(ctx, toId, []*model.MessageBox{inBox})

	var isBot = false
	for _, u := range pushUpdates.GetUsers() {
		if u.GetId() == toId {
			isBot = u.GetBot()
			break
		}
	}
	if isBot {
		err = sync_client.PushBotUpdates(ctx, inBox.SelfUserId, pushUpdates.To_Updates())
	} else {
		err = sync_client.PushUpdates(ctx, inBox.SelfUserId, pushUpdates.To_Updates())
	}
	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func (s *Service) SendChatMessageToInbox(ctx context.Context, fromId, chatId, dialogMessageId int32, messageDataId, clientRandomId int64, message *mtproto.Message) error {
	chatUserIdList, err := s.ChatFacade.GetChatParticipantIdList(ctx, chatId)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	if message.Action != nil &&
		message.Action.PredicateName == mtproto.Predicate_messageActionChatDeleteUser {
		chatUserIdList = append(chatUserIdList, message.Action.UserId)
	}

	var (
		migrateBox           *model.MessageBox
		updateChannel        *mtproto.Update
		updateNotifySettings *mtproto.Update
	)
	if message.GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
		migrateBox, _ = s.GetLastChannelMessage(ctx, fromId, message.GetAction().GetChannelId())
		migrateBox.PtsCount = 1
		migrateBox.Pts = int32(idgen.CurrentChannelPtsId(ctx, message.GetAction().GetChannelId()))

		updateChannel = mtproto.MakeTLUpdateChannel(&mtproto.Update{
			ChannelId: message.GetAction().GetChannelId(),
		}).To_Update()

		updateNotifySettings = mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
			Peer_NOTIFYPEER: mtproto.MakeTLNotifyPeer(&mtproto.NotifyPeer{
				Peer: mtproto.MakeTLPeerChannel(&mtproto.Peer{
					ChannelId: message.GetAction().GetChannelId(),
				}).To_Peer(),
			}).To_NotifyPeer(),
			NotifySettings: model.MakeDefaultPeerNotifySettings(model.PEER_CHANNEL),
		}).To_Update()

	}
	dmUsers := s.checkChatDmUser(ctx, fromId, chatId, model.PickDmUsers(message))
	if len(dmUsers) > 0 {
		chatUserIdList = dmUsers
	}
	for _, toId := range chatUserIdList {
		if toId == fromId {
			continue
		}

		inBox, err := s.MsgCore.SendChatMessageToInbox(ctx, fromId, chatId, toId, dialogMessageId, messageDataId, clientRandomId, message)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		var pushUpdates *mtproto.Updates

		if message.GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
			pushUpdates = s.makeUpdateNewMessageListUpdates(ctx, toId, []*model.MessageBox{migrateBox, inBox}).To_Updates()
			pushUpdates.Updates = append([]*mtproto.Update{updateChannel, updateNotifySettings}, pushUpdates.Updates...)
			pushUpdates.Updates = append(pushUpdates.Updates,
				mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
					FolderId:         nil,
					Peer_PEER:        model.MakePeerChat(chatId),
					MaxId:            inBox.MessageId,
					StillUnreadCount: 0,
					Pts_INT32:        int32(idgen.NextPtsId(context.Background(), toId)),
					PtsCount:         1,
				}).To_Update())
		} else {
			pushUpdates = s.makeUpdateNewMessageListUpdates(ctx, toId, []*model.MessageBox{inBox}).To_Updates()
		}

		err = sync_client.PushUpdates(ctx, toId, pushUpdates)
		if err != nil {
			log.Error(err.Error())
		}
	}

	return err
}

func (s *Service) SendUserMultiMessageToInbox(ctx context.Context, fromId, toId int32, dialogMessageId []int32, messageDataId, clientRandomId []int64, messages []*mtproto.Message) error {
	if fromId == toId {
		log.Warn("sendToSelfUser")
		return nil
	}

	inBoxList, err := s.MsgCore.SendUserMultiMessageToInbox(ctx, fromId, toId, dialogMessageId, messageDataId, clientRandomId, messages)
	if err != nil {
		return err
	}

	pushUpdates := s.makeUpdateNewMessageListUpdates(ctx, toId, inBoxList)
	err = sync_client.PushUpdates(ctx, toId, pushUpdates.To_Updates())
	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func (s *Service) SendChatMultiMessageToInbox(ctx context.Context, fromId, chatId int32, dialogMessageId []int32, messageDataId, clientRandomId []int64, messages []*mtproto.Message) error {
	chatUserIdList, err := s.GetChatParticipantIdList(ctx, chatId)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	for _, toId := range chatUserIdList {
		if toId == fromId {
			continue
		}

		inBoxList, err := s.MsgCore.SendChatMultiMessageToInbox(ctx, fromId, chatId, toId, dialogMessageId, messageDataId, clientRandomId, messages)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		pushUpdates := s.makeUpdateNewMessageListUpdates(ctx, toId, inBoxList)
		err = sync_client.PushUpdates(ctx, toId, pushUpdates.To_Updates())
		if err != nil {
			log.Error(err.Error())
		}
	}

	return err
}

func (s *Service) makeUpdateNewMessageListUpdates(ctx context.Context, selfUserId int32, boxList []*model.MessageBox) *mtproto.TLUpdates {
	var (
		messages      = make([]*mtproto.Message, 0, len(boxList))
		updateNewList = make([]*mtproto.Update, 0, len(boxList))
	)

	var users model.IDList
	for _, box := range boxList {
		if box == nil {
			continue
		}
		m := model.MessageUpdate(box.ToMessage(selfUserId))
		messages = append(messages, m)
		u := &mtproto.Update{
			Message_MESSAGE: m,
			Pts_INT32:       box.Pts,
			PtsCount:        box.PtsCount,
		}
		if box.MessageBoxType == model.MESSAGE_BOX_TYPE_CHANNEL {
			updateNewList = append(updateNewList, mtproto.MakeTLUpdateNewChannelMessage(u).To_Update())
		} else {
			updateNewList = append(updateNewList, mtproto.MakeTLUpdateNewMessage(u).To_Update())
		}
		if box.ReplyOwnerId != 0 {
			users.AddIfNot(box.ReplyOwnerId)
		}
	}

	userIdList, chatIdList, channelIdLIst := model.PickAllIdListByMessages(messages)
	userList := s.UserFacade.GetUserListByIdList(ctx, selfUserId, userIdList)
	chatList := s.ChatFacade.GetChatListByIdList(ctx, selfUserId, chatIdList)
	chatList = append(chatList, s.ChannelFacade.GetChannelListByIdList(ctx, selfUserId, channelIdLIst...)...)

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updateNewList,
		Users:   userList,
		Chats:   chatList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	})
}

func (s *Service) EditUserInboxMessage(ctx context.Context, fromId, toId int32, message *mtproto.Message) error {
	inBox, err := s.MsgCore.EditUserInboxMessage(ctx, fromId, toId, message)
	if err != nil {
		log.Errorf("editUserInboxMessage - error: %v", err)
		return err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
		Pts_INT32:       inBox.Pts,
		PtsCount:        inBox.PtsCount,
		Message_MESSAGE: model.MessageUpdate(inBox.Message),
	}).To_Update())
	pushUpdates := updatesHelper.ToPushUpdates(ctx, toId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	log.Debugf("pushUpdates - %s", pushUpdates.DebugString())
	return sync_client.PushUpdates(ctx, toId, pushUpdates)
}

func (s *Service) EditChatInboxMessage(ctx context.Context, fromId, chatId int32, message2 *mtproto.Message) error {
	chatUserIdList, err := s.GetChatParticipantIdList(ctx, chatId)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	for _, toId := range chatUserIdList {
		message := proto.Clone(message2).(*mtproto.Message)
		if toId == fromId {
			continue
		}

		inBox, err := s.MsgCore.EditUserInboxMessage(ctx, fromId, toId, message)
		if err != nil {
			log.Errorf("editUserInboxMessage - error: %v", err)
			return err
		} else if inBox == nil {
			log.Errorf("editUserInboxMessage - error: {from_id: %d, peer_id:%d}", fromId, toId)
			continue
		}

		updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
			Pts_INT32:       inBox.Pts,
			PtsCount:        inBox.PtsCount,
			Message_MESSAGE: model.MessageUpdate(inBox.Message),
		}).To_Update())
		pushUpdates := updatesHelper.ToPushUpdates(ctx, toId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
		log.Debugf("pushUpdates - %s", pushUpdates.DebugString())
		sync_client.PushUpdates(ctx, toId, pushUpdates)
	}
	return nil
}

func (s *Service) DeleteInboxMessages(ctx context.Context, fromId int32, id []int64) error {
	s.MsgCore.DeleteInboxMessages(ctx, fromId, id, func(ctx context.Context, userId int32, idList []int32) {
		sync_client.PushUpdates(ctx, userId, model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
			Messages:  idList,
			Pts_INT32: int32(idgen.NextNPtsId(ctx, userId, len(idList))),
			PtsCount:  int32(len(idList)),
		}).To_Update()))
	})
	return nil
}

func (s *Service) DeleteChatInboxMessages(ctx context.Context, fromId, chatId int32, id []int32) error {
	return nil
}

func (s *Service) DeleteUserInboxHistory(ctx context.Context, fromId, toId int32, justClear bool, maxId int32) error {
	var (
		pts, ptsCount int32
		peer          = &model.PeerUtil{PeerType: model.PEER_USER, PeerId: fromId}
	)

	lastMessageBox, deleteIds := s.MsgCore.GetLastMessageAndIdListByDialog(ctx, toId, peer)
	if len(deleteIds) == 0 ||
		len(deleteIds) == 1 &&
			lastMessageBox.Message.PredicateName == mtproto.Predicate_messageService &&
			lastMessageBox.Message.Action.PredicateName == mtproto.Predicate_messageActionHistoryClear {
		return nil
	}

	pts = int32(idgen.NextNPtsId(ctx, toId, len(deleteIds)+1))
	ptsCount = int32(len(deleteIds) + 1)

	if justClear {
		deleteIds = deleteIds[1:]
		if _, err := s.MsgCore.DeleteByMessageIdList(ctx, toId, deleteIds); err != nil {
			return err
		}

		clearHistoryMessage := mtproto.MakeTLMessageService(&mtproto.Message{
			Out:             true,
			Id:              lastMessageBox.MessageId,
			FromId_FLAGPEER: model.MakePeerUser(toId),
			PeerId:          model.MakePeerUser(fromId),
			ToId:            model.MakePeerUser(fromId),
			Date:            lastMessageBox.Message.GetDate(),
			Action:          mtproto.MakeTLMessageActionHistoryClear(nil).To_MessageAction(),
		}).To_Message()
		s.MsgCore.EditUserOutboxMessage(ctx, toId, fromId, clearHistoryMessage)
		updatesHelper := model.MakeUpdatesHelper(
			mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  deleteIds,
				Pts_INT32: pts - 2,
				PtsCount:  ptsCount - 2,
			}).To_Update(),
			mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
				Peer_PEER: peer.ToPeer(),
				MaxId:     lastMessageBox.MessageId,
				Pts_INT32: pts - 1,
				PtsCount:  1,
			}).To_Update(),
			mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
				Message_MESSAGE: model.MessageUpdate(clearHistoryMessage),
				Pts_INT32:       pts,
				PtsCount:        1,
			}).To_Update(),
		)
		sync_client.PushUpdates(ctx, toId, updatesHelper.ToPushUpdates(
			ctx,
			toId,
			s.UserFacade,
			s.ChatFacade,
			s.ChannelFacade))
	} else {
		if _, err := s.MsgCore.DeleteByMessageIdList(ctx, toId, deleteIds); err != nil {
			return err
		}

		updatesHelper := model.MakeUpdatesHelper(
			mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  deleteIds,
				Pts_INT32: pts - 2,
				PtsCount:  ptsCount - 2,
			}).To_Update(),
			mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
				Peer_PEER: peer.ToPeer(),
				MaxId:     lastMessageBox.MessageId,
				Pts_INT32: pts - 1,
				PtsCount:  1,
			}).To_Update(),
		)
		sync_client.PushUpdates(ctx, toId, updatesHelper.ToPushUpdates(
			ctx,
			toId,
			s.UserFacade,
			s.ChatFacade,
			s.ChannelFacade,
		))
	}
	return nil
}

func (s *Service) DeleteChatInboxHistory(ctx context.Context, fromId, chatId, maxId int32) error {
	return nil
}

func (s *Service) ReadUserMediaUnread(ctx context.Context, fromId int32, id []int32) error {
	peerIdList := s.MsgCore.GetPeerDialogMessageIdList(ctx, fromId, id)

	for peerId, idList := range peerIdList {
		for _, id2 := range idList {
			s.UpdateMediaUnread(ctx, peerId, id2)
		}

		if len(idList) > 0 {
			ptsCount := int32(len(idList))
			pts := int32(idgen.NextNPtsId(ctx, peerId, len(idList))) - ptsCount + 1
			sync_client.PushUpdates(ctx,
				peerId,
				model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
					Messages:  idList,
					Pts_INT32: pts,
					PtsCount:  ptsCount,
				}).To_Update()))
		}
	}
	return nil
}

func (s *Service) ReadChatMediaUnread(ctx context.Context, fromId, chatId int32, id []int32) error {
	peerIdList := s.MsgCore.GetPeerDialogMessageIdList(ctx, fromId, id)

	for peerId, idList := range peerIdList {
		for _, id2 := range idList {
			s.UpdateMediaUnread(ctx, peerId, id2)
		}

		if len(idList) > 0 {
			ptsCount := int32(len(idList))
			pts := int32(idgen.NextNPtsId(ctx, peerId, len(idList))) - ptsCount + 1
			sync_client.PushUpdates(ctx,
				peerId,
				model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
					Messages:  idList,
					Pts_INT32: pts,
					PtsCount:  ptsCount,
				}).To_Update()))
		}
	}
	return nil
}
