package model

import (
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
)

const (
	_maxChatId = 1073741824 - 1
)

const (
	ChatMemberNormal  = 0
	ChatMemberCreator = 1
	ChatMemberAdmin   = 2
	ChatMemberBanned  = 4
)

const (
	ChatMemberStateNormal   = 0 // normal
	ChatMemberStateLeft     = 1 // left
	ChatMemberStateKicked   = 2 // kicked
	ChatMemberStateMigrated = 3 // migrated

	//ChatMemberStateAdmin    = 1 // normal
	//ChatMemberStateCreator  = 2 // normal
	//ChatMemberStateBanned   = 3 // kicked
)

var (
	// Cache
	ExportedChatInviteEmpty = mtproto.MakeTLChatInviteEmpty(&mtproto.ExportedChatInvite{}).To_ExportedChatInvite()
)

func IsChatId(id int32) bool {
	return id <= _maxChatId
}

func SplitChatAndChannelIdList(idList []int32) (chatIdList, channelIdList []int32) {
	for _, id := range idList {
		if IsChatId(id) {
			chatIdList = append(chatIdList, id)
		} else {
			channelIdList = append(channelIdList, id)
		}
	}
	return
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
type ImmutableChat struct {
	Id                  int32                     `json:"id,omitempty"`
	Creator             int32                     `json:"creator,omitempty"`
	Title               string                    `json:"title,omitempty"`
	Photo               *mtproto.ChatPhoto        `json:"photo,omitempty"`
	Deactivated         int8                      `json:"deactivated,omitempty"`
	ParticipantsCount   int32                     `json:"participants_count,omitempty"`
	Date                int32                     `json:"date,omitempty"`
	Version             int32                     `json:"version,omitempty"`
	MigratedTo          *mtproto.InputChannel     `json:"migrated_to,omitempty"`
	DefaultBannedRights *mtproto.ChatBannedRights `json:"default_banned_rights,omitempty"`
	CanSetUsername      bool                      `json:"can_set_username,omitempty"`
	About               string                    `json:"about,omitempty"`
	Notice              string                    `json:"notice,omitempty"`
	ChatPhoto           *mtproto.Photo            `json:"chat_photo,omitempty"`
	Link                string                    `json:"link,omitempty"`
	//ExportedInvite      *mtproto.ExportedChatInvite `json:"exported_invite,omitempty"`
	BotInfo []*mtproto.BotInfo `json:"bot_info,omitempty"`
}

func (m *ImmutableChat) CanViewMessages() bool {
	return !m.DefaultBannedRights.GetViewMessages()
}

func (m *ImmutableChat) CanSendMessages() bool {
	return !m.DefaultBannedRights.GetSendMessages()
}

func (m *ImmutableChat) CanSendMedia() bool {
	return !m.DefaultBannedRights.GetSendMedia()
}

func (m *ImmutableChat) CanSendStickers() bool {
	return !m.DefaultBannedRights.GetSendStickers()
}

func (m *ImmutableChat) CanSendGifs() bool {
	return !m.DefaultBannedRights.GetSendGifs()
}

func (m *ImmutableChat) CanSendGames() bool {
	return !m.DefaultBannedRights.GetSendGames()
}

func (m *ImmutableChat) CanSendInline() bool {
	return !m.DefaultBannedRights.GetSendInline()
}

func (m *ImmutableChat) CanEmbedLinks() bool {
	return !m.DefaultBannedRights.GetEmbedLinks()
}

func (m *ImmutableChat) CanSendPolls() bool {
	return !m.DefaultBannedRights.GetSendPolls()
}

func (m *ImmutableChat) CanChangeInfo() bool {
	return !m.DefaultBannedRights.GetChangeInfo()
}

func (m *ImmutableChat) CanInviteUsers() bool {
	return !m.DefaultBannedRights.GetInviteUsers()
}

func (m *ImmutableChat) CanPinMessages() bool {
	return !m.DefaultBannedRights.GetPinMessages()
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MutableChat struct {
	Chat         *ImmutableChat
	Participants map[int32]*ImmutableChatParticipant
}

func (m *MutableChat) GetImmutableChatParticipant(id int32) (participant *ImmutableChatParticipant) {
	participant, _ = m.Participants[id]
	return
}

func (m *MutableChat) CheckParticipantExist(id int32) bool {
	_, ok := m.Participants[id]
	return ok
}

func (m *MutableChat) AddChatParticipant(p *ImmutableChatParticipant) {
	m.Participants[p.ChatParticipant.UserId] = p
}

func (m *MutableChat) ToChatParticipants(selfId int32) (participants *mtproto.ChatParticipants) {
	if selfId != 0 {
		me := m.GetImmutableChatParticipant(selfId)
		if !me.IsChatMemberStateNormal() {
			//
			participants = mtproto.MakeTLChatParticipantsForbidden(&mtproto.ChatParticipants{
				ChatId:          m.Chat.Id,
				SelfParticipant: me.ChatParticipant,
			}).To_ChatParticipants()
			return
		}
	}

	participants = mtproto.MakeTLChatParticipants(&mtproto.ChatParticipants{
		ChatId:       m.Chat.Id,
		Participants: make([]*mtproto.ChatParticipant, 0, len(m.Participants)),
		Version:      m.Chat.Version,
	}).To_ChatParticipants()

	for _, cp := range m.Participants {
		if cp.IsChatMemberStateNormal() {
			participants.Participants = append(participants.Participants, cp.ChatParticipant)
		}
	}
	return
}

func (m *MutableChat) ToUnsafeChatFull(id int32) (chatFull *mtproto.ChatFull) {
	chatFull = mtproto.MakeTLChatFull(&mtproto.ChatFull{
		CanSetUsername: false,
		Id:             m.Chat.Id,
		About:          m.Chat.About,
		Notice:         m.Chat.Notice,
		ChatPhoto:      m.Chat.ChatPhoto,
		Participants:   m.ToChatParticipants(id),
		NotifySettings: nil,
		ExportedInvite: ExportedChatInviteEmpty,
		BotInfo:        m.Chat.BotInfo,
		PinnedMsgId:    nil,
		FolderId:       nil,
	}).To_ChatFull()

	if m.Chat.Link != "" {
		chatFull.ExportedInvite = mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
			Link: env2.T_ME + "/joinchat?link=" + m.Chat.Link,
		}).To_ExportedChatInvite()
	}

	if me, ok := m.Participants[id]; ok {
		if me.IsChatMemberCreator() {
			chatFull.CanSetUsername = true
		}
		if me.PinnedMsgId > 0 {
			chatFull.PinnedMsgId = &types.Int32Value{Value: me.PinnedMsgId}
		}
		chatFull.FolderId = me.Dialog.FolderId
		chatFull.NotifySettings = me.Dialog.NotifySettings
	}
	return chatFull
}

func (m *MutableChat) ToUnsafeChat(id int32) (chat *mtproto.Chat) {
	var (
		ok bool
		me *ImmutableChatParticipant
	)

	if me, ok = m.Participants[id]; !ok {
		chat = mtproto.MakeTLChatEmpty(&mtproto.Chat{}).To_Chat()
		return
	}

	if me.IsChatMemberStateKicked() {
		chat = m.ToChatForbidden()
		return
	}

	chat = mtproto.MakeTLChat(&mtproto.Chat{
		Creator:                         me.IsChatMemberCreator(),
		Kicked:                          false,
		Left:                            me.IsChatMemberStateLeft(),
		Deactivated:                     false,
		Id:                              m.Chat.Id,
		Title:                           m.Chat.Title,
		Photo:                           m.Chat.Photo,
		ParticipantsCount_INT32:         m.Chat.ParticipantsCount,
		Date:                            m.Chat.Date,
		Version:                         m.Chat.Version,
		MigratedTo:                      m.Chat.MigratedTo,
		AdminRights_FLAGCHATADMINRIGHTS: me.AdminRights,
		DefaultBannedRights:             m.Chat.DefaultBannedRights,
	}).To_Chat()

	if chat.MigratedTo != nil {
		chat.Deactivated = true
		chat.ParticipantsCount_INT32 = 0
		chat.AdminRights_FLAGCHATADMINRIGHTS = nil
		chat.DefaultBannedRights = nil
	}
	return
}

func (m *MutableChat) ToChatForbidden() (chat *mtproto.Chat) {
	chat = mtproto.MakeTLChatForbidden(&mtproto.Chat{
		Id:    m.Chat.Id,
		Title: m.Chat.Title,
	}).To_Chat()
	return
}

func (m *MutableChat) ToChatParticipantIdList() []int32 {
	idList := make([]int32, 0, len(m.Participants))
	for _, cp := range m.Participants {
		if cp.IsChatMemberNormal() {
			idList = append(idList, cp.ChatParticipant.UserId)
		}
	}
	return idList
}

func (m *MutableChat) MakeMessageService(fromId int32, action *mtproto.MessageAction) *mtproto.Message {
	message := mtproto.MakeTLMessageService(&mtproto.Message{
		Out:             true,
		Mentioned:       false,
		MediaUnread:     false,
		Silent:          false,
		Post:            false,
		Legacy:          false,
		Id:              0,
		FromId_FLAGPEER: MakePeerUser(fromId),
		ToId:            MakePeerChat(m.Chat.Id),
		ReplyTo:         nil,
		Date:            int32(time.Now().Unix()),
		Action:          action,
	}).To_Message()
	return message
}

func (m *MutableChat) MakePinnedMessageService(fromId, id int32) *mtproto.Message {
	message := mtproto.MakeTLMessageService(&mtproto.Message{
		Out:             true,
		Mentioned:       false,
		MediaUnread:     false,
		Silent:          true,
		Post:            false,
		Legacy:          false,
		Id:              0,
		FromId_FLAGPEER: MakePeerUser(fromId),
		ToId:            MakePeerChat(m.Chat.Id),
		ReplyTo:         nil,
		Date:            int32(time.Now().Unix()),
		Action:          MakeMessageActionPinMessage(),
	}).To_Message()
	if id != 0 {
		message.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: id,
			}).To_MessageReplyHeader()
	}
	return message
}

func (m *MutableChat) Walk(visit func(userId int32, participant *ImmutableChatParticipant) error) {
	if visit == nil {
		return
	}
	for k, v := range m.Participants {
		visit(k, v)
	}
}
