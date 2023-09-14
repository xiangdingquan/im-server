package model

import (
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

type ImmutableChannelParticipant struct {
	Channel              *ImmutableChannel
	Creator              bool                        `json:"creator,omitempty"`
	State                int                         `json:"state,omitempty"`
	ChannelId            int32                       `json:"channel_id,omitempty"`
	Id                   int64                       `json:"id,omitempty"`
	UserId               int32                       `json:"user_id,omitempty"`
	Date                 int32                       `json:"date,omitempty"`
	InviterId            int32                       `json:"inviter_id_INT32,omitempty"`
	CanEdit              bool                        `json:"can_edit,omitempty"`
	PromotedBy           int32                       `json:"promoted_by,omitempty"`
	AdminRights          ChatAdminRights             `json:"admin_rights,omitempty"`
	KickedBy             int32                       `json:"kicked_by,omitempty"`
	BannedRights         ChatBannedRights            `json:"banned_rights,omitempty"`
	Rank                 string                      `json:"rank,omitempty"`
	Pinned               bool                        `json:"pinned,omitempty"`
	UnreadMark           bool                        `json:"unread_mark,omitempty"`
	TopMessage           int32                       `json:"top_message,omitempty"`
	ReadInboxMaxId       int32                       `json:"read_inbox_max_id,omitempty"`
	UnreadCount          int32                       `json:"unread_count,omitempty"`
	UnreadMentionsCount  int32                       `json:"unread_mentions_count,omitempty"`
	NotifySettings       *mtproto.PeerNotifySettings `json:"notify_settings,omitempty"`
	Draft                *mtproto.DraftMessage       `json:"draft,omitempty"`
	FolderId             int32                       `json:"folder_id,omitempty"`
	AvailableMinId       int32                       `json:"available_min_id,omitempty"`
	AvailableUpdatedDate int32                       `json:"available_updated_date,omitempty"`
	AvailableMinPts      int32                       `json:"available_min_pts,omitempty"`
	MigratedFromMaxId    int32                       `json:"migrated_from_max_id,omitempty"`
	Nickname             string                      `json:"nickname,omitempty"`
}

func (m *ImmutableChannelParticipant) TryGetChannelParticipantSelf(selfId int32) (*mtproto.ChannelParticipant, error) {
	if m.IsKicked() {
		return nil, mtproto.ErrChannelPrivate
	}

	if m.IsChatMemberBanned() {
		return mtproto.MakeTLChannelParticipantBanned(&mtproto.ChannelParticipant{
			Left:                             m.IsKicked() || m.IsLeft(),
			UserId:                           m.UserId,
			KickedBy:                         m.KickedBy,
			Date:                             m.Date,
			BannedRights_CHANNELBANNEDRIGHTS: m.BannedRights.ToChannelBannedRights(),
			BannedRights_CHATBANNEDRIGHTS:    m.BannedRights.ToChatBannedRights(),
			Nickname:                         m.Nickname,
		}).To_ChannelParticipant(), nil
	}

	if selfId == m.UserId {
		if m.IsLeft() {
			return nil, mtproto.ErrChannelPrivate
		}
		return mtproto.MakeTLChannelParticipantSelf(&mtproto.ChannelParticipant{
			UserId:          m.UserId,
			InviterId_INT32: m.InviterId,
			Date:            m.Date,
			Nickname:        m.Nickname,
		}).To_ChannelParticipant(), nil
	} else {
		if m.IsLeft() {
			return nil, mtproto.ErrUserNotParticipant
		}
		return mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
			UserId:          m.UserId,
			InviterId_INT32: m.InviterId,
			Date:            m.Date,
			Nickname:        m.Nickname,
		}).To_ChannelParticipant(), nil
	}
}

func (m *ImmutableChannelParticipant) ToUnsafeChannelParticipant(selfId int32) (p *mtproto.ChannelParticipant) {
	if m.IsChatMemberBanned() {
		p = mtproto.MakeTLChannelParticipantBanned(&mtproto.ChannelParticipant{
			Left:                             m.IsKicked() || m.IsLeft(),
			UserId:                           m.UserId,
			KickedBy:                         m.KickedBy,
			Date:                             m.Date,
			BannedRights_CHANNELBANNEDRIGHTS: m.BannedRights.ToChannelBannedRights(),
			BannedRights_CHATBANNEDRIGHTS:    m.BannedRights.ToChatBannedRights(),
			Nickname:                         m.Nickname,
		}).To_ChannelParticipant()
		return
	}

	if m.IsLeft() {
		return
	}

	if m.IsCreator() {
		p = mtproto.MakeTLChannelParticipantCreator(&mtproto.ChannelParticipant{
			UserId:                         m.UserId,
			AdminRights_CHANNELADMINRIGHTS: m.AdminRights.ToChannelAdminRights(),
			AdminRights_CHATADMINRIGHTS:    m.AdminRights.ToChatAdminRights(),
			Rank:                           nil,
			Nickname:                       m.Nickname,
		}).To_ChannelParticipant()

		if m.Rank != "" {
			p.Rank = &types.StringValue{Value: m.Rank}
		}
		return
	}

	if m.IsAdmin() {
		p = mtproto.MakeTLChannelParticipantAdmin(&mtproto.ChannelParticipant{
			CanEdit:                        false,
			Self:                           false, // selfId == m.UserId,
			UserId:                         m.UserId,
			InviterId_INT32:                m.InviterId,
			InviterId_FLAGINT32:            nil, // if self == true then InviterId_FLAGINT32 = InviterId
			PromotedBy:                     m.PromotedBy,
			Date:                           m.Date,
			AdminRights_CHANNELADMINRIGHTS: m.AdminRights.ToChannelAdminRights(),
			AdminRights_CHATADMINRIGHTS:    m.AdminRights.ToChatAdminRights(),
			Rank:                           nil,
			Nickname:                       m.Nickname,
		}).To_ChannelParticipant()

		if selfId == m.UserId {
			p.Self = true
			p.InviterId_FLAGINT32 = &types.Int32Value{Value: m.InviterId}
		}

		// removed by selfId
		if m.Channel.CreatorId == selfId {
			p.CanEdit = true
		}

		if m.Rank != "" {
			p.Rank = &types.StringValue{Value: m.Rank}
		}
		return
	}

	p = mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
		UserId:   m.UserId,
		Date:     m.Date,
		Nickname: m.Nickname,
	}).To_ChannelParticipant()
	return
}

func (m *ImmutableChannelParticipant) CanViewMessages(date int32) bool {
	return m.IsStateOk() && (m.Creator || m.AdminRights.HasAdminRights() || (m.Channel.canViewMessages() && m.BannedRights.CanViewMessages(date)))
}

func (m *ImmutableChannelParticipant) CanSendMessages(date int32) bool {
	return m.IsStateOk() && (m.Creator || m.AdminRights.HasAdminRights() || (m.Channel.canSendMessages() && m.BannedRights.CanSendMessages(date)))
}

func (m *ImmutableChannelParticipant) CanSendMedia(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendMedia() && m.BannedRights.CanSendMedia(date)))
}

func (m *ImmutableChannelParticipant) CanSendStickers(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendStickers() && m.BannedRights.CanSendStickers(date)))
}

func (m *ImmutableChannelParticipant) CanSendGifs(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendGifs() && m.BannedRights.CanSendGifs(date)))
}

func (m *ImmutableChannelParticipant) CanSendGames(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendGames() && m.BannedRights.CanSendGames(date)))
}

func (m *ImmutableChannelParticipant) CanSendInline(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendInline() && m.BannedRights.CanSendInline(date)))
}

func (m *ImmutableChannelParticipant) CanEmbedLinks(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canEmbedLinks() && m.BannedRights.CanEmbedLinks(date)))
}

func (m *ImmutableChannelParticipant) CanSendPolls(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canSendPolls() && m.BannedRights.CanSendPolls(date)))
}

func (m *ImmutableChannelParticipant) CanChangeInfo(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canChangeInfo() && m.BannedRights.CanChangeInfo(date)))
}

func (m *ImmutableChannelParticipant) CanInviteUsers(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canInviteUsers() && m.BannedRights.CanInviteUsers(date)))
}

func (m *ImmutableChannelParticipant) CanPinMessages(date int32) bool {
	return m.IsStateOk() && (m.AdminRights.HasAdminRights() || (m.Channel.canPinMessages() && m.BannedRights.CanPinMessage(date)))
}

func (m *ImmutableChannelParticipant) CanAdminChangeInfo() bool {
	return m.IsStateOk() && m.AdminRights.CanChangeInfo()
}

func (m *ImmutableChannelParticipant) CanAdminPostMessages() bool {
	return m.IsStateOk() && m.AdminRights.CanPostMessages()
}

func (m *ImmutableChannelParticipant) CanAdminEditMessages() bool {
	return m.IsStateOk() && m.AdminRights.CanEditMessages()
}

func (m *ImmutableChannelParticipant) CanAdminDeleteMessages() bool {
	return m.IsStateOk() && m.AdminRights.CanDeleteMessages()
}

func (m *ImmutableChannelParticipant) CanAdminBanUsers() bool {
	return m.IsStateOk() && m.AdminRights.CanBanUsers()
}

func (m *ImmutableChannelParticipant) CanAdminInviteUsers() bool {
	return m.IsStateOk() && m.AdminRights.CanInviteUsers()
}

func (m *ImmutableChannelParticipant) CanAdminPinMessages() bool {
	return m.IsStateOk() && m.AdminRights.CanPinMessages()
}

func (m *ImmutableChannelParticipant) CanAdminAddAdmins() bool {
	return m.IsStateOk() && m.AdminRights.CanAddAdmins()
}

func (m *ImmutableChannelParticipant) HasAnonymous() bool {
	return m.IsStateOk() && m.AdminRights.HasAnonymous()
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChannelParticipant) IsChatMemberNormal() bool {
	return m.IsStateOk() && !m.IsAdmin() && !m.IsCreator() && !m.BannedRights.IsBan(int32(time.Now().Unix()))
}

func (m *ImmutableChannelParticipant) IsChatMemberAdmin() bool {
	return m.IsStateOk() && m.AdminRights.HasAdminRights()
}

func (m *ImmutableChannelParticipant) IsChatMemberCreator() bool {
	return m.IsStateOk() && m.Creator
}

func (m *ImmutableChannelParticipant) IsChatMemberBanned() bool {
	return m.IsKicked() || (m.IsStateOk() && m.IsBan(int32(time.Now().Unix())))
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ImmutableChannelParticipant) IsCreator() bool {
	return m.Creator
}

func (m *ImmutableChannelParticipant) IsAdmin() bool {
	return m.AdminRights.HasAdminRights()
}

func (m *ImmutableChannelParticipant) IsCreatorOrAdmin() bool {
	return m.Creator || m.AdminRights.HasAdminRights()
}

func (m *ImmutableChannelParticipant) IsBan(date int32) bool {
	return m.BannedRights.IsBan(date)
}

func (m *ImmutableChannelParticipant) IsKicked() bool {
	return m.State == ChatMemberStateKicked
}

func (m *ImmutableChannelParticipant) IsLeft() bool {
	return m.State == ChatMemberStateLeft
}

func (m *ImmutableChannelParticipant) IsStateOk() bool {
	return m.State == ChatMemberStateNormal
}

func (m *ImmutableChannelParticipant) IsRestricted() bool {
	return m.BannedRights.IsRestrict(int32(time.Now().Unix()))
}

// /////////////////////////////////////////////////////////////////////////////////////////////
type ImmutableChannelParticipants []*ImmutableChannelParticipant

func (m ImmutableChannelParticipants) ToChannelParticipants(id int32) []*mtproto.ChannelParticipant {
	participants := make([]*mtproto.ChannelParticipant, 0, len(m))
	for _, p := range m {
		if p != nil {
			participants = append(participants, p.ToUnsafeChannelParticipant(id))
		}
	}
	return participants

}
