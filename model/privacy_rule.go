package model

import (
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	MinChannelId = 1073741824
)

const (
	RULE_TYPE_INVALID          = 0
	ALLOW_CONTACTS             = 1
	ALLOW_ALL                  = 2
	ALLOW_USERS                = 3
	DISALLOW_CONTACTS          = 4
	DISALLOW_ALL               = 5
	DISALLOW_USERS             = 6
	ALLOW_CHAT_PARTICIPANTS    = 7
	DISALLOW_CHAT_PARTICIPANTS = 8
)

const (
	KEY_TYPE_INVALID  = 0
	STATUS_TIMESTAMP  = 1 //
	CHAT_INVITE       = 2
	PHONE_CALL        = 3
	PHONE_P2P         = 4
	FORWARDS          = 5
	PROFILE_PHOTO     = 6
	PHONE_NUMBER      = 7
	ADDED_BY_PHONE    = 8
	ADDED_BY_USERNAME = 9
	SEND_MESSAGES     = 10
	MAX_KEY_TYPE      = 10
)

func FromInputPrivacyKeyType(k *mtproto.InputPrivacyKey) int {
	switch k.PredicateName {
	case mtproto.Predicate_inputPrivacyKeyStatusTimestamp:
		return STATUS_TIMESTAMP
	case mtproto.Predicate_inputPrivacyKeyChatInvite:
		return CHAT_INVITE
	case mtproto.Predicate_inputPrivacyKeyPhoneCall:
		return PHONE_CALL
	case mtproto.Predicate_inputPrivacyKeyPhoneP2P:
		return PHONE_P2P
	case mtproto.Predicate_inputPrivacyKeyForwards:
		return FORWARDS
	case mtproto.Predicate_inputPrivacyKeyProfilePhoto:
		return PROFILE_PHOTO
	case mtproto.Predicate_inputPrivacyKeyPhoneNumber:
		return PHONE_NUMBER
	case mtproto.Predicate_inputPrivacyKeyAddedByPhone:
		return ADDED_BY_PHONE
	case mtproto.Predicate_inputPrivacyKeyAddedByUsername:
		return ADDED_BY_USERNAME
	case mtproto.Predicate_inputPrivacyKeySendMessage:
		return SEND_MESSAGES
	}
	return KEY_TYPE_INVALID
}

func ToPrivacyKey(keyType int) (key *mtproto.PrivacyKey) {
	switch keyType {
	case STATUS_TIMESTAMP:
		key = mtproto.MakeTLPrivacyKeyStatusTimestamp(nil).To_PrivacyKey()
	case CHAT_INVITE:
		key = mtproto.MakeTLPrivacyKeyChatInvite(nil).To_PrivacyKey()
	case PHONE_CALL:
		key = mtproto.MakeTLPrivacyKeyPhoneCall(nil).To_PrivacyKey()
	case PHONE_P2P:
		key = mtproto.MakeTLPrivacyKeyPhoneP2P(nil).To_PrivacyKey()
	case FORWARDS:
		key = mtproto.MakeTLPrivacyKeyForwards(nil).To_PrivacyKey()
	case PROFILE_PHOTO:
		key = mtproto.MakeTLPrivacyKeyProfilePhoto(nil).To_PrivacyKey()
	case PHONE_NUMBER:
		key = mtproto.MakeTLPrivacyKeyPhoneNumber(nil).To_PrivacyKey()
	case ADDED_BY_PHONE:
		key = mtproto.MakeTLPrivacyKeyAddedByPhone(nil).To_PrivacyKey()
	case ADDED_BY_USERNAME:
		key = mtproto.MakeTLPrivacyKeyAddedByUsername(nil).To_PrivacyKey()
	case SEND_MESSAGES:
		key = mtproto.MakeTLPrivacyKeySendMessage(nil).To_PrivacyKey()
	default:
		panic("type is invalid")
	}
	return
}

func ToPrivacyRuleByInput(userSelfId int32, inputRule *mtproto.InputPrivacyRule) *mtproto.PrivacyRule {
	switch inputRule.PredicateName {
	case mtproto.Predicate_inputPrivacyValueAllowAll:
		return mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowContacts:
		return mtproto.MakeTLPrivacyValueAllowContacts(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowUsers:
		return mtproto.MakeTLPrivacyValueAllowUsers(&mtproto.PrivacyRule{
			Users: ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowAll:
		return mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowContacts:
		return mtproto.MakeTLPrivacyValueDisallowContacts(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowUsers:
		return mtproto.MakeTLPrivacyValueDisallowUsers(&mtproto.PrivacyRule{
			Users: ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowChatParticipants:
		return mtproto.MakeTLPrivacyValueAllowChatParticipants(&mtproto.PrivacyRule{
			Chats: inputRule.GetChats(),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowChatParticipants:
		return mtproto.MakeTLPrivacyValueDisallowChatParticipants(&mtproto.PrivacyRule{
			Chats: inputRule.GetChats(),
		}).To_PrivacyRule()
	default:
		log.Errorf("type is invalid")
	}
	return nil
}

func ToPrivacyRuleListByInput(userSelfId int32, inputRules []*mtproto.InputPrivacyRule) (rules []*mtproto.PrivacyRule) {
	rules = make([]*mtproto.PrivacyRule, 0, len(inputRules))
	for _, inputRule := range inputRules {
		rules = append(rules, ToPrivacyRuleByInput(userSelfId, inputRule))
	}
	return
}

// pick chat and channel
func PickAllIdListByRules(rules []*mtproto.PrivacyRule) (userIdList, chatIdList, channelIdList []int32) {
	userIdList = make([]int32, 0)
	chatIdList = make([]int32, 0)
	channelIdList = make([]int32, 0)
	for _, r := range rules {
		if len(r.Users) > 0 {
			userIdList = append(userIdList, r.Users...)
		}
		for _, id := range r.Chats {
			if id >= MinChannelId {
				channelIdList = append(channelIdList, id)
			} else {
				chatIdList = append(chatIdList, id)
			}
		}
	}
	return
}

func CheckPrivacyIsAllow(selfId int32,
	rules []*mtproto.PrivacyRule,
	checkId int32,
	cbContact func(id, checkId int32) bool,
	cbChat func(checkId int32, idList []int32) bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PredicateName {
		case mtproto.Predicate_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case mtproto.Predicate_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case mtproto.Predicate_privacyValueDisallowAll:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return cbContact(selfId, checkId)
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			}
		}
		return false
	}

	return false
}

func privacyIsAllow(rules []*mtproto.PrivacyRule, userId int32, isContact bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PredicateName {
		case mtproto.Predicate_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case mtproto.Predicate_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case mtproto.Predicate_privacyValueAllowUsers:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				return true
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return isContact
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				return true
			}
		}
		return false
	}

	return false
}
