package botapi

import (
	"encoding/json"
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type GetUpdatesReq struct {
	Offset         string `json:"offset,omitempty" form:"offset"`
	Limit          string `json:"limit,omitempty" form:"limit"`
	Timeout        string `json:"timeout,omitempty" form:"timeout"`
	AllowedUpdates string `json:"allowed_updates,omitempty" form:"allowed_updates"`
}

func (m *GetUpdatesReq) Method() string {
	return "getUpdates"
}

type GetUpdates2 struct {
	Offset         int32    `json:"offset,omitempty"`
	Limit          int32    `json:"limit,omitempty"`
	Timeout        int32    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (m *GetUpdates2) NewRequest() BotApiRequest {
	return new(GetUpdatesReq)
}

func (m *GetUpdates2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetUpdatesReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
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
	if len(req.Timeout) > 0 {
		if i, err2 := strconv.ParseInt(req.Timeout, 10, 32); err2 != nil {
			return err2
		} else {
			m.Timeout = int32(i)
		}
	}
	if len(req.AllowedUpdates) > 0 {
		err = json.Unmarshal([]byte(req.AllowedUpdates), &m.AllowedUpdates)
	}

	return
}

func (m *GetUpdates2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetUpdatesReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
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
	if len(req.Timeout) > 0 {
		if i, err2 := strconv.ParseInt(req.Timeout, 10, 32); err2 != nil {
			return err2
		} else {
			m.Timeout = int32(i)
		}
	}
	if len(req.AllowedUpdates) > 0 {
		err = json.Unmarshal([]byte(req.AllowedUpdates), &m.AllowedUpdates)
	}

	return
}

type SetWebhookReq struct {
	Url            string `json:"url" form:"url"`
	Certificate    string `json:"certificate,omitempty" form:"certificate"`
	MaxConnections string `json:"max_connections,omitempty" form:"max_connections"`
	AllowedUpdates string `json:"allowed_updates,omitempty" form:"allowed_updates"`
}

func (m *SetWebhookReq) Method() string {
	return "setWebhook"
}

type SetWebhook2 struct {
	Url            string    `json:"url,omitempty"`
	Certificate    InputFile `json:"certificate,omitempty"`
	MaxConnections int32     `json:"max_connections,omitempty"`
	AllowedUpdates []string  `json:"allowed_updates,omitempty"`
}

func (m *SetWebhook2) NewRequest() BotApiRequest {
	return new(SetWebhookReq)
}

func (m *SetWebhook2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetWebhookReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Url) > 0 {
		m.Url = req.Url
	}
	if len(req.Certificate) > 0 {
		m.Certificate = MakeInputFile(req.Certificate)
	}
	if len(req.MaxConnections) > 0 {
		if i, err2 := strconv.ParseInt(req.MaxConnections, 10, 32); err2 != nil {
			return err2
		} else {
			m.MaxConnections = int32(i)
		}
	}
	if len(req.AllowedUpdates) > 0 {
		err = json.Unmarshal([]byte(req.AllowedUpdates), &m.AllowedUpdates)
	}

	return
}

func (m *SetWebhook2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetWebhookReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Url) > 0 {
		m.Url = req.Url
	}
	if m.Certificate, err = MakeInputFile2(c, contentType, "certificate", req.Certificate); err != nil {
		return
	}

	if len(req.MaxConnections) > 0 {
		if i, err2 := strconv.ParseInt(req.MaxConnections, 10, 32); err2 != nil {
			return err2
		} else {
			m.MaxConnections = int32(i)
		}
	}
	if len(req.AllowedUpdates) > 0 {
		err = json.Unmarshal([]byte(req.AllowedUpdates), &m.AllowedUpdates)
	}

	return
}

type DeleteWebhookReq struct {
}

func (m *DeleteWebhookReq) Method() string {
	return "deleteWebhook"
}

type DeleteWebhook2 struct {
}

func (m *DeleteWebhook2) NewRequest() BotApiRequest {
	return new(DeleteWebhookReq)
}

func (m *DeleteWebhook2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*DeleteWebhookReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

func (m *DeleteWebhook2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*DeleteWebhookReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

type GetWebhookInfoReq struct {
}

func (m *GetWebhookInfoReq) Method() string {
	return "getWebhookInfo"
}

type GetWebhookInfo2 struct {
}

func (m *GetWebhookInfo2) NewRequest() BotApiRequest {
	return new(GetWebhookInfoReq)
}

func (m *GetWebhookInfo2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetWebhookInfoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}

func (m *GetWebhookInfo2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetWebhookInfoReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	return
}
