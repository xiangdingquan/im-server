package botapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/pkg/net/http/blademaster"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"

	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

type ChatAction string

const (
	Typing            ChatAction = "typing"
	UploadingPhoto    ChatAction = "upload_photo"
	UploadingVideo    ChatAction = "upload_video"
	UploadingAudio    ChatAction = "upload_audio"
	UploadingDocument ChatAction = "upload_document"
	UploadingVNote    ChatAction = "upload_video_note"
	RecordingVideo    ChatAction = "record_video"
	RecordingAudio    ChatAction = "record_audio"
	FindingLocation   ChatAction = "find_location"
)

type ParseMode string

const (
	ModeDefault    ParseMode = ""
	ModeMarkdown   ParseMode = "Markdown"
	ModeMarkdownV2 ParseMode = "MarkdownV2"
	ModeHTML       ParseMode = "HTML"
)

type EntityType string

const (
	EntityMention   EntityType = "mention"
	EntityTMention  EntityType = "text_mention"
	EntityHashtag   EntityType = "hashtag"
	EntityCommand   EntityType = "bot_command"
	EntityURL       EntityType = "url"
	EntityEmail     EntityType = "email"
	EntityBold      EntityType = "bold"
	EntityItalic    EntityType = "italic"
	EntityCode      EntityType = "code"
	EntityCodeBlock EntityType = "pre"
	EntityTextLink  EntityType = "text_link"
)

type ChatType string

const (
	ChatPrivate        ChatType = "private"
	ChatGroup          ChatType = "group"
	ChatSuperGroup     ChatType = "supergroup"
	ChatChannel        ChatType = "channel"
	ChatChannelPrivate ChatType = "privatechannel"
)

type ChatIdType int

const (
	ChatIdPrivate        ChatIdType = 0
	ChatIdGroup          ChatIdType = 1
	ChatIdSuperGroup     ChatIdType = 2
	ChatIdChannel        ChatIdType = 3
	ChatIdChannelPrivate ChatIdType = 4
)

func (c ChatIdType) ToChatId(id int32) int64 {
	return int64(c)<<32 | int64(id)
}

type MemberStatus string

const (
	Creator       MemberStatus = "creator"
	Administrator MemberStatus = "administrator"
	Member        MemberStatus = "member"
	Restricted    MemberStatus = "restricted"
	Left          MemberStatus = "left"
	Kicked        MemberStatus = "kicked"
)

type MaskFeature string

const (
	FeatureForehead MaskFeature = "forehead"
	FeatureEyes     MaskFeature = "eyes"
	FeatureMouth    MaskFeature = "mouth"
	FeatureChin     MaskFeature = "chin"
)

// --------------------------------------------------------------------------------------------
type BotApiRequest interface {
	Method() string
}

type BotApiMethod interface {
	NewRequest() BotApiRequest
	Decode(r BotApiRequest) (err error)
	Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error)
}

type BotApiResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type BotApiMessage interface {
}

func (m *GetUpdates2) CheckAllowUpdate(updateType string) bool {
	if len(m.AllowedUpdates) == 0 {
		return true
	}

	for _, v := range m.AllowedUpdates {
		switch v {
		case "message":
			return true
		case "edited_message":
			return true
		case "channel_post":
			return true
		case "edited_channel_post":
			return true
		case "inline_query":
			return true
		case "chosen_inline_result":
			return true
		case "callback_query":
			return true
		case "shipping_query":
			return true
		case "pre_checkout_query":
			return true
		case "poll":
			return true
		}
	}
	return false
}

type ChatId2 interface {
	IsChannelUserName() bool
}

type ChatID int64

func (c ChatID) IsChannelUserName() bool {
	return false
}

func (c ChatID) ToChatIdTypeId() (int32, int32) {
	return int32(c >> 32), int32(c & 0xffffffff)
}

type ChannelUsername string

func (c ChannelUsername) IsChannelUserName() bool {
	return true
}

func MakeChatId(v string) ChatId2 {
	if len(v) > 0 {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ChatID(i)
		} else {
			return ChannelUsername(v)
		}
	} else {
		return ChatID(0)
	}
}

type ReplyMarkup struct {
	InlineKeyboard  [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	Keyboard        [][]*KeyboardButton       `json:"keyboard,omitempty"`
	ResizeKeyboard  bool                      `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                      `json:"one_time_keyboard,omitempty"`
	Selective       bool                      `json:"selective,omitempty"`
	RemoveKeyboard  bool                      `json:"remove_keyboard,omitempty"`
	ForceReply      bool                      `json:"force_reply,omitempty"`
}

func (m *ReplyMarkup) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

type InputFile interface{}

type InputFileCloud struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
}

type InputFileURL struct {
	FileURL string `json:"file_url"`
}

type InputFileUpload struct {
	FileName   string `json:"file_name,omitempty"`
	FileUpload []byte `json:"-"`
}

func MakeInputFile(s string) InputFile {
	if strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://") {
		return InputFileURL{
			FileURL: s,
		}
	} else if len(s) <= 20 {
		return InputFileCloud{
			FileID: s,
		}
	} else {
		return InputFileUpload{
			FileUpload: hack.Bytes(s),
		}
	}
}

func MakeInputFile2(c *blademaster.Context, contentType string, key, v string) (InputFile, error) {
	switch contentType {
	case binding.MIMEMultipartPOSTForm:
		if len(v) > 0 {
			return InputFileUpload{
				FileUpload: hack.Bytes(v),
			}, nil
		} else {
			file, fileHeader, err := c.Request.FormFile(key)
			if err != nil {
				log.Warnf("upload.%s.file.illegal,err::%v", key, err.Error())
				return nil, nil
			}
			defer file.Close()

			buf := new(bytes.Buffer)
			if _, err = io.Copy(buf, file); err != nil {
				return nil, err
			}

			return InputFileUpload{
				FileName:   fileHeader.Filename,
				FileUpload: buf.Bytes(),
			}, nil
		}
	default:
		if strings.HasPrefix(v, "https://") || strings.HasPrefix(v, "http://") {
			return InputFileURL{
				FileURL: v,
			}, nil
		} else {
			return InputFileCloud{
				FileID: v,
			}, nil
		}
	}
}

func (f InputFileUpload) DebugSaveFile(filePath string) {
	ioutil.WriteFile(filePath+"/"+f.FileName, f.FileUpload, 0644)
}
