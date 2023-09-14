package botapi

import (
	"encoding/json"
)

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

func (m *ChosenInlineResult) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ChosenInlineResult) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

func (m *InlineQuery) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InlineQuery) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InlineQueryResult struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 string                `json:"url,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int32                 `json:"thumb_width,omitempty"`
	ThumbHeight         int32                 `json:"thumb_height,omitempty"`
	PhotoUrl            string                `json:"photo_url"`
	PhotoWidth          int32                 `json:"photo_width,omitempty"`
	PhotoHeight         int32                 `json:"photo_height,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	GifUrl              string                `json:"gif_url"`
	GifWidth            int32                 `json:"gif_width,omitempty"`
	GifHeight           int32                 `json:"gif_height,omitempty"`
	GifDuration         int32                 `json:"gif_duration,omitempty"`
	Mpeg4Url            string                `json:"mpeg4_url"`
	Mpeg4Width          int32                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int32                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int32                 `json:"mpeg4_duration,omitempty"`
	VideoUrl            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	VideoWidth          int32                 `json:"video_width,omitempty"`
	VideoHeight         int32                 `json:"video_height,omitempty"`
	VideoDuration       int32                 `json:"video_duration,omitempty"`
	AudioUrl            string                `json:"audio_url"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int32                 `json:"audio_duration,omitempty"`
	VoiceUrl            string                `json:"voice_url"`
	VoiceDuration       int32                 `json:"voice_duration,omitempty"`
	DocumentUrl         string                `json:"document_url"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	LivePeriod          int32                 `json:"live_period,omitempty"`
	Address             string                `json:"address"`
	FoursquareId        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
	GameShortName       string                `json:"game_short_name"`
	PhotoFileId         string                `json:"photo_file_id"`
	GifFileId           string                `json:"gif_file_id"`
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	StickerFileId       string                `json:"sticker_file_id"`
	DocumentFileId      string                `json:"document_file_id"`
	VideoFileId         string                `json:"video_file_id"`
	VoiceFileId         string                `json:"voice_file_id"`
	AudioFileId         string                `json:"audio_file_id"`
}

func (m *InlineQueryResult) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InlineQueryResult) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InputMessageContent struct {
	MessageText           string  `json:"message_text"`
	ParseMode             string  `json:"parse_mode"`
	DisableWebPagePreview bool    `json:"disable_web_page_preview,omitempty"`
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	LivePeriod            int32   `json:"live_period,omitempty"`
	Title                 string  `json:"title"`
	Address               string  `json:"address"`
	FoursquareId          string  `json:"foursquare_id,omitempty"`
	FoursquareType        string  `json:"foursquare_type,omitempty"`
	PhoneNumber           string  `json:"phone_number"`
	FirstName             string  `json:"first_name"`
	LastName              string  `json:"last_name,omitempty"`
	Vcard                 string  `json:"vcard,omitempty"`
}

func (m *InputMessageContent) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InputMessageContent) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
