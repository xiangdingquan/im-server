package model

import (
	"open.chat/mtproto"
)

const (
	// OK is returned on success.
	ADMIN_CHANGE_INFO     int32 = 1 << 0
	ADMIN_POST_MESSAGES   int32 = 1 << 1
	ADMIN_EDIT_MESSAGES   int32 = 1 << 2
	ADMIN_DELETE_MESSAGES int32 = 1 << 3
	ADMIN_BAN_USERS       int32 = 1 << 4
	ADMIN_INVITE_USERS    int32 = 1 << 5
	ADMIN_INVITE_LINK     int32 = 1 << 6
	ADMIN_PIN_MESSAGES    int32 = 1 << 7
	ADMIN_ADD_ADMINS      int32 = 1 << 8
	ADMIN_MANAGE_CALL     int32 = 1 << 9
	ADMIN_ANONYMOUS       int32 = 1 << 10
)

type ChatAdminRights int32

func MakeChatAdminRights(adminRights *mtproto.ChatAdminRights) ChatAdminRights {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= ADMIN_CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= ADMIN_POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= ADMIN_EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= ADMIN_DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= ADMIN_BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= ADMIN_INVITE_USERS
	}
	if adminRights.GetPinMessages() {
		rights |= ADMIN_PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADMIN_ADD_ADMINS
	}
	if adminRights.GetAnonymous() {
		rights |= ADMIN_ANONYMOUS
	}

	// FIX BAN_USERS
	if rights != 0 {
		rights |= ADMIN_BAN_USERS
	}

	return ChatAdminRights(rights)
}

func MakeChannelAdminRights(adminRights *mtproto.ChannelAdminRights) ChatAdminRights {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= ADMIN_CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= ADMIN_POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= ADMIN_EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= ADMIN_DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= ADMIN_BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= ADMIN_INVITE_USERS
	}
	if adminRights.GetInviteLink() {
		rights |= ADMIN_INVITE_LINK
	}
	if adminRights.GetPinMessages() {
		rights |= ADMIN_PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADMIN_ADD_ADMINS
	}
	if adminRights.GetManageCall() {
		rights |= ADMIN_MANAGE_CALL
	}

	return ChatAdminRights(rights)
}

func (m ChatAdminRights) ToChatAdminRights() *mtproto.ChatAdminRights {
	if int32(m) == 0 {
		return nil
	} else {
		return mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
			ChangeInfo:     int32(m)&ADMIN_CHANGE_INFO != 0,
			PostMessages:   int32(m)&ADMIN_POST_MESSAGES != 0,
			EditMessages:   int32(m)&ADMIN_EDIT_MESSAGES != 0,
			DeleteMessages: int32(m)&ADMIN_DELETE_MESSAGES != 0,
			BanUsers:       int32(m)&ADMIN_BAN_USERS != 0,
			InviteUsers:    int32(m)&ADMIN_INVITE_USERS != 0,
			PinMessages:    int32(m)&ADMIN_PIN_MESSAGES != 0,
			AddAdmins:      int32(m)&ADMIN_ADD_ADMINS != 0,
			Anonymous:      int32(m)&ADMIN_ANONYMOUS != 0,
		}).To_ChatAdminRights()
	}
}

func (m ChatAdminRights) ToChannelAdminRights() *mtproto.ChannelAdminRights {
	if int32(m) == 0 {
		return nil
	} else {
		return mtproto.MakeTLChannelAdminRights(&mtproto.ChannelAdminRights{
			ChangeInfo:     int32(m)&ADMIN_CHANGE_INFO != 0,
			PostMessages:   int32(m)&ADMIN_POST_MESSAGES != 0,
			EditMessages:   int32(m)&ADMIN_EDIT_MESSAGES != 0,
			DeleteMessages: int32(m)&ADMIN_DELETE_MESSAGES != 0,
			BanUsers:       int32(m)&ADMIN_BAN_USERS != 0,
			InviteUsers:    int32(m)&ADMIN_INVITE_USERS != 0,
			InviteLink:     int32(m)&ADMIN_INVITE_LINK != 0,
			PinMessages:    int32(m)&ADMIN_PIN_MESSAGES != 0,
			AddAdmins:      int32(m)&ADMIN_ADD_ADMINS != 0,
			ManageCall:     int32(m)&ADMIN_MANAGE_CALL != 0,
		}).To_ChannelAdminRights()
	}
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m ChatAdminRights) HasAdminRights() bool {
	return m != 0
}

func (m ChatAdminRights) CanChangeInfo() bool {
	return int32(m)&ADMIN_CHANGE_INFO != 0
}

func (m ChatAdminRights) CanPostMessages() bool {
	return int32(m)&ADMIN_POST_MESSAGES != 0
}

func (m ChatAdminRights) CanEditMessages() bool {
	return int32(m)&ADMIN_EDIT_MESSAGES != 0
}

func (m ChatAdminRights) CanDeleteMessages() bool {
	return int32(m)&ADMIN_DELETE_MESSAGES != 0
}

func (m ChatAdminRights) CanBanUsers() bool {
	return int32(m)&ADMIN_BAN_USERS != 0
}

func (m ChatAdminRights) CanInviteUsers() bool {
	return int32(m)&ADMIN_INVITE_USERS != 0
}

func (m ChatAdminRights) CanPinMessages() bool {
	return int32(m)&ADMIN_PIN_MESSAGES != 0
}

func (m ChatAdminRights) CanAddAdmins() bool {
	return int32(m)&ADMIN_ADD_ADMINS != 0
}

func (m ChatAdminRights) HasAnonymous() bool {
	return int32(m)&ADMIN_ANONYMOUS != 0
}

// ////////////////////////////////////////////////////////////////////////////////////////
func (m ChatAdminRights) DisallowMegagroup() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}

func (m ChatAdminRights) DisallowChat() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}

func FromChatAdminRights(adminRights *mtproto.ChatAdminRights) int32 {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= ADMIN_CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= ADMIN_POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= ADMIN_EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= ADMIN_DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= ADMIN_BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= ADMIN_INVITE_USERS
	}
	if adminRights.GetPinMessages() {
		rights |= ADMIN_PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADMIN_ADD_ADMINS
	}
	if adminRights.GetAnonymous() {
		rights |= ADMIN_ANONYMOUS
	}

	return rights
}

func ToChatAdminRights(rights int32) *mtproto.ChatAdminRights {
	if rights == 0 {
		return nil
	}

	adminRights := mtproto.MakeTLChatAdminRights(nil)

	if (rights & ADMIN_CHANGE_INFO) != 0 {
		adminRights.SetChangeInfo(true)
	}
	if (rights & ADMIN_POST_MESSAGES) != 0 {
		adminRights.SetPostMessages(true)
	}
	if (rights & ADMIN_EDIT_MESSAGES) != 0 {
		adminRights.SetEditMessages(true)
	}
	if (rights & ADMIN_DELETE_MESSAGES) != 0 {
		adminRights.SetDeleteMessages(true)
	}
	if (rights & ADMIN_BAN_USERS) != 0 {
		adminRights.SetBanUsers(true)
	}
	if (rights & ADMIN_INVITE_USERS) != 0 {
		adminRights.SetInviteUsers(true)
	}
	if (rights & ADMIN_PIN_MESSAGES) != 0 {
		adminRights.SetPinMessages(true)
	}
	if (rights & ADMIN_ADD_ADMINS) != 0 {
		adminRights.SetAddAdmins(true)
	}
	if (rights & ADMIN_ANONYMOUS) != 0 {
		adminRights.SetAnonymous(true)
	}

	return adminRights.To_ChatAdminRights()
}

func FromChannelAdminRights(adminRights *mtproto.ChannelAdminRights) int32 {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= ADMIN_CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= ADMIN_POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= ADMIN_EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= ADMIN_DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= ADMIN_BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= ADMIN_INVITE_USERS
	}
	if adminRights.GetInviteLink() {
		rights |= ADMIN_INVITE_LINK
	}
	if adminRights.GetPinMessages() {
		rights |= ADMIN_PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADMIN_ADD_ADMINS
	}
	if adminRights.GetManageCall() {
		rights |= ADMIN_MANAGE_CALL
	}

	return rights
}

func ToChannelAdminRights(rights int32) *mtproto.ChannelAdminRights {
	if rights == 0 {
		return nil
	}

	adminRights := mtproto.MakeTLChannelAdminRights(nil)

	if (rights & ADMIN_CHANGE_INFO) != 0 {
		adminRights.SetChangeInfo(true)
	}
	if (rights & ADMIN_POST_MESSAGES) != 0 {
		adminRights.SetPostMessages(true)
	}
	if (rights & ADMIN_EDIT_MESSAGES) != 0 {
		adminRights.SetEditMessages(true)
	}
	if (rights & ADMIN_DELETE_MESSAGES) != 0 {
		adminRights.SetDeleteMessages(true)
	}
	if (rights & ADMIN_BAN_USERS) != 0 {
		adminRights.SetBanUsers(true)
	}
	if (rights & ADMIN_INVITE_USERS) != 0 {
		adminRights.SetInviteUsers(true)
	}
	if (rights & ADMIN_INVITE_LINK) != 0 {
		adminRights.SetInviteLink(true)
	}
	if (rights & ADMIN_PIN_MESSAGES) != 0 {
		adminRights.SetPinMessages(true)
	}
	if (rights & ADMIN_ADD_ADMINS) != 0 {
		adminRights.SetAddAdmins(true)
	}
	if (rights & ADMIN_MANAGE_CALL) != 0 {
		adminRights.SetManageCall(true)
	}

	return adminRights.To_ChannelAdminRights()
}

func MakeDefaultChatAdminRights() *mtproto.ChatAdminRights {
	return mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
		ChangeInfo:     true,
		PostMessages:   false,
		EditMessages:   false,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      false,
		Anonymous:      false,
	}).To_ChatAdminRights()
}

func MakeChatCreatorAdminRights() *mtproto.ChatAdminRights {
	return mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
		ChangeInfo:     true,
		PostMessages:   true,
		EditMessages:   true,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      true,
		Anonymous:      false,
	}).To_ChatAdminRights()
}
