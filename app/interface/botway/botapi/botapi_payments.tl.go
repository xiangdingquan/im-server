package botapi

import (
	"encoding/json"
)

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int32  `json:"amount"`
}

func (m *LabeledPrice) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *LabeledPrice) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int32  `json:"total_amount"`
}

func (m *Invoice) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Invoice) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

func (m *ShippingAddress) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ShippingAddress) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

func (m *OrderInfo) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *OrderInfo) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ShippingOption struct {
	Id     string          `json:"id"`
	Title  string          `json:"title"`
	Prices []*LabeledPrice `json:"prices"`
}

func (m *ShippingOption) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ShippingOption) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int32      `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionId        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
}

func (m *SuccessfulPayment) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *SuccessfulPayment) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ShippingQuery struct {
	Id              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

func (m *ShippingQuery) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ShippingQuery) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int32      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

func (m *PreCheckoutQuery) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PreCheckoutQuery) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
