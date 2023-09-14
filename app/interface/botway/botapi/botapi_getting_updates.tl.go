package botapi

import (
	"encoding/json"
)

type Update struct {
	UpdateId           int32               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
}

func (m *Update) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Update) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type WebhookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int32    `json:"pending_update_count"`
	LastErrorDate        int32    `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       int32    `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

func (m *WebhookInfo) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *WebhookInfo) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
