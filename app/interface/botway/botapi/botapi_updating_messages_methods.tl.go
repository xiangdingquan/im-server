package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type EditMessageTextReq struct {
	ChatId                string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId             string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId       string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	Text                  string `json:"text" form:"text"`
	ParseMode             string `json:"parse_mode,omitempty" form:"parse_mode"`
	DisableWebPagePreview string `json:"disable_web_page_preview,omitempty" form:"disable_web_page_preview"`
	ReplyMarkup           string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *EditMessageTextReq) Method() string {
	return "editMessageText"
}

type EditMessageText2 struct {
	ChatId                ChatId2               `json:"chat_id,omitempty"`
	MessageId             int32                 `json:"message_id,omitempty"`
	InlineMessageId       string                `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *EditMessageText2) NewRequest() BotApiRequest {
	return new(EditMessageTextReq)
}

func (m *EditMessageText2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageTextReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Text) > 0 {
		m.Text = req.Text
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.DisableWebPagePreview) > 0 {
		if m.DisableWebPagePreview, err = strconv.ParseBool(req.DisableWebPagePreview); err != nil {
			return
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

func (m *EditMessageText2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageTextReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Text) > 0 {
		m.Text = req.Text
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.DisableWebPagePreview) > 0 {
		if m.DisableWebPagePreview, err = strconv.ParseBool(req.DisableWebPagePreview); err != nil {
			return
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

type EditMessageCaptionReq struct {
	ChatId          string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	Caption         string `json:"caption,omitempty" form:"caption"`
	ParseMode       string `json:"parse_mode,omitempty" form:"parse_mode"`
	ReplyMarkup     string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *EditMessageCaptionReq) Method() string {
	return "editMessageCaption"
}

type EditMessageCaption2 struct {
	ChatId          ChatId2               `json:"chat_id,omitempty"`
	MessageId       int32                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Caption         string                `json:"caption,omitempty"`
	ParseMode       string                `json:"parse_mode,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *EditMessageCaption2) NewRequest() BotApiRequest {
	return new(EditMessageCaptionReq)
}

func (m *EditMessageCaption2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageCaptionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *EditMessageCaption2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageCaptionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type EditMessageMediaReq struct {
	ChatId          string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	Media           string `json:"media" form:"media"`
	ReplyMarkup     string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *EditMessageMediaReq) Method() string {
	return "editMessageMedia"
}

type EditMessageMedia2 struct {
	ChatId          ChatId2               `json:"chat_id,omitempty"`
	MessageId       int32                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Media           *InputMedia           `json:"media,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *EditMessageMedia2) NewRequest() BotApiRequest {
	return new(EditMessageMediaReq)
}

func (m *EditMessageMedia2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageMediaReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Media) > 0 {
		m.Media = new(InputMedia)
		if err = m.Media.Decode([]byte(req.Media)); err != nil {
			return
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

func (m *EditMessageMedia2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageMediaReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.Media) > 0 {
		m.Media = new(InputMedia)
		if err = m.Media.Decode([]byte(req.Media)); err != nil {
			return
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

type EditMessageReplyMarkupReq struct {
	ChatId          string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	ReplyMarkup     string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *EditMessageReplyMarkupReq) Method() string {
	return "editMessageReplyMarkup"
}

type EditMessageReplyMarkup2 struct {
	ChatId          ChatId2               `json:"chat_id,omitempty"`
	MessageId       int32                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *EditMessageReplyMarkup2) NewRequest() BotApiRequest {
	return new(EditMessageReplyMarkupReq)
}

func (m *EditMessageReplyMarkup2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageReplyMarkupReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *EditMessageReplyMarkup2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageReplyMarkupReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type StopPollReq struct {
	ChatId      string `json:"chat_id" form:"chat_id"`
	MessageId   string `json:"message_id" form:"message_id"`
	ReplyMarkup string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *StopPollReq) Method() string {
	return "stopPoll"
}

type StopPoll2 struct {
	ChatId      ChatId2               `json:"chat_id,omitempty"`
	MessageId   int32                 `json:"message_id,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *StopPoll2) NewRequest() BotApiRequest {
	return new(StopPollReq)
}

func (m *StopPoll2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*StopPollReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
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

func (m *StopPoll2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*StopPollReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
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

type DeleteMessageReq struct {
	ChatId    string `json:"chat_id" form:"chat_id"`
	MessageId string `json:"message_id" form:"message_id"`
}

func (m *DeleteMessageReq) Method() string {
	return "deleteMessage"
}

type DeleteMessage2 struct {
	ChatId    ChatId2 `json:"chat_id,omitempty"`
	MessageId int32   `json:"message_id,omitempty"`
}

func (m *DeleteMessage2) NewRequest() BotApiRequest {
	return new(DeleteMessageReq)
}

func (m *DeleteMessage2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*DeleteMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}

	return
}

func (m *DeleteMessage2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*DeleteMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.MessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.MessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.MessageId = int32(i)
		}
	}

	return
}
