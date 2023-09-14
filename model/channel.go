package model

import (
	"math"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

const (
	_minChannelId = 1073741824
)

func IsChannelId(id int32) bool {
	return id >= _minChannelId
}

type ImmutableChannel struct {
	Id                   int32                        `json:"id,omitempty"`
	AccessHash           int64                        `json:"access_hash,omitempty"`
	SecretKeyId          int64                        `json:"secret_key_id,omitempty"`
	Title                string                       `json:"title,omitempty"`
	Username             string                       `json:"username,omitempty"`
	Photo                *mtproto.ChatPhoto           `json:"photo,omitempty"`
	CreatorId            int32                        `json:"creator_id,omitempty"`
	TopMessage           int32                        `json:"top_message,omitempty"`
	Broadcast            bool                         `json:"broadcast,omitempty"`
	Democracy            bool                         `json:"democracy,omitempty"`
	Verified             bool                         `json:"verified,omitempty"`
	Megagroup            bool                         `json:"megagroup,omitempty"`
	Signatures           bool                         `json:"signatures,omitempty"`
	Min                  bool                         `json:"min,omitempty"`
	Scam                 bool                         `json:"scam,omitempty"`
	HasLink              bool                         `json:"has_link,omitempty"`
	HasGeo               bool                         `json:"has_geo,omitempty"`
	SlowmodeEnabled      bool                         `json:"slowmode_enabled,omitempty"`
	Date                 int32                        `json:"date,omitempty"`
	Version              int32                        `json:"version,omitempty"`
	DefaultBannedRights  ChatBannedRights             `json:"default_banned_rights,omitempty"`
	ParticipantsCount    int32                        `json:"participants_count,omitempty"`
	About                string                       `json:"about,omitempty,omitempty"`
	Notice               string                       `json:"notice,omitempty"`
	HiddenPrehistory     bool                         `json:"hidden_prehistory,omitempty"`
	AdminsCount          int32                        `json:"admins_count,omitempty"`
	KickedCount          int32                        `json:"kicked_count,omitempty"`
	BannedCount          int32                        `json:"banned_count,omitempty"`
	OnlineCount          int32                        `json:"online_count,omitempty"`
	ChatPhoto            *mtproto.Photo               `json:"chat_photo,omitempty,omitempty"`
	Link                 string                       `json:"link,omitempty"`
	BotInfo              []*mtproto.BotInfo           `json:"bot_info,omitempty,omitempty"`
	MigratedFromChatId   int32                        `json:"migrated_from_chat_id,omitempty"`
	MigratedFromMaxId    int32                        `json:"migrated_from_max_id,omitempty"`
	PinnedMsgId          int32                        `json:"pinned_msg_id,omitempty"`
	StickerSet           *mtproto.StickerSet          `json:"sticker_set,omitempty"`
	LinkedChatId         int32                        `json:"linked_chat_id,omitempty"`
	Location             *mtproto.ChannelLocation     `json:"location,omitempty"`
	SlowmodeSeconds      int32                        `json:"slowmode_seconds,omitempty"`
	SlowmodeNextSendDate int32                        `json:"slowmode_next_send_date,omitempty"`
	ReadOutboxMaxId      int32                        `json:"read_outbox_max_id,omitempty"`
	Pts                  int32                        `json:"pts,omitempty"`
	Deleted              bool                         `json:"deleted,omitempty"`
	Restrictions         []*mtproto.RestrictionReason `json:"restrictions,omitempty"`
}

func (m *ImmutableChannel) canViewMessages() bool {
	return m.DefaultBannedRights.CanViewMessages(math.MaxInt32)
}

func (m *ImmutableChannel) canSendMessages() bool {
	return m.DefaultBannedRights.CanSendMessages(math.MaxInt32)
}

func (m *ImmutableChannel) canSendMedia() bool {
	return m.DefaultBannedRights.CanSendMedia(math.MaxInt32)
}

func (m *ImmutableChannel) canSendStickers() bool {
	return m.DefaultBannedRights.CanSendStickers(math.MaxInt32)
}

func (m *ImmutableChannel) canSendGifs() bool {
	return m.DefaultBannedRights.CanSendGifs(math.MaxInt32)
}

func (m *ImmutableChannel) canSendGames() bool {
	return m.DefaultBannedRights.CanSendGames(math.MaxInt32)
}

func (m *ImmutableChannel) canSendInline() bool {
	return m.DefaultBannedRights.CanSendInline(math.MaxInt32)
}

func (m *ImmutableChannel) canEmbedLinks() bool {
	return m.DefaultBannedRights.CanEmbedLinks(math.MaxInt32)
}

func (m *ImmutableChannel) canSendPolls() bool {
	return m.DefaultBannedRights.CanSendPolls(math.MaxInt32)
}

func (m *ImmutableChannel) canChangeInfo() bool {
	return m.DefaultBannedRights.CanChangeInfo(math.MaxInt32)
}

func (m *ImmutableChannel) canInviteUsers() bool {
	return m.DefaultBannedRights.CanInviteUsers(math.MaxInt32)
}

func (m *ImmutableChannel) canPinMessages() bool {
	return m.DefaultBannedRights.CanPinMessage(math.MaxInt32)
}

func (m *ImmutableChannel) IsDemocracy() bool {
	return m.Democracy
}

func (m *ImmutableChannel) IsSignatures() bool {
	return m.Signatures
}

func (m *ImmutableChannel) IsMegagroup() bool {
	return m.Megagroup
}

func (m *ImmutableChannel) IsBroadcast() bool {
	return m.Broadcast
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////
type MutableChannel struct {
	Channel      *ImmutableChannel
	Participants map[int32]*ImmutableChannelParticipant
}

func (m *MutableChannel) GetChannelId() int32 {
	return m.Channel.Id
}

func (m *MutableChannel) GetAccessHash() int64 {
	return m.Channel.AccessHash
}

func (m *MutableChannel) GetId() int32 {
	return m.Channel.Id
}

func (m *MutableChannel) GetImmutableChannelParticipant(id int32) (participant *ImmutableChannelParticipant) {
	participant = m.Participants[id]
	return
}

func (m *MutableChannel) AddChannelParticipant(p *ImmutableChannelParticipant) {
	m.Participants[p.UserId] = p
}

func (m *MutableChannel) ToUnsafeChat(selfUserId int32) (channel *mtproto.Chat) {
	var (
		me *ImmutableChannelParticipant
	)

	if m.Channel.Deleted {
		channel = m.ToChannelForbidden()
		return
	}

	me = m.Participants[selfUserId]
	if me != nil && me.IsKicked() {
		channel = m.ToChannelForbidden()
		return
	}

	channel = mtproto.MakeTLChannel(&mtproto.Chat{
		Creator:              me != nil && me.IsCreator(),
		Left:                 false,
		Broadcast:            m.Channel.Broadcast,
		Verified:             m.Channel.Verified,
		Megagroup:            m.Channel.Megagroup,
		Restricted:           false,
		Signatures:           m.Channel.Signatures,
		Min:                  m.Channel.Min,
		Scam:                 m.Channel.Scam,
		HasLink:              m.Channel.HasLink,
		HasGeo:               m.Channel.HasGeo,
		SlowmodeEnabled:      m.Channel.SlowmodeEnabled,
		Id:                   m.Channel.Id,
		AccessHash_FLAGINT64: &types.Int64Value{Value: m.Channel.AccessHash},
		Title:                m.Channel.Title,
		Username:             nil,
		Photo:                m.Channel.Photo,
		Date:                 m.Channel.Date,
		Version:              m.Channel.Version,
		RestrictionReason_FLAGVECTORRESTRICTIONREASON: nil,
		AdminRights_FLAGCHATADMINRIGHTS:               nil,
		BannedRights_FLAGCHATBANNEDRIGHTS:             nil,
		DefaultBannedRights:                           nil,
		ParticipantsCount_FLAGINT32:                   nil,
	}).To_Chat()

	// Left:                 me == nil || (me != nil && !me.IsStateOk()),

	if me == nil {
		channel.Left = true
	} else if me.IsLeft() {
		channel.Left = true
	}
	// Username
	if m.Channel.Username != "" {
		channel.Username = &types.StringValue{Value: m.Channel.Username}
	}

	// admin_rights
	if me != nil && me.IsCreatorOrAdmin() {
		if me.IsCreator() && m.Channel.IsBroadcast() {
		} else {
			channel.AdminRights_FLAGCHATADMINRIGHTS = me.AdminRights.ToChatAdminRights()
			channel.AdminRights_FLAGCHANNELADMINRIGHTS = me.AdminRights.ToChannelAdminRights()
		}
	}

	if me != nil {
		if me.IsCreatorOrAdmin() {
			if me.IsCreator() && m.Channel.IsBroadcast() {
			} else {
				channel.BannedRights_FLAGCHATBANNEDRIGHTS = me.BannedRights.ToChatBannedRights()
				channel.BannedRights_FLAGCHANNELBANNEDRIGHTS = me.BannedRights.ToChannelBannedRights()
			}
		} else {
			if !m.Channel.IsBroadcast() {
				channel.BannedRights_FLAGCHATBANNEDRIGHTS = me.BannedRights.ToChatBannedRights()
				channel.BannedRights_FLAGCHANNELBANNEDRIGHTS = me.BannedRights.ToChannelBannedRights()
			}
		}
	}

	if m.Channel.Megagroup {
		channel.DefaultBannedRights = m.Channel.DefaultBannedRights.ToChatBannedRights()
	}
	return
}

func (m *MutableChannel) ToChannelForbidden() (channel *mtproto.Chat) {
	channel = mtproto.MakeTLChannelForbidden(&mtproto.Chat{
		Broadcast: m.Channel.Broadcast,
		Megagroup: m.Channel.Megagroup,
		Id:        m.Channel.Id,
		Title:     m.Channel.Title,
		UntilDate: nil,
	}).To_Chat()
	return
}

func (m *MutableChannel) MakeMessageService(fromId int32, silent bool, replyToMsgId int32, action *mtproto.MessageAction) *mtproto.Message {
	message := mtproto.MakeTLMessageService(&mtproto.Message{
		Out:             true,
		Mentioned:       false,
		MediaUnread:     false,
		Silent:          silent,
		Post:            false,
		Legacy:          false,
		Id:              0,
		FromId_FLAGPEER: MakePeerUser(fromId),
		ToId:            MakePeerChannel(m.Channel.Id),
		ReplyTo:         nil,
		Date:            int32(time.Now().Unix()),
		Action:          action,
	}).To_Message()
	if m.Channel.IsBroadcast() {
		message.Post = true
	}
	if replyToMsgId != 0 {
		message.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: replyToMsgId,
			}).To_MessageReplyHeader()
	}
	return message
}

func (m *MutableChannel) Walk(visit func(userId int32, participant *ImmutableChannelParticipant) error) {
	if visit == nil {
		return
	}
	for k, v := range m.Participants {
		visit(k, v)
	}
}

func (m *MutableChannel) FetchAndWalk(fetch func() []*ImmutableChannelParticipant, visit func(participant *ImmutableChannelParticipant)) {
	if fetch == nil {
		return
	}

	for _, p := range fetch() {
		visit(p)
	}
}

type MutableChannels []*MutableChannel

func (m MutableChannels) ToChats(id int32) []*mtproto.Chat {
	channels := make([]*mtproto.Chat, 0, len(m))
	for _, c := range m {
		channels = append(channels, c.ToUnsafeChat(id))
	}
	return channels
}
