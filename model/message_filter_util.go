package model

import (
	"open.chat/mtproto"
)

type MediaType int8

const (
	Photo          = 1
	Video          = 2
	PhotoVideo     = 3
	MusicFile      = 4
	File           = 5
	VoiceFile      = 6
	Link           = 7
	ChatPhoto      = 8
	RoundVoiceFile = 9
	GIF            = 10
	RoundFile      = 11
)

const (
	MEDIA_EMPTY      MediaType = -1
	MEDIA_PHOTOVIDEO MediaType = 0
	MEDIA_FILE       MediaType = 1
	MEDIA_AUDIO      MediaType = 2
	MEDIA_URL        MediaType = 3
	MEDIA_MUSIC      MediaType = 4
	MEDIA_GIF        MediaType = 5
	MEDIA_PHONE_CALL MediaType = 6
)

func GetMediaType(message *mtproto.Message) MediaType {
	if message == nil {
		return MEDIA_EMPTY
	}

	switch message.GetPredicateName() {
	case mtproto.Predicate_messageService:
		switch message.GetAction().GetPredicateName() {
		case mtproto.Predicate_messageActionPhoneCall:
			return MEDIA_PHONE_CALL
		}
	case mtproto.Predicate_message:
		switch message.GetMedia().GetPredicateName() {
		case mtproto.Predicate_messageMediaPhoto:
			return MEDIA_PHOTOVIDEO
		case mtproto.Predicate_messageMediaDocument:
			if IsVoiceMessage(message) || IsRoundVideoMessage(message) {
				return MEDIA_AUDIO
			} else if IsVideoMessage(message) {
				return MEDIA_PHOTOVIDEO
			} else if IsStickerMessage(message) || IsAnimatedStickerMessage(message) {
				return MEDIA_EMPTY
			} else if IsNewGifMessage(message) {
				return MEDIA_GIF
			} else if IsMusicMessage(message) {
				return MEDIA_MUSIC
			} else {
				return MEDIA_FILE
			}
		}

		for _, entity := range message.GetEntities() {
			switch entity.PredicateName {
			case mtproto.Predicate_messageEntityUrl,
				mtproto.Predicate_messageEntityTextUrl,
				mtproto.Predicate_messageEntityEmail:
				return MEDIA_URL
			case mtproto.Predicate_messageActionPhoneCall:
				return MEDIA_PHONE_CALL
			}
		}
	}
	return MEDIA_EMPTY
}

type MessagesFilterType int8

const (
	FilterEmpty      MessagesFilterType = 0
	FilterPhotos     MessagesFilterType = 1
	FilterVideo      MessagesFilterType = 2
	FilterPhotoVideo MessagesFilterType = 3
	FilterDocument   MessagesFilterType = 4
	FilterUrl        MessagesFilterType = 5
	FilterGif        MessagesFilterType = 6
	FilterVoice      MessagesFilterType = 7
	FilterMusic      MessagesFilterType = 8
	FilterChatPhotos MessagesFilterType = 9
	FilterPhoneCalls MessagesFilterType = 10
	FilterRoundVoice MessagesFilterType = 11
	FilterRoundVideo MessagesFilterType = 12
	FilterMyMentions MessagesFilterType = 13
	FilterGeo        MessagesFilterType = 14
	FilterContacts   MessagesFilterType = 15
)

func FromMessagesFilter(filter *mtproto.MessagesFilter) MessagesFilterType {
	r := FilterEmpty
	switch filter.PredicateName {
	case mtproto.Predicate_inputMessagesFilterEmpty:
		r = FilterEmpty
	case mtproto.Predicate_inputMessagesFilterPhotos:
		r = FilterPhotos
	case mtproto.Predicate_inputMessagesFilterVideo:
		r = FilterVideo
	case mtproto.Predicate_inputMessagesFilterPhotoVideo:
		r = FilterPhotoVideo
	case mtproto.Predicate_inputMessagesFilterDocument:
		r = FilterDocument
	case mtproto.Predicate_inputMessagesFilterUrl:
		r = FilterUrl
	case mtproto.Predicate_inputMessagesFilterGif:
		r = FilterGif
	case mtproto.Predicate_inputMessagesFilterVoice:
		r = FilterVoice
	case mtproto.Predicate_inputMessagesFilterMusic:
		r = FilterMusic
	case mtproto.Predicate_inputMessagesFilterChatPhotos:
		r = FilterChatPhotos
	case mtproto.Predicate_inputMessagesFilterPhoneCalls:
		r = FilterPhoneCalls
	case mtproto.Predicate_inputMessagesFilterRoundVoice:
		r = FilterRoundVoice
	case mtproto.Predicate_inputMessagesFilterRoundVideo:
		r = FilterRoundVideo
	case mtproto.Predicate_inputMessagesFilterMyMentions:
		r = FilterMyMentions
	case mtproto.Predicate_inputMessagesFilterGeo:
		r = FilterGeo
	case mtproto.Predicate_inputMessagesFilterContacts:
		r = FilterContacts
	}

	return r
}

func GetMessagesFilterType(msg *mtproto.Message) MessagesFilterType {
	r := FilterEmpty

	switch msg.PredicateName {
	case mtproto.Predicate_message:
	case mtproto.Predicate_messageService:
		action := msg.Action
		switch action.PredicateName {
		case mtproto.Predicate_messageActionPhoneCall:
			r = FilterPhoneCalls
		}
	}
	return r
}
