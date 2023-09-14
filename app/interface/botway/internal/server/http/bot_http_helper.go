package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
	"open.chat/app/interface/botway/botapi"
	"open.chat/pkg/log"
	"strings"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type BotJSON struct {
	botapi.BotApiResponse
}

func writeJSON(w http.ResponseWriter, obj interface{}) (err error) {
	var jsonBytes []byte
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
	if jsonBytes, err = json.Marshal(obj); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

func writeStatusCode(w http.ResponseWriter, ecode int) {
	header := w.Header()
	header.Set("bots-status-code", strconv.FormatInt(int64(ecode), 10))
}

func (r BotJSON) Render(w http.ResponseWriter) error {
	return writeJSON(w, r)
}

func (r BotJSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
}

func botJSONError(c *bm.Context, errorCode int, description string) {
	code := http.StatusOK
	c.Render(code, BotJSON{botapi.BotApiResponse{
		Ok:          false,
		ErrorCode:   errorCode,
		Description: description,
	}})
}

func botJSON(c *bm.Context, data interface{}) error {
	code := http.StatusOK
	r, err := json.Marshal(data)
	if err != nil {
		botJSONError(c, 500, "Internal error")
	} else {
		c.Render(code, BotJSON{botapi.BotApiResponse{
			Ok:     true,
			Result: r,
		}})
	}
	return nil
}

func botHandlerHelper(c *bm.Context, req botapi.BotApiMethod, cb func(c *bm.Context, token string) (interface{}, error)) {
	token, _ := c.Params.Get("token")
	if token == "" {
		botJSONError(c, 401, "Unauthorized")
		return
	} else {
		log.Debugf("token: %s", token)
	}

	var (
		b           binding.Binding
		contentType = c.Request.Header.Get("Content-Type")
	)

	if c.Request.Method == "GET" {
		b = binding.Form
	} else {
		var stripContentTypeParam = func(contentType string) string {
			i := strings.Index(contentType, ";")
			if i != -1 {
				contentType = contentType[:i]
			}
			return contentType
		}

		contentType = stripContentTypeParam(contentType)
		switch contentType {
		case binding.MIMEJSON:
			b = binding.JSON
		case binding.MIMEXML, binding.MIMEXML2:
			b = binding.XML
		case binding.MIMEMultipartPOSTForm:
			b = binding.FormMultipart
		case binding.MIMEPOSTForm:
			b = binding.FormPost
		default:
			b = binding.Form
		}
	}

	bindingReq := req.NewRequest()
	if err := c.BindWith(bindingReq, b); err != nil {
		log.Errorf("bind form error: %v", err)
		botJSONError(c, 401, "Unauthorized")
		return
	}

	if err := req.Decode2(c, contentType, bindingReq); err != nil {
		log.Errorf("decode(%s) error: %v", bindingReq.Method(), err)
		botJSONError(c, 401, "Unauthorized")
		return
	}

	buf, _ := json.Marshal(req)
	log.Debugf("req(%s): %s", bindingReq.Method(), string(buf))

	res, err := cb(c, token)
	if err != nil {
		botJSONError(c, 401, "Unauthorized")
		return
	}

	dBuf, _ := json.Marshal(res)
	log.Debugf("res: %s", string(dBuf))
	botJSON(c, res)
}
