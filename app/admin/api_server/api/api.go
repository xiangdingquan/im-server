package api

import (
	"encoding/json"

	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

type JSONResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Ttl     int             `json:"ttl"`
	Data    json.RawMessage `json:"data"`
}

func (m *JSONResponse) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

var GVoidRsp = new(VoidRsp)

type VoidRsp struct {
}

func (m *VoidRsp) DebugString() string {
	return "{}"
}

type PushServiceNotificationRequest struct {
	PushIdList []int32                  `json:"push_id_list" form:"push_id_list"`
	Text       string                   `json:"text" form:"text"`
	Entities   []*mtproto.MessageEntity `json:"entities,omitempty" form:"entities"`
}

func (m *PushServiceNotificationRequest) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type PushServiceNotificationResponse struct {
	PushedIdList []int32 `json:"pushed_id_list"`
}

func (m *PushServiceNotificationResponse) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type CreatePredefinedUser struct {
	Phone     string `json:"phone" form:"phone"`
	FirstName string `json:"first_name,omitempty" form:"first_name"`
	LastName  string `json:"last_name,omitempty" form:"last_name"`
	Username  string `json:"username,omitempty" form:"username"`
	Code      string `json:"code" form:"code"`
	Verified  bool   `json:"verified,omitempty" form:"verified"`
}

func (m *CreatePredefinedUser) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type UpdatePredefinedUsername struct {
	Phone    string `json:"phone" form:"phone"`
	Username string `json:"username,omitempty" form:"username"`
}

func (m *UpdatePredefinedUsername) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type UpdatePredefinedProfile struct {
	Phone     string `json:"phone" form:"phone"`
	FirstName string `json:"first_name,omitempty" form:"first_name"`
	LastName  string `json:"last_name,omitempty" form:"last_name"`
	About     string `json:"about,omitempty" form:"about"`
}

func (m *UpdatePredefinedProfile) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type UpdatePredefinedVerified struct {
	Phone    string `json:"phone" form:"phone"`
	Verified bool   `json:"verified,omitempty" form:"verified"`
}

func (m *UpdatePredefinedVerified) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type UpdatePredefinedProfilePhoto struct {
	Phone string `json:"phone" form:"phone"`
	Photo string `json:"photo,omitempty" form:"photo"`
}

func (m *UpdatePredefinedProfilePhoto) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type UpdatePredefinedCode struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

func (m *UpdatePredefinedCode) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type GetPredefinedUser struct {
	Phone string `json:"phone" form:"phone"`
}

func (m *GetPredefinedUser) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type GetPredefinedUsers struct {
}

func (m *GetPredefinedUsers) DebugString() string {
	return "{}"
}

type ToggleBan struct {
	Phone      string `json:"phone" form:"phone"`
	Predefined bool   `json:"predefined,omitempty" form:"predefined"`
	Expires    int32  `json:"expires,omitempty" form:"expires"`
	Reason     string `json:"reason,omitempty" form:"reason"`
}

func (m *ToggleBan) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type CodeResult struct {
	Code string `json:"code"`
}

func (m *CodeResult) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type SendVerifyCode struct {
	Phone    string `json:"phone" form:"phone"`
	Code     string `json:"code" form:"code"`
	CodeHash string `json:"code_hash" form:"code_hash"`
}

func (m *SendVerifyCode) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type VerifyCode struct {
	Code      string `json:"code" form:"code"`
	CodeHash  string `json:"code_hash" form:"code_hash"`
	ExtraData string `json:"extra_data" form:"extra_data"`
}

func (m *VerifyCode) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type AddContact struct {
	Id                       string `json:"id" form:"id"`
	AddPhonePrivacyException bool   `json:"add_phone_privacy_exception" form:"add_phone_privacy_exception"`
	ContactId                string `json:"contact_id" form:"contact_id"`
	FirstName                string `json:"first_name" form:"first_name"`
	LastName                 string `json:"last_name" form:"last_name"`
	Phone                    string `json:"phone" form:"phone"`
}

func (m *AddContact) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type GetAllChats struct {
	Id int32 `json:"id" form:"id"`
}

func (m *GetAllChats) DebugString() string {
	b, _ := json.Marshal(m)
	return hack.String(b)
}

type PhotoProfileResponse struct {
	PushedIdList []int32 `json:"pushed_id_list"`
}
