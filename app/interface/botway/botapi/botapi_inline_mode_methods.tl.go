package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type AnswerInlineQueryReq struct {
	InlineQueryId     string `json:"inline_query_id" form:"inline_query_id"`
	Results           string `json:"results,omitempty" form:"results"`
	CacheTime         string `json:"cache_time,omitempty" form:"cache_time"`
	IsPersonal        string `json:"is_personal,omitempty" form:"is_personal"`
	NextOffset        string `json:"next_offset,omitempty" form:"next_offset"`
	SwitchPmText      string `json:"switch_pm_text,omitempty" form:"switch_pm_text"`
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty" form:"switch_pm_parameter"`
}

func (m *AnswerInlineQueryReq) Method() string {
	return "answerInlineQuery"
}

type AnswerInlineQuery2 struct {
	InlineQueryId     string               `json:"inline_query_id,omitempty"`
	Results           []*InlineQueryResult `json:"results,omitempty"`
	CacheTime         int32                `json:"cache_time,omitempty"`
	IsPersonal        bool                 `json:"is_personal,omitempty"`
	NextOffset        string               `json:"next_offset,omitempty"`
	SwitchPmText      string               `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string               `json:"switch_pm_parameter,omitempty"`
}

func (m *AnswerInlineQuery2) NewRequest() BotApiRequest {
	return new(AnswerInlineQueryReq)
}

func (m *AnswerInlineQuery2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*AnswerInlineQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.InlineQueryId) > 0 {
		m.InlineQueryId = req.InlineQueryId
	}
	if len(req.Results) > 0 {
	}
	if len(req.CacheTime) > 0 {
		if i, err2 := strconv.ParseInt(req.CacheTime, 10, 32); err2 != nil {
			return err2
		} else {
			m.CacheTime = int32(i)
		}
	}
	if len(req.IsPersonal) > 0 {
		if m.IsPersonal, err = strconv.ParseBool(req.IsPersonal); err != nil {
			return
		}
	}
	if len(req.NextOffset) > 0 {
		m.NextOffset = req.NextOffset
	}
	if len(req.SwitchPmText) > 0 {
		m.SwitchPmText = req.SwitchPmText
	}
	if len(req.SwitchPmParameter) > 0 {
		m.SwitchPmParameter = req.SwitchPmParameter
	}

	return
}

func (m *AnswerInlineQuery2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*AnswerInlineQueryReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.InlineQueryId) > 0 {
		m.InlineQueryId = req.InlineQueryId
	}
	if len(req.Results) > 0 {
	}
	if len(req.CacheTime) > 0 {
		if i, err2 := strconv.ParseInt(req.CacheTime, 10, 32); err2 != nil {
			return err2
		} else {
			m.CacheTime = int32(i)
		}
	}
	if len(req.IsPersonal) > 0 {
		if m.IsPersonal, err = strconv.ParseBool(req.IsPersonal); err != nil {
			return
		}
	}
	if len(req.NextOffset) > 0 {
		m.NextOffset = req.NextOffset
	}
	if len(req.SwitchPmText) > 0 {
		m.SwitchPmText = req.SwitchPmText
	}
	if len(req.SwitchPmParameter) > 0 {
		m.SwitchPmParameter = req.SwitchPmParameter
	}

	return
}
