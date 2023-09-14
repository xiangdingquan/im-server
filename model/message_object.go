package model

import "open.chat/mtproto"

const (
	MessageEmpty     = 0
	MessageTypePhoto = 1
	MessageTypePoll  = 17
)

func getMessageType(m *mtproto.Message) (mType int) {
	switch m.PredicateName {
	case mtproto.Predicate_message:
		if IsMediaEmpty(m) {
			mType = 0
		} else if m.Media.TtlSeconds != nil &&
			(m.Media.Photo_FLAGPHOTO.PredicateName == mtproto.Predicate_photoEmpty ||
				m.Media.Document.PredicateName == mtproto.Predicate_document) {
			mType = 10 // EncryptedPhoto
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaPhoto {
			mType = 1 // Photo
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaGeo ||
			m.Media.PredicateName == mtproto.Predicate_messageMediaVenue ||
			m.Media.PredicateName == mtproto.Predicate_messageMediaGeoLive {
			mType = 4 // Geo: messageMediaGeo || messageMediaVenue || messageMediaGeoLive
		} else if IsRoundVideoMessage(m) {
			mType = 5 // RoundVideo
		} else if IsVideoMessage(m) {
			mType = 3 // RoundVideo
		} else if IsVoiceMessage(m) {
			mType = 2 // Voice
		} else if IsMusicMessage(m) {
			mType = 14 // Music
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaContact {
			mType = 12 // Contact
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaPoll {
			mType = 17 // Poll
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaUnsupported {
			mType = 0 // Unsupported
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaDocument {
			if m.Media.Document != nil && m.Media.Document.MimeType != "" {
				if IsGifDocument(m.Media.Document) {
					mType = 8 //
				} else if m.Media.Document.MimeType == "image/webp" && IsStickerMessage(m) {
					mType = 13
				} else {
					mType = 9
				}
			} else {
				mType = 9
			}
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaGame {
			mType = 0
		} else if m.Media.PredicateName == mtproto.Predicate_messageMediaInvoice {
			mType = 0
		}
	case mtproto.Predicate_messageService:
		if m.Action.PredicateName == mtproto.Predicate_messageActionChatEditPhoto {
			mType = 11
		} else if m.Action.PredicateName == mtproto.Predicate_messageActionHistoryClear {
			mType = -1
		} else if m.Action.PredicateName == mtproto.Predicate_messageActionPhoneCall {
			mType = 16
		} else {
			mType = 10
		}
	case mtproto.Predicate_messageEmpty:
	default:
	}
	return
}

func IsMediaEmpty(message *mtproto.Message) bool {
	return message == nil ||
		message.GetMedia() == nil ||
		message.Media.PredicateName == mtproto.Predicate_messageMediaEmpty ||
		message.Media.PredicateName == mtproto.Predicate_messageMediaWebPage
}

func IsMediaEmptyWebpage(message *mtproto.Message) bool {
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaEmpty
}

func IsStickerMessage(message *mtproto.Message) bool {
	return IsStickerDocument(message.GetMedia().GetDocument())
}

func IsLocationMessage(message *mtproto.Message) bool {
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaGeo ||
		message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaGeoLive ||
		message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaVenue
}

func IsMaskMessage(message *mtproto.Message) bool {
	return IsMaskDocument(message.GetMedia().GetDocument())
}

func IsMusicMessage(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return IsMusicDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsMusicDocument(message.GetMedia().GetDocument())
}

func IsGifMessage(message *mtproto.Message) bool {
	return IsGifDocument(message.GetMedia().GetDocument())
}

func IsRoundVideoMessage(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return IsRoundVideoDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsRoundVideoDocument(message.GetMedia().GetDocument())
}

func IsPhoto(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return message.GetMedia().GetWebpage().GetPhoto().GetPredicateName() == mtproto.Predicate_photo &&
			message.GetMedia().GetWebpage().GetDocument().GetPredicateName() == mtproto.Predicate_document
	}
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaPhoto
}

func IsVoiceMessage(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return IsVoiceDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsVoiceDocument(message.GetMedia().GetDocument())
}

func IsNewGifMessage(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return IsNewGifDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsNewGifDocument(message.GetMedia().GetDocument())
}

func IsLiveLocationMessage(message *mtproto.Message) bool {
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaGeoLive
}

func IsVideoMessage(message *mtproto.Message) bool {
	if message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaWebPage {
		return IsVideoDocument(message.Media.Webpage.Document)
	}

	return IsVideoDocument(message.GetMedia().GetDocument())
}

func IsGameMessage(message *mtproto.Message) bool {
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaGame
}

func IsInvoiceMessage(message *mtproto.Message) bool {
	return message.GetMedia().GetPredicateName() == mtproto.Predicate_messageMediaInvoice
}

func IsForwardedMessage(message *mtproto.Message) bool {
	return message.GetFwdFrom() != nil
}

func IsAnimatedStickerMessage(message *mtproto.Message) bool {
	return false
}
