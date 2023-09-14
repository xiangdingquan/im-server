package model

import (
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	PEER_EMPTY           = 0
	PEER_SELF            = 1
	PEER_USER            = 2
	PEER_CHAT            = 3
	PEER_CHANNEL         = 4
	PEER_USERS           = 5
	PEER_CHATS           = 6
	PEER_ENCRYPTED_CHAT  = 7
	PEER_BROADCASTS      = 8
	PEER_ALL             = 9
	PEER_USER_MESSAGE    = 10
	PEER_CHANNEL_MESSAGE = 11
	PEER_UNKNOWN         = -1
)

type PeerUtil struct {
	selfId     int32
	PeerType   int32
	PeerId     int32
	AccessHash int64
}

func (p PeerUtil) String() (s string) {
	switch p.PeerType {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	}
	// return
}

func (p PeerUtil) CanDoSendMessage() bool {
	switch p.PeerType {
	case PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		return true
	default:
		return false
	}
}

func FromInputUser(selfId int32, user *mtproto.InputUser) (p *PeerUtil) {
	p = &PeerUtil{}
	switch user.PredicateName {
	case mtproto.Predicate_inputUserEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.Predicate_inputUserSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case mtproto.Predicate_inputUser:
		p.PeerType = PEER_USER
		p.PeerId = user.UserId
		p.AccessHash = user.GetAccessHash()
	case mtproto.Predicate_inputUserFromMessage:
	default:
		p.PeerType = PEER_EMPTY
	}
	return
}

func FromInputPeer(peer *mtproto.InputPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case mtproto.Predicate_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.Predicate_inputPeerSelf:
		p.PeerType = PEER_SELF
	case mtproto.Predicate_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
		p.AccessHash = peer.GetAccessHash()
	case mtproto.Predicate_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case mtproto.Predicate_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
		p.AccessHash = peer.AccessHash
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return
}

func FromInputPeer2(selfId int32, peer *mtproto.InputPeer) (p *PeerUtil) {
	p = &PeerUtil{
		selfId: selfId,
	}
	switch peer.PredicateName {
	case mtproto.Predicate_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case mtproto.Predicate_inputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case mtproto.Predicate_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
		p.AccessHash = peer.AccessHash
	case mtproto.Predicate_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case mtproto.Predicate_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
		p.AccessHash = peer.AccessHash
	default:
		p.PeerType = PEER_UNKNOWN
	}
	return

}

func (p *PeerUtil) ToInputPeer() (peer *mtproto.InputPeer) {
	switch p.PeerType {
	case PEER_EMPTY:
		peer = mtproto.MakeTLInputPeerEmpty(&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerEmpty,
		}).To_InputPeer()
	case PEER_SELF:
		peer = mtproto.MakeTLInputPeerSelf(&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerSelf,
		}).To_InputPeer()
	case PEER_USER:
		peer = mtproto.MakeTLInputPeerUser(&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerUser,
			UserId:      p.PeerId,
			AccessHash:  p.AccessHash,
		}).To_InputPeer()
	case PEER_CHAT:
		peer = mtproto.MakeTLInputPeerChat(&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerChat,
			ChatId:      p.PeerId,
		}).To_InputPeer()
	case PEER_CHANNEL:
		peer = mtproto.MakeTLInputPeerChannel(&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerChannel,
			ChannelId:   p.PeerId,
			AccessHash:  p.AccessHash,
		}).To_InputPeer()
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", p))
	}
	return
}

func FromPeer(peer *mtproto.Peer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case mtproto.Predicate_peerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
	case mtproto.Predicate_peerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case mtproto.Predicate_peerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", peer))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer *mtproto.Peer) {
	switch p.PeerType {
	case PEER_SELF:
		if p.PeerId != 0 {
			peer = mtproto.MakeTLPeerUser(&mtproto.Peer{
				UserId: p.PeerId,
			}).To_Peer()
		} else if p.selfId != 0 {
			peer = mtproto.MakeTLPeerUser(&mtproto.Peer{
				UserId: p.selfId,
			}).To_Peer()
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", p))
		}
	case PEER_USER:
		peer = mtproto.MakeTLPeerUser(&mtproto.Peer{
			UserId: p.PeerId,
		}).To_Peer()
	case PEER_CHAT:
		peer = mtproto.MakeTLPeerChat(&mtproto.Peer{
			ChatId: p.PeerId,
		}).To_Peer()
	case PEER_CHANNEL:
		peer = mtproto.MakeTLPeerChannel(&mtproto.Peer{
			ChannelId: p.PeerId,
		}).To_Peer()
	default:
		peer = nil
	}
	return
}

func (p *PeerUtil) IsEmpty() bool {
	return p.PeerType == PEER_EMPTY
}

func (p *PeerUtil) IsSelf() bool {
	return p.PeerType == PEER_SELF
}

func (p *PeerUtil) IsUser() bool {
	return p.PeerType == PEER_SELF || p.PeerType == PEER_USER
}

func FromInputNotifyPeer(selfId int32, peer *mtproto.InputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case mtproto.Predicate_inputNotifyPeer:
		p = FromInputPeer2(selfId, peer.Peer)
	case mtproto.Predicate_inputNotifyUsers:
		p.PeerType = PEER_USERS
	case mtproto.Predicate_inputNotifyChats:
		p.PeerType = PEER_CHATS
	case mtproto.Predicate_inputNotifyBroadcasts:
		p.PeerType = PEER_BROADCASTS
	case mtproto.Predicate_inputNotifyAll:
		p.PeerType = PEER_ALL
	default:
		log.Errorf("fromInputNotifyPeer: invalid peer - %v", peer)
		p.PeerType = PEER_UNKNOWN
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer *mtproto.InputNotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = mtproto.MakeTLInputNotifyPeer(&mtproto.InputNotifyPeer{
			Constructor: mtproto.CRC32_inputNotifyPeer,
			Peer:        p.ToInputPeer(),
		}).To_InputNotifyPeer()
	case PEER_USERS:
		peer = mtproto.MakeTLInputNotifyUsers(&mtproto.InputNotifyPeer{
			Constructor: mtproto.CRC32_inputNotifyUsers,
		}).To_InputNotifyPeer()
	case PEER_CHATS:
		peer = mtproto.MakeTLInputNotifyChats(&mtproto.InputNotifyPeer{
			Constructor: mtproto.CRC32_inputNotifyChats,
		}).To_InputNotifyPeer()
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToNotifyPeer() (peer *mtproto.NotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = mtproto.MakeTLNotifyPeer(&mtproto.NotifyPeer{
			Peer: p.ToPeer(),
		}).To_NotifyPeer()
	case PEER_USERS:
		peer = mtproto.MakeTLNotifyUsers(&mtproto.NotifyPeer{}).To_NotifyPeer()
	case PEER_CHATS:
		peer = mtproto.MakeTLNotifyChats(&mtproto.NotifyPeer{}).To_NotifyPeer()
	case PEER_BROADCASTS:
		peer = mtproto.MakeTLNotifyBroadcasts(&mtproto.NotifyPeer{}).To_NotifyPeer()
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}

func ToPeerByTypeAndID(peerType int8, peerId int32) (peer *mtproto.Peer) {
	switch peerType {
	case PEER_USER:
		peer = mtproto.MakeTLPeerUser(&mtproto.Peer{
			UserId: peerId,
		}).To_Peer()
	case PEER_CHAT:
		peer = mtproto.MakeTLPeerChat(&mtproto.Peer{
			ChatId: peerId,
		}).To_Peer()
	case PEER_CHANNEL:
		peer = mtproto.MakeTLPeerChannel(&mtproto.Peer{
			ChannelId: peerId,
		}).To_Peer()
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerId))
	}
	return
}

func PickAllIdListByPeers(peers []*mtproto.Peer) (idList, chatIdList, channelIdList []int32) {
	for _, p := range peers {
		switch p.PredicateName {
		case mtproto.Predicate_peerUser:
			idList = append(idList, p.UserId)
		case mtproto.Predicate_peerChat:
			chatIdList = append(chatIdList, p.ChatId)
		case mtproto.Predicate_peerChannel:
			channelIdList = append(channelIdList, p.ChannelId)
		}
	}
	if idList == nil {
		idList = []int32{}
	}
	if chatIdList == nil {
		chatIdList = []int32{}
	}
	if channelIdList == nil {
		channelIdList = []int32{}
	}
	return
}

func MakePeerUser(peerId int32) *mtproto.Peer {
	return mtproto.MakeTLPeerUser(&mtproto.Peer{
		UserId: peerId,
	}).To_Peer()
}

func MakePeerChat(peerId int32) *mtproto.Peer {
	return mtproto.MakeTLPeerChat(&mtproto.Peer{
		ChatId: peerId,
	}).To_Peer()
}

func MakePeerChannel(peerId int32) *mtproto.Peer {
	return mtproto.MakeTLPeerChannel(&mtproto.Peer{
		ChannelId: peerId,
	}).To_Peer()
}

func MakeUserPeerUtil(peerId int32) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_USER,
		PeerId:   peerId,
	}
}

func MakeChatPeerUtil(peerId int32) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_CHAT,
		PeerId:   peerId,
	}
}

func MakeChannelPeerUtil(peerId int32) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_CHANNEL,
		PeerId:   peerId,
	}
}
