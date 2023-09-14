package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type SetPassportDataErrorsReq struct {
	UserId string `json:"user_id" form:"user_id"`
	Errors string `json:"errors" form:"errors"`
}

func (m *SetPassportDataErrorsReq) Method() string {
	return "setPassportDataErrors"
}

type SetPassportDataErrors2 struct {
	UserId int32                   `json:"user_id,omitempty"`
	Errors []*PassportElementError `json:"errors,omitempty"`
}

func (m *SetPassportDataErrors2) NewRequest() BotApiRequest {
	return new(SetPassportDataErrorsReq)
}

func (m *SetPassportDataErrors2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetPassportDataErrorsReq)
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
	if len(req.Errors) > 0 {
	}

	return
}

func (m *SetPassportDataErrors2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetPassportDataErrorsReq)
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
	if len(req.Errors) > 0 {
	}

	return
}
