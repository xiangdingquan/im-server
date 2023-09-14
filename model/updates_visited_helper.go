package model

import (
	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func makeMessageByUpdateShortMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) (message *mtproto.Message) {
	var (
		fromId, peerId int32
	)

	if shortMessage.GetOut() {
		fromId = userId
		peerId = shortMessage.GetUserId()
	} else {
		fromId = shortMessage.GetUserId()
		peerId = userId
	}

	message = mtproto.MakeTLMessage(&mtproto.Message{
		Out:             shortMessage.GetOut(),
		Mentioned:       shortMessage.GetMentioned(),
		MediaUnread:     shortMessage.GetMediaUnread(),
		Silent:          shortMessage.GetSilent(),
		Id:              shortMessage.GetId(),
		FromId_FLAGPEER: MakePeerUser(fromId),
		PeerId:          MakePeerUser(peerId),
		ToId:            MakePeerUser(peerId),
		Message:         shortMessage.GetMessage(),
		Date:            shortMessage.GetDate(),
		FwdFrom:         shortMessage.GetFwdFrom(),
		ViaBotId:        shortMessage.GetViaBotId(),
		ReplyTo:         nil,
		Entities:        shortMessage.GetEntities(),
	}).To_Message()
	if shortMessage.GetReplyTo() != nil {
		message.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: shortMessage.GetReplyTo().ReplyToMsgId,
			}).To_MessageReplyHeader()
	}
	return
}

func makeMessageByUpdateShortChatMessage(shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Message {
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             shortMessage.GetOut(),
		Mentioned:       shortMessage.GetMentioned(),
		MediaUnread:     shortMessage.GetMediaUnread(),
		Silent:          shortMessage.GetSilent(),
		Id:              shortMessage.GetId(),
		FromId_FLAGPEER: MakePeerUser(shortMessage.GetFromId()),
		PeerId:          MakePeerChat(shortMessage.GetChatId()),
		ToId:            MakePeerChat(shortMessage.GetChatId()),
		Message:         shortMessage.GetMessage(),
		Date:            shortMessage.GetDate(),
		FwdFrom:         shortMessage.GetFwdFrom(),
		ViaBotId:        shortMessage.GetViaBotId(),
		ReplyTo:         nil,
		Entities:        shortMessage.GetEntities(),
	}).To_Message()
	if shortMessage.GetReplyTo() != nil {
		message.ReplyToMsgId = &types.Int32Value{Value: shortMessage.GetReplyTo().ReplyToMsgId}
		message.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: shortMessage.GetReplyTo().ReplyToMsgId,
			}).To_MessageReplyHeader()
	}
	return message
}

func makeNewUpdateMessageByShortMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) *mtproto.Update {
	return mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Message_MESSAGE: makeMessageByUpdateShortMessage(userId, shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	}).To_Update()
}

func makeNewUpdateMessageByShortChatMessage(userId int32, shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Update {
	return mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Message_MESSAGE: makeMessageByUpdateShortChatMessage(shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	}).To_Update()
}

type UpdateVisitedFunc func(userId int32, update *mtproto.Update, users []*mtproto.User, chats []*mtproto.Chat, date int32)

func VisitUpdates(userId int32, updates *mtproto.Updates, handlers map[string]UpdateVisitedFunc) {
	if handlers == nil {
		log.Warnf("VisitUpdates - handlers is nil")
		return
	}

	switch updates.PredicateName {
	case mtproto.Predicate_updatesTooLong:
	case mtproto.Predicate_updateShortMessage:
		if vF, ok := handlers[mtproto.Predicate_updateNewMessage]; ok {
			vF(userId,
				makeNewUpdateMessageByShortMessage(userId, updates.To_UpdateShortMessage()),
				nil,
				nil,
				updates.Date)
		}
	case mtproto.Predicate_updateShortChatMessage:
		if vF, ok := handlers[mtproto.Predicate_updateNewMessage]; ok {
			vF(userId,
				makeNewUpdateMessageByShortChatMessage(userId, updates.To_UpdateShortChatMessage()),
				nil,
				nil,
				updates.Date)
		}
	case mtproto.Predicate_updateShort:
		if vF, ok := handlers[updates.Update.GetPredicateName()]; ok {
			vF(userId, updates.Update, nil, nil, updates.Date)
		}
	case mtproto.Predicate_updatesCombined:
		for _, update := range updates.Updates {
			if vF, ok := handlers[update.PredicateName]; ok {
				vF(userId, update, updates.Users, updates.Chats, updates.Date)
			}
		}
	case mtproto.Predicate_updates:
		for _, update := range updates.Updates {
			if vF, ok := handlers[update.PredicateName]; ok {
				vF(userId, update, updates.Users, updates.Chats, updates.Date)
			}
		}
	case mtproto.Predicate_updateShortSentMessage:
	}
}
