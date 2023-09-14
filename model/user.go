package model

import (
	"time"

	"open.chat/app/pkg/env2"
	"open.chat/pkg/phonenumber"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
)

type UserData struct {
	Id                int32                     `json:"id,omitempty"`
	UserType          UserType                  `json:"user_type,omitempty"`
	AccessHash        int64                     `json:"access_hash,omitempty"`
	SecretKeyId       int64                     `json:"secret_key_id,omitempty"`
	FirstName         string                    `json:"first_name,omitempty"`
	LastName          string                    `json:"last_name,omitempty"`
	Username          string                    `json:"username,omitempty"`
	Phone             string                    `json:"phone,omitempty"`
	CountryCode       string                    `json:"country_code,omitempty"`
	Country           string                    `json:"country"`
	Province          string                    `json:"province"`
	City              string                    `json:"city"`
	CityCode          string                    `json:"cityCode"`
	Gender            int32                     `json:"gender"`
	Birth             string                    `json:"birth"`
	Verified          bool                      `json:"verified,omitempty"`
	About             string                    `json:"about,omitempty"`
	State             int32                     `json:"state,omitempty"`
	IsBot             bool                      `json:"is_bot,omitempty"`
	IsVirtualUser     bool                      `json:"is_virtual_user,omitempty"`
	IsInternal        bool                      `json:"is_internal,omitempty"`
	AccountDaysTtl    int32                     `json:"account_days_ttl,omitempty"`
	Photo             *mtproto.Photo            `json:"photo,omitempty"`
	ProfilePhoto      *mtproto.UserProfilePhoto `json:"profile_photo,omitempty"`
	Min               bool                      `json:"min,omitempty"`
	Restricted        bool                      `json:"restricted,omitempty"`
	RestrictionReason string                    `json:"restriction_reason,omitempty"`
	Deleted           bool                      `json:"deleted,omitempty"`
	DeleteReason      string                    `json:"delete_reason,omitempty"`
	LastSeenAt        int32                     `json:"last_seen_at,omitempty"`
}

type BotData struct {
	Id                   int32  `json:"id,omitempty"`
	BotId                int32  `json:"bot_id,omitempty"`
	BotType              int    `json:"bot_type,omitempty"`
	CreatorUserId        int32  `json:"creator_user_id,omitempty"`
	Token                string `json:"token,omitempty"`
	Description          string `json:"description,omitempty"`
	BotChatHistory       bool   `json:"bot_chat_history,omitempty"`
	BotNochats           bool   `json:"bot_nochats,omitempty"`
	Verified             bool   `json:"verified,omitempty"`
	BotInlineGeo         bool   `json:"bot_inline_geo,omitempty"`
	BotInfoVersion       int32  `json:"bot_info_version,omitempty"`
	BotInlinePlaceholder string `json:"bot_inline_placeholder,omitempty"`
}

type ImmutableUser struct {
	User         *UserData                      `json:"user,omitempty"`
	Bot          *BotData                       `json:"bot,omitempty"`
	Contacts     map[int32]*Contact             `json:"contacts,omitempty"`
	PrivacyRules map[int][]*mtproto.PrivacyRule `json:"privacy_rules,omitempty"`
}

func MakeImmutableUser(user *UserData, bot *BotData, contacts map[int32]*Contact, rules map[int][]*mtproto.PrivacyRule) *ImmutableUser {
	return &ImmutableUser{
		User:         user,
		Bot:          bot,
		Contacts:     contacts,
		PrivacyRules: rules,
	}
}
func (m *ImmutableUser) ResetSetContacts(contacts map[int32]*Contact) {
	m.Contacts = contacts
}

func (m *ImmutableUser) ResetSetPrivacyRules(rules map[int][]*mtproto.PrivacyRule) {
	m.PrivacyRules = rules
}

func (m *ImmutableUser) ID() int32 {
	return m.User.Id
}

func (m *ImmutableUser) AccessHash() int64 {
	return m.User.AccessHash
}

func (m *ImmutableUser) Username() string {
	return m.User.Username
}

func (m *ImmutableUser) Deleted() bool {
	return m.User.Deleted
}

func (m *ImmutableUser) DebugString() string {
	return "{}"
}

func (m *ImmutableUser) Restricted() bool {
	return m.User.Restricted
}

func (m *ImmutableUser) LastSeenAt() int32 {
	return m.User.LastSeenAt
}

func (m *ImmutableUser) ProfilePhoto() *mtproto.UserProfilePhoto {
	return m.User.ProfilePhoto
}

func (m *ImmutableUser) Photo() *mtproto.Photo {
	return m.User.Photo
}

func (m *ImmutableUser) About() string {
	return m.User.About
}

func (m *ImmutableUser) IsBot() bool {
	return m.User.IsBot
}

func (m *ImmutableUser) CheckContact(cId int32) bool {
	_, ok := m.Contacts[cId]
	return ok
}

func (m *ImmutableUser) ToImmutableUser(selfUser *ImmutableUser) *mtproto.User {
	user := mtproto.MakeTLUser(&mtproto.User{
		Self:                         false,
		Contact:                      false,
		MutualContact:                false,
		Deleted:                      m.User.Deleted,
		Bot:                          m.User.IsBot,
		BotChatHistory:               false,
		BotNochats:                   false,
		Verified:                     m.User.Verified,
		Restricted:                   false,
		Min:                          false,
		BotInlineGeo:                 false,
		Support:                      IsSupportId(m.User.Id),
		Scam:                         false,
		Id:                           m.User.Id,
		AccessHash:                   &types.Int64Value{Value: m.User.AccessHash},
		FirstName:                    nil,
		LastName:                     nil,
		Username:                     nil,
		Phone:                        nil,
		Photo:                        nil,
		Status:                       nil,
		BotInfoVersion:               nil,
		RestrictionReason_FLAGSTRING: nil,
		BotInlinePlaceholder:         nil,
		LangCode:                     nil,
		RestrictionReason_FLAGVECTORRESTRICTIONREASON: nil,
	}).To_User()

	if m.User.Deleted {
		return user
	}

	if m.User.Username != "" {
		user.Username = &types.StringValue{Value: m.User.Username}
	}

	isSelf := m.User.Id == selfUser.User.Id
	if isSelf {
		user.Self = true
		user.Contact = true
		user.MutualContact = true
		if m.User.FirstName != "" {
			user.FirstName = &types.StringValue{Value: m.User.FirstName}
		}
		if m.User.LastName != "" {
			user.LastName = &types.StringValue{Value: m.User.LastName}
		}
		if !phonenumber.IsVirtualUser(m.User.Phone) && !phonenumber.IsNotPhoneUser(m.User.Phone) {
			user.Phone = &types.StringValue{Value: m.User.Phone}
		}
		user.Photo = m.User.ProfilePhoto
		user.Status = mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
			WasOnline: int32(time.Now().Unix()),
		}).To_UserStatus()
	} else {
		// contact
		if c, ok := selfUser.Contacts[m.User.Id]; ok {
			user.Contact = true
			user.MutualContact = c.MutualContact
			if c.FirstName != "" {
				user.FirstName = &types.StringValue{Value: c.FirstName}
			}
			if c.LastName != "" {
				user.LastName = &types.StringValue{Value: c.LastName}
			}
		} else {
			if m.User.FirstName != "" {
				user.FirstName = &types.StringValue{Value: m.User.FirstName}
			}
			if m.User.LastName != "" {
				user.LastName = &types.StringValue{Value: m.User.LastName}
			}
		}

		// phone
		if m.checkPrivacy(PHONE_NUMBER, selfUser.User.Id) {
			if !phonenumber.IsVirtualUser(m.User.Phone) && !phonenumber.IsNotPhoneUser(m.User.Phone) {
				user.Phone = &types.StringValue{Value: m.User.Phone}
			}
		}

		// photo
		if m.checkPrivacy(PROFILE_PHOTO, selfUser.User.Id) {
			user.Photo = m.User.ProfilePhoto
		} else {
			user.Photo = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
		}

		allowTimestamp := false
		if env2.PredefinedUser {
			if selfUser.User.Verified {
				allowTimestamp = true
			}
			user.Phone = nil
		} else {
			// status
			allowTimestamp = m.checkPrivacy(STATUS_TIMESTAMP, selfUser.User.Id)
		}
		user.Status = MakeUserStatus(m.User.LastSeenAt, allowTimestamp)
	}

	if m.IsBot() {
		user.Phone = nil
		user.BotChatHistory = m.Bot.BotChatHistory
		user.BotNochats = m.Bot.BotNochats
		user.BotInlineGeo = m.Bot.BotInlineGeo
		user.BotInfoVersion = &types.Int32Value{Value: m.Bot.BotInfoVersion}
		if m.Bot.BotInlinePlaceholder != "" {
			user.BotInlinePlaceholder = &types.StringValue{Value: m.Bot.BotInlinePlaceholder}
		}
	}

	return user
}

func (m *ImmutableUser) checkPrivacy(keyType int, id int32) bool {
	_, isContact := m.Contacts[id]
	if p, ok := m.PrivacyRules[keyType]; ok {
		return privacyIsAllow(p, id, isContact)
	}
	return false
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////
type MutableUsers map[int32]*ImmutableUser

func NewMutableUsers() MutableUsers {
	return make(map[int32]*ImmutableUser)
}

func (m MutableUsers) GetImmutableUser(id int32) (u *ImmutableUser, ok bool) {
	u, ok = m[id]
	return
}

func (m MutableUsers) CheckExistUser(id int32) bool {
	_, ok := m[id]
	return ok
}

func (m MutableUsers) ToUnsafeUserSelf(userId int32) (user *mtproto.User, err error) {
	if me, ok := m[userId]; ok {
		user = me.ToImmutableUser(me)
	} else {
		err = mtproto.ErrUserIdInvalid
	}
	return
}

func (m MutableUsers) ToUnsafeUser(selfId, userId int32) (user *mtproto.User, err error) {
	var (
		selfUser, me *ImmutableUser
		ok           bool
	)

	if selfUser, ok = m[selfId]; !ok {
		err = mtproto.ErrInternelServerError
		return
	}

	if me, ok = m[userId]; !ok {
		err = mtproto.ErrUserIdInvalid
		return
	}

	user = me.ToImmutableUser(selfUser)
	return
}

func (m MutableUsers) GetUsersByIdList(selfId int32, idList []int32) (users []*mtproto.User) {
	var (
		selfUser, user *ImmutableUser
		ok             bool
	)

	users = make([]*mtproto.User, 0, len(idList))
	if selfUser, ok = m[selfId]; !ok {
		return
	}

	for _, id := range idList {
		if user, ok = m[id]; !ok {
			continue
		}
		users = append(users, user.ToImmutableUser(selfUser))
	}

	return
}
