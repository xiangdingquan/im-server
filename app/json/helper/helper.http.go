package helper

import (
	"context"
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/render"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/http_client"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type (
	HttpApiConfig struct {
		Url string
	}
)

var HttpApi *HttpApiConfig = nil

func stripContentTypeParam(contentType string) string {
	i := strings.Index(contentType, ";")
	if i != -1 {
		contentType = contentType[:i]
	}
	return contentType
}

func DoHttpJson(c *bm.Context, req interface{}, cb func(context.Context, interface{}) *ResultJSON) {
	var (
		contentType = c.Request.Header.Get("Content-Type")
		r           render.Render
	)
	if stripContentTypeParam(contentType) != binding.MIMEJSON {
		c.Render(http.StatusBadRequest, r)
		return
	}
	if err := binding.JSON.Bind(c.Request, req); err != nil {
		c.Error = ecode.RequestErr
		r = &ResultJSON{
			Code: ecode.RequestErr.Code(),
			Msg:  err.Error(),
		}
		c.Render(http.StatusOK, r)
		return
	}
	log.Infof("request:%s,%s", c.Request.URL, logger.JsonDebugData(req))
	res := cb(context.Background(), req)
	log.Infof("response:%s", logger.JsonDebugData(res))
	c.Render(http.StatusOK, res)
}

func DoHttpUpload(c *bm.Context, cb func(context.Context, []byte) *ResultJSON) {
	var err error
	var l int = 1024
	file := make([]byte, 0)
	buf := make([]byte, l)
	for l >= 1024 {
		l, err = c.Request.Body.Read(buf)
		if err != nil && err == io.EOF {
			err = nil
		}
		if l > 0 {
			file = append(file, buf[:l]...)
		}
	}
	if err != nil {
		c.Error = ecode.RequestErr
		c.Render(http.StatusOK, &ResultJSON{
			Code: ecode.RequestErr.Code(),
			Msg:  err.Error(),
		})
		return
	}
	log.Infof("request:%s,upload[%d]", c.Request.URL, len(file))
	res := cb(context.Background(), file)
	log.Infof("response:%s", logger.JsonDebugData(res))
	c.Render(http.StatusOK, res)
}

func DoHttpDownload(c *bm.Context, req interface{}, cb func(context.Context, interface{}) *[]byte) {
	var (
		contentType = c.Request.Header.Get("Content-Type")
		r           render.Render
	)
	if stripContentTypeParam(contentType) != binding.MIMEJSON {
		c.Render(http.StatusBadRequest, r)
		return
	}
	if err := binding.JSON.Bind(c.Request, req); err != nil {
		c.Error = ecode.RequestErr
		c.Render(http.StatusOK, &ResultJSON{
			Code: ecode.RequestErr.Code(),
			Msg:  err.Error(),
		})
		return
	}
	log.Infof("request:%s,%s", c.Request.URL, logger.JsonDebugData(req))
	res := cb(context.Background(), req)
	if res == nil || len(*res) == 0 {
		log.Infof("response:download[0]")
		c.Render(http.StatusNotFound, &ResultJSON{
			Code: 0,
			Msg:  "resource not exist",
		})
	} else {
		log.Infof("response:download[%d]", len(*res))
		c.Render(http.StatusOK, render.Data{
			ContentType: "application/octet-stream",
			Data:        [][]byte{*res},
		})
	}
}

func DoHttpMultipleDownload(c *bm.Context, req interface{}, cb func(context.Context, interface{}) [][]byte) {
	var (
		contentType = c.Request.Header.Get("Content-Type")
		r           render.Render
	)
	if stripContentTypeParam(contentType) != binding.MIMEJSON {
		c.Render(http.StatusBadRequest, r)
		return
	}
	if err := binding.JSON.Bind(c.Request, req); err != nil {
		c.Error = ecode.RequestErr
		c.Render(http.StatusOK, &ResultJSON{
			Code: ecode.RequestErr.Code(),
			Msg:  err.Error(),
		})
		return
	}
	log.Infof("request:%s,%s", c.Request.URL, logger.JsonDebugData(req))
	res := cb(context.Background(), req)
	if res == nil || len(res) == 0 {
		log.Infof("response:download[0]")
		c.Render(http.StatusNotFound, &ResultJSON{
			Code: 0,
			Msg:  "resource not exist",
		})
	} else {
		log.Infof("response:download[%d]", len(res))
		c.Render(http.StatusOK, render.Data{
			ContentType: "application/octet-stream",
			Data:        res,
		})
	}
}

// DoHandler .
func DoHandler1(c *bm.Context, req interface{}, cb func(context.Context, interface{}) *ResultJSON) {
	var (
		b           binding.Binding
		contentType = c.Request.Header.Get("Content-Type")
		r           render.Render
	)

	if c.Request.Method == "GET" {
		b = binding.Form
	} else {
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
		default: //case MIMEPOSTForm, MIMEMultipartPOSTForm:
			b = binding.Form
		}
	}

	var err error
	if req != nil {
		err = b.Bind(c.Request, req)
		log.Infof("request:%s,%s", c.Request.URL, logger.JsonDebugData(req))
	} else {
		req = make([]byte, 0)
		if b == binding.Form {
			var l int = 1024
			buf := make([]byte, l)
			for l >= 1024 {
				l, err = c.Request.Body.Read(buf)
				if err != nil && err == io.EOF {
					err = nil
				}
				if l > 0 {
					req = append(req.([]byte), buf[:l]...)
				}
			}
		}
		log.Infof("request:%s,data[%d]", c.Request.URL, len(req.([]byte)))
	}

	if err != nil {
		c.Error = ecode.RequestErr
		r = &ResultJSON{
			Code: ecode.RequestErr.Code(),
			Msg:  err.Error(),
		}
		c.Render(http.StatusOK, r)
		return
	}

	res := cb(context.Background(), req)
	log.Infof("response:%s", logger.JsonDebugData(res))
	c.Render(http.StatusOK, res)
}

// DefaultMetadata .
func DefaultMetadata(ctx context.Context, uid uint32, authId int64) (context.Context, error) {
	metaData := &grpc_util.RpcMetadata{
		ServerId:    env.Hostname,
		ClientAddr:  env.Hostname,
		AuthId:      authId,
		SessionId:   0,
		ReceiveTime: time.Now().Unix(),
		UserId:      int32(uid),
		ClientMsgId: 0,
		IsBot:       false,
		Layer:       109,
		IsAdmin:     true,
	}
	return grpc_util.RpcMetadataToIncoming(ctx, metaData)
}

func WebInterface(jsRoute string, body interface{}, result interface{}) (code int, msg string, err error) {
	req := http_client.Post(HttpApi.Url + "/" + jsRoute)
	req.JSONBody(body)
	res := &ResultJSON{}
	err = req.ToJSON(res)
	if err != nil {
		log.Errorf("%s:{%s}", jsRoute, err.Error())
		return
	}
	code = res.Code
	msg = res.Msg
	err = res.GetJSONData(result)
	if err != nil {
		log.Errorf("%s:{%s}", jsRoute, err.Error())
	}
	return
}

func UrlToInputFile(fileUrl string, cb func(fileId int64, filePart int32, bytes []byte) error) (*mtproto.InputFile, error) {
	var (
		fileUpload []byte
		err        error
	)

	req := http_client.Get(fileUrl)
	if strings.HasSuffix(fileUrl, "https://") {
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	req.SetTimeout(5*time.Second, 60*time.Second)
	fileUpload, err = req.Bytes()
	if err != nil {
		log.Errorf("getFileError: %v", err)
	}

	file := mtproto.MakeTLInputFile(&mtproto.InputFile{
		Id:          rand.Int63(),
		Parts:       0,
		Name:        fileUrl,
		Md5Checksum: "",
	}).To_InputFile()

	for i := 0; i < len(fileUpload); i = i + 512*1024 {
		if i+512*1024 > len(fileUpload) {
			file.Parts = int32(i + 1)
			err = cb(file.Id, int32(i), fileUpload[i:len(fileUpload)-i])
		} else {
			err = cb(file.Id, int32(i), fileUpload[i:i+512*1024])
		}
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	return file, nil
}
