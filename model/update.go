package model

import (
	"time"

	"open.chat/mtproto"
)

type UpdatesLogic struct {
	ownerUserId int32
	message     *mtproto.Message
	updates     []*mtproto.Update
	users       []*mtproto.User
	chats       []*mtproto.Chat
	date        int32
}

// ///////////////////////////////////////////////////////////////////////////////////////
func NewUpdatesLogic(userId int32) *UpdatesLogic {
	return &UpdatesLogic{
		ownerUserId: userId,
		updates:     make([]*mtproto.Update, 0),
		users:       make([]*mtproto.User, 0),
		chats:       make([]*mtproto.Chat, 0),
		date:        int32(time.Now().Unix()),
	}
}

func NewUpdatesLogicByMessage(userId int32, message *mtproto.Message) *UpdatesLogic {
	return &UpdatesLogic{
		ownerUserId: userId,
		message:     message,
	}
}

func NewUpdatesLogicByUpdate(userId int32, update *mtproto.Update) *UpdatesLogic {
	return &UpdatesLogic{
		ownerUserId: userId,
		updates:     []*mtproto.Update{update},
		users:       make([]*mtproto.User, 0),
		chats:       make([]*mtproto.Chat, 0),
		date:        int32(time.Now().Unix()),
	}
}

func NewUpdatesLogicByUpdates(userId int32, updateList []*mtproto.Update) *UpdatesLogic {
	return &UpdatesLogic{
		ownerUserId: userId,
		updates:     updateList,
		users:       make([]*mtproto.User, 0),
		chats:       make([]*mtproto.Chat, 0),
		date:        int32(time.Now().Unix()),
	}
}

func messageToUpdateShortMessage(message2 *mtproto.Message) (shortMessage *mtproto.TLUpdateShortMessage) {
	var (
		userId int32
	)

	switch message2.PredicateName {
	case mtproto.Predicate_message:
		message := message2.To_Message()
		if message.GetOut() {
			userId = message.GetToId().GetUserId()
		} else {
			userId = message.GetFromId_FLAGPEER().GetUserId()
		}
		shortMessage = mtproto.MakeTLUpdateShortMessage(&mtproto.Updates{
			Out:          message.GetOut(),
			Mentioned:    message.GetMentioned(),
			MediaUnread:  message.GetMediaUnread(),
			Silent:       message.GetSilent(),
			Id:           message.GetId(),
			UserId:       userId,
			Message:      message.GetMessage(),
			Date:         message.GetDate(),
			FwdFrom:      message.GetFwdFrom(),
			ViaBotId:     message.GetViaBotId(),
			ReplyToMsgId: message.GetReplyToMsgId(),
			ReplyTo:      message.GetReplyTo(),
			Entities:     message.GetEntities(),
		})
	case mtproto.Predicate_messageService:
	default:
	}
	return
}

func messageToUpdateShortChatMessage(message2 *mtproto.Message) (shortMessage *mtproto.TLUpdateShortMessage) {
	return
}

func messageToUpdateShortSentMessage(message2 *mtproto.Message) (sentMessage *mtproto.TLUpdateShortSentMessage) {
	switch message2.PredicateName {
	case mtproto.Predicate_message:
		message := message2.To_Message()
		sentMessage = mtproto.MakeTLUpdateShortSentMessage(&mtproto.Updates{
			Out: message.GetOut(),
			Id:  message.GetId(),
			// Pts:,
			// PtsCount,
			Date:     message.GetDate(),
			Media:    message.GetMedia(),
			Entities: message.GetEntities(),
		})
	case mtproto.Predicate_messageService:
	default:
	}
	return
}

// ///////////////////////////////////////////////////////////////////////////////////////
func (m *UpdatesLogic) ToUpdateTooLong() *mtproto.Updates {
	return mtproto.MakeTLUpdatesTooLong(nil).To_Updates()
}

func (m *UpdatesLogic) ToUpdateShortMessage() *mtproto.Updates {
	if m.message == nil {
	}

	shortMessage := messageToUpdateShortMessage(m.message)
	return shortMessage.To_Updates()
}

func (m *UpdatesLogic) ToUpdateShortChatMessage() *mtproto.Updates {
	if m.message == nil {
	}

	shortMessage := messageToUpdateShortChatMessage(m.message)
	return shortMessage.To_Updates()
}

func (m *UpdatesLogic) ToUpdateShort() *mtproto.Updates {
	if len(m.updates) != 1 {
	}

	updateShort := mtproto.MakeTLUpdateShort(&mtproto.Updates{
		Update: m.updates[0],
		Date:   m.date,
	})
	return updateShort.To_Updates()
}

func (m *UpdatesLogic) ToUpdatesCombined() *mtproto.Updates {
	updates := mtproto.MakeTLUpdatesCombined(&mtproto.Updates{
		Updates: m.updates,
		Users:   m.users,
		Chats:   m.chats,
		Date:    m.date,
	})
	return updates.To_Updates()
}

func (m *UpdatesLogic) ToUpdates() *mtproto.Updates {
	updates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: m.updates,
		Users:   m.users,
		Chats:   m.chats,
		Date:    m.date,
	})
	return updates.To_Updates()
}

func (m *UpdatesLogic) ToUpdateShortSentMessage() *mtproto.Updates {
	if m.message == nil {
	}

	sentMessage := messageToUpdateShortSentMessage(m.message)
	return sentMessage.To_Updates()
}

// ///////////////////////////////////////////////////////////////////////////////////////
func (m *UpdatesLogic) AddUpdateNewMessage(pts, ptsCount int32, message *mtproto.Message) {
	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Message_MESSAGE: MessageUpdate(message),
		Pts_INT32:       pts,
		PtsCount:        ptsCount,
	})
	m.updates = append(m.updates, updateNewMessage.To_Update())
}

func (m *UpdatesLogic) AddUpdateNewChannelMessage(pts, ptsCount int32, message *mtproto.Message) {
	updateNewChannelMessage := mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
		Message_MESSAGE: MessageUpdate(message),
		Pts_INT32:       pts,
		PtsCount:        ptsCount,
	})
	m.updates = append(m.updates, updateNewChannelMessage.To_Update())
}

func (m *UpdatesLogic) AddUpdateMessageId(messageId int32, randomId int64) {
	updateMessageID := mtproto.MakeTLUpdateMessageID(&mtproto.Update{
		Id_INT32: messageId,
		RandomId: randomId,
	})

	updates := []*mtproto.Update{updateMessageID.To_Update()}
	m.updates = append(updates, m.updates...)
}

func (m *UpdatesLogic) PushTopUpdateMessageId(messageId int32, randomId int64) {
	updateMessageID := mtproto.MakeTLUpdateMessageID(&mtproto.Update{
		Id_INT32: messageId,
		RandomId: randomId,
	})

	updates2 := make([]*mtproto.Update, 0, 1+len(m.updates))
	updates2 = append(updates2, updateMessageID.To_Update())
	m.updates = append(updates2, m.updates...)
}

// ///////////////////////////////////////////////////////////////////////////////////////
func (m *UpdatesLogic) AddUpdates(updateList []*mtproto.Update) {
	m.updates = append(m.updates, updateList...)
}

func (m *UpdatesLogic) AddUpdate(update *mtproto.Update) {
	m.updates = append(m.updates, update)
}

func (m *UpdatesLogic) AddChats(chatList []*mtproto.Chat) {
	m.chats = append(m.chats, chatList...)
}

func (m *UpdatesLogic) AddChat(chat *mtproto.Chat) {
	m.chats = append(m.chats, chat)
}

func (m *UpdatesLogic) AddUsers(userList []*mtproto.User) {
	m.users = append(m.users, userList...)
}

func (m *UpdatesLogic) AddUser(user *mtproto.User) {
	m.users = append(m.users, user)
}

func AddUpdateMessageId(upds *mtproto.Updates, messageId int32, randomId int64) *mtproto.Updates {
	updateMessageID := mtproto.MakeTLUpdateMessageID(&mtproto.Update{
		Id_INT32: messageId,
		RandomId: randomId,
	})
	upds.Updates = append([]*mtproto.Update{updateMessageID.To_Update()}, upds.Updates...)
	return upds
}

func MakeUpdateChannel(channelId int32) *mtproto.Update {
	update := mtproto.MakeTLUpdateChannel(&mtproto.Update{
		ChannelId: channelId,
	})
	return update.To_Update()
}
