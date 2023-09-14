package service

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"time"

	"github.com/gogo/protobuf/types"
	inbox_client "open.chat/app/messenger/msg/inbox/client"
	"open.chat/app/messenger/msg/internal/core"
	"open.chat/app/messenger/msg/internal/dao"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	_ "open.chat/app/service/biz_service/channel/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	_ "open.chat/app/service/biz_service/chat/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	_ "open.chat/app/service/biz_service/private/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
	idgen "open.chat/app/service/idgen/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Service struct {
	conf *Config
	*core.MsgCore

	inboxClient *inbox_client.InboxClient

	username_facade.UsernameFacade
	private_facade.PrivateFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	user_client.UserFacade
}

func New() *Service {
	var (
		ac  = &Config{}
		err error
		s   = new(Service)
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s.conf = ac
	s.MsgCore = core.New(dao.New())

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)

	s.inboxClient, err = inbox_client.New("inbox")
	checkErr(err)

	media_client.New()
	sync_client.New()

	return s
}

func (s *Service) Ping(ctx context.Context) (err error) {
	// return s.Dao.Ping(ctx)
	return nil
}

// Close close the resources.
func (s *Service) Close() error {
	return nil
}

func makeSender(o *msgpb.Sender) *msgpb.Sender {
	return &msgpb.Sender{
		Id:        o.Id,
		Type:      o.Type,
		AuthKeyId: o.AuthKeyId,
	}
}

func (s *Service) SendUserMessage(ctx context.Context, r *msgpb.UserMessage) (*mtproto.Updates, error) {
	log.Debugf("sendUserMessage - %s", logger.JsonDebugData(r))
	sendMe := r.From.Id == r.PeerUserId
	if !sendMe {
	}

	return s.sendUserMessage(ctx, r.From.Id, r.From.AuthKeyId, r.PeerUserId, r.RandomId, r.Message, func(did int32, mid int64, inboxMsg *mtproto.Message) error {
		inBox := &msgpb.InboxUserMessage{
			From:            makeSender(r.From),
			PeerUserId:      r.PeerUserId,
			RandomId:        r.RandomId,
			DialogMessageId: did,
			MessageDataId:   mid,
			Message:         inboxMsg,
		}
		if !sendMe {
			blocked := s.UserFacade.IsBlockedByUser(ctx, r.PeerUserId, r.From.Id)
			if blocked {
				return nil
			}
		}
		return s.inboxClient.SendUserMessageToInbox(ctx, inBox)
	})
}

func (s *Service) SendChatMessage(ctx context.Context, r *msgpb.ChatMessage) (*mtproto.Updates, error) {
	log.Debugf("sendChatMessage - %s", logger.JsonDebugData(r))
	return s.sendChatMessage(ctx, r.From.Id, r.From.AuthKeyId, r.PeerChatId, r.RandomId, r.Message, func(did int32, mid int64, inboxMsg *mtproto.Message) error {
		inBox := &msgpb.InboxChatMessage{
			From:            makeSender(r.From),
			PeerChatId:      r.PeerChatId,
			RandomId:        r.RandomId,
			DialogMessageId: did,
			MessageDataId:   mid,
			Message:         inboxMsg,
		}
		return s.inboxClient.SendChatMessageToInbox(ctx, inBox)
	})
}

func (s *Service) SendUserMultiMessage(ctx context.Context, r *msgpb.UserMultiMessage) (*mtproto.Updates, error) {
	log.Debugf("sendUserMultiMessage - %s", logger.JsonDebugData(r))
	sendMe := r.From.Id == r.PeerUserId
	if !sendMe {
	}

	return s.sendUserMultiMessage(ctx, r.From.Id, r.From.AuthKeyId, r.PeerUserId, r.RandomId, r.Message, func(didList []int32, midList []int64, inboxMsgList []*mtproto.Message) error {
		inBox := &msgpb.InboxUserMultiMessage{
			From:            makeSender(r.From),
			PeerUserId:      r.PeerUserId,
			RandomId:        r.RandomId,
			DialogMessageId: didList,
			MessageDataId:   midList,
			Message:         inboxMsgList,
		}
		return s.inboxClient.SendUserMultiMessageToInbox(ctx, inBox)
	})
}

func (s *Service) SendChatMultiMessage(ctx context.Context, r *msgpb.ChatMultiMessage) (*mtproto.Updates, error) {
	log.Debugf("sendChatMultiMessage - %s", logger.JsonDebugData(r))
	return s.sendChatMultiMessage(ctx, r.From.Id, r.From.AuthKeyId, r.PeerChatId, r.RandomId, r.Message, func(didList []int32, midList []int64, inboxMsgList []*mtproto.Message) error {
		inBox := &msgpb.InboxChatMultiMessage{
			From:            makeSender(r.From),
			PeerChatId:      r.PeerChatId,
			RandomId:        r.RandomId,
			DialogMessageId: didList,
			MessageDataId:   midList,
			Message:         inboxMsgList,
		}
		return s.inboxClient.SendChatMultiMessageToInbox(ctx, inBox)
	})

}

func (s *Service) PushUserMessage(ctx context.Context, r *msgpb.UserMessage) (*mtproto.Bool, error) {
	log.Debugf("pushUserMessage - %s", logger.JsonDebugData(r))
	sendMe := r.From.Id == r.PeerUserId
	if !sendMe {
	}

	if r.From.Type == 0 {
		err := s.pushUserMessage(ctx, r.From.Id, r.PeerUserId, r.RandomId, r.Message, r.From.Type, func(did int32, mid int64, inboxMsg *mtproto.Message) error {
			inBox := &msgpb.InboxUserMessage{
				From:            makeSender(r.From),
				PeerUserId:      r.PeerUserId,
				RandomId:        r.RandomId,
				DialogMessageId: did,
				MessageDataId:   mid,
				Message:         inboxMsg,
			}
			return s.inboxClient.SendUserMessageToInbox(ctx, inBox)
		})
		if err != nil {
			return nil, err
		}
	} else {
		dialogId := model.MakeDialogId(r.From.Id, model.PEER_USER, r.PeerUserId)
		dialogMessageId := int32(idgen.NextMessageDataId(ctx, dialogId))
		inBox := &msgpb.InboxUserMessage{
			From:            makeSender(r.From),
			PeerUserId:      r.PeerUserId,
			RandomId:        r.RandomId,
			DialogMessageId: dialogMessageId,
			MessageDataId:   idgen.GetUUID(),
			Message:         r.Message,
		}
		s.inboxClient.SendUserMessageToInbox(ctx, inBox)
	}

	return mtproto.ToBool(true), nil
}

func (s *Service) EditUserMessage(ctx context.Context, r *msgpb.UserMessage) (*mtproto.Updates, error) {
	return nil, nil
}

func (s *Service) EditChatMessage(ctx context.Context, r *msgpb.ChatMessage) (*mtproto.Updates, error) {
	return nil, nil
}

func (s *Service) sendUserOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	var users model.MutableUsers = s.UserFacade.GetMutableUsers(ctx, r.From.Id, r.PeerId)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}
	peerUser, _ := users.GetImmutableUser(r.PeerId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}
	if s.UserFacade.IsBlockedByUser(ctx, r.PeerId, r.From.Id) {
		err = mtproto.ErrYouBlockedUser
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}

	reply, err = s.SendUserMessage(ctx, &msgpb.UserMessage{
		From:       r.From,
		PeerUserId: r.PeerId,
		RandomId:   r.Message.RandomId,
		Message:    r.Message.Message,
	})
	return
}

func (s *Service) sendChatOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	var users model.MutableUsers = s.UserFacade.GetMutableUsers(ctx, r.From.Id)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}

	reply, err = s.SendChatMessage(ctx, &msgpb.ChatMessage{
		From:       r.From,
		PeerChatId: r.PeerId,
		RandomId:   r.Message.RandomId,
		Message:    r.Message.Message,
	})
	return
}

func (s *Service) sendUserOutgoingMultiMessage(ctx context.Context, r *msgpb.OutgoingMultiMessage) (reply *mtproto.Updates, err error) {
	var users model.MutableUsers = s.UserFacade.GetMutableUsers(ctx, r.From.Id, r.PeerId)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMultiMessage - error: %v", err)
		return
	}
	peerUser, _ := users.GetImmutableUser(r.PeerId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("sendUserOutgoingMultiMessage - error: %v", err)
		return
	}
	if s.UserFacade.IsBlockedByUser(ctx, r.PeerId, r.From.Id) {
		err = mtproto.ErrYouBlockedUser
		log.Errorf("sendUserOutgoingMultiMessage - error: %v", err)
		return
	}

	userMultiMessage := &msgpb.UserMultiMessage{
		From:       r.From,
		PeerUserId: r.PeerId,
		RandomId:   nil,
		Message:    nil,
	}
	for _, m := range r.MultiMessage {
		userMultiMessage.RandomId = append(userMultiMessage.RandomId, m.RandomId)
		userMultiMessage.Message = append(userMultiMessage.Message, m.Message)
	}

	reply, err = s.SendUserMultiMessage(ctx, userMultiMessage)
	return
}

func (s *Service) sendChatOutgoingMultiMessage(ctx context.Context, r *msgpb.OutgoingMultiMessage) (reply *mtproto.Updates, err error) {
	var users model.MutableUsers = s.UserFacade.GetMutableUsers(ctx, r.From.Id)
	sender, _ := users.GetImmutableUser(r.From.Id)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		log.Errorf("sendUserOutgoingMessage - error: %v", err)
		return
	}

	chatMultiMessage := &msgpb.ChatMultiMessage{
		From:       r.From,
		PeerChatId: r.PeerId,
		RandomId:   nil,
		Message:    nil,
	}
	for _, m := range r.MultiMessage {
		chatMultiMessage.RandomId = append(chatMultiMessage.RandomId, m.RandomId)
		chatMultiMessage.Message = append(chatMultiMessage.Message, m.Message)
	}

	reply, err = s.SendChatMultiMessage(ctx, chatMultiMessage)
	return
}

func (s *Service) SendMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	log.Debugf("SendMessage - request: %s", logger.JsonDebugData(r))

	if r.From == nil || r.Message == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("SendMessage - %v", err)
		return
	}

	message := r.Message.GetMessage()
	if message != nil && message.GetTtlSeconds().GetValue() == 0 && message.Message != "" {
		msg, _ := model.ParseJsonMessage(message.Message)
		ephemeralMsg := struct {
			Text      string `json:"text"`
			CountDown uint32 `json:"countDown"`
		}{}
		if msg != nil && msg.Parse(&ephemeralMsg) {
			message.TtlSeconds = mtproto.MakeFlagsInt32(int32(ephemeralMsg.CountDown))
		}
	}

	if r.Message.GetScheduleDate().GetValue() != 0 {
		reply, err = s.sendScheduledMessage(ctx, r)
	} else {
		switch r.PeerType {
		case model.PEER_USER:
			reply, err = s.sendUserOutgoingMessage(ctx, r)
		case model.PEER_CHAT:
			reply, err = s.sendChatOutgoingMessage(ctx, r)
		case model.PEER_CHANNEL:
			reply, err = s.sendChannelOutgoingMessage(ctx, r)
		default:
			err = mtproto.ErrPeerIdInvalid
		}
	}

	if err != nil {
		log.Errorf("sendMessage - error: %v", err)
	} else {
		log.Debugf("sendMessage - reply: %v", reply.DebugString())
	}

	return
}

func (s *Service) SendMultiMessage(ctx context.Context, r *msgpb.OutgoingMultiMessage) (reply *mtproto.Updates, err error) {
	log.Debugf("SendMultiMessage - request: %s", logger.JsonDebugData(r))

	if r.From == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("SendMedia - %v", err)
		return
	}

	switch r.PeerType {
	case model.PEER_USER:
		reply, err = s.sendUserOutgoingMultiMessage(ctx, r)
	case model.PEER_CHAT:
		reply, err = s.sendChatOutgoingMultiMessage(ctx, r)
	case model.PEER_CHANNEL:
		reply, err = s.sendChannelOutgoingMultiMessage(ctx, r)
	default:
		err = mtproto.ErrPeerIdInvalid
		return
	}

	return
}

func (s *Service) PushMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Bool, err error) {
	log.Debugf("PushMessage - request: %s", logger.JsonDebugData(r))

	if r.From == nil || r.Message == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("SendMedia - %v", err)
		return
	}

	if len(r.Message.Message.Message) > 4000 {
		err = mtproto.ErrMediaCaptionTooLong
		log.Errorf("SendMedia - %v", err)
		return
	}

	switch r.PeerType {
	case model.PEER_USER:
		_, err = s.sendUserOutgoingMessage(ctx, r)
	case model.PEER_CHAT:
		_, err = s.sendChatOutgoingMessage(ctx, r)
	case model.PEER_CHANNEL:
		_, err = s.sendChannelOutgoingMessage(ctx, r)
	default:
		err = mtproto.ErrPeerIdInvalid
		return
	}

	reply = mtproto.BoolTrue
	return
}

func (s *Service) editUserOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	outBox, err2 := s.EditUserOutboxMessage(ctx, r.From.Id, r.PeerId, r.Message.Message)
	if err2 != nil {
		err = err2
		return
	}

	inBox := &msgpb.InboxUserEditMessage{
		From:       makeSender(r.From),
		PeerUserId: r.PeerId,
		Message:    outBox.Message,
	}
	s.inboxClient.EditUserMessageToInbox(ctx, inBox)

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())
	syncNotMe := updatesHelper.ToSyncNotMeUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	err = sync_client.SyncUpdatesNotMe(ctx, r.From.Id, r.From.AuthKeyId, syncNotMe)
	if err != nil {
		return
	}

	reply = updatesHelper.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	return
}

func (s *Service) editChatOutgoingMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	outBox, err2 := s.MsgCore.EditChatOutboxMessage(ctx, r.From.Id, r.PeerId, r.Message.Message)
	if err2 != nil {
		err = err2
		return
	}

	inBox := &msgpb.InboxChatEditMessage{
		From:       makeSender(r.From),
		PeerChatId: r.PeerId,
		Message:    outBox.Message,
	}
	s.inboxClient.EditChatMessageToInbox(ctx, inBox)

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())
	syncNotMe := updatesHelper.ToSyncNotMeUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	err = sync_client.SyncUpdatesNotMe(ctx, r.From.Id, r.From.AuthKeyId, syncNotMe)
	if err != nil {
		return
	}

	reply = updatesHelper.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	return
}

func (s *Service) EditMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	log.Debugf("EditMessage - request: %s", logger.JsonDebugData(r))

	if r.From == nil || r.Message == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("EditMessage - %v", err)
		return
	}

	if r.Message.Message.EditDate == nil {
		r.Message.Message.EditDate = &types.Int32Value{Value: int32(time.Now().Unix())}
	}
	switch r.PeerType {
	case model.PEER_USER:
		reply, err = s.editUserOutgoingMessage(ctx, r)
	case model.PEER_CHAT:
		reply, err = s.editChatOutgoingMessage(ctx, r)
	case model.PEER_CHANNEL:
		reply, err = s.editChannelOutgoingMessage(ctx, r)
	default:
		err = mtproto.ErrPeerIdInvalid
		return
	}

	return
}

func (s *Service) DeleteMessages(ctx context.Context, r *msgpb.DeleteMessagesRequest) (reply *mtproto.Messages_AffectedMessages, err error) {
	log.Debugf("DeleteMessages - request: %s", logger.JsonDebugData(r))
	switch r.PeerType {
	case model.PEER_EMPTY:
		if len(r.Id) == 0 {
			reply = mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
				Pts:      int32(idgen.CurrentPtsId(ctx, r.PeerId)),
				PtsCount: 0,
			}).To_Messages_AffectedMessages()
		} else {
			reply, err = s.deleteUserMessages(ctx, r)
		}
	case model.PEER_CHANNEL:
		if len(r.Id) == 0 {
			reply = mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
				Pts:      int32(idgen.CurrentChannelPtsId(ctx, r.PeerId)),
				PtsCount: 0,
			}).To_Messages_AffectedMessages()
		} else {
			reply, err = s.deleteChannelUserMessages(ctx, r)
		}
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("DeleteMessages - error: %v", err)
		return
	}

	log.Debugf("DeleteMessages - reply: %s", reply.DebugString())
	return
}

func (s *Service) deleteUserMessages(ctx context.Context, r *msgpb.DeleteMessagesRequest) (reply *mtproto.Messages_AffectedMessages, err error) {
	if r.From == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("DeleteMessages - %v", err)
		return
	}

	var (
		pts, ptsCount int32
		msgDataIdList []int64
	)

	if msgDataIdList, err = s.MsgCore.DeleteMessages(ctx, r.From.Id, r.Id); err != nil {
		log.Errorf("DeleteMessages - %v", err)
		return
	}

	pts = int32(idgen.NextNPtsId(ctx, r.From.Id, int(len(r.Id))))
	ptsCount = int32(len(r.Id))

	sync_client.SyncUpdatesNotMe(ctx,
		r.From.Id,
		r.From.AuthKeyId,
		model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
			Messages:  r.Id,
			Pts_INT32: pts,
			PtsCount:  ptsCount,
		}).To_Update()))

	if r.Revoke {
		if err2 := s.inboxClient.DeleteMessagesToInbox(ctx, &msgpb.InboxDeleteMessages{
			From: r.From,
			Id:   msgDataIdList,
		}); err2 != nil {
			log.Errorf("DeleteMessages - %v", err)
		}
	}

	reply = mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages()

	return
}

func (s *Service) DeleteHistory(ctx context.Context, r *msgpb.DeleteHistoryRequest) (reply *mtproto.Messages_AffectedHistory, err error) {
	log.Debugf("DeleteHistory - request: %s", logger.JsonDebugData(r))

	if r.From == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("DeleteMessages - %v", err)
		return
	}

	if r.ChannelId == 0 {
		reply, err = s.deleteUserHistory(ctx, r)
	} else {
		reply, err = s.deleteChannelUserHistory(ctx, r)
	}

	log.Debugf("DeleteHistory - reply: %s", reply.DebugString())
	return
}

func (s *Service) deleteUserHistory(ctx context.Context, r *msgpb.DeleteHistoryRequest) (reply *mtproto.Messages_AffectedHistory, err error) {
	var (
		pts, ptsCount int32
		peer          = &model.PeerUtil{PeerType: r.PeerType, PeerId: r.PeerId}
	)

	lastMessageBox, deleteIds := s.MsgCore.GetLastMessageAndIdListByDialog(ctx, r.From.Id, peer)
	if len(deleteIds) == 0 ||
		len(deleteIds) == 1 &&
			lastMessageBox.Message.PredicateName == mtproto.Predicate_messageService &&
			lastMessageBox.Message.Action.PredicateName == mtproto.Predicate_messageActionHistoryClear {
		pts = int32(idgen.CurrentPtsId(ctx, r.From.Id))
		ptsCount = 0
		reply = mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
			Pts:      pts,
			PtsCount: ptsCount,
			Offset:   0,
		}).To_Messages_AffectedHistory()
		return
	}

	pts = int32(idgen.NextNPtsId(ctx, r.From.Id, len(deleteIds)+1))
	ptsCount = int32(len(deleteIds) + 1)
	if r.JustClear {
		deleteIds = deleteIds[1:]
		if _, err = s.MsgCore.DeleteByMessageIdList(ctx, r.From.Id, deleteIds); err != nil {
			return nil, err
		}

		clearHistoryMessage := mtproto.MakeTLMessageService(&mtproto.Message{
			Out:             true,
			Id:              lastMessageBox.MessageId,
			FromId_FLAGPEER: model.MakePeerUser(r.From.Id),
			PeerId:          peer.ToPeer(),
			ToId:            peer.ToPeer(),
			Date:            lastMessageBox.Message.GetDate(),
			Action:          mtproto.MakeTLMessageActionHistoryClear(nil).To_MessageAction(),
		}).To_Message()
		s.MsgCore.EditUserOutboxMessage(ctx, r.From.Id, r.PeerId, clearHistoryMessage)
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
		sync_client.SyncUpdatesNotMe(ctx, r.From.Id, r.From.AuthKeyId, updatesHelper.ToPushUpdates(
			ctx,
			r.From.Id,
			s.UserFacade,
			s.ChatFacade,
			s.ChannelFacade))
	} else {
		if _, err = s.MsgCore.DeleteByMessageIdList(ctx, r.From.Id, deleteIds); err != nil {
			return nil, err
		}
		_, err = s.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
			"deleted": 1,
		}, r.From.Id, r.PeerId)

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
		sync_client.SyncUpdatesNotMe(ctx, r.From.Id, r.From.AuthKeyId, updatesHelper.ToPushUpdates(
			ctx,
			r.From.Id,
			s.UserFacade,
			s.ChatFacade,
			s.ChannelFacade))
	}

	reply = mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
	}).To_Messages_AffectedHistory()

	if r.Revoke {
		switch peer.PeerType {
		case model.PEER_USER:
			s.inboxClient.DeleteUserHistoryToInbox(ctx, &msgpb.InboxUserDeleteHistory{
				From:       r.From,
				PeerUserId: r.PeerId,
				JustClear:  r.JustClear,
				MaxId:      r.MaxId,
			})
		case model.PEER_CHAT:
			s.inboxClient.DeleteChatHistoryToInbox(ctx, &msgpb.InboxChatDeleteHistory{
				From:       r.From,
				PeerChatId: r.PeerId,
				MaxId:      r.MaxId,
			})
		}
	}

	return
}
