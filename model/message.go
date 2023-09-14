package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/mention"
)

const (
	MESSAGE_TYPE_UNKNOWN          = 0
	MESSAGE_TYPE_MESSAGE_EMPTY    = 1
	MESSAGE_TYPE_MESSAGE          = 2
	MESSAGE_TYPE_MESSAGE_44F9B43D = 2
	MESSAGE_TYPE_MESSAGE_SERVICE  = 3
	MESSAGE_TYPE_MESSAGE_452C0E65 = 4
)

const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
	MESSAGE_BOX_TYPE_CHANNEL  = 2
)

func MakeDialogId(fromId, peerType, peerId int32) (did int64) {
	switch peerType {
	case PEER_SELF:
		did = int64(fromId)<<32 | int64(fromId)
	case PEER_USER:
		if fromId <= peerId {
			did = int64(fromId)<<32 | int64(peerId)
		} else {
			did = int64(peerId)<<32 | int64(fromId)
		}
	case PEER_CHAT:
		did = int64(-peerId)
	case PEER_CHANNEL:
		did = int64(-peerId)
	default:
		log.Warnf("invalid peer{%d, %d, %d}", fromId, peerType, peerId)
	}
	return
}

func GetPeerIdByDialogId(userId int32, did int64) int32 {
	if did < 0 {
		return int32(-did)
	} else {
		id1 := int32(did & 0xffffffff)
		id2 := int32(did >> 32)
		if userId == id1 {
			return id2
		} else {
			return id1
		}
	}
}

func GetPeerByDialogId(userId int32, did int64) *mtproto.Peer {
	if did < 0 {
		return MakePeerChat(int32(-did))
	} else {
		id1 := int32(did & 0xffffffff)
		id2 := int32(did >> 32)
		if userId == id1 {
			return MakePeerUser(id2)
		} else {
			return MakePeerUser(id1)
		}
	}
}

func PickAllIdListByMessages(messageList []*mtproto.Message) (users, chats, channels IDList) {
	for _, m := range messageList {
		peers := make([]*PeerUtil, 0, 3)
		if m.FromId_FLAGPEER != nil {
			peers = append(peers, FromPeer(m.FromId_FLAGPEER))
		}
		if m.PeerId != nil {
			peers = append(peers, FromPeer(m.PeerId))
		}
		if m.ToId != nil {
			peers = append(peers, FromPeer(m.ToId))
		}
		if m.FwdFrom != nil && m.FwdFrom.FromId_FLAGPEER != nil {
			peers = append(peers, FromPeer(m.FwdFrom.FromId_FLAGPEER))
		}
		for _, peer := range peers {
			switch peer.PeerType {
			case PEER_USER:
				users.AddIfNot(peer.PeerId)
			case PEER_CHAT:
				chats.AddIfNot(peer.PeerId)
			case PEER_CHANNEL:
				channels.AddIfNot(peer.PeerId)
			}
		}
		if m.Action != nil {
			if m.Action.UserId != 0 {
				users.AddIfNot(m.Action.UserId)
			}
			if m.Action.InviterId != 0 {
				users.AddIfNot(m.Action.InviterId)
			}
			if len(m.Action.Users) != 0 {
				users.AddIfNot(m.Action.Users...)
			}
			if m.Action.ChatId != 0 {
				chats.AddIfNot(m.Action.ChatId)
			}
			if m.Action.ChannelId != 0 {
				channels.AddIfNot(m.Action.ChannelId)
			}
		}
		if m.FwdFrom != nil {
			if m.FwdFrom.ChannelId != nil {
				channels.AddIfNot(m.FwdFrom.ChannelId.Value)
			}
		}
	}
	return
}

func CheckHasMention(entities []*mtproto.MessageEntity, userId int32) bool {
	for _, e := range entities {
		// check name
		switch e.PredicateName {
		case mtproto.Predicate_messageEntityMentionName:
			if e.UserId_INT32 == userId {
				return true
			}
		case mtproto.Predicate_messageEntityMention:
			if e.UserId_INT32 == userId {
				return true
			}
		}

		// check crc32
		switch e.GetConstructor() {
		case mtproto.CRC32_messageEntityMentionName:
			if e.UserId_INT32 == userId {
				return true
			}
		case mtproto.CRC32_messageEntityMention:
			if e.UserId_INT32 == userId {
				return true
			}
		}
	}
	return false
}

func CheckHasMediaUnread(m *mtproto.Message) bool {
	return IsVoiceMessage(m) || IsRoundVideoMessage(m)
}

func EncodeMessage(message *mtproto.Message) (messageType int, messageData []byte) {
	switch message.PredicateName {
	case mtproto.Predicate_messageEmpty:
		messageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.Predicate_message:
		messageType = MESSAGE_TYPE_MESSAGE
	case mtproto.Predicate_messageService:
		messageType = MESSAGE_TYPE_MESSAGE_SERVICE
	default:
		log.Warnf("invalid name or clazzId: %s", message.DebugString())
	}

	messageData, _ = json.Marshal(message)
	return
}

func DecodeMessage(messageType int, messageData []byte) (message *mtproto.Message, err error) {
	message = &mtproto.Message{}

	switch messageType {
	case MESSAGE_TYPE_MESSAGE_EMPTY:
	case MESSAGE_TYPE_MESSAGE:
		err = json.Unmarshal(messageData, message)
		if err != nil {
			log.Errorf("decodeMessage - Unmarshal message(%s)error: %v", messageData, err)
			return
		}
	case MESSAGE_TYPE_MESSAGE_SERVICE:
		err = json.Unmarshal(messageData, message)
		if err != nil {
			log.Errorf("decodeMessage - Unmarshal message(%s)error: %v", messageData, err)
			return nil, err
		}
	default:
		err = fmt.Errorf("decodeMessage - Invalid messageType, db's data error, message(%s)", messageData)
		log.Error(err.Error())
		return nil, err
	}

	return message, nil
}

type MessageBox struct {
	SelfUserId        int32            `json:"user_id"`
	SendUserId        int32            `json:"send_user_id"`
	MessageId         int32            `json:"message_id"`
	DialogId          int64            `json:"dialog_id"`
	DialogMessageId   int32            `json:"dialog_message_id"`
	MessageDataId     int64            `json:"message_data_id"`
	RandomId          int64            `json:"random_id"`
	Pts               int32            `json:"pts"`
	PtsCount          int32            `json:"pts_count"`
	MessageFilterType int8             `json:"message_filter_type"`
	MessageBoxType    int8             `json:"message_box_type"`
	MessageType       int8             `json:"message_type"`
	Message           *mtproto.Message `json:"message"`
	Views             int32            `json:"views"`
	ReplyOwnerId      int32            `json:"reply_owner_id"`
	TtlSeconds        int32            `json:"ttl_seconds"`
	DmUsers           []int32          `json:"dm_users"`
}

func (m *MessageBox) ToMessage(toUserId int32) *mtproto.Message {
	switch m.MessageBoxType {
	case MESSAGE_BOX_TYPE_OUTGOING:
		m.Message.Out = true
	case MESSAGE_BOX_TYPE_INCOMING:
		m.Message.Out = false
	case MESSAGE_BOX_TYPE_CHANNEL:
		if m.SendUserId == toUserId {
			m.Message.Out = true
		} else {
			m.Message.Out = false
		}
	default:
	}

	if !m.Message.Out && m.Message.PeerId != nil && m.Message.PeerId.PredicateName == mtproto.Predicate_peerUser {
		m.Message.PeerId.UserId = m.Message.FromId_FLAGPEER.UserId
	}

	m.Message.Id = m.MessageId
	if m.Message.Mentioned {
		m.Message.MediaUnread = true
	}
	return m.Message
}

func PickDmUsers(message *mtproto.Message) []int32 {
	var dmUsers IDList
	text := message.GetMessage()
	if len(text) < 3 {
		return dmUsers
	}
	idxList := mention.EncodeStringToUTF16Index(text)
	for _, entitie := range message.Entities {
		if entitie.PredicateName == mtproto.Predicate_messageEntityMentionName || entitie.PredicateName == mtproto.Predicate_messageEntityMention {
			offset := entitie.GetOffset()
			if offset >= 2 && int32(len(idxList)) > offset {
				start := -1
				end := -1
				for i, idx := range idxList {
					if start < 0 {
						if int32(idx) == offset-2 {
							start = i
						}
					} else if end < 0 {
						if int32(idx) == offset+1 {
							end = i
							break
						}
					}
				}
				if start >= 0 && end > start {
					flag := text[start:end]
					if strings.ToLower(flag) == "dm@" {
						dmUsers.AddIfNot(entitie.UserId_INT32)
					}
				}
			}
		}
	}
	return dmUsers
}
