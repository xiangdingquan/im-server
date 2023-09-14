package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type SendInvoiceReq struct {
	ChatId                    string `json:"chat_id" form:"chat_id"`
	Title                     string `json:"title" form:"title"`
	Description               string `json:"description" form:"description"`
	Payload                   string `json:"payload" form:"payload"`
	ProviderToken             string `json:"provider_token" form:"provider_token"`
	StartParameter            string `json:"start_parameter" form:"start_parameter"`
	Currency                  string `json:"currency" form:"currency"`
	Prices                    string `json:"prices" form:"prices"`
	ProviderData              string `json:"provider_data,omitempty" form:"provider_data"`
	PhotoUrl                  string `json:"photo_url,omitempty" form:"photo_url"`
	PhotoSize                 string `json:"photo_size,omitempty" form:"photo_size"`
	PhotoWidth                string `json:"photo_width,omitempty" form:"photo_width"`
	PhotoHeight               string `json:"photo_height,omitempty" form:"photo_height"`
	NeedName                  string `json:"need_name,omitempty" form:"need_name"`
	NeedPhoneNumber           string `json:"need_phone_number,omitempty" form:"need_phone_number"`
	NeedEmail                 string `json:"need_email,omitempty" form:"need_email"`
	NeedShippingAddress       string `json:"need_shipping_address,omitempty" form:"need_shipping_address"`
	SendPhoneNumberToProvider string `json:"send_phone_number_to_provider,omitempty" form:"send_phone_number_to_provider"`
	SendEmailToProvider       string `json:"send_email_to_provider,omitempty" form:"send_email_to_provider"`
	IsFlexible                string `json:"is_flexible,omitempty" form:"is_flexible"`
	DisableNotification       string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId          string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup               string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendInvoiceReq) Method() string {
	return "sendInvoice"
}

type SendInvoice2 struct {
	ChatId                    int64                 `json:"chat_id,omitempty"`
	Title                     string                `json:"title,omitempty"`
	Description               string                `json:"description,omitempty"`
	Payload                   string                `json:"payload,omitempty"`
	ProviderToken             string                `json:"provider_token,omitempty"`
	StartParameter            string                `json:"start_parameter,omitempty"`
	Currency                  string                `json:"currency,omitempty"`
	Prices                    []*LabeledPrice       `json:"prices,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoUrl                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int32                 `json:"photo_size,omitempty"`
	PhotoWidth                int32                 `json:"photo_width,omitempty"`
	PhotoHeight               int32                 `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId          int32                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *SendInvoice2) NewRequest() BotApiRequest {
	return new(SendInvoiceReq)
}

func (m *SendInvoice2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendInvoiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
			return
		}
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.Description) > 0 {
		m.Description = req.Description
	}
	if len(req.Payload) > 0 {
		m.Payload = req.Payload
	}
	if len(req.ProviderToken) > 0 {
		m.ProviderToken = req.ProviderToken
	}
	if len(req.StartParameter) > 0 {
		m.StartParameter = req.StartParameter
	}
	if len(req.Currency) > 0 {
		m.Currency = req.Currency
	}
	if len(req.Prices) > 0 {
	}
	if len(req.ProviderData) > 0 {
		m.ProviderData = req.ProviderData
	}
	if len(req.PhotoUrl) > 0 {
		m.PhotoUrl = req.PhotoUrl
	}
	if len(req.PhotoSize) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoSize, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoSize = int32(i)
		}
	}
	if len(req.PhotoWidth) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoWidth, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoWidth = int32(i)
		}
	}
	if len(req.PhotoHeight) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoHeight, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoHeight = int32(i)
		}
	}
	if len(req.NeedName) > 0 {
		if m.NeedName, err = strconv.ParseBool(req.NeedName); err != nil {
			return
		}
	}
	if len(req.NeedPhoneNumber) > 0 {
		if m.NeedPhoneNumber, err = strconv.ParseBool(req.NeedPhoneNumber); err != nil {
			return
		}
	}
	if len(req.NeedEmail) > 0 {
		if m.NeedEmail, err = strconv.ParseBool(req.NeedEmail); err != nil {
			return
		}
	}
	if len(req.NeedShippingAddress) > 0 {
		if m.NeedShippingAddress, err = strconv.ParseBool(req.NeedShippingAddress); err != nil {
			return
		}
	}
	if len(req.SendPhoneNumberToProvider) > 0 {
		if m.SendPhoneNumberToProvider, err = strconv.ParseBool(req.SendPhoneNumberToProvider); err != nil {
			return
		}
	}
	if len(req.SendEmailToProvider) > 0 {
		if m.SendEmailToProvider, err = strconv.ParseBool(req.SendEmailToProvider); err != nil {
			return
		}
	}
	if len(req.IsFlexible) > 0 {
		if m.IsFlexible, err = strconv.ParseBool(req.IsFlexible); err != nil {
			return
		}
	}
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.ReplyToMessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.ReplyToMessageId = int32(i)
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendInvoice2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendInvoiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
			return
		}
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.Description) > 0 {
		m.Description = req.Description
	}
	if len(req.Payload) > 0 {
		m.Payload = req.Payload
	}
	if len(req.ProviderToken) > 0 {
		m.ProviderToken = req.ProviderToken
	}
	if len(req.StartParameter) > 0 {
		m.StartParameter = req.StartParameter
	}
	if len(req.Currency) > 0 {
		m.Currency = req.Currency
	}
	if len(req.Prices) > 0 {
	}
	if len(req.ProviderData) > 0 {
		m.ProviderData = req.ProviderData
	}
	if len(req.PhotoUrl) > 0 {
		m.PhotoUrl = req.PhotoUrl
	}
	if len(req.PhotoSize) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoSize, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoSize = int32(i)
		}
	}
	if len(req.PhotoWidth) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoWidth, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoWidth = int32(i)
		}
	}
	if len(req.PhotoHeight) > 0 {
		if i, err2 := strconv.ParseInt(req.PhotoHeight, 10, 32); err2 != nil {
			return err2
		} else {
			m.PhotoHeight = int32(i)
		}
	}
	if len(req.NeedName) > 0 {
		if m.NeedName, err = strconv.ParseBool(req.NeedName); err != nil {
			return
		}
	}
	if len(req.NeedPhoneNumber) > 0 {
		if m.NeedPhoneNumber, err = strconv.ParseBool(req.NeedPhoneNumber); err != nil {
			return
		}
	}
	if len(req.NeedEmail) > 0 {
		if m.NeedEmail, err = strconv.ParseBool(req.NeedEmail); err != nil {
			return
		}
	}
	if len(req.NeedShippingAddress) > 0 {
		if m.NeedShippingAddress, err = strconv.ParseBool(req.NeedShippingAddress); err != nil {
			return
		}
	}
	if len(req.SendPhoneNumberToProvider) > 0 {
		if m.SendPhoneNumberToProvider, err = strconv.ParseBool(req.SendPhoneNumberToProvider); err != nil {
			return
		}
	}
	if len(req.SendEmailToProvider) > 0 {
		if m.SendEmailToProvider, err = strconv.ParseBool(req.SendEmailToProvider); err != nil {
			return
		}
	}
	if len(req.IsFlexible) > 0 {
		if m.IsFlexible, err = strconv.ParseBool(req.IsFlexible); err != nil {
			return
		}
	}
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.ReplyToMessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.ReplyToMessageId = int32(i)
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type AnswerShippingQueryReq struct {
	ShippingQueryId string `json:"shipping_query_id" form:"shipping_query_id"`
	Ok              string `json:"ok" form:"ok"`
	ShippingOptions string `json:"shipping_options,omitempty" form:"shipping_options"`
	ErrorMessage    string `json:"error_message,omitempty" form:"error_message"`
}

func (m *AnswerShippingQueryReq) Method() string {
	return "answerShippingQuery"
}

type AnswerShippingQuery2 struct {
	ShippingQueryId string            `json:"shipping_query_id,omitempty"`
	Ok              bool              `json:"ok,omitempty"`
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string            `json:"error_message,omitempty"`
}

func (m *AnswerShippingQuery2) NewRequest() BotApiRequest {
	return new(AnswerShippingQueryReq)
}

func (m *AnswerShippingQuery2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*AnswerShippingQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ShippingQueryId) > 0 {
		m.ShippingQueryId = req.ShippingQueryId
	}
	if len(req.Ok) > 0 {
		if m.Ok, err = strconv.ParseBool(req.Ok); err != nil {
			return
		}
	}
	if len(req.ShippingOptions) > 0 {
	}
	if len(req.ErrorMessage) > 0 {
		m.ErrorMessage = req.ErrorMessage
	}

	return
}

func (m *AnswerShippingQuery2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*AnswerShippingQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ShippingQueryId) > 0 {
		m.ShippingQueryId = req.ShippingQueryId
	}
	if len(req.Ok) > 0 {
		if m.Ok, err = strconv.ParseBool(req.Ok); err != nil {
			return
		}
	}
	if len(req.ShippingOptions) > 0 {
	}
	if len(req.ErrorMessage) > 0 {
		m.ErrorMessage = req.ErrorMessage
	}

	return
}

type AnswerPreCheckoutQueryReq struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id" form:"pre_checkout_query_id"`
	Ok                 string `json:"ok" form:"ok"`
	ErrorMessage       string `json:"error_message,omitempty" form:"error_message"`
}

func (m *AnswerPreCheckoutQueryReq) Method() string {
	return "answerPreCheckoutQuery"
}

type AnswerPreCheckoutQuery2 struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id,omitempty"`
	Ok                 bool   `json:"ok,omitempty"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

func (m *AnswerPreCheckoutQuery2) NewRequest() BotApiRequest {
	return new(AnswerPreCheckoutQueryReq)
}

func (m *AnswerPreCheckoutQuery2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*AnswerPreCheckoutQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.PreCheckoutQueryId) > 0 {
		m.PreCheckoutQueryId = req.PreCheckoutQueryId
	}
	if len(req.Ok) > 0 {
		if m.Ok, err = strconv.ParseBool(req.Ok); err != nil {
			return
		}
	}
	if len(req.ErrorMessage) > 0 {
		m.ErrorMessage = req.ErrorMessage
	}

	return
}

func (m *AnswerPreCheckoutQuery2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*AnswerPreCheckoutQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.PreCheckoutQueryId) > 0 {
		m.PreCheckoutQueryId = req.PreCheckoutQueryId
	}
	if len(req.Ok) > 0 {
		if m.Ok, err = strconv.ParseBool(req.Ok); err != nil {
			return
		}
	}
	if len(req.ErrorMessage) > 0 {
		m.ErrorMessage = req.ErrorMessage
	}

	return
}
