package model

import (
	"open.chat/mtproto"
)

type ImmutableChatParticipant struct {
	Id              int32                    `json:"id,omitempty,omitempty"`
	State           int                      `json:"state,omitempty,omitempty"`
	ChatParticipant *mtproto.ChatParticipant `json:"chat_participant,omitempty"`
	AdminRights     *mtproto.ChatAdminRights `json:"admin_rights,omitempty"`
	PinnedMsgId     int32                    `json:"pinned_msg_id,omitempty"`
	Dialog          *mtproto.Dialog          `json:"dialog,omitempty"`
	Chat            *ImmutableChat
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChatParticipant) IsChatMemberNormal() bool {
	return m.ChatParticipant.PredicateName == mtproto.Predicate_chatParticipant
}

func (m *ImmutableChatParticipant) IsChatMemberAdmin() bool {
	return m.ChatParticipant.PredicateName == mtproto.Predicate_chatParticipantAdmin
}

func (m *ImmutableChatParticipant) IsChatMemberCreator() bool {
	return m.ChatParticipant.PredicateName == mtproto.Predicate_chatParticipantCreator
}

func (m *ImmutableChatParticipant) IsChatCreatorOrAdmin() bool {
	return m.IsChatMemberCreator() || m.IsChatMemberAdmin()
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChatParticipant) IsChatMemberStateNormal() bool {
	return m.State == ChatMemberStateNormal
}

func (m *ImmutableChatParticipant) IsChatMemberStateLeft() bool {
	return m.State == ChatMemberStateLeft
}

func (m *ImmutableChatParticipant) IsChatMemberStateKicked() bool {
	return m.State == ChatMemberStateKicked
}

func (m *ImmutableChatParticipant) IsChatMemberStateMigrated() bool {
	return m.State == ChatMemberStateMigrated
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChatParticipant) CanViewMessages() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanViewMessages()
}

func (m *ImmutableChatParticipant) CanSendMessages() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendMessages()
}

func (m *ImmutableChatParticipant) CanSendMedia() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendMedia()
}

func (m *ImmutableChatParticipant) CanSendStickers() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendStickers()
}

func (m *ImmutableChatParticipant) CanSendGifs() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendGifs()
}

func (m *ImmutableChatParticipant) CanSendGames() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendGames()
}

func (m *ImmutableChatParticipant) CanSendInline() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendInline()
}

func (m *ImmutableChatParticipant) CanEmbedLinks() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanEmbedLinks()
}

func (m *ImmutableChatParticipant) CanSendPolls() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanSendPolls()
}

// merge ChatAdminRights
func (m *ImmutableChatParticipant) CanChangeInfo() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanChangeInfo()
}

func (m *ImmutableChatParticipant) CanInviteUsers() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanInviteUsers()
}

func (m *ImmutableChatParticipant) CanPinMessages() bool {
	if m.IsChatCreatorOrAdmin() {
		return true
	}
	return m.Chat.CanPinMessages()
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChatParticipant) CanAdminChangeInfo() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetChangeInfo()
}

func (m *ImmutableChatParticipant) CanAdminPostMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetPinMessages()
}

func (m *ImmutableChatParticipant) CanAdminEditMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetEditMessages()
}

func (m *ImmutableChatParticipant) CanAdminDeleteMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetDeleteMessages()
}

func (m *ImmutableChatParticipant) CanAdminBanUsers() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetBanUsers()
}

func (m *ImmutableChatParticipant) CanAdminInviteUsers() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetInviteUsers()
}

func (m *ImmutableChatParticipant) CanAdminPinMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetPinMessages()
}

func (m *ImmutableChatParticipant) CanAdminAddAdmins() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetAddAdmins()
}

func (m *ImmutableChatParticipant) HasAnonymous() bool {
	return m.IsChatCreatorOrAdmin() && m.AdminRights.GetAnonymous()
}
