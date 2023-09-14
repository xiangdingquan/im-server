package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type SendGameReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	GameShortName       string `json:"game_short_name" form:"game_short_name"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendGameReq) Method() string {
	return "sendGame"
}

type SendGame2 struct {
	ChatId              int64                 `json:"chat_id,omitempty"`
	GameShortName       string                `json:"game_short_name,omitempty"`
	DisableNotification bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *SendGame2) NewRequest() BotApiRequest {
	return new(SendGameReq)
}

func (m *SendGame2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendGameReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
			return
		}
	}
	if len(req.GameShortName) > 0 {
		m.GameShortName = req.GameShortName
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

func (m *SendGame2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendGameReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
			return
		}
	}
	if len(req.GameShortName) > 0 {
		m.GameShortName = req.GameShortName
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

type SetGameScoreReq struct {
	UserId             string `json:"user_id" form:"user_id"`
	Score              string `json:"score" form:"score"`
	Force              string `json:"force,omitempty" form:"force"`
	DisableEditMessage string `json:"disable_edit_message,omitempty" form:"disable_edit_message"`
	ChatId             string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId          string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId    string `json:"inline_message_id,omitempty" form:"inline_message_id"`
}

func (m *SetGameScoreReq) Method() string {
	return "setGameScore"
}

type SetGameScore2 struct {
	UserId             int32  `json:"user_id,omitempty"`
	Score              int32  `json:"score,omitempty"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatId             int64  `json:"chat_id,omitempty"`
	MessageId          int32  `json:"message_id,omitempty"`
	InlineMessageId    string `json:"inline_message_id,omitempty"`
}

func (m *SetGameScore2) NewRequest() BotApiRequest {
	return new(SetGameScoreReq)
}

func (m *SetGameScore2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetGameScoreReq)
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
	if len(req.Score) > 0 {
		if i, err2 := strconv.ParseInt(req.Score, 10, 32); err2 != nil {
			return err2
		} else {
			m.Score = int32(i)
		}
	}
	if len(req.Force) > 0 {
		if m.Force, err = strconv.ParseBool(req.Force); err != nil {
			return
		}
	}
	if len(req.DisableEditMessage) > 0 {
		if m.DisableEditMessage, err = strconv.ParseBool(req.DisableEditMessage); err != nil {
			return
		}
	}
	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
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
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}

	return
}

func (m *SetGameScore2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetGameScoreReq)
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
	if len(req.Score) > 0 {
		if i, err2 := strconv.ParseInt(req.Score, 10, 32); err2 != nil {
			return err2
		} else {
			m.Score = int32(i)
		}
	}
	if len(req.Force) > 0 {
		if m.Force, err = strconv.ParseBool(req.Force); err != nil {
			return
		}
	}
	if len(req.DisableEditMessage) > 0 {
		if m.DisableEditMessage, err = strconv.ParseBool(req.DisableEditMessage); err != nil {
			return
		}
	}
	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
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
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}

	return
}

type GetGameHighScoresReq struct {
	UserId          string `json:"user_id" form:"user_id"`
	ChatId          string `json:"chat_id,omitempty" form:"chat_id"`
	MessageId       string `json:"message_id,omitempty" form:"message_id"`
	InlineMessageId string `json:"inline_message_id,omitempty" form:"inline_message_id"`
}

func (m *GetGameHighScoresReq) Method() string {
	return "getGameHighScores"
}

type GetGameHighScores2 struct {
	UserId          int32  `json:"user_id,omitempty"`
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int32  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

func (m *GetGameHighScores2) NewRequest() BotApiRequest {
	return new(GetGameHighScoresReq)
}

func (m *GetGameHighScores2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetGameHighScoresReq)
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
	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
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
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}

	return
}

func (m *GetGameHighScores2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetGameHighScoresReq)
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
	if len(req.ChatId) > 0 {
		if m.ChatId, err = strconv.ParseInt(req.ChatId, 10, 64); err != nil {
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
	if len(req.InlineMessageId) > 0 {
		m.InlineMessageId = req.InlineMessageId
	}

	return
}
