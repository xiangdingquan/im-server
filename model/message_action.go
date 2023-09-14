package model

import (
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

func MakeMessageActionEmpty() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionEmpty(nil).To_MessageAction()
}

func MakeMessageActionChatCreate(title string, users []int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatCreate(&mtproto.MessageAction{
		Title: title,
		Users: users,
	}).To_MessageAction()
}

func MakeMessageActionChatEditTitle(title string) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatEditTitle(&mtproto.MessageAction{
		Title: title,
	}).To_MessageAction()
}

func MakeMessageActionChatEditPhoto(photo *mtproto.Photo) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatEditPhoto(&mtproto.MessageAction{
		Photo: photo,
	}).To_MessageAction()
}

func MakeMessageActionChatDeletePhoto() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatDeletePhoto(&mtproto.MessageAction{}).To_MessageAction()
}

func MakeMessageActionChatAddUser(users ...int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatAddUser(&mtproto.MessageAction{
		Users: users,
	}).To_MessageAction()
}

func MakeMessageActionChatDeleteUser(userId int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatDeleteUser(&mtproto.MessageAction{
		UserId: userId,
	}).To_MessageAction()
}

func MakeMessageActionChatJoinByLink(inviterId int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatJoinedByLink(&mtproto.MessageAction{
		InviterId: inviterId,
	}).To_MessageAction()
}

func MakeMessageActionChannelCreate(title string) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChannelCreate(&mtproto.MessageAction{
		Title: title,
	}).To_MessageAction()
}

func MakeMessageActionChatMigrateTo(channelId int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChatMigrateTo(&mtproto.MessageAction{
		ChannelId: channelId,
	}).To_MessageAction()
}

func MakeMessageActionChannelMigrateFrom(title string, chatId int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionChannelMigrateFrom(&mtproto.MessageAction{
		Title:  title,
		ChatId: chatId,
	}).To_MessageAction()
}

func MakeMessageActionPinMessage() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionPinMessage(nil).To_MessageAction()
}

func MakeMessageActionHistoryClear() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionHistoryClear(nil).To_MessageAction()
}

func MakeMessageActionGameScore(gameId int64, score int32) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionGameScore(&mtproto.MessageAction{
		GameId: gameId,
		Score:  score,
	}).To_MessageAction()
}

func MakeMessageActionPaymentSentMe(currency string, totalAmount int64, payload []byte, info *mtproto.PaymentRequestedInfo, shippingOptionId string, charge *mtproto.PaymentCharge) *mtproto.MessageAction {
	action := mtproto.MakeTLMessageActionGameScore(&mtproto.MessageAction{
		Currency:         currency,
		TotalAmount:      totalAmount,
		Payload:          payload,
		Info:             info,
		ShippingOptionId: nil,
		Charge:           charge,
	}).To_MessageAction()

	if shippingOptionId != "" {
		action.ShippingOptionId = &types.StringValue{Value: shippingOptionId}
	}
	return action
}

func MakeMessageActionPaymentSent(currency string, totalAmount int64) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionPaymentSent(&mtproto.MessageAction{
		Currency:    currency,
		TotalAmount: totalAmount,
	}).To_MessageAction()
}

func MakeMessageActionPhoneCall(video bool, callId int64, reason *mtproto.PhoneCallDiscardReason, duration int32) *mtproto.MessageAction {
	action := mtproto.MakeTLMessageActionPhoneCall(&mtproto.MessageAction{
		Video:    video,
		CallId:   callId,
		Reason:   reason,
		Duration: nil,
	}).To_MessageAction()
	if duration > 0 {
		action.Duration = &types.Int32Value{Value: duration}
	}
	return action
}

func MakeMessageActionScreenshotTaken() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionScreenshotTaken(nil).To_MessageAction()
}

func MakeMessageActionCustomAction(message string) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionCustomAction(&mtproto.MessageAction{
		Message: message,
	}).To_MessageAction()
}

func MakeMessageActionBotAllowed(domain string) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionBotAllowed(&mtproto.MessageAction{
		Domain: domain,
	}).To_MessageAction()
}

func MakeMessageActionSecureValuesSentMe(values []*mtproto.SecureValue, credentials *mtproto.SecureCredentialsEncrypted) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionSecureValuesSentMe(&mtproto.MessageAction{
		Values:      values,
		Credentials: credentials,
	}).To_MessageAction()
}

func MakeMessageActionSecureValuesSent(types []*mtproto.SecureValueType) *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionSecureValuesSent(&mtproto.MessageAction{
		Types: types,
	}).To_MessageAction()
}

func MakeMessageActionContactSignUp() *mtproto.MessageAction {
	return mtproto.MakeTLMessageActionContactSignUp(nil).To_MessageAction()
}

func MakeContactSignUpMessage(fromId, toId int32) *mtproto.Message {
	return mtproto.MakeTLMessageService(&mtproto.Message{
		Out:             true,
		Mentioned:       false,
		MediaUnread:     false,
		Silent:          false,
		Post:            false,
		Legacy:          false,
		Id:              0,
		FromId_FLAGPEER: MakePeerUser(fromId),
		ToId:            MakePeerUser(toId),
		ReplyTo:         nil,
		Date:            int32(time.Now().Unix()),
		Action:          MakeMessageActionContactSignUp(),
	}).To_Message()
}
