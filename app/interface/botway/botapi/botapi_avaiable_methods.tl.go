package botapi

import (
	"encoding/json"
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type GetMeReq struct {
}

func (m *GetMeReq) Method() string {
	return "getMe"
}

type GetMe2 struct {
}

func (m *GetMe2) NewRequest() BotApiRequest {
	return new(GetMeReq)
}

func (m *GetMe2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetMeReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

func (m *GetMe2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetMeReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

type SendMessageReq struct {
	ChatId                string `json:"chat_id" form:"chat_id"`
	Text                  string `json:"text" form:"text"`
	ParseMode             string `json:"parse_mode,omitempty" form:"parse_mode"`
	DisableWebPagePreview string `json:"disable_web_page_preview,omitempty" form:"disable_web_page_preview"`
	DisableNotification   string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId      string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup           string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendMessageReq) Method() string {
	return "sendMessage"
}

type SendMessage2 struct {
	ChatId                ChatId2      `json:"chat_id,omitempty"`
	Text                  string       `json:"text,omitempty"`
	ParseMode             string       `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId      int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendMessage2) NewRequest() BotApiRequest {
	return new(SendMessageReq)
}

func (m *SendMessage2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendMessage2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type ForwardMessageReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	FromChatId          string `json:"from_chat_id" form:"from_chat_id"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	MessageId           string `json:"message_id" form:"message_id"`
}

func (m *ForwardMessageReq) Method() string {
	return "forwardMessage"
}

type ForwardMessage2 struct {
	ChatId              ChatId2 `json:"chat_id,omitempty"`
	FromChatId          ChatId2 `json:"from_chat_id,omitempty"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	MessageId           int32   `json:"message_id,omitempty"`
}

func (m *ForwardMessage2) NewRequest() BotApiRequest {
	return new(ForwardMessageReq)
}

func (m *ForwardMessage2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*ForwardMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.FromChatId) > 0 {
		m.FromChatId = MakeChatId(req.FromChatId)
	}
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
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

func (m *ForwardMessage2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*ForwardMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.FromChatId) > 0 {
		m.FromChatId = MakeChatId(req.FromChatId)
	}

	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
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

type SendPhotoReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Photo               string `json:"photo" form:"photo"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendPhotoReq) Method() string {
	return "sendPhoto"
}

type SendPhoto2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Photo               InputFile    `json:"photo,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendPhoto2) NewRequest() BotApiRequest {
	return new(SendPhotoReq)
}

func (m *SendPhoto2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Photo) > 0 {
		m.Photo = MakeInputFile(req.Photo)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendPhoto2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Photo, err = MakeInputFile2(c, contentType, "photo", req.Photo); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendAudioReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Audio               string `json:"audio" form:"audio"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	Duration            string `json:"duration,omitempty" form:"duration"`
	Performer           string `json:"performer,omitempty" form:"performer"`
	Title               string `json:"title,omitempty" form:"title"`
	Thumb               string `json:"thumb,omitempty" form:"thumb"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendAudioReq) Method() string {
	return "sendAudio"
}

type SendAudio2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Audio               InputFile    `json:"audio,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	Duration            int32        `json:"duration,omitempty"`
	Performer           string       `json:"performer,omitempty"`
	Title               string       `json:"title,omitempty"`
	Thumb               InputFile    `json:"thumb,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendAudio2) NewRequest() BotApiRequest {
	return new(SendAudioReq)
}

func (m *SendAudio2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendAudioReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Audio) > 0 {
		m.Audio = MakeInputFile(req.Audio)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Performer) > 0 {
		m.Performer = req.Performer
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendAudio2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendAudioReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Audio, err = MakeInputFile2(c, contentType, "audio", req.Audio); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Performer) > 0 {
		m.Performer = req.Performer
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendDocumentReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Document            string `json:"document" form:"document"`
	Thumb               string `json:"thumb,omitempty" form:"thumb"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendDocumentReq) Method() string {
	return "sendDocument"
}

type SendDocument2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Document            InputFile    `json:"document,omitempty"`
	Thumb               InputFile    `json:"thumb,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendDocument2) NewRequest() BotApiRequest {
	return new(SendDocumentReq)
}

func (m *SendDocument2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendDocumentReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Document) > 0 {
		m.Document = MakeInputFile(req.Document)
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendDocument2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendDocumentReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Document, err = MakeInputFile2(c, contentType, "document", req.Document); err != nil {
		return
	}

	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendVideoReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Video               string `json:"video" form:"video"`
	Duration            string `json:"duration,omitempty" form:"duration"`
	Width               string `json:"width,omitempty" form:"width"`
	Height              string `json:"height,omitempty" form:"height"`
	Thumb               string `json:"thumb,omitempty" form:"thumb"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	SupportsStreaming   string `json:"supports_streaming,omitempty" form:"supports_streaming"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendVideoReq) Method() string {
	return "sendVideo"
}

type SendVideo2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Video               InputFile    `json:"video,omitempty"`
	Duration            int32        `json:"duration,omitempty"`
	Width               int32        `json:"width,omitempty"`
	Height              int32        `json:"height,omitempty"`
	Thumb               InputFile    `json:"thumb,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	SupportsStreaming   bool         `json:"supports_streaming,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendVideo2) NewRequest() BotApiRequest {
	return new(SendVideoReq)
}

func (m *SendVideo2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendVideoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Video) > 0 {
		m.Video = MakeInputFile(req.Video)
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Width) > 0 {
		if i, err2 := strconv.ParseInt(req.Width, 10, 32); err2 != nil {
			return err2
		} else {
			m.Width = int32(i)
		}
	}
	if len(req.Height) > 0 {
		if i, err2 := strconv.ParseInt(req.Height, 10, 32); err2 != nil {
			return err2
		} else {
			m.Height = int32(i)
		}
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.SupportsStreaming) > 0 {
		if m.SupportsStreaming, err = strconv.ParseBool(req.SupportsStreaming); err != nil {
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendVideo2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendVideoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Video, err = MakeInputFile2(c, contentType, "video", req.Video); err != nil {
		return
	}

	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Width) > 0 {
		if i, err2 := strconv.ParseInt(req.Width, 10, 32); err2 != nil {
			return err2
		} else {
			m.Width = int32(i)
		}
	}
	if len(req.Height) > 0 {
		if i, err2 := strconv.ParseInt(req.Height, 10, 32); err2 != nil {
			return err2
		} else {
			m.Height = int32(i)
		}
	}
	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.SupportsStreaming) > 0 {
		if m.SupportsStreaming, err = strconv.ParseBool(req.SupportsStreaming); err != nil {
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendAnimationReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Animation           string `json:"animation" form:"animation"`
	Duration            string `json:"duration,omitempty" form:"duration"`
	Width               string `json:"width,omitempty" form:"width"`
	Height              string `json:"height,omitempty" form:"height"`
	Thumb               string `json:"thumb,omitempty" form:"thumb"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendAnimationReq) Method() string {
	return "sendAnimation"
}

type SendAnimation2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Animation           InputFile    `json:"animation,omitempty"`
	Duration            int32        `json:"duration,omitempty"`
	Width               int32        `json:"width,omitempty"`
	Height              int32        `json:"height,omitempty"`
	Thumb               InputFile    `json:"thumb,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendAnimation2) NewRequest() BotApiRequest {
	return new(SendAnimationReq)
}

func (m *SendAnimation2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendAnimationReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Animation) > 0 {
		m.Animation = MakeInputFile(req.Animation)
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Width) > 0 {
		if i, err2 := strconv.ParseInt(req.Width, 10, 32); err2 != nil {
			return err2
		} else {
			m.Width = int32(i)
		}
	}
	if len(req.Height) > 0 {
		if i, err2 := strconv.ParseInt(req.Height, 10, 32); err2 != nil {
			return err2
		} else {
			m.Height = int32(i)
		}
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendAnimation2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendAnimationReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Animation, err = MakeInputFile2(c, contentType, "animation", req.Animation); err != nil {
		return
	}

	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Width) > 0 {
		if i, err2 := strconv.ParseInt(req.Width, 10, 32); err2 != nil {
			return err2
		} else {
			m.Width = int32(i)
		}
	}
	if len(req.Height) > 0 {
		if i, err2 := strconv.ParseInt(req.Height, 10, 32); err2 != nil {
			return err2
		} else {
			m.Height = int32(i)
		}
	}
	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendVoiceReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Voice               string `json:"voice" form:"voice"`
	Caption             string `json:"caption,omitempty" form:"caption"`
	ParseMode           string `json:"parse_mode,omitempty" form:"parse_mode"`
	Duration            string `json:"duration,omitempty" form:"duration"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendVoiceReq) Method() string {
	return "sendVoice"
}

type SendVoice2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Voice               InputFile    `json:"voice,omitempty"`
	Caption             string       `json:"caption,omitempty"`
	ParseMode           string       `json:"parse_mode,omitempty"`
	Duration            int32        `json:"duration,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendVoice2) NewRequest() BotApiRequest {
	return new(SendVoiceReq)
}

func (m *SendVoice2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendVoiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Voice) > 0 {
		m.Voice = MakeInputFile(req.Voice)
	}
	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendVoice2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendVoiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Voice, err = MakeInputFile2(c, contentType, "voice", req.Voice); err != nil {
		return
	}

	if len(req.Caption) > 0 {
		m.Caption = req.Caption
	}
	if len(req.ParseMode) > 0 {
		m.ParseMode = req.ParseMode
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendVideoNoteReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	VideoNote           string `json:"video_note" form:"video_note"`
	Duration            string `json:"duration,omitempty" form:"duration"`
	Length              string `json:"length,omitempty" form:"length"`
	Thumb               string `json:"thumb,omitempty" form:"thumb"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendVideoNoteReq) Method() string {
	return "sendVideoNote"
}

type SendVideoNote2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	VideoNote           InputFile    `json:"video_note,omitempty"`
	Duration            int32        `json:"duration,omitempty"`
	Length              int32        `json:"length,omitempty"`
	Thumb               InputFile    `json:"thumb,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendVideoNote2) NewRequest() BotApiRequest {
	return new(SendVideoNoteReq)
}

func (m *SendVideoNote2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendVideoNoteReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.VideoNote) > 0 {
		m.VideoNote = MakeInputFile(req.VideoNote)
	}
	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Length) > 0 {
		if i, err2 := strconv.ParseInt(req.Length, 10, 32); err2 != nil {
			return err2
		} else {
			m.Length = int32(i)
		}
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendVideoNote2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendVideoNoteReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.VideoNote, err = MakeInputFile2(c, contentType, "video_note", req.VideoNote); err != nil {
		return
	}

	if len(req.Duration) > 0 {
		if i, err2 := strconv.ParseInt(req.Duration, 10, 32); err2 != nil {
			return err2
		} else {
			m.Duration = int32(i)
		}
	}
	if len(req.Length) > 0 {
		if i, err2 := strconv.ParseInt(req.Length, 10, 32); err2 != nil {
			return err2
		} else {
			m.Length = int32(i)
		}
	}
	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendMediaGroupReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Media               string `json:"media" form:"media"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
}

func (m *SendMediaGroupReq) Method() string {
	return "sendMediaGroup"
}

type SendMediaGroup2 struct {
	ChatId              ChatId2       `json:"chat_id,omitempty"`
	Media               []*InputMedia `json:"media,omitempty"`
	DisableNotification bool          `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32         `json:"reply_to_message_id,omitempty"`
}

func (m *SendMediaGroup2) NewRequest() BotApiRequest {
	return new(SendMediaGroupReq)
}

func (m *SendMediaGroup2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendMediaGroupReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Media) > 0 {
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

	return
}

func (m *SendMediaGroup2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendMediaGroupReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Media) > 0 {
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

	return
}

type SendLocationReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Latitude            string `json:"latitude" form:"latitude"`
	Longitude           string `json:"longitude" form:"longitude"`
	LivePeriod          string `json:"live_period,omitempty" form:"live_period"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendLocationReq) Method() string {
	return "sendLocation"
}

type SendLocation2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Latitude            float64      `json:"latitude,omitempty"`
	Longitude           float64      `json:"longitude,omitempty"`
	LivePeriod          int32        `json:"live_period,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendLocation2) NewRequest() BotApiRequest {
	return new(SendLocationReq)
}

func (m *SendLocation2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendLocationReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.LivePeriod) > 0 {
		if i, err2 := strconv.ParseInt(req.LivePeriod, 10, 32); err2 != nil {
			return err2
		} else {
			m.LivePeriod = int32(i)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendLocation2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendLocationReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.LivePeriod) > 0 {
		if i, err2 := strconv.ParseInt(req.LivePeriod, 10, 32); err2 != nil {
			return err2
		} else {
			m.LivePeriod = int32(i)
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type EditMessageLiveLocationReq struct {
	ChatId          string `json:"chat_id" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	Latitude        string `json:"latitude" form:"latitude"`
	Longitude       string `json:"longitude" form:"longitude"`
	ReplyMarkup     string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *EditMessageLiveLocationReq) Method() string {
	return "editMessageLiveLocation"
}

type EditMessageLiveLocation2 struct {
	ChatId          ChatId2               `json:"chat_id,omitempty"`
	MessageId       int32                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	Latitude        float64               `json:"latitude,omitempty"`
	Longitude       float64               `json:"longitude,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *EditMessageLiveLocation2) NewRequest() BotApiRequest {
	return new(EditMessageLiveLocationReq)
}

func (m *EditMessageLiveLocation2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageLiveLocationReq)
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

func (m *EditMessageLiveLocation2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*EditMessageLiveLocationReq)
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

type StopMessageLiveLocationReq struct {
	ChatId          string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
	ReplyMarkup     string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *StopMessageLiveLocationReq) Method() string {
	return "stopMessageLiveLocation"
}

type StopMessageLiveLocation2 struct {
	ChatId          ChatId2               `json:"chat_id,omitempty"`
	MessageId       int32                 `json:"message_id,omitempty"`
	InlineMessageId string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *StopMessageLiveLocation2) NewRequest() BotApiRequest {
	return new(StopMessageLiveLocationReq)
}

func (m *StopMessageLiveLocation2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*StopMessageLiveLocationReq)
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

func (m *StopMessageLiveLocation2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*StopMessageLiveLocationReq)
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

type SendVenueReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Latitude            string `json:"latitude" form:"latitude"`
	Longitude           string `json:"longitude" form:"longitude"`
	Title               string `json:"title" form:"title"`
	Address             string `json:"address" form:"address"`
	FoursquareId        string `json:"foursquare_id,omitempty" form:"foursquare_id"`
	FoursquareType      string `json:"foursquare_type,omitempty" form:"foursquare_type"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup" form:"reply_markup"`
}

func (m *SendVenueReq) Method() string {
	return "sendVenue"
}

type SendVenue2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	Latitude            float64      `json:"latitude,omitempty"`
	Longitude           float64      `json:"longitude,omitempty"`
	Title               string       `json:"title,omitempty"`
	Address             string       `json:"address,omitempty"`
	FoursquareId        string       `json:"foursquare_id,omitempty"`
	FoursquareType      string       `json:"foursquare_type,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendVenue2) NewRequest() BotApiRequest {
	return new(SendVenueReq)
}

func (m *SendVenue2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendVenueReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.Address) > 0 {
		m.Address = req.Address
	}
	if len(req.FoursquareId) > 0 {
		m.FoursquareId = req.FoursquareId
	}
	if len(req.FoursquareType) > 0 {
		m.FoursquareType = req.FoursquareType
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendVenue2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendVenueReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.Address) > 0 {
		m.Address = req.Address
	}
	if len(req.FoursquareId) > 0 {
		m.FoursquareId = req.FoursquareId
	}
	if len(req.FoursquareType) > 0 {
		m.FoursquareType = req.FoursquareType
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendContactReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	PhoneNumber         string `json:"phone_number" form:"phone_number"`
	FirstName           string `json:"first_name" form:"first_name"`
	LastName            string `json:"last_name,omitempty" form:"last_name"`
	Vcard               string `json:"vcard,omitempty" form:"vcard"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendContactReq) Method() string {
	return "sendContact"
}

type SendContact2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	PhoneNumber         string       `json:"phone_number,omitempty"`
	FirstName           string       `json:"first_name,omitempty"`
	LastName            string       `json:"last_name,omitempty"`
	Vcard               string       `json:"vcard,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendContact2) NewRequest() BotApiRequest {
	return new(SendContactReq)
}

func (m *SendContact2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendContactReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.PhoneNumber) > 0 {
		m.PhoneNumber = req.PhoneNumber
	}
	if len(req.FirstName) > 0 {
		m.FirstName = req.FirstName
	}
	if len(req.LastName) > 0 {
		m.LastName = req.LastName
	}
	if len(req.Vcard) > 0 {
		m.Vcard = req.Vcard
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendContact2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendContactReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.PhoneNumber) > 0 {
		m.PhoneNumber = req.PhoneNumber
	}
	if len(req.FirstName) > 0 {
		m.FirstName = req.FirstName
	}
	if len(req.LastName) > 0 {
		m.LastName = req.LastName
	}
	if len(req.Vcard) > 0 {
		m.Vcard = req.Vcard
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendPollReq struct {
	ChatId                string `json:"chat_id" form:"chat_id"`
	Question              string `json:"question" form:"question"`
	Options               string `json:"options" form:"options"`
	IsAnonymous           string `json:"is_anonymous,omitempty" form:"is_anonymous"`
	Type                  string `json:"type,omitempty" form:"type"`
	AllowsMultipleAnswers string `json:"allows_multiple_answers,omitempty" form:"allows_multiple_answers"`
	CorrectOptionId       string `json:"correct_option_id,omitempty" form:"correct_option_id"`
	IsClosed              string `json:"is_closed,omitempty" form:"is_closed"`
	DisableNotification   string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId      string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup           string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendPollReq) Method() string {
	return "sendPoll"
}

type SendPoll2 struct {
	ChatId                ChatId2      `json:"chat_id,omitempty"`
	Question              string       `json:"question,omitempty"`
	Options               []string     `json:"options,omitempty"`
	IsAnonymous           bool         `json:"is_anonymous,omitempty"`
	Type                  string       `json:"type,omitempty"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId       int32        `json:"correct_option_id,omitempty"`
	IsClosed              bool         `json:"is_closed,omitempty"`
	DisableNotification   bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId      int32        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendPoll2) NewRequest() BotApiRequest {
	return new(SendPollReq)
}

func (m *SendPoll2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendPollReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Question) > 0 {
		m.Question = req.Question
	}
	if len(req.Options) > 0 {
		err = json.Unmarshal([]byte(req.Options), &m.Options)
	}
	if len(req.IsAnonymous) > 0 {
		if m.IsAnonymous, err = strconv.ParseBool(req.IsAnonymous); err != nil {
			return
		}
	}
	if len(req.Type) > 0 {
		m.Type = req.Type
	}
	if len(req.AllowsMultipleAnswers) > 0 {
		if m.AllowsMultipleAnswers, err = strconv.ParseBool(req.AllowsMultipleAnswers); err != nil {
			return
		}
	}
	if len(req.CorrectOptionId) > 0 {
		if i, err2 := strconv.ParseInt(req.CorrectOptionId, 10, 32); err2 != nil {
			return err2
		} else {
			m.CorrectOptionId = int32(i)
		}
	}
	if len(req.IsClosed) > 0 {
		if m.IsClosed, err = strconv.ParseBool(req.IsClosed); err != nil {
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendPoll2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendPollReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Question) > 0 {
		m.Question = req.Question
	}
	if len(req.Options) > 0 {
		err = json.Unmarshal([]byte(req.Options), &m.Options)
	}
	if len(req.IsAnonymous) > 0 {
		if m.IsAnonymous, err = strconv.ParseBool(req.IsAnonymous); err != nil {
			return
		}
	}
	if len(req.Type) > 0 {
		m.Type = req.Type
	}
	if len(req.AllowsMultipleAnswers) > 0 {
		if m.AllowsMultipleAnswers, err = strconv.ParseBool(req.AllowsMultipleAnswers); err != nil {
			return
		}
	}
	if len(req.CorrectOptionId) > 0 {
		if i, err2 := strconv.ParseInt(req.CorrectOptionId, 10, 32); err2 != nil {
			return err2
		} else {
			m.CorrectOptionId = int32(i)
		}
	}
	if len(req.IsClosed) > 0 {
		if m.IsClosed, err = strconv.ParseBool(req.IsClosed); err != nil {
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
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendDiceReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendDiceReq) Method() string {
	return "sendDice"
}

type SendDice2 struct {
	ChatId              ChatId2      `json:"chat_id,omitempty"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageId    bool         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *SendDice2) NewRequest() BotApiRequest {
	return new(SendDiceReq)
}

func (m *SendDice2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendDiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if m.ReplyToMessageId, err = strconv.ParseBool(req.ReplyToMessageId); err != nil {
			return
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendDice2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendDiceReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if m.ReplyToMessageId, err = strconv.ParseBool(req.ReplyToMessageId); err != nil {
			return
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(ReplyMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type SendChatActionReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
	Action string `json:"action" form:"action"`
}

func (m *SendChatActionReq) Method() string {
	return "sendChatAction"
}

type SendChatAction2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
	Action string  `json:"action,omitempty"`
}

func (m *SendChatAction2) NewRequest() BotApiRequest {
	return new(SendChatActionReq)
}

func (m *SendChatAction2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendChatActionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Action) > 0 {
		m.Action = req.Action
	}

	return
}

func (m *SendChatAction2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendChatActionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Action) > 0 {
		m.Action = req.Action
	}

	return
}

type GetUserProfilePhotosReq struct {
	UserId string `json:"user_id" form:"user_id"`
	Offset string `json:"offset,omitempty" form:"offset"`
	Limit  string `json:"limit,omitempty" form:"limit"`
}

func (m *GetUserProfilePhotosReq) Method() string {
	return "getUserProfilePhotos"
}

type GetUserProfilePhotos2 struct {
	UserId int32 `json:"user_id,omitempty"`
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

func (m *GetUserProfilePhotos2) NewRequest() BotApiRequest {
	return new(GetUserProfilePhotosReq)
}

func (m *GetUserProfilePhotos2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetUserProfilePhotosReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Offset) > 0 {
		if i, err2 := strconv.ParseInt(req.Offset, 10, 32); err2 != nil {
			return err2
		} else {
			m.Offset = int32(i)
		}
	}
	if len(req.Limit) > 0 {
		if i, err2 := strconv.ParseInt(req.Limit, 10, 32); err2 != nil {
			return err2
		} else {
			m.Limit = int32(i)
		}
	}

	return
}

func (m *GetUserProfilePhotos2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetUserProfilePhotosReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Offset) > 0 {
		if i, err2 := strconv.ParseInt(req.Offset, 10, 32); err2 != nil {
			return err2
		} else {
			m.Offset = int32(i)
		}
	}
	if len(req.Limit) > 0 {
		if i, err2 := strconv.ParseInt(req.Limit, 10, 32); err2 != nil {
			return err2
		} else {
			m.Limit = int32(i)
		}
	}

	return
}

type GetFileReq struct {
	FileId string `json:"file_id" form:"file_id"`
}

func (m *GetFileReq) Method() string {
	return "getFile"
}

type GetFile2 struct {
	FileId string `json:"file_id,omitempty"`
}

func (m *GetFile2) NewRequest() BotApiRequest {
	return new(GetFileReq)
}

func (m *GetFile2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetFileReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.FileId) > 0 {
		m.FileId = req.FileId
	}

	return
}

func (m *GetFile2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetFileReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.FileId) > 0 {
		m.FileId = req.FileId
	}

	return
}

type KickChatMemberReq struct {
	ChatId    string `json:"chat_id" form:"chat_id"`
	UserId    string `json:"user_id" form:"user_id"`
	UntilDate string `json:"until_date,omitempty" form:"until_date"`
}

func (m *KickChatMemberReq) Method() string {
	return "kickChatMember"
}

type KickChatMember2 struct {
	ChatId    ChatId2 `json:"chat_id,omitempty"`
	UserId    int32   `json:"user_id,omitempty"`
	UntilDate int32   `json:"until_date,omitempty"`
}

func (m *KickChatMember2) NewRequest() BotApiRequest {
	return new(KickChatMemberReq)
}

func (m *KickChatMember2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*KickChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.UntilDate) > 0 {
		if i, err2 := strconv.ParseInt(req.UntilDate, 10, 32); err2 != nil {
			return err2
		} else {
			m.UntilDate = int32(i)
		}
	}

	return
}

func (m *KickChatMember2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*KickChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.UntilDate) > 0 {
		if i, err2 := strconv.ParseInt(req.UntilDate, 10, 32); err2 != nil {
			return err2
		} else {
			m.UntilDate = int32(i)
		}
	}

	return
}

type UnbanChatMemberReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
	UserId string `json:"user_id" form:"user_id"`
}

func (m *UnbanChatMemberReq) Method() string {
	return "unbanChatMember"
}

type UnbanChatMember2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
	UserId int32   `json:"user_id,omitempty"`
}

func (m *UnbanChatMember2) NewRequest() BotApiRequest {
	return new(UnbanChatMemberReq)
}

func (m *UnbanChatMember2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*UnbanChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}

	return
}

func (m *UnbanChatMember2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*UnbanChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}

	return
}

type RestrictChatMemberReq struct {
	ChatId      string `json:"chat_id" form:"chat_id"`
	UserId      string `json:"user_id" form:"user_id"`
	Permissions string `json:"permissions" form:"permissions"`
	UntilDate   string `json:"until_date,omitempty" form:"until_date"`
}

func (m *RestrictChatMemberReq) Method() string {
	return "restrictChatMember"
}

type RestrictChatMember2 struct {
	ChatId      ChatId2          `json:"chat_id,omitempty"`
	UserId      int32            `json:"user_id,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	UntilDate   int32            `json:"until_date,omitempty"`
}

func (m *RestrictChatMember2) NewRequest() BotApiRequest {
	return new(RestrictChatMemberReq)
}

func (m *RestrictChatMember2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*RestrictChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Permissions) > 0 {
		m.Permissions = new(ChatPermissions)
		if err = m.Permissions.Decode([]byte(req.Permissions)); err != nil {
			return
		}
	}
	if len(req.UntilDate) > 0 {
		if i, err2 := strconv.ParseInt(req.UntilDate, 10, 32); err2 != nil {
			return err2
		} else {
			m.UntilDate = int32(i)
		}
	}

	return
}

func (m *RestrictChatMember2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*RestrictChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Permissions) > 0 {
		m.Permissions = new(ChatPermissions)
		if err = m.Permissions.Decode([]byte(req.Permissions)); err != nil {
			return
		}
	}
	if len(req.UntilDate) > 0 {
		if i, err2 := strconv.ParseInt(req.UntilDate, 10, 32); err2 != nil {
			return err2
		} else {
			m.UntilDate = int32(i)
		}
	}

	return
}

type PromoteChatMemberReq struct {
	ChatId             string `json:"chat_id" form:"chat_id"`
	UserId             string `json:"user_id" form:"user_id"`
	CanChangeInfo      string `json:"can_change_info,omitempty" form:"can_change_info"`
	CanPostMessages    string `json:"can_post_messages,omitempty" form:"can_post_messages"`
	CanEditMessages    string `json:"can_edit_messages,omitempty" form:"can_edit_messages"`
	CanDeleteMessages  string `json:"can_delete_messages,omitempty" form:"can_delete_messages"`
	CanInviteUsers     string `json:"can_invite_users,omitempty" form:"can_invite_users"`
	CanRestrictMembers string `json:"can_restrict_members,omitempty" form:"can_restrict_members"`
	CanPinMessages     string `json:"can_pin_messages,omitempty" form:"can_pin_messages"`
	CanPromoteMembers  string `json:"can_promote_members,omitempty" form:"can_promote_members"`
}

func (m *PromoteChatMemberReq) Method() string {
	return "promoteChatMember"
}

type PromoteChatMember2 struct {
	ChatId             ChatId2 `json:"chat_id,omitempty"`
	UserId             int32   `json:"user_id,omitempty"`
	CanChangeInfo      bool    `json:"can_change_info,omitempty"`
	CanPostMessages    bool    `json:"can_post_messages,omitempty"`
	CanEditMessages    bool    `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool    `json:"can_delete_messages,omitempty"`
	CanInviteUsers     bool    `json:"can_invite_users,omitempty"`
	CanRestrictMembers bool    `json:"can_restrict_members,omitempty"`
	CanPinMessages     bool    `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  bool    `json:"can_promote_members,omitempty"`
}

func (m *PromoteChatMember2) NewRequest() BotApiRequest {
	return new(PromoteChatMemberReq)
}

func (m *PromoteChatMember2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*PromoteChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.CanChangeInfo) > 0 {
		if m.CanChangeInfo, err = strconv.ParseBool(req.CanChangeInfo); err != nil {
			return
		}
	}
	if len(req.CanPostMessages) > 0 {
		if m.CanPostMessages, err = strconv.ParseBool(req.CanPostMessages); err != nil {
			return
		}
	}
	if len(req.CanEditMessages) > 0 {
		if m.CanEditMessages, err = strconv.ParseBool(req.CanEditMessages); err != nil {
			return
		}
	}
	if len(req.CanDeleteMessages) > 0 {
		if m.CanDeleteMessages, err = strconv.ParseBool(req.CanDeleteMessages); err != nil {
			return
		}
	}
	if len(req.CanInviteUsers) > 0 {
		if m.CanInviteUsers, err = strconv.ParseBool(req.CanInviteUsers); err != nil {
			return
		}
	}
	if len(req.CanRestrictMembers) > 0 {
		if m.CanRestrictMembers, err = strconv.ParseBool(req.CanRestrictMembers); err != nil {
			return
		}
	}
	if len(req.CanPinMessages) > 0 {
		if m.CanPinMessages, err = strconv.ParseBool(req.CanPinMessages); err != nil {
			return
		}
	}
	if len(req.CanPromoteMembers) > 0 {
		if m.CanPromoteMembers, err = strconv.ParseBool(req.CanPromoteMembers); err != nil {
			return
		}
	}

	return
}

func (m *PromoteChatMember2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*PromoteChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.CanChangeInfo) > 0 {
		if m.CanChangeInfo, err = strconv.ParseBool(req.CanChangeInfo); err != nil {
			return
		}
	}
	if len(req.CanPostMessages) > 0 {
		if m.CanPostMessages, err = strconv.ParseBool(req.CanPostMessages); err != nil {
			return
		}
	}
	if len(req.CanEditMessages) > 0 {
		if m.CanEditMessages, err = strconv.ParseBool(req.CanEditMessages); err != nil {
			return
		}
	}
	if len(req.CanDeleteMessages) > 0 {
		if m.CanDeleteMessages, err = strconv.ParseBool(req.CanDeleteMessages); err != nil {
			return
		}
	}
	if len(req.CanInviteUsers) > 0 {
		if m.CanInviteUsers, err = strconv.ParseBool(req.CanInviteUsers); err != nil {
			return
		}
	}
	if len(req.CanRestrictMembers) > 0 {
		if m.CanRestrictMembers, err = strconv.ParseBool(req.CanRestrictMembers); err != nil {
			return
		}
	}
	if len(req.CanPinMessages) > 0 {
		if m.CanPinMessages, err = strconv.ParseBool(req.CanPinMessages); err != nil {
			return
		}
	}
	if len(req.CanPromoteMembers) > 0 {
		if m.CanPromoteMembers, err = strconv.ParseBool(req.CanPromoteMembers); err != nil {
			return
		}
	}

	return
}

type SetChatAdministratorCustomTitleReq struct {
	ChatId      string `json:"chat_id" form:"chat_id"`
	UserId      string `json:"user_id" form:"user_id"`
	CustomTitle string `json:"custom_title" form:"custom_title"`
}

func (m *SetChatAdministratorCustomTitleReq) Method() string {
	return "setChatAdministratorCustomTitle"
}

type SetChatAdministratorCustomTitle2 struct {
	ChatId      ChatId2 `json:"chat_id,omitempty"`
	UserId      int32   `json:"user_id,omitempty"`
	CustomTitle string  `json:"custom_title,omitempty"`
}

func (m *SetChatAdministratorCustomTitle2) NewRequest() BotApiRequest {
	return new(SetChatAdministratorCustomTitleReq)
}

func (m *SetChatAdministratorCustomTitle2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatAdministratorCustomTitleReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.CustomTitle) > 0 {
		m.CustomTitle = req.CustomTitle
	}

	return
}

func (m *SetChatAdministratorCustomTitle2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatAdministratorCustomTitleReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.CustomTitle) > 0 {
		m.CustomTitle = req.CustomTitle
	}

	return
}

type SetChatPermissionsReq struct {
	ChatId      string `json:"chat_id" form:"chat_id"`
	Permissions string `json:"permissions" form:"permissions"`
}

func (m *SetChatPermissionsReq) Method() string {
	return "setChatPermissions"
}

type SetChatPermissions2 struct {
	ChatId      ChatId2          `json:"chat_id,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
}

func (m *SetChatPermissions2) NewRequest() BotApiRequest {
	return new(SetChatPermissionsReq)
}

func (m *SetChatPermissions2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatPermissionsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Permissions) > 0 {
		m.Permissions = new(ChatPermissions)
		if err = m.Permissions.Decode([]byte(req.Permissions)); err != nil {
			return
		}
	}

	return
}

func (m *SetChatPermissions2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatPermissionsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Permissions) > 0 {
		m.Permissions = new(ChatPermissions)
		if err = m.Permissions.Decode([]byte(req.Permissions)); err != nil {
			return
		}
	}

	return
}

type ExportChatInviteLinkReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *ExportChatInviteLinkReq) Method() string {
	return "exportChatInviteLink"
}

type ExportChatInviteLink2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *ExportChatInviteLink2) NewRequest() BotApiRequest {
	return new(ExportChatInviteLinkReq)
}

func (m *ExportChatInviteLink2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*ExportChatInviteLinkReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *ExportChatInviteLink2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*ExportChatInviteLinkReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type SetChatPhotoReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
	Photo  string `json:"photo" form:"photo"`
}

func (m *SetChatPhotoReq) Method() string {
	return "setChatPhoto"
}

type SetChatPhoto2 struct {
	ChatId ChatId2   `json:"chat_id,omitempty"`
	Photo  InputFile `json:"photo,omitempty"`
}

func (m *SetChatPhoto2) NewRequest() BotApiRequest {
	return new(SetChatPhotoReq)
}

func (m *SetChatPhoto2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Photo) > 0 {
		m.Photo = MakeInputFile(req.Photo)
	}

	return
}

func (m *SetChatPhoto2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Photo, err = MakeInputFile2(c, contentType, "photo", req.Photo); err != nil {
		return
	}

	return
}

type DeleteChatPhotoReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *DeleteChatPhotoReq) Method() string {
	return "deleteChatPhoto"
}

type DeleteChatPhoto2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *DeleteChatPhoto2) NewRequest() BotApiRequest {
	return new(DeleteChatPhotoReq)
}

func (m *DeleteChatPhoto2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*DeleteChatPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *DeleteChatPhoto2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*DeleteChatPhotoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type SetChatTitleReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
	Title  string `json:"title" form:"title"`
}

func (m *SetChatTitleReq) Method() string {
	return "setChatTitle"
}

type SetChatTitle2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
	Title  string  `json:"title,omitempty"`
}

func (m *SetChatTitle2) NewRequest() BotApiRequest {
	return new(SetChatTitleReq)
}

func (m *SetChatTitle2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatTitleReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}

	return
}

func (m *SetChatTitle2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatTitleReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Title) > 0 {
		m.Title = req.Title
	}

	return
}

type SetChatDescriptionReq struct {
	ChatId      string `json:"chat_id" form:"chat_id"`
	Description string `json:"description,omitempty" form:"description"`
}

func (m *SetChatDescriptionReq) Method() string {
	return "setChatDescription"
}

type SetChatDescription2 struct {
	ChatId      ChatId2 `json:"chat_id,omitempty"`
	Description string  `json:"description,omitempty"`
}

func (m *SetChatDescription2) NewRequest() BotApiRequest {
	return new(SetChatDescriptionReq)
}

func (m *SetChatDescription2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatDescriptionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Description) > 0 {
		m.Description = req.Description
	}

	return
}

func (m *SetChatDescription2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatDescriptionReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.Description) > 0 {
		m.Description = req.Description
	}

	return
}

type PinChatMessageReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	MessageId           string `json:"message_id" form:"message_id"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
}

func (m *PinChatMessageReq) Method() string {
	return "pinChatMessage"
}

type PinChatMessage2 struct {
	ChatId              ChatId2 `json:"chat_id,omitempty"`
	MessageId           int32   `json:"message_id,omitempty"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
}

func (m *PinChatMessage2) NewRequest() BotApiRequest {
	return new(PinChatMessageReq)
}

func (m *PinChatMessage2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*PinChatMessageReq)
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
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}

	return
}

func (m *PinChatMessage2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*PinChatMessageReq)
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
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}

	return
}

type UnpinChatMessageReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *UnpinChatMessageReq) Method() string {
	return "unpinChatMessage"
}

type UnpinChatMessage2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *UnpinChatMessage2) NewRequest() BotApiRequest {
	return new(UnpinChatMessageReq)
}

func (m *UnpinChatMessage2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*UnpinChatMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *UnpinChatMessage2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*UnpinChatMessageReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type LeaveChatReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *LeaveChatReq) Method() string {
	return "leaveChat"
}

type LeaveChat2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *LeaveChat2) NewRequest() BotApiRequest {
	return new(LeaveChatReq)
}

func (m *LeaveChat2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*LeaveChatReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *LeaveChat2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*LeaveChatReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type GetChatReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *GetChatReq) Method() string {
	return "getChat"
}

type GetChat2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *GetChat2) NewRequest() BotApiRequest {
	return new(GetChatReq)
}

func (m *GetChat2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetChatReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *GetChat2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetChatReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type GetChatAdministratorsReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *GetChatAdministratorsReq) Method() string {
	return "getChatAdministrators"
}

type GetChatAdministrators2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *GetChatAdministrators2) NewRequest() BotApiRequest {
	return new(GetChatAdministratorsReq)
}

func (m *GetChatAdministrators2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetChatAdministratorsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *GetChatAdministrators2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetChatAdministratorsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type GetChatMembersCountReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *GetChatMembersCountReq) Method() string {
	return "getChatMembersCount"
}

type GetChatMembersCount2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *GetChatMembersCount2) NewRequest() BotApiRequest {
	return new(GetChatMembersCountReq)
}

func (m *GetChatMembersCount2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetChatMembersCountReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *GetChatMembersCount2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetChatMembersCountReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type GetChatMemberReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
	UserId string `json:"user_id" form:"user_id"`
}

func (m *GetChatMemberReq) Method() string {
	return "getChatMember"
}

type GetChatMember2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
	UserId int32   `json:"user_id,omitempty"`
}

func (m *GetChatMember2) NewRequest() BotApiRequest {
	return new(GetChatMemberReq)
}

func (m *GetChatMember2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}

	return
}

func (m *GetChatMember2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetChatMemberReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}

	return
}

type SetChatStickerSetReq struct {
	ChatId         string `json:"chat_id" form:"chat_id"`
	StickerSetName string `json:"sticker_set_name" form:"sticker_set_name"`
}

func (m *SetChatStickerSetReq) Method() string {
	return "setChatStickerSet"
}

type SetChatStickerSet2 struct {
	ChatId         ChatId2 `json:"chat_id,omitempty"`
	StickerSetName string  `json:"sticker_set_name,omitempty"`
}

func (m *SetChatStickerSet2) NewRequest() BotApiRequest {
	return new(SetChatStickerSetReq)
}

func (m *SetChatStickerSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetChatStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.StickerSetName) > 0 {
		m.StickerSetName = req.StickerSetName
	}

	return
}

func (m *SetChatStickerSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetChatStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if len(req.StickerSetName) > 0 {
		m.StickerSetName = req.StickerSetName
	}

	return
}

type DeleteChatStickerSetReq struct {
	ChatId string `json:"chat_id" form:"chat_id"`
}

func (m *DeleteChatStickerSetReq) Method() string {
	return "deleteChatStickerSet"
}

type DeleteChatStickerSet2 struct {
	ChatId ChatId2 `json:"chat_id,omitempty"`
}

func (m *DeleteChatStickerSet2) NewRequest() BotApiRequest {
	return new(DeleteChatStickerSetReq)
}

func (m *DeleteChatStickerSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*DeleteChatStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

func (m *DeleteChatStickerSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*DeleteChatStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	return
}

type AnswerCallbackQueryReq struct {
	CallbackQueryId string `json:"callback_query_id" form:"callback_query_id"`
	Text            string `json:"text,omitempty" form:"text"`
	ShowAlert       string `json:"show_alert,omitempty" form:"show_alert"`
	Url             string `json:"url,omitempty" form:"url"`
}

func (m *AnswerCallbackQueryReq) Method() string {
	return "answerCallbackQuery"
}

type AnswerCallbackQuery2 struct {
	CallbackQueryId string `json:"callback_query_id,omitempty"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
}

func (m *AnswerCallbackQuery2) NewRequest() BotApiRequest {
	return new(AnswerCallbackQueryReq)
}

func (m *AnswerCallbackQuery2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*AnswerCallbackQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.CallbackQueryId) > 0 {
		m.CallbackQueryId = req.CallbackQueryId
	}
	if len(req.Text) > 0 {
		m.Text = req.Text
	}
	if len(req.ShowAlert) > 0 {
		if m.ShowAlert, err = strconv.ParseBool(req.ShowAlert); err != nil {
			return
		}
	}
	if len(req.Url) > 0 {
		m.Url = req.Url
	}

	return
}

func (m *AnswerCallbackQuery2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*AnswerCallbackQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.CallbackQueryId) > 0 {
		m.CallbackQueryId = req.CallbackQueryId
	}
	if len(req.Text) > 0 {
		m.Text = req.Text
	}
	if len(req.ShowAlert) > 0 {
		if m.ShowAlert, err = strconv.ParseBool(req.ShowAlert); err != nil {
			return
		}
	}
	if len(req.Url) > 0 {
		m.Url = req.Url
	}

	return
}

type SetMyCommandsReq struct {
	Commands string `json:"commands" form:"commands"`
}

func (m *SetMyCommandsReq) Method() string {
	return "setMyCommands"
}

type SetMyCommands2 struct {
	Commands []*BotCommand `json:"commands,omitempty"`
}

func (m *SetMyCommands2) NewRequest() BotApiRequest {
	return new(SetMyCommandsReq)
}

func (m *SetMyCommands2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetMyCommandsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Commands) > 0 {
	}

	return
}

func (m *SetMyCommands2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetMyCommandsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Commands) > 0 {
	}

	return
}

type GetMyCommandsReq struct {
}

func (m *GetMyCommandsReq) Method() string {
	return "getMyCommands"
}

type GetMyCommands2 struct {
}

func (m *GetMyCommands2) NewRequest() BotApiRequest {
	return new(GetMyCommandsReq)
}

func (m *GetMyCommands2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetMyCommandsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

func (m *GetMyCommands2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetMyCommandsReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}
