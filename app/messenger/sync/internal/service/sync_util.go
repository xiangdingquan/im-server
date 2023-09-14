package service

import (
	"open.chat/model"
	"open.chat/mtproto"
)

type SyncType int

const (
	syncTypeUser      SyncType = 1
	syncTypeUserNotMe SyncType = 2
	syncTypeUserMe    SyncType = 3
)

func updateShortMessageToMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) *mtproto.Message {
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

	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             shortMessage.GetOut(),
		Mentioned:       shortMessage.GetMentioned(),
		MediaUnread:     shortMessage.GetMediaUnread(),
		Silent:          shortMessage.GetSilent(),
		Id:              shortMessage.GetId(),
		FromId_FLAGPEER: model.MakePeerUser(fromId),
		PeerId:          model.MakePeerUser(peerId),
		ToId:            model.MakePeerUser(peerId),
		Message:         shortMessage.GetMessage(),
		Date:            shortMessage.GetDate(),
		FwdFrom:         shortMessage.GetFwdFrom(),
		ViaBotId:        shortMessage.GetViaBotId(),
		ReplyTo:         nil,
		Entities:        shortMessage.GetEntities(),
	}).To_Message()
	if shortMessage.GetReplyToMsgId() != nil {
		message.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: shortMessage.GetReplyToMsgId().GetValue(),
			}).To_MessageReplyHeader()
	}
	return message
}

func updateShortChatMessageToMessage(shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Message {
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             shortMessage.GetOut(),
		Mentioned:       shortMessage.GetMentioned(),
		MediaUnread:     shortMessage.GetMediaUnread(),
		Silent:          shortMessage.GetSilent(),
		Id:              shortMessage.GetId(),
		FromId_FLAGPEER: model.MakePeerUser(shortMessage.GetFromId()),
		PeerId:          model.MakePeerChat(shortMessage.GetChatId()),
		ToId:            model.MakePeerChat(shortMessage.GetChatId()),
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
	return message
}

func updateShortToUpdateNewMessage(userId int32, shortMessage *mtproto.TLUpdateShortMessage) *mtproto.Update {
	updateNew := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Message_MESSAGE: updateShortMessageToMessage(userId, shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	})
	return updateNew.To_Update()
}

func updateShortChatToUpdateNewMessage(userId int32, shortMessage *mtproto.TLUpdateShortChatMessage) *mtproto.Update {
	updateNew := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Message_MESSAGE: updateShortChatMessageToMessage(shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	})
	return updateNew.To_Update()
}
