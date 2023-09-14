package model

import (
	"math"

	"open.chat/mtproto"
)

const (
	// OK is returned on success.
	BAN_VIEW_MESSAGES int32 = 1 << 0
	BAN_SEND_MESSAGES int32 = 1 << 1
	BAN_SEND_MEDIA    int32 = 1 << 2
	BAN_SEND_STICKERS int32 = 1 << 3
	BAN_SEND_GIFS     int32 = 1 << 4
	BAN_SEND_GAMES    int32 = 1 << 5
	BAN_SEND_INLINE   int32 = 1 << 6
	BAN_EMBED_LINKS   int32 = 1 << 7
	BAN_SEND_POLLS    int32 = 1 << 8
	BAN_CHANGE_INFO   int32 = 1 << 9
	BAN_INVITE_USERS  int32 = 1 << 10
	BAN_PIN_MESSAGES  int32 = 1 << 11
	BAN_BANNED_ALL    int32 = 1<<12 - 1
	// UNTIL_DATE 		BannedRights = 1 << 32
)

type ChatBannedRights struct {
	Rights    int32 `json:"rights"`
	UntilDate int32 `json:"until_date"`
}

func MakeChatBannedRights(bannedRights *mtproto.ChatBannedRights) ChatBannedRights {
	var (
		rights    int32 = 0
		untilDate int32 = 0
	)

	if bannedRights.GetViewMessages() {
		rights |= BAN_VIEW_MESSAGES
	}
	if bannedRights.GetSendMessages() {
		rights |= BAN_SEND_MESSAGES
	}
	if bannedRights.GetSendMedia() {
		rights |= BAN_SEND_MEDIA
	}
	if bannedRights.GetSendStickers() {
		rights |= BAN_SEND_STICKERS
	}
	if bannedRights.GetSendGifs() {
		rights |= BAN_SEND_GIFS
	}
	if bannedRights.GetSendGames() {
		rights |= BAN_SEND_GAMES
	}
	if bannedRights.GetSendInline() {
		rights |= BAN_SEND_INLINE
	}
	if bannedRights.GetEmbedLinks() {
		rights |= BAN_EMBED_LINKS
	}
	if bannedRights.GetSendPolls() {
		rights |= BAN_SEND_POLLS
	}
	if bannedRights.GetChangeInfo() {
		rights |= BAN_CHANGE_INFO
	}
	if bannedRights.GetInviteUsers() {
		rights |= BAN_INVITE_USERS
	}
	if bannedRights.GetPinMessages() {
		rights |= BAN_PIN_MESSAGES
	}

	untilDate = bannedRights.GetUntilDate()
	if untilDate == 0 {
		untilDate = math.MaxInt32
	}

	return ChatBannedRights{
		Rights:    rights,
		UntilDate: untilDate,
	}
}

func MakeChannelBannedRights(bannedRights *mtproto.ChannelBannedRights) ChatBannedRights {
	var (
		rights    int32 = 0
		untilDate int32 = 0
	)

	if bannedRights.GetViewMessages() {
		rights |= BAN_VIEW_MESSAGES
	}
	if bannedRights.GetSendMessages() {
		rights |= BAN_SEND_MESSAGES
	}
	if bannedRights.GetSendMedia() {
		rights |= BAN_SEND_MEDIA
	}
	if bannedRights.GetSendStickers() {
		rights |= BAN_SEND_STICKERS
	}
	if bannedRights.GetSendGifs() {
		rights |= BAN_SEND_GIFS
	}
	if bannedRights.GetSendGames() {
		rights |= BAN_SEND_GAMES
	}
	if bannedRights.GetSendInline() {
		rights |= BAN_SEND_INLINE
	}
	if bannedRights.GetEmbedLinks() {
		rights |= BAN_EMBED_LINKS
	}

	untilDate = bannedRights.GetUntilDate()
	if untilDate == 0 {
		untilDate = math.MaxInt32
	}

	return ChatBannedRights{
		Rights:    rights,
		UntilDate: untilDate,
	}
}

func (m ChatBannedRights) ToChatBannedRights() *mtproto.ChatBannedRights {
	return mtproto.MakeTLChatBannedRights(&mtproto.ChatBannedRights{
		ViewMessages: (m.Rights & BAN_VIEW_MESSAGES) != 0,
		SendMessages: (m.Rights & BAN_SEND_MESSAGES) != 0,
		SendMedia:    (m.Rights & BAN_SEND_MEDIA) != 0,
		SendStickers: (m.Rights & BAN_SEND_STICKERS) != 0,
		SendGifs:     (m.Rights & BAN_SEND_GIFS) != 0,
		SendGames:    (m.Rights & BAN_SEND_GAMES) != 0,
		SendInline:   (m.Rights & BAN_SEND_INLINE) != 0,
		EmbedLinks:   (m.Rights & BAN_EMBED_LINKS) != 0,
		SendPolls:    (m.Rights & BAN_SEND_POLLS) != 0,
		ChangeInfo:   (m.Rights & BAN_CHANGE_INFO) != 0,
		InviteUsers:  (m.Rights & BAN_INVITE_USERS) != 0,
		PinMessages:  (m.Rights & BAN_PIN_MESSAGES) != 0,
		UntilDate:    m.UntilDate,
	}).To_ChatBannedRights()
}

func (m ChatBannedRights) ToChannelBannedRights() *mtproto.ChannelBannedRights {
	return mtproto.MakeTLChannelBannedRights(&mtproto.ChannelBannedRights{
		ViewMessages: (m.Rights & BAN_VIEW_MESSAGES) != 0,
		SendMessages: (m.Rights & BAN_SEND_MESSAGES) != 0,
		SendMedia:    (m.Rights & BAN_SEND_MEDIA) != 0,
		SendStickers: (m.Rights & BAN_SEND_STICKERS) != 0,
		SendGifs:     (m.Rights & BAN_SEND_GIFS) != 0,
		SendGames:    (m.Rights & BAN_SEND_GAMES) != 0,
		SendInline:   (m.Rights & BAN_SEND_INLINE) != 0,
		EmbedLinks:   (m.Rights & BAN_EMBED_LINKS) != 0,
		UntilDate:    m.UntilDate,
	}).To_ChannelBannedRights()
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m ChatBannedRights) NoBanRights() bool {
	return m.Rights == 0 && m.UntilDate == math.MaxInt32
}

func (m ChatBannedRights) IsKick() bool {
	return m.Rights&BAN_VIEW_MESSAGES != 0 && m.UntilDate == math.MaxInt32
}

func (m ChatBannedRights) IsBan(date int32) bool {
	return m.Rights != 0 && date <= m.UntilDate
}

func (m ChatBannedRights) IsRestrict(date int32) bool {
	return !m.IsKick() && m.Rights != 0 && date <= m.UntilDate
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m ChatBannedRights) CanViewMessages(date int32) bool {
	return m.Rights&BAN_VIEW_MESSAGES == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendMessages(date int32) bool {
	return m.Rights&BAN_SEND_MESSAGES == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendMedia(date int32) bool {
	return m.Rights&BAN_SEND_MEDIA == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendStickers(date int32) bool {
	return m.Rights&BAN_SEND_STICKERS == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendGifs(date int32) bool {
	return m.Rights&BAN_SEND_GIFS == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendGames(date int32) bool {
	return m.Rights&BAN_SEND_GAMES == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendInline(date int32) bool {
	return m.Rights&BAN_SEND_INLINE == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanEmbedLinks(date int32) bool {
	return m.Rights&BAN_EMBED_LINKS == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanSendPolls(date int32) bool {
	return m.Rights&BAN_SEND_POLLS == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanChangeInfo(date int32) bool {
	return m.Rights&BAN_CHANGE_INFO == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanInviteUsers(date int32) bool {
	return m.Rights&BAN_INVITE_USERS == 0 || date >= m.UntilDate
}

func (m ChatBannedRights) CanPinMessage(date int32) bool {
	return m.Rights&BAN_PIN_MESSAGES == 0 || date >= m.UntilDate
}

// ////////////////////////////////////////////////////////////////////////////////////////
func (m ChatAdminRights) DisallowChannel() bool {
	// return m.CanPostMessages() || m.CanEditMessages()
	return false
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func FromChannelBannedRights(bannedRights *mtproto.TLChannelBannedRights) (int32, int32) {
	var (
		rights    int32 = 0
		untilDate int32 = 0
	)

	if bannedRights.GetViewMessages() {
		rights |= BAN_VIEW_MESSAGES
	}
	if bannedRights.GetSendMessages() {
		rights |= BAN_SEND_MESSAGES
	}
	if bannedRights.GetSendMedia() {
		rights |= BAN_SEND_MEDIA
	}
	if bannedRights.GetSendStickers() {
		rights |= BAN_SEND_STICKERS
	}
	if bannedRights.GetSendGifs() {
		rights |= BAN_SEND_GIFS
	}
	if bannedRights.GetSendGames() {
		rights |= BAN_SEND_GAMES
	}
	if bannedRights.GetSendInline() {
		rights |= BAN_SEND_INLINE
	}
	if bannedRights.GetEmbedLinks() {
		rights |= BAN_EMBED_LINKS
	}

	untilDate = bannedRights.GetUntilDate()
	if rights == 0 {
		untilDate = 0
	} else if untilDate == 0 {
		untilDate = math.MaxInt32
	}

	return rights, untilDate
}

func FromChatBannedRights(bannedRights *mtproto.TLChatBannedRights) (int32, int32) {
	var (
		rights    int32 = 0
		untilDate int32 = 0
	)

	if bannedRights.GetViewMessages() {
		rights |= BAN_VIEW_MESSAGES
	}
	if bannedRights.GetSendMessages() {
		rights |= BAN_SEND_MESSAGES
	}
	if bannedRights.GetSendMedia() {
		rights |= BAN_SEND_MEDIA
	}
	if bannedRights.GetSendStickers() {
		rights |= BAN_SEND_STICKERS
	}
	if bannedRights.GetSendGifs() {
		rights |= BAN_SEND_GIFS
	}
	if bannedRights.GetSendGames() {
		rights |= BAN_SEND_GAMES
	}
	if bannedRights.GetSendInline() {
		rights |= BAN_SEND_INLINE
	}
	if bannedRights.GetEmbedLinks() {
		rights |= BAN_EMBED_LINKS
	}
	if bannedRights.GetSendPolls() {
		rights |= BAN_SEND_POLLS
	}
	if bannedRights.GetChangeInfo() {
		rights |= BAN_CHANGE_INFO
	}
	if bannedRights.GetInviteUsers() {
		rights |= BAN_INVITE_USERS
	}
	if bannedRights.GetPinMessages() {
		rights |= BAN_PIN_MESSAGES
	}

	untilDate = bannedRights.GetUntilDate()
	if rights == 0 {
		untilDate = 0
	} else if untilDate == 0 {
		untilDate = math.MaxInt32
	}

	return rights, untilDate
}

func ToChatBannedRights(rights, untilDate int32) *mtproto.ChatBannedRights {
	if rights == 0 && untilDate == 0 {
		return nil
	}

	bannedRights := mtproto.MakeTLChatBannedRights(nil)

	if (rights & BAN_VIEW_MESSAGES) != 0 {
		bannedRights.SetViewMessages(true)
	}
	if (rights & BAN_SEND_MESSAGES) != 0 {
		bannedRights.SetSendMessages(true)
	}
	if (rights & BAN_SEND_MEDIA) != 0 {
		bannedRights.SetSendMedia(true)
	}
	if (rights & BAN_SEND_STICKERS) != 0 {
		bannedRights.SetSendStickers(true)
	}
	if (rights & BAN_SEND_GIFS) != 0 {
		bannedRights.SetSendGifs(true)
	}
	if (rights & BAN_SEND_GAMES) != 0 {
		bannedRights.SetSendGames(true)
	}
	if (rights & BAN_SEND_INLINE) != 0 {
		bannedRights.SetSendInline(true)
	}
	if (rights & BAN_EMBED_LINKS) != 0 {
		bannedRights.SetEmbedLinks(true)
	}
	if (rights & BAN_SEND_POLLS) != 0 {
		bannedRights.SetSendPolls(true)
	}
	if (rights & BAN_CHANGE_INFO) != 0 {
		bannedRights.SetChangeInfo(true)
	}
	if (rights & BAN_INVITE_USERS) != 0 {
		bannedRights.SetInviteUsers(true)
	}
	if (rights & BAN_PIN_MESSAGES) != 0 {
		bannedRights.SetPinMessages(true)
	}

	bannedRights.SetUntilDate(untilDate)
	return bannedRights.To_ChatBannedRights()
}

func ToChannelBannedRights(rights, untilDate int32) *mtproto.ChannelBannedRights {
	if rights == 0 && untilDate == 0 {
		return nil
	}

	bannedRights := mtproto.MakeTLChannelBannedRights(nil)

	if (rights & BAN_VIEW_MESSAGES) != 0 {
		bannedRights.SetViewMessages(true)
	}
	if (rights & BAN_SEND_MESSAGES) != 0 {
		bannedRights.SetSendMessages(true)
	}
	if (rights & BAN_SEND_MEDIA) != 0 {
		bannedRights.SetSendMedia(true)
	}
	if (rights & BAN_SEND_STICKERS) != 0 {
		bannedRights.SetSendStickers(true)
	}
	if (rights & BAN_SEND_GIFS) != 0 {
		bannedRights.SetSendGifs(true)
	}
	if (rights & BAN_SEND_GAMES) != 0 {
		bannedRights.SetSendGames(true)
	}
	if (rights & BAN_SEND_INLINE) != 0 {
		bannedRights.SetSendInline(true)
	}
	if (rights & BAN_EMBED_LINKS) != 0 {
		bannedRights.SetEmbedLinks(true)
	}

	bannedRights.SetUntilDate(untilDate)
	return bannedRights.To_ChannelBannedRights()
}
